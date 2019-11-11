// Code generated by go-swagger; DO NOT EDIT.

package users

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetUsersHealth creates a new GetUsersHealthRequest object
// with the default values initialized.
func (c *Client) GetUsersHealth() *GetUsersHealthRequest {

	return &GetUsersHealthRequest{

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetUsersHealthRequest struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetUsersHealthRequest) FromJson(j string) *GetUsersHealthRequest {

	return o
}

//////////////////
// WithContext adds the context to the get users health Request
func (o *GetUsersHealthRequest) WithContext(ctx context.Context) *GetUsersHealthRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get users health Request
func (o *GetUsersHealthRequest) WithHTTPClient(client *http.Client) *GetUsersHealthRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetUsersHealthRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
