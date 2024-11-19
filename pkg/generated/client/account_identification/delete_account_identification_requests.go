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
	"github.com/go-openapi/swag"
)

// Client.DeleteAccountIdentification creates a new DeleteAccountIdentificationRequest object
// with the default values initialized.
func (c *Client) DeleteAccountIdentification() *DeleteAccountIdentificationRequest {
	var ()
	return &DeleteAccountIdentificationRequest{

		AccountID: c.Defaults.GetStrfmtUUID("DeleteAccountIdentification", "account_id"),

		IdentificationID: c.Defaults.GetStrfmtUUID("DeleteAccountIdentification", "identification_id"),

		Version: c.Defaults.GetInt64("DeleteAccountIdentification", "version"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type DeleteAccountIdentificationRequest struct {

	/*AccountID      Account Id      */

	AccountID strfmt.UUID

	/*IdentificationID      Account Identification Id      */

	IdentificationID strfmt.UUID

	/*Version      Version      */

	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *DeleteAccountIdentificationRequest) FromJson(j string) (*DeleteAccountIdentificationRequest, error) {

	return o, nil
}

func (o *DeleteAccountIdentificationRequest) WithAccountID(accountID strfmt.UUID) *DeleteAccountIdentificationRequest {

	o.AccountID = accountID

	return o
}

func (o *DeleteAccountIdentificationRequest) WithIdentificationID(identificationID strfmt.UUID) *DeleteAccountIdentificationRequest {

	o.IdentificationID = identificationID

	return o
}

func (o *DeleteAccountIdentificationRequest) WithVersion(version int64) *DeleteAccountIdentificationRequest {

	o.Version = version

	return o
}

// ////////////////
// WithContext adds the context to the delete account identification Request
func (o *DeleteAccountIdentificationRequest) WithContext(ctx context.Context) *DeleteAccountIdentificationRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the delete account identification Request
func (o *DeleteAccountIdentificationRequest) WithHTTPClient(client *http.Client) *DeleteAccountIdentificationRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *DeleteAccountIdentificationRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
