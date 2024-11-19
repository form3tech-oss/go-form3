// Code generated by go-swagger; DO NOT EDIT.

package accounts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// Client.ModifyAccount creates a new ModifyAccountRequest object
// with the default values initialized.
func (c *Client) ModifyAccount() *ModifyAccountRequest {
	var ()
	return &ModifyAccountRequest{

		AccountAmendment: models.AccountAmendmentWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("ModifyAccount", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ModifyAccountRequest struct {

	/*AccountAmendRequest*/

	*models.AccountAmendment

	/*ID      Account Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ModifyAccountRequest) FromJson(j string) (*ModifyAccountRequest, error) {

	var m models.AccountAmendment
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.AccountAmendment = &m

	return o, nil
}

func (o *ModifyAccountRequest) WithAccountAmendRequest(accountAmendRequest models.AccountAmendment) *ModifyAccountRequest {

	o.AccountAmendment = &accountAmendRequest

	return o
}

func (o *ModifyAccountRequest) WithoutAccountAmendRequest() *ModifyAccountRequest {

	o.AccountAmendment = &models.AccountAmendment{}

	return o
}

func (o *ModifyAccountRequest) WithID(id strfmt.UUID) *ModifyAccountRequest {

	o.ID = id

	return o
}

// ////////////////
// WithContext adds the context to the modify account Request
func (o *ModifyAccountRequest) WithContext(ctx context.Context) *ModifyAccountRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the modify account Request
func (o *ModifyAccountRequest) WithHTTPClient(client *http.Client) *ModifyAccountRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ModifyAccountRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.AccountAmendment != nil {
		if err := r.SetBodyParam(o.AccountAmendment); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
