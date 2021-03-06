// Code generated by go-swagger; DO NOT EDIT.

package account_amendment

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetAccountAmendment creates a new GetAccountAmendmentRequest object
// with the default values initialized.
func (c *Client) GetAccountAmendment() *GetAccountAmendmentRequest {
	var ()
	return &GetAccountAmendmentRequest{

		ID: c.Defaults.GetStrfmtUUID("GetAccountAmendment", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetAccountAmendmentRequest struct {

	/*ID      Account Amendment ID      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetAccountAmendmentRequest) FromJson(j string) (*GetAccountAmendmentRequest, error) {

	return o, nil
}

func (o *GetAccountAmendmentRequest) WithID(id strfmt.UUID) *GetAccountAmendmentRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the get account amendment Request
func (o *GetAccountAmendmentRequest) WithContext(ctx context.Context) *GetAccountAmendmentRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get account amendment Request
func (o *GetAccountAmendmentRequest) WithHTTPClient(client *http.Client) *GetAccountAmendmentRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetAccountAmendmentRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
