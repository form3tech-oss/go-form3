// Code generated by go-swagger; DO NOT EDIT.

package mandates

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

// Client.CreateMandateSubmission creates a new CreateMandateSubmissionRequest object
// with the default values initialized.
func (c *Client) CreateMandateSubmission() *CreateMandateSubmissionRequest {
	var ()
	return &CreateMandateSubmissionRequest{

		MandateSubmissionCreation: models.MandateSubmissionCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreateMandateSubmission", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateMandateSubmissionRequest struct {

	/*MandateSubmissionCreationRequest*/

	*models.MandateSubmissionCreation

	/*ID      Mandate Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateMandateSubmissionRequest) FromJson(j string) (*CreateMandateSubmissionRequest, error) {

	var m models.MandateSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.MandateSubmissionCreation = &m

	return o, nil
}

func (o *CreateMandateSubmissionRequest) WithMandateSubmissionCreationRequest(mandateSubmissionCreationRequest models.MandateSubmissionCreation) *CreateMandateSubmissionRequest {

	o.MandateSubmissionCreation = &mandateSubmissionCreationRequest

	return o
}

func (o *CreateMandateSubmissionRequest) WithoutMandateSubmissionCreationRequest() *CreateMandateSubmissionRequest {

	o.MandateSubmissionCreation = &models.MandateSubmissionCreation{}

	return o
}

func (o *CreateMandateSubmissionRequest) WithID(id strfmt.UUID) *CreateMandateSubmissionRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create mandate submission Request
func (o *CreateMandateSubmissionRequest) WithContext(ctx context.Context) *CreateMandateSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create mandate submission Request
func (o *CreateMandateSubmissionRequest) WithHTTPClient(client *http.Client) *CreateMandateSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateMandateSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.MandateSubmissionCreation != nil {
		if err := r.SetBodyParam(o.MandateSubmissionCreation); err != nil {
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
