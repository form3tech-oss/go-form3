// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.GetDirectDebitReturnReversal creates a new GetDirectDebitReturnReversalRequest object
// with the default values initialized.
func (c *Client) GetDirectDebitReturnReversal() *GetDirectDebitReturnReversalRequest {
	var ()
	return &GetDirectDebitReturnReversalRequest{

		ID: c.Defaults.GetStrfmtUUID("GetDirectDebitReturnReversal", "id"),

		ReturnID: c.Defaults.GetStrfmtUUID("GetDirectDebitReturnReversal", "returnId"),

		ReversalID: c.Defaults.GetStrfmtUUID("GetDirectDebitReturnReversal", "reversalId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetDirectDebitReturnReversalRequest struct {

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	/*ReturnID      Return Id      */

	ReturnID strfmt.UUID

	/*ReversalID      Reversal Id      */

	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetDirectDebitReturnReversalRequest) FromJson(j string) (*GetDirectDebitReturnReversalRequest, error) {

	return o, nil
}

func (o *GetDirectDebitReturnReversalRequest) WithID(id strfmt.UUID) *GetDirectDebitReturnReversalRequest {

	o.ID = id

	return o
}

func (o *GetDirectDebitReturnReversalRequest) WithReturnID(returnID strfmt.UUID) *GetDirectDebitReturnReversalRequest {

	o.ReturnID = returnID

	return o
}

func (o *GetDirectDebitReturnReversalRequest) WithReversalID(reversalID strfmt.UUID) *GetDirectDebitReturnReversalRequest {

	o.ReversalID = reversalID

	return o
}

// ////////////////
// WithContext adds the context to the get direct debit return reversal Request
func (o *GetDirectDebitReturnReversalRequest) WithContext(ctx context.Context) *GetDirectDebitReturnReversalRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get direct debit return reversal Request
func (o *GetDirectDebitReturnReversalRequest) WithHTTPClient(client *http.Client) *GetDirectDebitReturnReversalRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetDirectDebitReturnReversalRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
