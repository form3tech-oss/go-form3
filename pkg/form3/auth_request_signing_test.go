package form3_test

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/form3tech-oss/go-form3/v7/pkg/form3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestRequestSigningAuth(t *testing.T) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	const keyID = "9cdc44af-479c-45a7-9973-07b5b8608d11"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.Header.Get("User-agent"), "go-form3-client")
		assert.Contains(t, r.Header.Get("Authorization"), fmt.Sprintf("Signature keyId=\"%s\"", keyID))
		assert.Contains(t, r.Header.Get("Authorization"), "algorithm=\"rsa-sha256\"")
		assert.Contains(t, r.Header.Get("Authorization"), "headers=\"(request-target) host date\"")

		reg, err := regexp.Compile("signature=\"(.+)\"")
		require.NoError(t, err)

		signature := reg.FindStringSubmatch(r.Header.Get("Authorization"))[1]

		_, err = base64.StdEncoding.DecodeString(signature)
		assert.NoError(t, err)
	}))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, bytes.NewBuffer([]byte("this is the body of a signed request")))
	require.NoError(t, err)

	u, err := uuid.Parse("9cdc44af-479c-45a7-9973-07b5b8608d11")
	require.NoError(t, err)

	tr := form3.NewRequestSigningTransport(
		form3.WithPrivateKey(privKey),
		form3.WithPublicKeyID(u))

	resp, err := tr.RoundTrip(req)
	defer resp.Body.Close()
	require.NoError(t, err)
}

func TestWithUserAgent(t *testing.T) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	pubKeyID, err := uuid.Parse("9cdc44af-479c-45a7-9973-07b5b8608d11")
	require.NoError(t, err)

	tests := []struct {
		name      string
		opts      []form3.RequestSigningOption
		validator func(t *testing.T) func(w http.ResponseWriter, r *http.Request)
	}{
		{
			"Default",
			[]form3.RequestSigningOption{},
			func(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
				return func(w http.ResponseWriter, r *http.Request) {
					assert.Contains(t, r.Header.Get("User-Agent"), form3.UserAgent)
				}
			},
		},
		{
			"Custom User Agent",
			[]form3.RequestSigningOption{form3.WithUserAgent("Service-v1")},
			func(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
				return func(w http.ResponseWriter, r *http.Request) {
					assert.Contains(t, r.Header.Get("User-Agent"), "Service-v1")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(tt.validator(t)))
			defer ts.Close()

			req, err := http.NewRequest(http.MethodGet, ts.URL, bytes.NewBuffer([]byte("dummy body")))
			require.NoError(t, err)

			opts := []form3.RequestSigningOption{
				form3.WithPrivateKey(privKey),
				form3.WithPublicKeyID(pubKeyID),
			}
			tr := form3.NewRequestSigningTransport(append(opts, tt.opts...)...)

			resp, err := tr.RoundTrip(req)
			defer resp.Body.Close()
			require.NoError(t, err)
		})
	}
}

func TestRequestSigningTransportPOSTJSONBodySigning(t *testing.T) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	pubKeyID, err := uuid.Parse("9cdc44af-479c-45a7-9973-07b5b8608d11")
	require.NoError(t, err)

	body := []byte(`{"data":{"type":"payments"}}`)
	req, err := http.NewRequest(http.MethodPost, "https://example.com/v1/payments", bytes.NewReader(body))
	require.NoError(t, err)

	tr := form3.NewRequestSigningTransport(
		form3.WithPrivateKey(privKey),
		form3.WithPublicKeyID(pubKeyID),
		form3.WithUnderlyingRequestSigningTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
			msgToSign := assertSignedRequest(t, req, &privKey.PublicKey)

			assert.Equal(t, form3.ReqMimeType, req.Header.Get("Content-Type"))
			assert.Equal(t, expectedDigest(body), req.Header.Get("Digest"))
			assert.Equal(t, int64(len(body)), req.ContentLength)
			assert.Contains(t, req.Header.Get("Authorization"), `headers="(request-target) host date content-type digest content-length"`)
			assert.Contains(t, msgToSign, "content-type: "+form3.ReqMimeType)
			assert.Contains(t, msgToSign, fmt.Sprintf("content-length: %d", len(body)))

			gotBody, err := io.ReadAll(req.Body)
			require.NoError(t, err)
			assert.Equal(t, body, gotBody)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(nil)),
				Header:     make(http.Header),
				Request:    req,
			}, nil
		})),
	)

	resp, err := tr.RoundTrip(req)
	require.NoError(t, err)
	defer resp.Body.Close()
}

