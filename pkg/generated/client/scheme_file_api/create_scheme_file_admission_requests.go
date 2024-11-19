// Code generated by go-swagger; DO NOT EDIT.

package scheme_file_api

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

// Client.CreateSchemeFileAdmission creates a new CreateSchemeFileAdmissionRequest object
// with the default values initialized.
func (c *Client) CreateSchemeFileAdmission() *CreateSchemeFileAdmissionRequest {
	var ()
	return &CreateSchemeFileAdmissionRequest{

		SchemeFileAdmissionCreation: models.SchemeFileAdmissionCreationWithDefaults(c.Defaults),

		SchemeFileID: c.Defaults.GetStrfmtUUID("CreateSchemeFileAdmission", "scheme_file_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateSchemeFileAdmissionRequest struct {

	/*SchemeFileAdmissionCreationRequest*/

	*models.SchemeFileAdmissionCreation

	/*SchemeFileID      Scheme File Id      */

	SchemeFileID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateSchemeFileAdmissionRequest) FromJson(j string) (*CreateSchemeFileAdmissionRequest, error) {

	var m models.SchemeFileAdmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.SchemeFileAdmissionCreation = &m

	return o, nil
}

func (o *CreateSchemeFileAdmissionRequest) WithSchemeFileAdmissionCreationRequest(schemeFileAdmissionCreationRequest models.SchemeFileAdmissionCreation) *CreateSchemeFileAdmissionRequest {

	o.SchemeFileAdmissionCreation = &schemeFileAdmissionCreationRequest

	return o
}

func (o *CreateSchemeFileAdmissionRequest) WithoutSchemeFileAdmissionCreationRequest() *CreateSchemeFileAdmissionRequest {

	o.SchemeFileAdmissionCreation = &models.SchemeFileAdmissionCreation{}

	return o
}

func (o *CreateSchemeFileAdmissionRequest) WithSchemeFileID(schemeFileID strfmt.UUID) *CreateSchemeFileAdmissionRequest {

	o.SchemeFileID = schemeFileID

	return o
}

// ////////////////
// WithContext adds the context to the create scheme file admission Request
func (o *CreateSchemeFileAdmissionRequest) WithContext(ctx context.Context) *CreateSchemeFileAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create scheme file admission Request
func (o *CreateSchemeFileAdmissionRequest) WithHTTPClient(client *http.Client) *CreateSchemeFileAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateSchemeFileAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.SchemeFileAdmissionCreation != nil {
		if err := r.SetBodyParam(o.SchemeFileAdmissionCreation); err != nil {
			return err
		}
	}

	// path param scheme_file_id
	if err := r.SetPathParam("scheme_file_id", o.SchemeFileID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
