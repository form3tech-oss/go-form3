// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package payments

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetPositions creates a new GetPositionsRequest object
// with the default values initialized.
func (c *Client) GetPositions() *GetPositionsRequest {

	return &GetPositionsRequest{

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPositionsRequest struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPositionsRequest) FromJson(j string) *GetPositionsRequest {

	return o
}

//////////////////
// WithContext adds the context to the get positions Request
func (o *GetPositionsRequest) WithContext(ctx context.Context) *GetPositionsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get positions Request
func (o *GetPositionsRequest) WithHTTPClient(client *http.Client) *GetPositionsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPositionsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
