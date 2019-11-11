// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// Client.CreateDirectDebitReturn creates a new CreateDirectDebitReturnRequest object
// with the default values initialized.
func (c *Client) CreateDirectDebitReturn() *CreateDirectDebitReturnRequest {
	var ()
	return &CreateDirectDebitReturnRequest{

		DirectDebitReturnCreation: models.DirectDebitReturnCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreateDirectDebitReturn", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateDirectDebitReturnRequest struct {

	/*ReturnCreationRequest*/

	*models.DirectDebitReturnCreation

	/*ID      DirectDebit Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateDirectDebitReturnRequest) FromJson(j string) *CreateDirectDebitReturnRequest {

	var m models.DirectDebitReturnCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.DirectDebitReturnCreation = &m

	return o
}

func (o *CreateDirectDebitReturnRequest) WithReturnCreationRequest(returnCreationRequest models.DirectDebitReturnCreation) *CreateDirectDebitReturnRequest {

	o.DirectDebitReturnCreation = &returnCreationRequest

	return o
}

func (o *CreateDirectDebitReturnRequest) WithoutReturnCreationRequest() *CreateDirectDebitReturnRequest {

	o.DirectDebitReturnCreation = &models.DirectDebitReturnCreation{}

	return o
}

func (o *CreateDirectDebitReturnRequest) WithID(id strfmt.UUID) *CreateDirectDebitReturnRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create direct debit return Request
func (o *CreateDirectDebitReturnRequest) WithContext(ctx context.Context) *CreateDirectDebitReturnRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create direct debit return Request
func (o *CreateDirectDebitReturnRequest) WithHTTPClient(client *http.Client) *CreateDirectDebitReturnRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateDirectDebitReturnRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.DirectDebitReturnCreation != nil {
		if err := r.SetBodyParam(o.DirectDebitReturnCreation); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
