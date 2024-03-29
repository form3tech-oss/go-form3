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

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// Client.PostDirectdebitsIDDecisionsDecisionIDAdmissions creates a new PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest object
// with the default values initialized.
func (c *Client) PostDirectdebitsIDDecisionsDecisionIDAdmissions() *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest {
	var ()
	return &PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest{

		DirectDebitDecisionAdmissionCreation: models.DirectDebitDecisionAdmissionCreationWithDefaults(c.Defaults),

		DecisionID: c.Defaults.GetStrfmtUUID("PostDirectdebitsIDDecisionsDecisionIDAdmissions", "decisionId"),

		ID: c.Defaults.GetStrfmtUUID("PostDirectdebitsIDDecisionsDecisionIDAdmissions", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest struct {

	/*DirectDebitAdmissionSubmissionCreationRequest*/

	*models.DirectDebitDecisionAdmissionCreation

	/*DecisionID      Direct Debit decision id      */

	DecisionID strfmt.UUID

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest) FromJson(j string) (*PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest, error) {

	var m models.DirectDebitDecisionAdmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.DirectDebitDecisionAdmissionCreation = &m

	return o, nil
}

func (o *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest) WithDirectDebitAdmissionSubmissionCreationRequest(directDebitAdmissionSubmissionCreationRequest models.DirectDebitDecisionAdmissionCreation) *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest {

	o.DirectDebitDecisionAdmissionCreation = &directDebitAdmissionSubmissionCreationRequest

	return o
}

func (o *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest) WithoutDirectDebitAdmissionSubmissionCreationRequest() *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest {

	o.DirectDebitDecisionAdmissionCreation = &models.DirectDebitDecisionAdmissionCreation{}

	return o
}

func (o *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest) WithDecisionID(decisionID strfmt.UUID) *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest {

	o.DecisionID = decisionID

	return o
}

func (o *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest) WithID(id strfmt.UUID) *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the post directdebits ID decisions decision ID admissions Request
func (o *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest) WithContext(ctx context.Context) *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the post directdebits ID decisions decision ID admissions Request
func (o *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest) WithHTTPClient(client *http.Client) *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *PostDirectdebitsIDDecisionsDecisionIDAdmissionsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.DirectDebitDecisionAdmissionCreation != nil {
		if err := r.SetBodyParam(o.DirectDebitDecisionAdmissionCreation); err != nil {
			return err
		}
	}

	// path param decisionId
	if err := r.SetPathParam("decisionId", o.DecisionID.String()); err != nil {
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
