// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionID creates a new GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest object
// with the default values initialized.
func (c *Client) GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionID() *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest {
	var ()
	return &GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest{

		DecisionID: c.Defaults.GetStrfmtUUID("GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionID", "decisionId"),

		ID: c.Defaults.GetStrfmtUUID("GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionID", "id"),

		SubmissionID: c.Defaults.GetStrfmtUUID("GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionID", "submissionId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest struct {

	/*DecisionID      Direct Debit decision id      */

	DecisionID strfmt.UUID

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	/*SubmissionID      Direct Debit decision submission id      */

	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest) FromJson(j string) (*GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest, error) {

	return o, nil
}

func (o *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest) WithDecisionID(decisionID strfmt.UUID) *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest {

	o.DecisionID = decisionID

	return o
}

func (o *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest) WithID(id strfmt.UUID) *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest {

	o.ID = id

	return o
}

func (o *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest) WithSubmissionID(submissionID strfmt.UUID) *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest {

	o.SubmissionID = submissionID

	return o
}

//////////////////
// WithContext adds the context to the get transaction directdebits ID decisions decision ID submissions submission ID Request
func (o *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest) WithContext(ctx context.Context) *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get transaction directdebits ID decisions decision ID submissions submission ID Request
func (o *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest) WithHTTPClient(client *http.Client) *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param decisionId
	if err := r.SetPathParam("decisionId", o.DecisionID.String()); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param submissionId
	if err := r.SetPathParam("submissionId", o.SubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
