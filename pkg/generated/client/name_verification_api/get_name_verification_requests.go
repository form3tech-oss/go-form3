// Code generated by go-swagger; DO NOT EDIT.

package name_verification_api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.GetNameVerification creates a new GetNameVerificationRequest object
// with the default values initialized.
func (c *Client) GetNameVerification() *GetNameVerificationRequest {
	var ()
	return &GetNameVerificationRequest{

		ID: c.Defaults.GetStrfmtUUID("GetNameVerification", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetNameVerificationRequest struct {

	/*ID      Name Verification ID      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetNameVerificationRequest) FromJson(j string) (*GetNameVerificationRequest, error) {

	return o, nil
}

func (o *GetNameVerificationRequest) WithID(id strfmt.UUID) *GetNameVerificationRequest {

	o.ID = id

	return o
}

// ////////////////
// WithContext adds the context to the get name verification Request
func (o *GetNameVerificationRequest) WithContext(ctx context.Context) *GetNameVerificationRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get name verification Request
func (o *GetNameVerificationRequest) WithHTTPClient(client *http.Client) *GetNameVerificationRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetNameVerificationRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
