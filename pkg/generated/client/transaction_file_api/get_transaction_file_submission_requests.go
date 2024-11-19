// Code generated by go-swagger; DO NOT EDIT.

package transaction_file_api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.GetTransactionFileSubmission creates a new GetTransactionFileSubmissionRequest object
// with the default values initialized.
func (c *Client) GetTransactionFileSubmission() *GetTransactionFileSubmissionRequest {
	var ()
	return &GetTransactionFileSubmissionRequest{

		TransactionFileID: c.Defaults.GetStrfmtUUID("GetTransactionFileSubmission", "transaction_file_id"),

		TransactionFileSubmissionID: c.Defaults.GetStrfmtUUID("GetTransactionFileSubmission", "transaction_file_submission_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetTransactionFileSubmissionRequest struct {

	/*TransactionFileID      Transaction File Id      */

	TransactionFileID strfmt.UUID

	/*TransactionFileSubmissionID      Transaction File Submission Id      */

	TransactionFileSubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetTransactionFileSubmissionRequest) FromJson(j string) (*GetTransactionFileSubmissionRequest, error) {

	return o, nil
}

func (o *GetTransactionFileSubmissionRequest) WithTransactionFileID(transactionFileID strfmt.UUID) *GetTransactionFileSubmissionRequest {

	o.TransactionFileID = transactionFileID

	return o
}

func (o *GetTransactionFileSubmissionRequest) WithTransactionFileSubmissionID(transactionFileSubmissionID strfmt.UUID) *GetTransactionFileSubmissionRequest {

	o.TransactionFileSubmissionID = transactionFileSubmissionID

	return o
}

// ////////////////
// WithContext adds the context to the get transaction file submission Request
func (o *GetTransactionFileSubmissionRequest) WithContext(ctx context.Context) *GetTransactionFileSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get transaction file submission Request
func (o *GetTransactionFileSubmissionRequest) WithHTTPClient(client *http.Client) *GetTransactionFileSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetTransactionFileSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param transaction_file_id
	if err := r.SetPathParam("transaction_file_id", o.TransactionFileID.String()); err != nil {
		return err
	}

	// path param transaction_file_submission_id
	if err := r.SetPathParam("transaction_file_submission_id", o.TransactionFileSubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
