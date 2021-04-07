// Code generated by go-swagger; DO NOT EDIT.

package account_amendment

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetAccountAmendmentSubmission creates a new GetAccountAmendmentSubmissionRequest object
// with the default values initialized.
func (c *Client) GetAccountAmendmentSubmission() *GetAccountAmendmentSubmissionRequest {
	var ()
	return &GetAccountAmendmentSubmissionRequest{

		ID: c.Defaults.GetStrfmtUUID("GetAccountAmendmentSubmission", "id"),

		SubmissionID: c.Defaults.GetStrfmtUUID("GetAccountAmendmentSubmission", "submissionId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetAccountAmendmentSubmissionRequest struct {

	/*ID      Account Amendment ID      */

	ID strfmt.UUID

	/*SubmissionID      Account Amendment Submission ID      */

	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetAccountAmendmentSubmissionRequest) FromJson(j string) *GetAccountAmendmentSubmissionRequest {

	return o
}

func (o *GetAccountAmendmentSubmissionRequest) WithID(id strfmt.UUID) *GetAccountAmendmentSubmissionRequest {

	o.ID = id

	return o
}

func (o *GetAccountAmendmentSubmissionRequest) WithSubmissionID(submissionID strfmt.UUID) *GetAccountAmendmentSubmissionRequest {

	o.SubmissionID = submissionID

	return o
}

//////////////////
// WithContext adds the context to the get account amendment submission Request
func (o *GetAccountAmendmentSubmissionRequest) WithContext(ctx context.Context) *GetAccountAmendmentSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get account amendment submission Request
func (o *GetAccountAmendmentSubmissionRequest) WithHTTPClient(client *http.Client) *GetAccountAmendmentSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetAccountAmendmentSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
