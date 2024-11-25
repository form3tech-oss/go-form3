// Code generated by go-swagger; DO NOT EDIT.

package users

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.CreateUserCredentials creates a new CreateUserCredentialsRequest object
// with the default values initialized.
func (c *Client) CreateUserCredentials() *CreateUserCredentialsRequest {
	var ()
	return &CreateUserCredentialsRequest{

		UserID: c.Defaults.GetStrfmtUUID("CreateUserCredentials", "user_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateUserCredentialsRequest struct {

	/*UserID      User Id      */

	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateUserCredentialsRequest) FromJson(j string) (*CreateUserCredentialsRequest, error) {

	return o, nil
}

func (o *CreateUserCredentialsRequest) WithUserID(userID strfmt.UUID) *CreateUserCredentialsRequest {

	o.UserID = userID

	return o
}

// ////////////////
// WithContext adds the context to the create user credentials Request
func (o *CreateUserCredentialsRequest) WithContext(ctx context.Context) *CreateUserCredentialsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create user credentials Request
func (o *CreateUserCredentialsRequest) WithHTTPClient(client *http.Client) *CreateUserCredentialsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateUserCredentialsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
