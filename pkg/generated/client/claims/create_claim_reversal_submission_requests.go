// Code generated by go-swagger; DO NOT EDIT.

package claims

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// Client.CreateClaimReversalSubmission creates a new CreateClaimReversalSubmissionRequest object
// with the default values initialized.
func (c *Client) CreateClaimReversalSubmission() *CreateClaimReversalSubmissionRequest {
	var ()
	return &CreateClaimReversalSubmissionRequest{

		ClaimReversalSubmissionCreation: models.ClaimReversalSubmissionCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreateClaimReversalSubmission", "id"),

		ReversalID: c.Defaults.GetStrfmtUUID("CreateClaimReversalSubmission", "reversalId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateClaimReversalSubmissionRequest struct {

	/*ClaimReversalSubmissionCreationRequest*/

	*models.ClaimReversalSubmissionCreation

	/*ID      Claim Id      */

	ID strfmt.UUID

	/*ReversalID      Claim Reversal Id      */

	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateClaimReversalSubmissionRequest) FromJson(j string) *CreateClaimReversalSubmissionRequest {

	var m models.ClaimReversalSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.ClaimReversalSubmissionCreation = &m

	return o
}

func (o *CreateClaimReversalSubmissionRequest) WithClaimReversalSubmissionCreationRequest(claimReversalSubmissionCreationRequest models.ClaimReversalSubmissionCreation) *CreateClaimReversalSubmissionRequest {

	o.ClaimReversalSubmissionCreation = &claimReversalSubmissionCreationRequest

	return o
}

func (o *CreateClaimReversalSubmissionRequest) WithoutClaimReversalSubmissionCreationRequest() *CreateClaimReversalSubmissionRequest {

	o.ClaimReversalSubmissionCreation = &models.ClaimReversalSubmissionCreation{}

	return o
}

func (o *CreateClaimReversalSubmissionRequest) WithID(id strfmt.UUID) *CreateClaimReversalSubmissionRequest {

	o.ID = id

	return o
}

func (o *CreateClaimReversalSubmissionRequest) WithReversalID(reversalID strfmt.UUID) *CreateClaimReversalSubmissionRequest {

	o.ReversalID = reversalID

	return o
}

//////////////////
// WithContext adds the context to the create claim reversal submission Request
func (o *CreateClaimReversalSubmissionRequest) WithContext(ctx context.Context) *CreateClaimReversalSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create claim reversal submission Request
func (o *CreateClaimReversalSubmissionRequest) WithHTTPClient(client *http.Client) *CreateClaimReversalSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateClaimReversalSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.ClaimReversalSubmissionCreation != nil {
		if err := r.SetBodyParam(o.ClaimReversalSubmissionCreation); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
