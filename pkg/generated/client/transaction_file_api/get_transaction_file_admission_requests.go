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

// Client.GetTransactionFileAdmission creates a new GetTransactionFileAdmissionRequest object
// with the default values initialized.
func (c *Client) GetTransactionFileAdmission() *GetTransactionFileAdmissionRequest {
	var ()
	return &GetTransactionFileAdmissionRequest{

		TransactionFileAdmissionID: c.Defaults.GetStrfmtUUID("GetTransactionFileAdmission", "transaction_file_admission_id"),

		TransactionFileID: c.Defaults.GetStrfmtUUID("GetTransactionFileAdmission", "transaction_file_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetTransactionFileAdmissionRequest struct {

	/*TransactionFileAdmissionID      Transaction File Admission Id      */

	TransactionFileAdmissionID strfmt.UUID

	/*TransactionFileID      Transaction File Id      */

	TransactionFileID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetTransactionFileAdmissionRequest) FromJson(j string) (*GetTransactionFileAdmissionRequest, error) {

	return o, nil
}

func (o *GetTransactionFileAdmissionRequest) WithTransactionFileAdmissionID(transactionFileAdmissionID strfmt.UUID) *GetTransactionFileAdmissionRequest {

	o.TransactionFileAdmissionID = transactionFileAdmissionID

	return o
}

func (o *GetTransactionFileAdmissionRequest) WithTransactionFileID(transactionFileID strfmt.UUID) *GetTransactionFileAdmissionRequest {

	o.TransactionFileID = transactionFileID

	return o
}

// ////////////////
// WithContext adds the context to the get transaction file admission Request
func (o *GetTransactionFileAdmissionRequest) WithContext(ctx context.Context) *GetTransactionFileAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get transaction file admission Request
func (o *GetTransactionFileAdmissionRequest) WithHTTPClient(client *http.Client) *GetTransactionFileAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetTransactionFileAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param transaction_file_admission_id
	if err := r.SetPathParam("transaction_file_admission_id", o.TransactionFileAdmissionID.String()); err != nil {
		return err
	}

	// path param transaction_file_id
	if err := r.SetPathParam("transaction_file_id", o.TransactionFileID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
