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

// Client.DeleteUserRole creates a new DeleteUserRoleRequest object
// with the default values initialized.
func (c *Client) DeleteUserRole() *DeleteUserRoleRequest {
	var ()
	return &DeleteUserRoleRequest{

		RoleID: c.Defaults.GetStrfmtUUID("DeleteUserRole", "role_id"),

		UserID: c.Defaults.GetStrfmtUUID("DeleteUserRole", "user_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type DeleteUserRoleRequest struct {

	/*RoleID      Role Id      */

	RoleID strfmt.UUID

	/*UserID      User Id      */

	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *DeleteUserRoleRequest) FromJson(j string) (*DeleteUserRoleRequest, error) {

	return o, nil
}

func (o *DeleteUserRoleRequest) WithRoleID(roleID strfmt.UUID) *DeleteUserRoleRequest {

	o.RoleID = roleID

	return o
}

func (o *DeleteUserRoleRequest) WithUserID(userID strfmt.UUID) *DeleteUserRoleRequest {

	o.UserID = userID

	return o
}

// ////////////////
// WithContext adds the context to the delete user role Request
func (o *DeleteUserRoleRequest) WithContext(ctx context.Context) *DeleteUserRoleRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the delete user role Request
func (o *DeleteUserRoleRequest) WithHTTPClient(client *http.Client) *DeleteUserRoleRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *DeleteUserRoleRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID.String()); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
