// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/a_c_e"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/account_amendment"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/account_identification"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/account_request"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/account_validation"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/accounts"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/audit"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/branch_identification"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/branches"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/claims"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/direct_debits"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/lhv_gateway"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/mandates"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/name_verification_api"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/oauth2"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/organisations"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/payments"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/platformsecurityapi"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/query_api"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/reports"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/roles"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/scheme_file_api"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/scheme_messages"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/subscriptions"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/transaction_file_api"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/client/users"
)

// Default form3 public API HTTP client.
// var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.form3.tech"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

/*
// NewHTTPClient creates a new form3 public API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Form3PublicAPI {
  return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new form3 public API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Form3PublicAPI {
  // ensure nullable parameters have default
  if cfg == nil {
    cfg = DefaultTransportConfig()
  }

  // create transport and client
  transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
  return New(transport, formats)
}
*/

// New creates a new form3 public API client
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Form3PublicAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Form3PublicAPI)
	cli.Transport = transport

	cli.Ace = a_c_e.New(transport, formats, defaults)

	cli.AccountAmendment = account_amendment.New(transport, formats, defaults)

	cli.AccountIdentification = account_identification.New(transport, formats, defaults)

	cli.AccountRequest = account_request.New(transport, formats, defaults)

	cli.AccountValidation = account_validation.New(transport, formats, defaults)

	cli.Accounts = accounts.New(transport, formats, defaults)

	cli.Audit = audit.New(transport, formats, defaults)

	cli.BranchIdentification = branch_identification.New(transport, formats, defaults)

	cli.Branches = branches.New(transport, formats, defaults)

	cli.Claims = claims.New(transport, formats, defaults)

	cli.DirectDebits = direct_debits.New(transport, formats, defaults)

	cli.LhvGateway = lhv_gateway.New(transport, formats, defaults)

	cli.Mandates = mandates.New(transport, formats, defaults)

	cli.NameVerificationAPI = name_verification_api.New(transport, formats, defaults)

	cli.Oauth2 = oauth2.New(transport, formats, defaults)

	cli.Organisations = organisations.New(transport, formats, defaults)

	cli.Payments = payments.New(transport, formats, defaults)

	cli.Platformsecurityapi = platformsecurityapi.New(transport, formats, defaults)

	cli.QueryAPI = query_api.New(transport, formats, defaults)

	cli.Reports = reports.New(transport, formats, defaults)

	cli.Roles = roles.New(transport, formats, defaults)

	cli.SchemeFileAPI = scheme_file_api.New(transport, formats, defaults)

	cli.SchemeMessages = scheme_messages.New(transport, formats, defaults)

	cli.Subscriptions = subscriptions.New(transport, formats, defaults)

	cli.TransactionFileAPI = transaction_file_api.New(transport, formats, defaults)

	cli.Users = users.New(transport, formats, defaults)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Form3PublicAPI is a client for form3 public API
type Form3PublicAPI struct {
	Ace *a_c_e.Client

	AccountAmendment *account_amendment.Client

	AccountIdentification *account_identification.Client

	AccountRequest *account_request.Client

	AccountValidation *account_validation.Client

	Accounts *accounts.Client

	Audit *audit.Client

	BranchIdentification *branch_identification.Client

	Branches *branches.Client

	Claims *claims.Client

	DirectDebits *direct_debits.Client

	LhvGateway *lhv_gateway.Client

	Mandates *mandates.Client

	NameVerificationAPI *name_verification_api.Client

	Oauth2 *oauth2.Client

	Organisations *organisations.Client

	Payments *payments.Client

	Platformsecurityapi *platformsecurityapi.Client

	QueryAPI *query_api.Client

	Reports *reports.Client

	Roles *roles.Client

	SchemeFileAPI *scheme_file_api.Client

	SchemeMessages *scheme_messages.Client

	Subscriptions *subscriptions.Client

	TransactionFileAPI *transaction_file_api.Client

	Users *users.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Form3PublicAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Ace.SetTransport(transport)

	c.AccountAmendment.SetTransport(transport)

	c.AccountIdentification.SetTransport(transport)

	c.AccountRequest.SetTransport(transport)

	c.AccountValidation.SetTransport(transport)

	c.Accounts.SetTransport(transport)

	c.Audit.SetTransport(transport)

	c.BranchIdentification.SetTransport(transport)

	c.Branches.SetTransport(transport)

	c.Claims.SetTransport(transport)

	c.DirectDebits.SetTransport(transport)

	c.LhvGateway.SetTransport(transport)

	c.Mandates.SetTransport(transport)

	c.NameVerificationAPI.SetTransport(transport)

	c.Oauth2.SetTransport(transport)

	c.Organisations.SetTransport(transport)

	c.Payments.SetTransport(transport)

	c.Platformsecurityapi.SetTransport(transport)

	c.QueryAPI.SetTransport(transport)

	c.Reports.SetTransport(transport)

	c.Roles.SetTransport(transport)

	c.SchemeFileAPI.SetTransport(transport)

	c.SchemeMessages.SetTransport(transport)

	c.Subscriptions.SetTransport(transport)

	c.TransactionFileAPI.SetTransport(transport)

	c.Users.SetTransport(transport)

}
