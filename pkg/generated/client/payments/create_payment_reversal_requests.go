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

// Client.CreatePaymentReversal creates a new CreatePaymentReversalRequest object
// with the default values initialized.
func (c *Client) CreatePaymentReversal() *CreatePaymentReversalRequest {
	var ()
	return &CreatePaymentReversalRequest{

		ReversalCreation: models.ReversalCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreatePaymentReversal", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreatePaymentReversalRequest struct {

	/*ReversalCreationRequest*/

	*models.ReversalCreation

	/*ID      Payment Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreatePaymentReversalRequest) FromJson(j string) *CreatePaymentReversalRequest {

	var m models.ReversalCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.ReversalCreation = &m

	return o
}

func (o *CreatePaymentReversalRequest) WithReversalCreationRequest(reversalCreationRequest models.ReversalCreation) *CreatePaymentReversalRequest {

	o.ReversalCreation = &reversalCreationRequest

	return o
}

func (o *CreatePaymentReversalRequest) WithoutReversalCreationRequest() *CreatePaymentReversalRequest {

	o.ReversalCreation = &models.ReversalCreation{}

	return o
}

func (o *CreatePaymentReversalRequest) WithID(id strfmt.UUID) *CreatePaymentReversalRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create payment reversal Request
func (o *CreatePaymentReversalRequest) WithContext(ctx context.Context) *CreatePaymentReversalRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create payment reversal Request
func (o *CreatePaymentReversalRequest) WithHTTPClient(client *http.Client) *CreatePaymentReversalRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreatePaymentReversalRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.ReversalCreation != nil {
		if err := r.SetBodyParam(o.ReversalCreation); err != nil {
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
