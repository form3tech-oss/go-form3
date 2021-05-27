package form3

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	genClient "github.com/form3tech-oss/go-form3/v2/pkg/generated/client"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	ReqMimeType = "application/vnd.api+json"
)

var ErrPEMDecode = errors.New("failed to decode PEM block containing private key")

type ErrMissingEnvVariable struct {
	env string
}

func (e *ErrMissingEnvVariable) Error() string {
	return fmt.Sprintf("missing environment variable %s", e.env)
}

type F3 struct {
	genClient.Form3PublicAPI
	Defaults *ClientDefaults
}

func New(host, pubKeyID string, privateKey *rsa.PrivateKey, orgID string) (*F3, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	config := NewRequestSigningClientConfig(pubKeyID, privateKey)
	httpClient := NewRequestSigningHTTPClient(config)

	return newF3(u, httpClient, orgID), nil
}

func NewFromEnv() (*F3, error) {
	privKeyInPEM, err := getEnv("FORM3_PRIVATE_KEY")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(privKeyInPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, ErrPEMDecode
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}

	host, err := getEnv("FORM3_HOST")
	if err != nil {
		return nil, err
	}

	publicKeyID, err := getEnv("FORM3_PUBLIC_KEY_ID")
	if err != nil {
		return nil, err
	}

	organisationID, err := getEnv("FORM3_ORGANISATION_ID")
	if err != nil {
		return nil, err
	}

	return New(
		host,
		publicKeyID,
		privateKey,
		organisationID,
	)
}

func NewWithTokenBasedAuth(host, clientID, secret, orgID string) (*F3, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	config := NewTokenBasedClientConfig(clientID, secret, u)
	httpClient := NewTokenBasedHTTPClient(config)

	return newF3(u, httpClient, orgID), nil
}

func NewWithTokenBasedAuthFromEnv() (*F3, error) {
	host, err := getEnv("FORM3_HOST")
	if err != nil {
		return nil, err
	}

	clientID, err := getEnv("FORM3_CLIENT_ID")
	if err != nil {
		return nil, err
	}

	clientSecret, err := getEnv("FORM3_CLIENT_SECRET")
	if err != nil {
		return nil, err
	}

	organisationID, err := getEnv("FORM3_ORGANISATION_ID")
	if err != nil {
		return nil, err
	}

	return NewWithTokenBasedAuth(host, clientID, clientSecret, organisationID)
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
		Form3PublicAPI: *pubClient,
		Defaults:       defaults,
	}
}

func getEnv(name string) (string, error) {
	env, ok := os.LookupEnv(name)
	if !ok {
		return "", &ErrMissingEnvVariable{env: env}
	}
	return env, nil
}
