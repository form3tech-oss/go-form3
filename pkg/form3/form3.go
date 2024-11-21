package form3

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	genClient "github.com/form3tech-oss/go-form3/v7/pkg/generated/client"
	"github.com/go-openapi/runtime"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

const (
	ReqMimeType = "application/vnd.api+json"
	UserAgent   = "go-form3-client"
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

	defaults   *ClientDefaults
	orgID      uuid.UUID
	baseURL    url.URL
	httpClient *http.Client
}

type Option func(*F3)

func New(opts ...Option) (*F3, error) {
	b, err := url.Parse(fmt.Sprintf("https://%s", genClient.DefaultHost))
	if err != nil {
		return nil, fmt.Errorf("could not parse URL: %w", err)
	}

	dt := http.DefaultTransport.(*http.Transport).Clone()
	c := &http.Client{
		Transport: dt,
	}

	f3 := &F3{
		defaults:   nil,
		httpClient: c,
		baseURL:    *b,
	}

	for _, o := range opts {
		o(f3)
	}

	rt := rc.NewWithClient(f3.baseURL.Host, "/v1", []string{f3.baseURL.Scheme}, f3.httpClient)
	rt.Producers["application/vnd.api+json"] = runtime.JSONProducer()
	rt.Consumers["application/vnd.api+json"] = runtime.JSONConsumer()

	defaults := NewClientDefaults()
	orgUUID := strfmt.UUID(f3.orgID.String())
	defaults.OrganisationID = &orgUUID

	pubClient := genClient.New(rt, strfmt.Default, defaults)
	f3.Form3PublicAPI = *pubClient

	return f3, nil
}

func WithHTTPClient(c *http.Client) Option {
	return func(f3 *F3) {
		f3.httpClient = c
	}
}

func WithBaseURL(u url.URL) Option {
	return func(f3 *F3) {
		f3.baseURL = u
	}
}

func WithOrganisationID(u uuid.UUID) Option {
	return func(f3 *F3) {
		f3.orgID = u
	}
}

func WithRequestSigningTransport(opts ...RequestSigningOption) Option {
	return func(f3 *F3) {
		f3.httpClient.Transport = NewRequestSigningTransport(opts...)
	}
}

func WithTokenTransport(opts ...TokenOption) Option {
	return func(f3 *F3) {
		f3.httpClient.Transport = NewTokenTransport(opts...)
	}
}

func NewFromEnv() (*F3, error) {
	privKeyInPEM, err := getEnv("FORM3_PRIVATE_KEY")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(privKeyInPEM))
	if block == nil || (block.Type != "RSA PRIVATE KEY" && block.Type != "PRIVATE KEY") {
		return nil, ErrPEMDecode
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("error parsing privateKey for both PKCS1 and PKCS8: %w", err)
		}
		var ok bool
		privateKey, ok = parsedKey.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("parsed privateKey from PKCS8, but it is not compatible with rsa.PrivateKey")
		}
	}

	publicKeyID, err := getEnv("FORM3_PUBLIC_KEY_ID")
	if err != nil {
		return nil, err
	}
	p, err := uuid.Parse(publicKeyID)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key ID: %w", err)
	}
	opts := []Option{WithRequestSigningTransport(WithPublicKeyID(p), WithPrivateKey(privateKey))}

	host, err := getEnv("FORM3_HOST")
	if err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the base URL: %w", err)
	}

	opts = append(opts, WithBaseURL(*baseURL))

	organisationID, err := getEnv("FORM3_ORGANISATION_ID")
	if err != nil {
		return nil, err
	}

	u, err := uuid.Parse(organisationID)
	if err != nil {
		return nil, fmt.Errorf("could not parse UUID: %w", err)
	}

	return New(append(opts, WithOrganisationID(u))...)
}

func getEnv(name string) (string, error) {
	env, ok := os.LookupEnv(name)
	if !ok {
		return "", &ErrMissingEnvVariable{env: name}
	}

	return env, nil
}
