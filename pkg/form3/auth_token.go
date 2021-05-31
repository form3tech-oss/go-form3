package form3

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// #nosec G101
	tokenAuthEndpoint = "/v1/oauth2/token"
)

type TokenClientConfig struct {
	clientID            string
	clientSecret        string
	hostURL             *url.URL
	initialToken        string
	underlyingTransport http.RoundTripper
}

type tokenTransport struct {
	config              *TokenClientConfig
	token               string
	underlyingTransport http.RoundTripper
}

type authJSONResponse struct {
	AccessToken string `json:"access_token"`
}

func NewTokenClientConfig(clientID, clientSecret string, hostURL *url.URL) *TokenClientConfig {
	return &TokenClientConfig{
		clientID:            clientID,
		clientSecret:        clientSecret,
		hostURL:             hostURL,
		underlyingTransport: http.DefaultTransport,
	}
}

func (c *TokenClientConfig) WithInitialToken(token string) *TokenClientConfig {
	c.initialToken = token

	return c
}

func (c *TokenClientConfig) WithUnderlyingTransport(underlyingTransport http.RoundTripper) *TokenClientConfig {
	c.underlyingTransport = underlyingTransport

	return c
}

func (c *TokenClientConfig) WithCertificates(certificates []tls.Certificate) *TokenClientConfig {
	c.underlyingTransport = &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: certificates,
			RootCAs:      nil,
		},
	}

	return c
}

func NewTokenHTTPClient(config *TokenClientConfig) *http.Client {
	transport := &tokenTransport{underlyingTransport: config.underlyingTransport, config: config}
	transport.token = config.initialToken

	h := &http.Client{Transport: transport}

	return h
}

func (t *tokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
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
		authRequest, err := http.NewRequest(http.MethodPost, t.config.hostURL.String()+tokenAuthEndpoint, nil)
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
