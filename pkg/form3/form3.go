package form3

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"net/url"
	"os"

	genClient "github.com/form3tech-oss/go-form3/pkg/generated/client"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	ReqMimeType = "application/vnd.api+json"
)

type F3 struct {
	genClient.Form3Public
	Defaults *ClientDefaults
}

func New(host, pubKeyID string, privateKey *rsa.PrivateKey, orgID string) *F3 {
	u, err := url.Parse(host)
	if err != nil {
		panic(err)
	}

	config := NewRequestSigningClientConfig(pubKeyID, privateKey)
	httpClient := NewRequestSigningHTTPClient(config)

	return newF3(u, httpClient, orgID)
}

func NewFromEnv() *F3 {
	privKeyInPEM := getEnv("FORM3_PRIVATE_KEY")

	block, _ := pem.Decode([]byte(privKeyInPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		panic("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return New(
		getEnv("FORM3_HOST"),
		getEnv("FORM3_PUBLIC_KEY_ID"),
		privateKey,
		getEnv("FORM3_ORGANISATION_ID"),
	)
}

func NewWithTokenBasedAuth(host, clientID, secret, orgID string) *F3 {
	u, err := url.Parse(host)
	if err != nil {
		panic(err)
	}

	config := NewTokenBasedClientConfig(clientID, secret, u)
	httpClient := NewTokenBasedHttpClient(config)

	return newF3(u, httpClient, orgID)
}

func NewWithTokenBasedAuthFromEnv() *F3 {
	return NewWithTokenBasedAuth(
		getEnv("FORM3_HOST"),
		getEnv("FORM3_CLIENT_ID"),
		getEnv("FORM3_CLIENT_SECRET"),
		getEnv("FORM3_ORGANISATION_ID"),
	)
}

func newF3(u *url.URL, c *http.Client, orgID string) *F3 {
	c.Transport = &addUserAgent{c.Transport}

	rt := rc.NewWithClient(u.Host, "/v1", []string{u.Scheme}, c)
	rt.Context = context.Background()

	defaults := NewClientDefaults()
	orgUUID := strfmt.UUID(orgID)
	defaults.OrganisationId = &orgUUID

	pubClient := genClient.New(rt, strfmt.Default, defaults)

	return &F3{
		Form3Public: *pubClient,
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
