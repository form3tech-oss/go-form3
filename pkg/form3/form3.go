package form3

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/form3tech-oss/go-form3/pkg/generated/client"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type F3 struct {
	client.Form3Public
	Defaults *TestDefaults
}

type Config struct {
	Host           string
	ClientId       string
	ClientSecret   string
	OrganisationId string
	HttpDebug      bool
	schemes        []string
}

func NewConfig(host, clientID, secret, organisationID string) Config {
	var schemes []string
	if strings.HasPrefix(host, "https://") {
		schemes = append(schemes, "https")
	} else if strings.HasPrefix(host, "http://") {
		schemes = append(schemes, "http")
	} else {
		panic("Unsupported host protocol")
	}

	config := Config{
		Host:           host,
		ClientId:       clientID,
		ClientSecret:   secret,
		OrganisationId: organisationID,
		HttpDebug:      false,
		schemes:        schemes,
	}
	return config
}

func New() *F3 {
	return NewWithConfig(
		NewConfig(
			getEnv("FORM3_HOST"),
			getEnv("FORM3_CLIENT_ID"),
			getEnv("FORM3_CLIENT_SECRET"),
			getEnv("FORM3_ORGANISATION_ID"),
		),
	)
}

func NewWithConfig(config Config) *F3 {
	u, _ := url.Parse(config.Host)

	httpclientconf := NewClientConfig(config.ClientId, config.ClientSecret, u)
	httpClient := NewHttpClient(httpclientconf)
	httpClient.Transport = &AddHeaderTransport{httpClient.Transport}

	rt := rc.NewWithClient(config.Host, "/v1", config.schemes, httpClient)
	rt.Context = context.Background()

	rt.SetDebug(config.HttpDebug)

	defaults := NewTestDefaults()
	orgId := strfmt.UUID(config.OrganisationId)
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
