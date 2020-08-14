// Code generated by go-swagger; DO NOT EDIT.

package payments

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

// Client.CreatePaymentReturn creates a new CreatePaymentReturnRequest object
// with the default values initialized.
func (c *Client) CreatePaymentReturn() *CreatePaymentReturnRequest {
	var ()
	return &CreatePaymentReturnRequest{

		ReturnCreation: models.ReturnCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreatePaymentReturn", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreatePaymentReturnRequest struct {

	/*ReturnCreationRequest*/

	*models.ReturnCreation

	/*ID      Payment Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreatePaymentReturnRequest) FromJson(j string) *CreatePaymentReturnRequest {

	var m models.ReturnCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.ReturnCreation = &m

	return o
}

func (o *CreatePaymentReturnRequest) WithReturnCreationRequest(returnCreationRequest models.ReturnCreation) *CreatePaymentReturnRequest {

	o.ReturnCreation = &returnCreationRequest

	return o
}

func (o *CreatePaymentReturnRequest) WithoutReturnCreationRequest() *CreatePaymentReturnRequest {

	o.ReturnCreation = &models.ReturnCreation{}

	return o
}

func (o *CreatePaymentReturnRequest) WithID(id strfmt.UUID) *CreatePaymentReturnRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create payment return Request
func (o *CreatePaymentReturnRequest) WithContext(ctx context.Context) *CreatePaymentReturnRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create payment return Request
func (o *CreatePaymentReturnRequest) WithHTTPClient(client *http.Client) *CreatePaymentReturnRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreatePaymentReturnRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.ReturnCreation != nil {
		if err := r.SetBodyParam(o.ReturnCreation); err != nil {
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
