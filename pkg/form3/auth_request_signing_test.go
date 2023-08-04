package form3_test

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/form3tech-oss/go-form3/v6/pkg/form3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
