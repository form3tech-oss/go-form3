// Code generated by go-swagger; DO NOT EDIT.

package transaction_file_api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// Client.CreateTransactionFileAdmission creates a new CreateTransactionFileAdmissionRequest object
// with the default values initialized.
func (c *Client) CreateTransactionFileAdmission() *CreateTransactionFileAdmissionRequest {
	var ()
	return &CreateTransactionFileAdmissionRequest{

		TransactionFileAdmissionCreation: models.TransactionFileAdmissionCreationWithDefaults(c.Defaults),

		TransactionFileID: c.Defaults.GetStrfmtUUID("CreateTransactionFileAdmission", "transaction_file_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateTransactionFileAdmissionRequest struct {

	/*TransactionFileAdmissionCreationRequest*/

	*models.TransactionFileAdmissionCreation

	/*TransactionFileID      Transaction File Id      */

	TransactionFileID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateTransactionFileAdmissionRequest) FromJson(j string) (*CreateTransactionFileAdmissionRequest, error) {

	var m models.TransactionFileAdmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.TransactionFileAdmissionCreation = &m

	return o, nil
}

func (o *CreateTransactionFileAdmissionRequest) WithTransactionFileAdmissionCreationRequest(transactionFileAdmissionCreationRequest models.TransactionFileAdmissionCreation) *CreateTransactionFileAdmissionRequest {

	o.TransactionFileAdmissionCreation = &transactionFileAdmissionCreationRequest

	return o
}

func (o *CreateTransactionFileAdmissionRequest) WithoutTransactionFileAdmissionCreationRequest() *CreateTransactionFileAdmissionRequest {

	o.TransactionFileAdmissionCreation = &models.TransactionFileAdmissionCreation{}

	return o
}

func (o *CreateTransactionFileAdmissionRequest) WithTransactionFileID(transactionFileID strfmt.UUID) *CreateTransactionFileAdmissionRequest {

	o.TransactionFileID = transactionFileID

	return o
}

// ////////////////
// WithContext adds the context to the create transaction file admission Request
func (o *CreateTransactionFileAdmissionRequest) WithContext(ctx context.Context) *CreateTransactionFileAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create transaction file admission Request
func (o *CreateTransactionFileAdmissionRequest) WithHTTPClient(client *http.Client) *CreateTransactionFileAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateTransactionFileAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.TransactionFileAdmissionCreation != nil {
		if err := r.SetBodyParam(o.TransactionFileAdmissionCreation); err != nil {
			return err
		}
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
