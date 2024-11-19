// Code generated by go-swagger; DO NOT EDIT.

package payment_reads

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.GetPaymentRecallDecisionSubmission creates a new GetPaymentRecallDecisionSubmissionRequest object
// with the default values initialized.
func (c *Client) GetPaymentRecallDecisionSubmission() *GetPaymentRecallDecisionSubmissionRequest {
	var ()
	return &GetPaymentRecallDecisionSubmissionRequest{

		DecisionID: c.Defaults.GetStrfmtUUID("GetPaymentRecallDecisionSubmission", "decisionId"),

		ID: c.Defaults.GetStrfmtUUID("GetPaymentRecallDecisionSubmission", "id"),

		RecallID: c.Defaults.GetStrfmtUUID("GetPaymentRecallDecisionSubmission", "recallId"),

		SubmissionID: c.Defaults.GetStrfmtUUID("GetPaymentRecallDecisionSubmission", "submissionId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPaymentRecallDecisionSubmissionRequest struct {

	/*DecisionID      Decision Id      */

	DecisionID strfmt.UUID

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*RecallID      Recall Id      */

	RecallID strfmt.UUID

	/*SubmissionID      Submission Id      */

	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPaymentRecallDecisionSubmissionRequest) FromJson(j string) (*GetPaymentRecallDecisionSubmissionRequest, error) {

	return o, nil
}

func (o *GetPaymentRecallDecisionSubmissionRequest) WithDecisionID(decisionID strfmt.UUID) *GetPaymentRecallDecisionSubmissionRequest {

	o.DecisionID = decisionID

	return o
}

func (o *GetPaymentRecallDecisionSubmissionRequest) WithID(id strfmt.UUID) *GetPaymentRecallDecisionSubmissionRequest {

	o.ID = id

	return o
}

func (o *GetPaymentRecallDecisionSubmissionRequest) WithRecallID(recallID strfmt.UUID) *GetPaymentRecallDecisionSubmissionRequest {

	o.RecallID = recallID

	return o
}

func (o *GetPaymentRecallDecisionSubmissionRequest) WithSubmissionID(submissionID strfmt.UUID) *GetPaymentRecallDecisionSubmissionRequest {

	o.SubmissionID = submissionID

	return o
}

// ////////////////
// WithContext adds the context to the get payment recall decision submission Request
func (o *GetPaymentRecallDecisionSubmissionRequest) WithContext(ctx context.Context) *GetPaymentRecallDecisionSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get payment recall decision submission Request
func (o *GetPaymentRecallDecisionSubmissionRequest) WithHTTPClient(client *http.Client) *GetPaymentRecallDecisionSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPaymentRecallDecisionSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param recallId
	if err := r.SetPathParam("recallId", o.RecallID.String()); err != nil {
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
