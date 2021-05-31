package form3_test

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenAuth(t *testing.T) {
	if os.Getenv("FORM3_CLIENT_ID") == "" || os.Getenv("FORM3_CLIENT_SECRET") == "" {
		t.Skip("WARN: FORM3_CLIENT_ID or FORM3_CLIENT_SECRET not set, skipping test")
	}

	f3, err := form3.NewWithTokenAuthFromEnv()
	require.NoError(t, err)

	req := f3.Users.GetUser()
	req.UserID = strfmt.UUID(os.Getenv("FORM3_CLIENT_ID"))

	resp, err := req.Do()
	require.Nil(t, err)
	assert.Equal(t, resp.Data.ID.String(), os.Getenv("FORM3_CLIENT_ID"))
}

func TestMTLS(t *testing.T) {
	ts := NewLocalHTTPSTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		assert.NotEmpty(t, r.TLS.PeerCertificates)
	}))

	defer ts.Close()

	u, err := url.Parse(ts.URL)
	require.NoError(t, err)

	cert, err := tls.LoadX509KeyPair("../../test/client.crt", "../../test/ca.key")
	require.NoError(t, err)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            nil,
			InsecureSkipVerify: true, // required because we're listening on an IP address as opposed to a domain name
		},
	}

	co := form3.NewTokenClientConfig(os.Getenv("FORM3_CLIENT_ID"), os.Getenv("FORM3_CLIENT_SECRET"), u)
	co.WithUnderlyingTransport(transport)

	c := form3.NewTokenHTTPClient(co)

	f3 := form3.NewF3(u, c, os.Getenv("FORM3_ORGANISATION_ID"))
	req := f3.Users.GetUser()
	_, err = req.Do()
	require.NoError(t, err)
}

func NewLocalHTTPSTestServer(t *testing.T, handler http.Handler) *httptest.Server {
	t.Helper()

	ts := httptest.NewUnstartedServer(handler)

	cert, err := tls.LoadX509KeyPair("../../test/server.crt", "../../test/ca.key")
	require.NoError(t, err)

	caCert, err := ioutil.ReadFile("../../test/root.crt")
	require.NoError(t, err)

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	ts.TLS = &tls.Config{
		ClientCAs:    caCertPool,
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	ts.StartTLS()

	return ts
}
