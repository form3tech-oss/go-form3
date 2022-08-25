// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

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

	"github.com/form3tech-oss/go-form3/v5/pkg/generated/models"
)

// Client.CreateDirectDebitReturnSubmission creates a new CreateDirectDebitReturnSubmissionRequest object
// with the default values initialized.
func (c *Client) CreateDirectDebitReturnSubmission() *CreateDirectDebitReturnSubmissionRequest {
	var ()
	return &CreateDirectDebitReturnSubmissionRequest{

		DirectDebitReturnSubmissionCreation: models.DirectDebitReturnSubmissionCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreateDirectDebitReturnSubmission", "id"),

		ReturnID: c.Defaults.GetStrfmtUUID("CreateDirectDebitReturnSubmission", "returnId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateDirectDebitReturnSubmissionRequest struct {

	/*ReturnSubmissionCreationRequest*/

	*models.DirectDebitReturnSubmissionCreation

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	/*ReturnID      Return Id      */

	ReturnID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateDirectDebitReturnSubmissionRequest) FromJson(j string) (*CreateDirectDebitReturnSubmissionRequest, error) {

	var m models.DirectDebitReturnSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.DirectDebitReturnSubmissionCreation = &m

	return o, nil
}

func (o *CreateDirectDebitReturnSubmissionRequest) WithReturnSubmissionCreationRequest(returnSubmissionCreationRequest models.DirectDebitReturnSubmissionCreation) *CreateDirectDebitReturnSubmissionRequest {

	o.DirectDebitReturnSubmissionCreation = &returnSubmissionCreationRequest

	return o
}

func (o *CreateDirectDebitReturnSubmissionRequest) WithoutReturnSubmissionCreationRequest() *CreateDirectDebitReturnSubmissionRequest {

	o.DirectDebitReturnSubmissionCreation = &models.DirectDebitReturnSubmissionCreation{}

	return o
}

func (o *CreateDirectDebitReturnSubmissionRequest) WithID(id strfmt.UUID) *CreateDirectDebitReturnSubmissionRequest {

	o.ID = id

	return o
}

func (o *CreateDirectDebitReturnSubmissionRequest) WithReturnID(returnID strfmt.UUID) *CreateDirectDebitReturnSubmissionRequest {

	o.ReturnID = returnID

	return o
}

//////////////////
// WithContext adds the context to the create direct debit return submission Request
func (o *CreateDirectDebitReturnSubmissionRequest) WithContext(ctx context.Context) *CreateDirectDebitReturnSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create direct debit return submission Request
func (o *CreateDirectDebitReturnSubmissionRequest) WithHTTPClient(client *http.Client) *CreateDirectDebitReturnSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateDirectDebitReturnSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.DirectDebitReturnSubmissionCreation != nil {
		if err := r.SetBodyParam(o.DirectDebitReturnSubmissionCreation); err != nil {
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
