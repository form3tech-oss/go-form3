// Code generated by go-swagger; DO NOT EDIT.

package scheme_messages

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.GetSchemeMessage creates a new GetSchemeMessageRequest object
// with the default values initialized.
func (c *Client) GetSchemeMessage() *GetSchemeMessageRequest {
	var ()
	return &GetSchemeMessageRequest{

		ID: c.Defaults.GetStrfmtUUID("GetSchemeMessage", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetSchemeMessageRequest struct {

	/*ID      Scheme Message Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetSchemeMessageRequest) FromJson(j string) (*GetSchemeMessageRequest, error) {

	return o, nil
}

func (o *GetSchemeMessageRequest) WithID(id strfmt.UUID) *GetSchemeMessageRequest {

	o.ID = id

	return o
}

// ////////////////
// WithContext adds the context to the get scheme message Request
func (o *GetSchemeMessageRequest) WithContext(ctx context.Context) *GetSchemeMessageRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get scheme message Request
func (o *GetSchemeMessageRequest) WithHTTPClient(client *http.Client) *GetSchemeMessageRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetSchemeMessageRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
