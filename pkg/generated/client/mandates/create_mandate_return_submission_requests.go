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
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// Client.CreateMandateReturnSubmission creates a new CreateMandateReturnSubmissionRequest object
// with the default values initialized.
func (c *Client) CreateMandateReturnSubmission() *CreateMandateReturnSubmissionRequest {
	var ()
	return &CreateMandateReturnSubmissionRequest{

		MandateReturnSubmissionCreation: models.MandateReturnSubmissionCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreateMandateReturnSubmission", "id"),

		ReturnID: c.Defaults.GetStrfmtUUID("CreateMandateReturnSubmission", "returnId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateMandateReturnSubmissionRequest struct {

	/*ReturnSubmissionCreationRequest*/

	*models.MandateReturnSubmissionCreation

	/*ID      Mandate Id      */

	ID strfmt.UUID

	/*ReturnID      Return Id      */

	ReturnID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateMandateReturnSubmissionRequest) FromJson(j string) (*CreateMandateReturnSubmissionRequest, error) {

	var m models.MandateReturnSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.MandateReturnSubmissionCreation = &m

	return o, nil
}

func (o *CreateMandateReturnSubmissionRequest) WithReturnSubmissionCreationRequest(returnSubmissionCreationRequest models.MandateReturnSubmissionCreation) *CreateMandateReturnSubmissionRequest {

	o.MandateReturnSubmissionCreation = &returnSubmissionCreationRequest

	return o
}

func (o *CreateMandateReturnSubmissionRequest) WithoutReturnSubmissionCreationRequest() *CreateMandateReturnSubmissionRequest {

	o.MandateReturnSubmissionCreation = &models.MandateReturnSubmissionCreation{}

	return o
}

func (o *CreateMandateReturnSubmissionRequest) WithID(id strfmt.UUID) *CreateMandateReturnSubmissionRequest {

	o.ID = id

	return o
}

func (o *CreateMandateReturnSubmissionRequest) WithReturnID(returnID strfmt.UUID) *CreateMandateReturnSubmissionRequest {

	o.ReturnID = returnID

	return o
}

// ////////////////
// WithContext adds the context to the create mandate return submission Request
func (o *CreateMandateReturnSubmissionRequest) WithContext(ctx context.Context) *CreateMandateReturnSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create mandate return submission Request
func (o *CreateMandateReturnSubmissionRequest) WithHTTPClient(client *http.Client) *CreateMandateReturnSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateMandateReturnSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.MandateReturnSubmissionCreation != nil {
		if err := r.SetBodyParam(o.MandateReturnSubmissionCreation); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param returnId
	if err := r.SetPathParam("returnId", o.ReturnID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
