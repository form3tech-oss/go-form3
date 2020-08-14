// Code generated by go-swagger; DO NOT EDIT.

package payments

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetPaymentAdviceSubmission creates a new GetPaymentAdviceSubmissionRequest object
// with the default values initialized.
func (c *Client) GetPaymentAdviceSubmission() *GetPaymentAdviceSubmissionRequest {
	var ()
	return &GetPaymentAdviceSubmissionRequest{

		AdviceID: c.Defaults.GetStrfmtUUID("GetPaymentAdviceSubmission", "adviceId"),

		ID: c.Defaults.GetStrfmtUUID("GetPaymentAdviceSubmission", "id"),

		SubmissionID: c.Defaults.GetStrfmtUUID("GetPaymentAdviceSubmission", "submissionId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPaymentAdviceSubmissionRequest struct {

	/*AdviceID      Advice Id      */

	AdviceID strfmt.UUID

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*SubmissionID      Submission Id      */

	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPaymentAdviceSubmissionRequest) FromJson(j string) *GetPaymentAdviceSubmissionRequest {

	return o
}

func (o *GetPaymentAdviceSubmissionRequest) WithAdviceID(adviceID strfmt.UUID) *GetPaymentAdviceSubmissionRequest {

	o.AdviceID = adviceID

	return o
}

func (o *GetPaymentAdviceSubmissionRequest) WithID(id strfmt.UUID) *GetPaymentAdviceSubmissionRequest {

	o.ID = id

	return o
}

func (o *GetPaymentAdviceSubmissionRequest) WithSubmissionID(submissionID strfmt.UUID) *GetPaymentAdviceSubmissionRequest {

	o.SubmissionID = submissionID

	return o
}

//////////////////
// WithContext adds the context to the get payment advice submission Request
func (o *GetPaymentAdviceSubmissionRequest) WithContext(ctx context.Context) *GetPaymentAdviceSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get payment advice submission Request
func (o *GetPaymentAdviceSubmissionRequest) WithHTTPClient(client *http.Client) *GetPaymentAdviceSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPaymentAdviceSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param adviceId
	if err := r.SetPathParam("adviceId", o.AdviceID.String()); err != nil {
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
