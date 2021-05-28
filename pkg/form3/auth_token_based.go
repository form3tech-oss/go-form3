package form3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// #nosec G101
	tokenBasedAuthEndpoint = "/v1/oauth2/token"
)

type TokenBasedClientConfig struct {
	clientID            string
	clientSecret        string
	hostURL             *url.URL
	initialToken        string
	underlyingTransport http.RoundTripper
}

type tokenBasedTransport struct {
	config              *TokenBasedClientConfig
	token               string
	underlyingTransport http.RoundTripper
}

type authJSONResponse struct {
	AccessToken string `json:"access_token"`
}

func NewTokenBasedClientConfig(clientID, clientSecret string, hostURL *url.URL) *TokenBasedClientConfig {
	return &TokenBasedClientConfig{
		clientID:            clientID,
		clientSecret:        clientSecret,
		hostURL:             hostURL,
		underlyingTransport: http.DefaultTransport,
	}
}

func (c *TokenBasedClientConfig) WithinitialToken(token string) *TokenBasedClientConfig {
	c.initialToken = token
	return c
}

func (c *TokenBasedClientConfig) WithUnderlyingTransport(underlyingTransport http.RoundTripper) *TokenBasedClientConfig {
	c.underlyingTransport = underlyingTransport
	return c
}

func NewTokenBasedHTTPClient(config *TokenBasedClientConfig) *http.Client {
	transport := &tokenBasedTransport{underlyingTransport: config.underlyingTransport, config: config}
	transport.token = config.initialToken

	h := &http.Client{Transport: transport}

	return h
}

func (t *tokenBasedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.token))
	}

	var requestBodyClone io.Reader

	if req.Body != nil {
		originalRequestBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("could read request body: %w", err)
		}

		req.Body = ioutil.NopCloser(bytes.NewBuffer(originalRequestBody))
		requestBodyClone = bytes.NewBuffer(originalRequestBody)
	}

	res, err := t.underlyingTransport.RoundTrip(req)
	if res != nil && (res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden) {
		authRequest, err := http.NewRequest(http.MethodPost, t.config.hostURL.String()+tokenBasedAuthEndpoint, nil)
		if err != nil {
			return nil, fmt.Errorf("could not build auth request, error: %w", err)
		}
		authRequest.SetBasicAuth(t.config.clientID, t.config.clientSecret)

		authResponse, err := t.underlyingTransport.RoundTrip(authRequest)
		if err != nil {
			return nil, fmt.Errorf("error authenticating, error: %w", err)
		}
		defer authResponse.Body.Close()

		if authResponse.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("non 200 status code getting token, got status code: %d", authResponse.StatusCode)
		}

		authBody, err := ioutil.ReadAll(authResponse.Body)
		if err != nil {
			return nil, fmt.Errorf("could not read auth response body, error: %w", err)
		}

		var authJSON authJSONResponse
		err = json.Unmarshal(authBody, &authJSON)
		if err != nil {
			return nil, fmt.Errorf("could not parse auth json response, error: %w", err)
		}

		t.token = authJSON.AccessToken
		req.Header.Del("Authorization")

		retryRequest, err := http.NewRequest(req.Method, req.URL.String(), requestBodyClone)
		if err != nil {
			return nil, fmt.Errorf("could not build authenticated request, error: %w", err)
		}

		retryRequest.Header = req.Header
		retryRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.token))

		return t.underlyingTransport.RoundTrip(retryRequest)
	}

	return res, err
}
