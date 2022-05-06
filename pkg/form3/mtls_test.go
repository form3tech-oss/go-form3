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

	"github.com/form3tech-oss/go-form3/v4/pkg/form3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMTLS(t *testing.T) {
	ts := NewLocalHTTPSTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		assert.NotEmpty(t, r.TLS.PeerCertificates)
	}))

	defer ts.Close()

	u, err := url.Parse(ts.URL)
	require.NoError(t, err)

	cert, err := tls.LoadX509KeyPair("../../test/client.crt", "../../test/client.key")
	require.NoError(t, err)

	uu, err := uuid.Parse("9cdc44af-479c-45a7-9973-07b5b8608d11")
	require.NoError(t, err)

	f3, err := form3.New(
		form3.WithBaseURL(*u),
		form3.WithTokenTransport(
			form3.WithClientID(uu),
			form3.WithClientSecret(os.Getenv("FORM3_CLIENT_SECRET")),
			form3.WithUnderlyingTokenTransport(
				form3.NewMTLSTransport(
					form3.WithCertificate(cert),
					form3.WithIgnoreSkipVerify(true)))))
	require.NoError(t, err)

	req := f3.Users.GetUser()
	_, err = req.Do()
	require.NoError(t, err)
}

func NewLocalHTTPSTestServer(t *testing.T, handler http.Handler) *httptest.Server {
	t.Helper()

	ts := httptest.NewUnstartedServer(handler)

	cert, err := tls.LoadX509KeyPair("../../test/server.crt", "../../test/server.key")
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
