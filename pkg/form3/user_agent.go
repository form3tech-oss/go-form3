package form3

import "net/http"

type addUserAgent struct {
	underlyingTransport http.RoundTripper
}

func (a *addUserAgent) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "go-form3-client")

	return a.underlyingTransport.RoundTrip(req)
}
