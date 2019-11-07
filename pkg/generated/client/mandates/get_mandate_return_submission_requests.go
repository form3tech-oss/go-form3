// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package mandates

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetMandateReturnSubmission creates a new GetMandateReturnSubmissionRequest object
// with the default values initialized.
func (c *Client) GetMandateReturnSubmission() *GetMandateReturnSubmissionRequest {
	var ()
	return &GetMandateReturnSubmissionRequest{

		ID: c.Defaults.GetStrfmtUUID("GetMandateReturnSubmission", "id"),

		ReturnID: c.Defaults.GetStrfmtUUID("GetMandateReturnSubmission", "returnId"),

		SubmissionID: c.Defaults.GetStrfmtUUID("GetMandateReturnSubmission", "submissionId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetMandateReturnSubmissionRequest struct {

	/*ID      Mandate Id      */

	ID strfmt.UUID

	/*ReturnID      Return Id      */

	ReturnID strfmt.UUID

	/*SubmissionID      Submission Id      */

	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetMandateReturnSubmissionRequest) FromJson(j string) *GetMandateReturnSubmissionRequest {

	return o
}

func (o *GetMandateReturnSubmissionRequest) WithID(id strfmt.UUID) *GetMandateReturnSubmissionRequest {

	o.ID = id

	return o
}

func (o *GetMandateReturnSubmissionRequest) WithReturnID(returnID strfmt.UUID) *GetMandateReturnSubmissionRequest {

	o.ReturnID = returnID

	return o
}

func (o *GetMandateReturnSubmissionRequest) WithSubmissionID(submissionID strfmt.UUID) *GetMandateReturnSubmissionRequest {

	o.SubmissionID = submissionID

	return o
}

//////////////////
// WithContext adds the context to the get mandate return submission Request
func (o *GetMandateReturnSubmissionRequest) WithContext(ctx context.Context) *GetMandateReturnSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get mandate return submission Request
func (o *GetMandateReturnSubmissionRequest) WithHTTPClient(client *http.Client) *GetMandateReturnSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetMandateReturnSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param returnId
	if err := r.SetPathParam("returnId", o.ReturnID.String()); err != nil {
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