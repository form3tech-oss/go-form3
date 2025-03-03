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

// Client.GetPaymentSubmissionTask creates a new GetPaymentSubmissionTaskRequest object
// with the default values initialized.
func (c *Client) GetPaymentSubmissionTask() *GetPaymentSubmissionTaskRequest {
	var ()
	return &GetPaymentSubmissionTaskRequest{

		ID: c.Defaults.GetStrfmtUUID("GetPaymentSubmissionTask", "id"),

		SubmissionID: c.Defaults.GetStrfmtUUID("GetPaymentSubmissionTask", "submissionId"),

		TaskID: c.Defaults.GetStrfmtUUID("GetPaymentSubmissionTask", "taskId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPaymentSubmissionTaskRequest struct {

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*SubmissionID      Submission Id      */

	SubmissionID strfmt.UUID

	/*TaskID      Payment Submission Task Id      */

	TaskID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPaymentSubmissionTaskRequest) FromJson(j string) (*GetPaymentSubmissionTaskRequest, error) {

	return o, nil
}

func (o *GetPaymentSubmissionTaskRequest) WithID(id strfmt.UUID) *GetPaymentSubmissionTaskRequest {

	o.ID = id

	return o
}

func (o *GetPaymentSubmissionTaskRequest) WithSubmissionID(submissionID strfmt.UUID) *GetPaymentSubmissionTaskRequest {

	o.SubmissionID = submissionID

	return o
}

func (o *GetPaymentSubmissionTaskRequest) WithTaskID(taskID strfmt.UUID) *GetPaymentSubmissionTaskRequest {

	o.TaskID = taskID

	return o
}

// ////////////////
// WithContext adds the context to the get payment submission task Request
func (o *GetPaymentSubmissionTaskRequest) WithContext(ctx context.Context) *GetPaymentSubmissionTaskRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get payment submission task Request
func (o *GetPaymentSubmissionTaskRequest) WithHTTPClient(client *http.Client) *GetPaymentSubmissionTaskRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPaymentSubmissionTaskRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param taskId
	if err := r.SetPathParam("taskId", o.TaskID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
