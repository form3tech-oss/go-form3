// Code generated by go-swagger; DO NOT EDIT.

package accounts

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.DeleteAccount creates a new DeleteAccountRequest object
// with the default values initialized.
func (c *Client) DeleteAccount() *DeleteAccountRequest {
	var ()
	return &DeleteAccountRequest{

		ID: c.Defaults.GetStrfmtUUID("DeleteAccount", "id"),

		Version: c.Defaults.GetInt64("DeleteAccount", "version"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type DeleteAccountRequest struct {

	/*ID      Account Id      */

	ID strfmt.UUID

	/*Version      Version      */

	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *DeleteAccountRequest) FromJson(j string) *DeleteAccountRequest {

	return o
}

func (o *DeleteAccountRequest) WithID(id strfmt.UUID) *DeleteAccountRequest {

	o.ID = id

	return o
}

func (o *DeleteAccountRequest) WithVersion(version int64) *DeleteAccountRequest {

	o.Version = version

	return o
}

//////////////////
// WithContext adds the context to the delete account Request
func (o *DeleteAccountRequest) WithContext(ctx context.Context) *DeleteAccountRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the delete account Request
func (o *DeleteAccountRequest) WithHTTPClient(client *http.Client) *DeleteAccountRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *DeleteAccountRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
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