func TestRequestSigningTransportPOSTMultipartBodySigning(t *testing.T) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	pubKeyID, err := uuid.Parse("9cdc44af-479c-45a7-9973-07b5b8608d11")
	require.NoError(t, err)

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	require.NoError(t, writer.WriteField("name", "value"))
	require.NoError(t, writer.Close())

	body := buf.Bytes()
	contentType := writer.FormDataContentType()

	req, err := http.NewRequest(http.MethodPost, "https://example.com/v1/files", bytes.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", contentType)
	req.ContentLength = -1

	tr := form3.NewRequestSigningTransport(
		form3.WithPrivateKey(privKey),
		form3.WithPublicKeyID(pubKeyID),
		form3.WithUnderlyingRequestSigningTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
			msgToSign := assertSignedRequest(t, req, &privKey.PublicKey)

			assert.Equal(t, contentType, req.Header.Get("Content-Type"))
			assert.Equal(t, expectedDigest(body), req.Header.Get("Digest"))
			assert.Equal(t, int64(len(body)), req.ContentLength)
			assert.Contains(t, msgToSign, "content-type: "+contentType)
			assert.Contains(t, msgToSign, fmt.Sprintf("content-length: %d", len(body)))
			assert.NotContains(t, msgToSign, "content-type: "+form3.ReqMimeType)

			gotBody, err := io.ReadAll(req.Body)
			require.NoError(t, err)
			assert.Equal(t, body, gotBody)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(nil)),
				Header:     make(http.Header),
				Request:    req,
			}, nil
		})),
	)

	resp, err := tr.RoundTrip(req)
	require.NoError(t, err)
	defer resp.Body.Close()
}

func assertSignedRequest(t *testing.T, req *http.Request, publicKey *rsa.PublicKey) string {
	t.Helper()

	msgToSign, signature := signedMessageAndSignature(t, req)
	hashed := sha256.Sum256([]byte(msgToSign))

	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	require.NoError(t, err)

	return msgToSign
}

func signedMessageAndSignature(t *testing.T, req *http.Request) (string, []byte) {
	t.Helper()

	authHeader := req.Header.Get("Authorization")
	reg := regexp.MustCompile(`headers="([^"]+)".*signature="([^"]+)"`)
	matches := reg.FindStringSubmatch(authHeader)
	require.Len(t, matches, 3, "authorization header missing signed headers or signature")

	headers := strings.Split(matches[1], " ")
	lines := make([]string, 0, len(headers))
	for _, header := range headers {
		switch header {
		case "(request-target)":
			lines = append(lines, fmt.Sprintf("(request-target): %s %s", strings.ToLower(req.Method), req.URL.RequestURI()))
		case "host":
			lines = append(lines, "host: "+req.URL.Host)
		case "date":
			lines = append(lines, "date: "+req.Header.Get("Date"))
		case "content-type":
			lines = append(lines, "content-type: "+req.Header.Get("Content-Type"))
		case "digest":
			lines = append(lines, "digest: SHA-256="+req.Header.Get("Digest"))
		case "content-length":
			lines = append(lines, fmt.Sprintf("content-length: %d", req.ContentLength))
		default:
			t.Fatalf("unexpected signed header %q", header)
		}
	}

	signature, err := base64.StdEncoding.DecodeString(matches[2])
	require.NoError(t, err)

	return strings.Join(lines, "\n"), signature
}

func expectedDigest(body []byte) string {
	sum := sha256.Sum256(body)
	return base64.StdEncoding.EncodeToString(sum[:])
}
