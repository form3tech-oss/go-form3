// Code generated by go-swagger; DO NOT EDIT.

package organisations

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Client.ListUnits creates a new ListUnitsRequest object
// with the default values initialized.
func (c *Client) ListUnits() *ListUnitsRequest {
	var ()
	return &ListUnitsRequest{

		FilterChildOrganisationID: c.Defaults.GetStrfmtUUIDPtr("ListUnits", "filter[child_organisation_id]"),

		FilterOrganisationIds: make([]strfmt.UUID, 0),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListUnitsRequest struct {

	/*FilterChildOrganisationID      Child org id      */

	FilterChildOrganisationID *strfmt.UUID

	/*FilterOrganisationIds      Organisation ids      */

	FilterOrganisationIds []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ListUnitsRequest) FromJson(j string) (*ListUnitsRequest, error) {

	return o, nil
}

func (o *ListUnitsRequest) WithFilterChildOrganisationID(filterChildOrganisationID strfmt.UUID) *ListUnitsRequest {

	o.FilterChildOrganisationID = &filterChildOrganisationID

	return o
}

func (o *ListUnitsRequest) WithoutFilterChildOrganisationID() *ListUnitsRequest {

	o.FilterChildOrganisationID = nil

	return o
}

func (o *ListUnitsRequest) WithFilterOrganisationIds(filterOrganisationIds []strfmt.UUID) *ListUnitsRequest {

	o.FilterOrganisationIds = filterOrganisationIds

	return o
}

func (o *ListUnitsRequest) WithoutFilterOrganisationIds() *ListUnitsRequest {

	o.FilterOrganisationIds = nil

	return o
}

// ////////////////
// WithContext adds the context to the list units Request
func (o *ListUnitsRequest) WithContext(ctx context.Context) *ListUnitsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list units Request
func (o *ListUnitsRequest) WithHTTPClient(client *http.Client) *ListUnitsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListUnitsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterChildOrganisationID != nil {

		// query param filter[child_organisation_id]
		var qrFilterChildOrganisationID strfmt.UUID
		if o.FilterChildOrganisationID != nil {
			qrFilterChildOrganisationID = *o.FilterChildOrganisationID
		}
		qFilterChildOrganisationID := qrFilterChildOrganisationID.String()
		if qFilterChildOrganisationID != "" {
			if err := r.SetQueryParam("filter[child_organisation_id]", qFilterChildOrganisationID); err != nil {
				return err
			}
		}

	}

	var valuesFilterOrganisationIds []string
	for _, v := range o.FilterOrganisationIds {
		valuesFilterOrganisationIds = append(valuesFilterOrganisationIds, v.String())
	}

	joinedFilterOrganisationIds := swag.JoinByFormat(valuesFilterOrganisationIds, "")
	// query array param filter[organisation_ids]
	if err := r.SetQueryParam("filter[organisation_ids]", joinedFilterOrganisationIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
