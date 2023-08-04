package form3

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type RequestSigningOption func(*RequestSigningTransport)

type RequestSigningTransport struct {
	pubKeyID            uuid.UUID
	privateKey          *rsa.PrivateKey
	userAgent           string
	underlyingTransport http.RoundTripper
}

func WithUserAgent(ua string) RequestSigningOption {
	return func(t *RequestSigningTransport) {
		t.userAgent = ua
	}
}

func WithPrivateKey(key *rsa.PrivateKey) RequestSigningOption {
	return func(t *RequestSigningTransport) {
		t.privateKey = key
	}
}

func WithPublicKeyID(keyID uuid.UUID) RequestSigningOption {
	return func(t *RequestSigningTransport) {
		t.pubKeyID = keyID
	}
}

func WithUnderlyingRequestSigningTransport(tr http.RoundTripper) RequestSigningOption {
	return func(t *RequestSigningTransport) {
		t.underlyingTransport = tr
	}
}

func NewRequestSigningTransport(opts ...RequestSigningOption) *RequestSigningTransport {
	t := &RequestSigningTransport{
		underlyingTransport: http.DefaultTransport,
		userAgent:           UserAgent,
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

func (t *RequestSigningTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", t.userAgent)

	var headers []string
	hasher := sha256.New()
	dt := time.Now().UTC().Format(time.RFC1123)

	msgToSign := fmt.Sprintf(`(request-target): %s %s
host: %s
date: %s`,
		strings.ToLower(req.Method),
		req.URL.RequestURI(),
		req.URL.Host,
		dt,
	)

	headers = append(headers, "(request-target)", "host", "date")

	req.Header.Set("Date", dt)

	// Calculate the digest from payload for POST, PUT, PATCH requests
	if (req.Method == http.MethodPost ||
		req.Method == http.MethodPut ||
		req.Method == http.MethodPatch) &&
		req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading request body: %w", err)
		}

		if _, err := hasher.Write(body); err != nil {
			return nil, fmt.Errorf("error hashing: %w", err)
		}
		digest := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

		msgToSign += fmt.Sprintf(`
content-type: %s
digest: SHA-256=%s
content-length: %d`,
			ReqMimeType,
			digest,
			req.ContentLength,
		)

		headers = append(headers, "content-type", "digest", "content-length")

		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		req.Header.Set("Content-Type", ReqMimeType)
		req.Header.Set("Digest", digest)
	}

	hasher.Reset()
	if _, err := hasher.Write([]byte(msgToSign)); err != nil {
		return nil, fmt.Errorf("error hashing: %w", err)
	}
	hashedMsgToSign := hasher.Sum(nil)

	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		t.privateKey,
		crypto.SHA256,
		hashedMsgToSign,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message %s: %w", msgToSign, err)
	}

	req.Header.Set(
		"Authorization",
		fmt.Sprintf(
			`Signature keyId="%s",algorithm="rsa-sha256",headers="%s", signature="%s"`,
			t.pubKeyID,
			strings.Join(headers, " "),
			base64.StdEncoding.EncodeToString(signature),
		),
	)

	return t.underlyingTransport.RoundTrip(req)
}
