// Code generated by go-swagger; DO NOT EDIT.

package account_identification

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.GetAccountIdentification creates a new GetAccountIdentificationRequest object
// with the default values initialized.
func (c *Client) GetAccountIdentification() *GetAccountIdentificationRequest {
	var ()
	return &GetAccountIdentificationRequest{

		AccountID: c.Defaults.GetStrfmtUUID("GetAccountIdentification", "account_id"),

		IdentificationID: c.Defaults.GetStrfmtUUID("GetAccountIdentification", "identification_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetAccountIdentificationRequest struct {

	/*AccountID      Account Id      */

	AccountID strfmt.UUID

	/*IdentificationID      Account Identification Id      */

	IdentificationID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetAccountIdentificationRequest) FromJson(j string) (*GetAccountIdentificationRequest, error) {

	return o, nil
}

func (o *GetAccountIdentificationRequest) WithAccountID(accountID strfmt.UUID) *GetAccountIdentificationRequest {

	o.AccountID = accountID

	return o
}

func (o *GetAccountIdentificationRequest) WithIdentificationID(identificationID strfmt.UUID) *GetAccountIdentificationRequest {

	o.IdentificationID = identificationID

	return o
}

// ////////////////
// WithContext adds the context to the get account identification Request
func (o *GetAccountIdentificationRequest) WithContext(ctx context.Context) *GetAccountIdentificationRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get account identification Request
func (o *GetAccountIdentificationRequest) WithHTTPClient(client *http.Client) *GetAccountIdentificationRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetAccountIdentificationRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_id
	if err := r.SetPathParam("account_id", o.AccountID.String()); err != nil {
		return err
	}

	// path param identification_id
	if err := r.SetPathParam("identification_id", o.IdentificationID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}