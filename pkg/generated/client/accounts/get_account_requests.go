// Code generated by go-swagger; DO NOT EDIT.

package accounts

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.GetAccount creates a new GetAccountRequest object
// with the default values initialized.
func (c *Client) GetAccount() *GetAccountRequest {
	var ()
	return &GetAccountRequest{

		ID: c.Defaults.GetStrfmtUUID("GetAccount", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetAccountRequest struct {

	/*ID      Account Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetAccountRequest) FromJson(j string) (*GetAccountRequest, error) {

	return o, nil
}

func (o *GetAccountRequest) WithID(id strfmt.UUID) *GetAccountRequest {

	o.ID = id

	return o
}

// ////////////////
// WithContext adds the context to the get account Request
func (o *GetAccountRequest) WithContext(ctx context.Context) *GetAccountRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get account Request
func (o *GetAccountRequest) WithHTTPClient(client *http.Client) *GetAccountRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetAccountRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
