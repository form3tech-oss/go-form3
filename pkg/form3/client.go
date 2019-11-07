package form3

import (
	"bytes"
	"io"
	"net/http"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ClientConfig struct {
	ClientId            string
	ClientSecret        string
	HostUrl             *url.URL
	InitialToken        string
	authEndpoint        string
	UnderlyingTransport http.RoundTripper
}

type transport struct {
	config              *ClientConfig
	token               string
	underlyingTransport http.RoundTripper
}

type authJsonResponse struct {
	AccessToken string `json:"access_token"`
}

func NewClientConfig(clientId, clientSecret string, hostUrl *url.URL) *ClientConfig {
	return &ClientConfig{
		ClientId:            clientId,
		ClientSecret:        clientSecret,
		HostUrl:             hostUrl,
		authEndpoint:        "/v1/oauth2/token",
		UnderlyingTransport: http.DefaultTransport,
	}
}

func (c *ClientConfig) WithInitialToken(token string) *ClientConfig {
	c.InitialToken = token
	return c
}

func (c *ClientConfig) WithUnderlyingTransport(underlyingTransport http.RoundTripper) *ClientConfig {
	c.UnderlyingTransport = underlyingTransport
	return c
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {

	if t.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.token))
	}

	var requestBodyClone io.Reader
	if req.Body != nil {
		originalRequestBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(originalRequestBody))
		requestBodyClone = bytes.NewBuffer(originalRequestBody)
	}

	res, err := t.underlyingTransport.RoundTrip(req)
	if res != nil && (res.StatusCode == 401 || res.StatusCode == 403) {
		authRequest, err := http.NewRequest("POST", t.config.HostUrl.String()+t.config.authEndpoint, nil)
		if err != nil {
			return nil, fmt.Errorf("could not build auth request, error: %v", err)
		}
		authRequest.SetBasicAuth(t.config.ClientId, t.config.ClientSecret)

		authResponse, err := t.underlyingTransport.RoundTrip(authRequest)
		if err != nil {
			return nil, fmt.Errorf("error authenticating, error: %v", err)
		}
		defer authResponse.Body.Close()

		if authResponse.StatusCode != 200 {
			return nil, fmt.Errorf("non 200 status code getting token, status code: %d", authResponse.StatusCode)
		}

		authBody, err := ioutil.ReadAll(authResponse.Body)
		if err != nil {
			return nil, fmt.Errorf("could not read auth response body, error: %v", err)
		}

		var authJson authJsonResponse
		err = json.Unmarshal(authBody, &authJson)
		if err != nil {
			return nil, fmt.Errorf("could not parse auth json response, error: %v", err)
		}

		t.token = authJson.AccessToken
		req.Header.Del("Authorization")

		retryRequest, err := http.NewRequest(req.Method, req.URL.String(), requestBodyClone)
		if err != nil {
			return nil, fmt.Errorf("could not build authenticated request, error: %v", err)
		}
		retryRequest.Header = req.Header
		retryRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.token))

		return t.underlyingTransport.RoundTrip(retryRequest)
	}

	return res, err

}

func NewHttpClient(config *ClientConfig) *http.Client {
	transport := &transport{underlyingTransport: config.UnderlyingTransport, config: config}
	transport.token = config.InitialToken

	h := &http.Client{Transport: transport}

	return h
}

func resetRequestBody(req *http.Request) (*http.Request, error) {
	newReq := *req
	body, err := req.GetBody()
	newReq.Body = body
	if err != nil {
		return nil, fmt.Errorf("could not reset request body, error: %v", err)
	}
	req = &newReq
	return req, err
}

