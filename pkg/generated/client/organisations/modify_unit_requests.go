// Code generated by go-swagger; DO NOT EDIT.

package organisations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// Client.ModifyUnit creates a new ModifyUnitRequest object
// with the default values initialized.
func (c *Client) ModifyUnit() *ModifyUnitRequest {
	var ()
	return &ModifyUnitRequest{

		OrganisationUpdate: models.OrganisationUpdateWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("ModifyUnit", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ModifyUnitRequest struct {

	/*OrganisationUpdateRequest*/

	*models.OrganisationUpdate

	/*ID      Organisation Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ModifyUnitRequest) FromJson(j string) (*ModifyUnitRequest, error) {

	var m models.OrganisationUpdate
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.OrganisationUpdate = &m

	return o, nil
}

func (o *ModifyUnitRequest) WithOrganisationUpdateRequest(organisationUpdateRequest models.OrganisationUpdate) *ModifyUnitRequest {

	o.OrganisationUpdate = &organisationUpdateRequest

	return o
}

func (o *ModifyUnitRequest) WithoutOrganisationUpdateRequest() *ModifyUnitRequest {

	o.OrganisationUpdate = &models.OrganisationUpdate{}

	return o
}

func (o *ModifyUnitRequest) WithID(id strfmt.UUID) *ModifyUnitRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the modify unit Request
func (o *ModifyUnitRequest) WithContext(ctx context.Context) *ModifyUnitRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the modify unit Request
func (o *ModifyUnitRequest) WithHTTPClient(client *http.Client) *ModifyUnitRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ModifyUnitRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.OrganisationUpdate != nil {
		if err := r.SetBodyParam(o.OrganisationUpdate); err != nil {
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
