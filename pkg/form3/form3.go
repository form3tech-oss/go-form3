package form3

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/form3tech-oss/go-form3/pkg/generated/client"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type F3 struct {
	client.Form3Public
	Defaults *TestDefaults
}

func New() *F3 {
	return NewWithConfig(
		getEnv("FORM3_HOST"),
		getEnv("FORM3_CLIENT_ID"),
		getEnv("FORM3_CLIENT_SECRET"),
		getEnv("FORM3_ORGANISATION_ID"),
	)
}

func NewWithConfig(host, clientID, secret, organisationID string) *F3 {
	u, _ := url.Parse("https://" + host)

	httpclientconf := NewClientConfig(clientID, secret, u)
	httpClient := NewHttpClient(httpclientconf)
	httpClient.Transport = &AddHeaderTransport{httpClient.Transport}

	rt := rc.NewWithClient(host, "/v1", []string{"https"}, httpClient)
	rt.Context = context.Background()

	debug, hasDebug := os.LookupEnv("HTTP_DEBUG")
	rt.SetDebug(hasDebug && debug == "1")

	defaults := NewTestDefaults()
	orgId := strfmt.UUID(organisationID)
	defaults.OrganisationId = &orgId
	cli := client.New(rt, strfmt.Default, defaults)

	return &F3{
		Form3Public: *cli,
		Defaults:    defaults,
	}
}

func getEnv(name string) string {
	env, ok := os.LookupEnv(name)
	if !ok {
		panic("Missing required environment variable " + name)
	}
	return env
}

type AddHeaderTransport struct {
	underlyingTransport http.RoundTripper
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "go-form3-client")
	return adt.underlyingTransport.RoundTrip(req)
}
