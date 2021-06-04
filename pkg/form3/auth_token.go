package form3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

// #nosec G101
const tokenBasedAuthEndpoint = "/v1/oauth2/token"

type authJSONResponse struct {
	AccessToken string `json:"access_token"`
}

type TokenOption func(*TokenTransport)

type TokenTransport struct {
	clientSecret        string
	clientID            uuid.UUID
	token               string
	underlyingTransport http.RoundTripper
}

// Deprecated: token based authentication is deprecated and will be removed at
// some point in the future in favour of request signing.
func WithClientID(clientID uuid.UUID) TokenOption {
	return func(t *TokenTransport) {
		t.clientID = clientID
	}
}

// Deprecated: token based authentication is deprecated and will be removed at
// some point in the future in favour of request signing.
func WithClientSecret(secret string) TokenOption {
	return func(t *TokenTransport) {
		t.clientSecret = secret
	}
}

// Deprecated: token based authentication is deprecated and will be removed at
// some point in the future in favour of request signing.
func WithInitialToken(token string) TokenOption {
	return func(t *TokenTransport) {
		t.token = token
	}
}

// Deprecated: token based authentication is deprecated and will be removed at
// some point in the future in favour of request signing.
func WithUnderlyingTokenTransport(tr http.RoundTripper) TokenOption {
	return func(t *TokenTransport) {
		t.underlyingTransport = tr
	}
}

// Deprecated: token based authentication is deprecated and will be removed at
// some point in the future in favour of request signing.
func NewTokenTransport(opts ...TokenOption) *TokenTransport {
	t := &TokenTransport{
		underlyingTransport: http.DefaultTransport,
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

func (t *TokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", UserAgent)

	if t.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.token))
	}

	var requestBodyClone io.Reader
	if req.Body != nil {
		originalRequestBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("could not read original request body: %w", err)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(originalRequestBody))
		requestBodyClone = bytes.NewBuffer(originalRequestBody)
	}

	res, err := t.underlyingTransport.RoundTrip(req)
	if res != nil && (res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden) {
		authRequest, err := http.NewRequest(http.MethodPost, req.URL.Scheme+"://"+req.URL.Host+tokenBasedAuthEndpoint, nil)
		if err != nil {
			return nil, fmt.Errorf("could not build auth request: %w", err)
		}
		authRequest.SetBasicAuth(t.clientID.String(), t.clientSecret)

		authResponse, err := t.underlyingTransport.RoundTrip(authRequest)
		if err != nil {
			return nil, fmt.Errorf("error authenticating: %w", err)
		}
		defer authResponse.Body.Close()

		if authResponse.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("non 200 status code getting token, status code: %d", authResponse.StatusCode)
		}

		authBody, err := ioutil.ReadAll(authResponse.Body)
		if err != nil {
			return nil, fmt.Errorf("could not read auth response body: %w", err)
		}

		var authJSON authJSONResponse
		err = json.Unmarshal(authBody, &authJSON)
		if err != nil {
			return nil, fmt.Errorf("could not parse auth json response: %w", err)
		}

		t.token = authJSON.AccessToken
		req.Header.Del("Authorization")

		retryRequest, err := http.NewRequest(req.Method, req.URL.String(), requestBodyClone)
		if err != nil {
			return nil, fmt.Errorf("could not build authenticated request: %w", err)
		}
		retryRequest.Header = req.Header
		retryRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.token))

		return t.underlyingTransport.RoundTrip(retryRequest)
	}

	return res, err
}
