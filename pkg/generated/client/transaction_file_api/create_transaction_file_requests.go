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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// Client.CreateTransactionFile creates a new CreateTransactionFileRequest object
// with the default values initialized.
func (c *Client) CreateTransactionFile() *CreateTransactionFileRequest {
	var ()
	return &CreateTransactionFileRequest{

		TransactionFileCreation: models.TransactionFileCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateTransactionFileRequest struct {

	/*TransactionFileCreationRequest*/

	*models.TransactionFileCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateTransactionFileRequest) FromJson(j string) (*CreateTransactionFileRequest, error) {

	var m models.TransactionFileCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.TransactionFileCreation = &m

	return o, nil
}

func (o *CreateTransactionFileRequest) WithTransactionFileCreationRequest(transactionFileCreationRequest models.TransactionFileCreation) *CreateTransactionFileRequest {

	o.TransactionFileCreation = &transactionFileCreationRequest

	return o
}

func (o *CreateTransactionFileRequest) WithoutTransactionFileCreationRequest() *CreateTransactionFileRequest {

	o.TransactionFileCreation = &models.TransactionFileCreation{}

	return o
}

//////////////////
// WithContext adds the context to the create transaction file Request
func (o *CreateTransactionFileRequest) WithContext(ctx context.Context) *CreateTransactionFileRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create transaction file Request
func (o *CreateTransactionFileRequest) WithHTTPClient(client *http.Client) *CreateTransactionFileRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateTransactionFileRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.TransactionFileCreation != nil {
		if err := r.SetBodyParam(o.TransactionFileCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}