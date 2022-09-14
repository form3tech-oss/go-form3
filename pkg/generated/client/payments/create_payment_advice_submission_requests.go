// Code generated by go-swagger; DO NOT EDIT.

package payments

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

// Client.CreatePaymentAdviceSubmission creates a new CreatePaymentAdviceSubmissionRequest object
// with the default values initialized.
func (c *Client) CreatePaymentAdviceSubmission() *CreatePaymentAdviceSubmissionRequest {
	var ()
	return &CreatePaymentAdviceSubmissionRequest{

		AdviceSubmissionCreation: models.AdviceSubmissionCreationWithDefaults(c.Defaults),

		AdviceID: c.Defaults.GetStrfmtUUID("CreatePaymentAdviceSubmission", "adviceId"),

		ID: c.Defaults.GetStrfmtUUID("CreatePaymentAdviceSubmission", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreatePaymentAdviceSubmissionRequest struct {

	/*AdviceSubmissionCreationRequest*/

	*models.AdviceSubmissionCreation

	/*AdviceID      Advice Id      */

	AdviceID strfmt.UUID

	/*ID      Payment Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreatePaymentAdviceSubmissionRequest) FromJson(j string) (*CreatePaymentAdviceSubmissionRequest, error) {

	var m models.AdviceSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.AdviceSubmissionCreation = &m

	return o, nil
}

func (o *CreatePaymentAdviceSubmissionRequest) WithAdviceSubmissionCreationRequest(adviceSubmissionCreationRequest models.AdviceSubmissionCreation) *CreatePaymentAdviceSubmissionRequest {

	o.AdviceSubmissionCreation = &adviceSubmissionCreationRequest

	return o
}

func (o *CreatePaymentAdviceSubmissionRequest) WithoutAdviceSubmissionCreationRequest() *CreatePaymentAdviceSubmissionRequest {

	o.AdviceSubmissionCreation = &models.AdviceSubmissionCreation{}

	return o
}

func (o *CreatePaymentAdviceSubmissionRequest) WithAdviceID(adviceID strfmt.UUID) *CreatePaymentAdviceSubmissionRequest {

	o.AdviceID = adviceID

	return o
}

func (o *CreatePaymentAdviceSubmissionRequest) WithID(id strfmt.UUID) *CreatePaymentAdviceSubmissionRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create payment advice submission Request
func (o *CreatePaymentAdviceSubmissionRequest) WithContext(ctx context.Context) *CreatePaymentAdviceSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create payment advice submission Request
func (o *CreatePaymentAdviceSubmissionRequest) WithHTTPClient(client *http.Client) *CreatePaymentAdviceSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreatePaymentAdviceSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.AdviceSubmissionCreation != nil {
		if err := r.SetBodyParam(o.AdviceSubmissionCreation); err != nil {
			return err
		}
	}

	// path param adviceId
	if err := r.SetPathParam("adviceId", o.AdviceID.String()); err != nil {
		return err
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
