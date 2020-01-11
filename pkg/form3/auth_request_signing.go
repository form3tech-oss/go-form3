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
)

type RequestSigningClientConfig struct {
	pubKeyID            string
	privateKey          *rsa.PrivateKey
	underlyingTransport http.RoundTripper
}

type requestSigningTransport struct {
	config              *RequestSigningClientConfig
	underlyingTransport http.RoundTripper
}

func NewRequestSigningClientConfig(pubKeyID string, privateKey *rsa.PrivateKey) *RequestSigningClientConfig {
	return &RequestSigningClientConfig{
		pubKeyID:            pubKeyID,
		privateKey:          privateKey,
		underlyingTransport: http.DefaultTransport,
	}
}

func (c *RequestSigningClientConfig) WithUnderlyingTransport(underlyingTransport http.RoundTripper) *RequestSigningClientConfig {
	c.underlyingTransport = underlyingTransport
	return c
}

func NewRequestSigningHTTPClient(config *RequestSigningClientConfig) *http.Client {
	transport := &requestSigningTransport{
		underlyingTransport: config.underlyingTransport,
		config:              config,
	}

	h := &http.Client{Transport: transport}

	return h
}

func (t *requestSigningTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	var headers []string
	hasher := sha256.New()
	dt := time.Now().Format(time.RFC1123)

	msgToSign := fmt.Sprintf(`(request-target): %s %s
Host: %s
Date: %s
Accept: %s`,
		strings.ToLower(req.Method),
		req.URL.RequestURI(),
		req.Host,
		dt,
		ReqMimeType,
	)

	headers = append(headers, "(request-target)", "Host", "Date", "Accept")

	// Calculate the digest from payload for POST, PUT, PATCH requests
	if (req.Method == http.MethodPost ||
		req.Method == http.MethodPut ||
		req.Method == http.MethodPatch) &&
		req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}

		if _, err := hasher.Write(body); err != nil {
			return nil, err
		}
		digest := hasher.Sum(nil)

		msgToSign += fmt.Sprintf(`
Content-Type: %s
Digest: SHA-256=%s
Content-Length: %d`,
			ReqMimeType,
			base64.StdEncoding.EncodeToString(digest),
			req.ContentLength,
		)

		headers = append(headers, "Content-Type", "Digest", "Content-Length")

		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		req.Header.Set("Content-Type", ReqMimeType)

	}

	hasher.Reset()
	if _, err := hasher.Write([]byte(msgToSign)); err != nil {
		return nil, err
	}
	hashedMsgToSign := hasher.Sum(nil)

	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		t.config.privateKey,
		crypto.SHA256,
		hashedMsgToSign[:],
	)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message %s: %w", msgToSign, err)
	}

	req.Header.Set(
		"Authorization",
		fmt.Sprintf(
			`Signature: keyId="%s",algorithm="rsa-sha256",headers="%s", signature="%s"`,
			t.config.pubKeyID,
			strings.Join(headers, " "),
			base64.StdEncoding.EncodeToString(signature),
		),
	)

	return t.underlyingTransport.RoundTrip(req)
}
