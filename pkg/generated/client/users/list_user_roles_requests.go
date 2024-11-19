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

// Client.ListUserRoles creates a new ListUserRolesRequest object
// with the default values initialized.
func (c *Client) ListUserRoles() *ListUserRolesRequest {
	var ()
	return &ListUserRolesRequest{

		UserID: c.Defaults.GetStrfmtUUID("ListUserRoles", "user_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListUserRolesRequest struct {

	/*UserID      User Id      */

	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ListUserRolesRequest) FromJson(j string) (*ListUserRolesRequest, error) {

	return o, nil
}

func (o *ListUserRolesRequest) WithUserID(userID strfmt.UUID) *ListUserRolesRequest {

	o.UserID = userID

	return o
}

// ////////////////
// WithContext adds the context to the list user roles Request
func (o *ListUserRolesRequest) WithContext(ctx context.Context) *ListUserRolesRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list user roles Request
func (o *ListUserRolesRequest) WithHTTPClient(client *http.Client) *ListUserRolesRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListUserRolesRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
