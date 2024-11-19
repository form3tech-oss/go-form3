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

// Client.ListPublicKeysForUser creates a new ListPublicKeysForUserRequest object
// with the default values initialized.
func (c *Client) ListPublicKeysForUser() *ListPublicKeysForUserRequest {
	var ()
	return &ListPublicKeysForUserRequest{

		UserID: c.Defaults.GetStrfmtUUID("ListPublicKeysForUser", "user_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListPublicKeysForUserRequest struct {

	/*UserID      User Id      */

	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ListPublicKeysForUserRequest) FromJson(j string) (*ListPublicKeysForUserRequest, error) {

	return o, nil
}

func (o *ListPublicKeysForUserRequest) WithUserID(userID strfmt.UUID) *ListPublicKeysForUserRequest {

	o.UserID = userID

	return o
}

// ////////////////
// WithContext adds the context to the list public keys for user Request
func (o *ListPublicKeysForUserRequest) WithContext(ctx context.Context) *ListPublicKeysForUserRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list public keys for user Request
func (o *ListPublicKeysForUserRequest) WithHTTPClient(client *http.Client) *ListPublicKeysForUserRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListPublicKeysForUserRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
