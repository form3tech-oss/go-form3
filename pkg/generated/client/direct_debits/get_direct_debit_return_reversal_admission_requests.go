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

// Client.GetDirectDebitReturnReversalAdmission creates a new GetDirectDebitReturnReversalAdmissionRequest object
// with the default values initialized.
func (c *Client) GetDirectDebitReturnReversalAdmission() *GetDirectDebitReturnReversalAdmissionRequest {
	var ()
	return &GetDirectDebitReturnReversalAdmissionRequest{

		AdmissionID: c.Defaults.GetStrfmtUUID("GetDirectDebitReturnReversalAdmission", "admissionId"),

		ID: c.Defaults.GetStrfmtUUID("GetDirectDebitReturnReversalAdmission", "id"),

		ReturnID: c.Defaults.GetStrfmtUUID("GetDirectDebitReturnReversalAdmission", "returnId"),

		ReversalID: c.Defaults.GetStrfmtUUID("GetDirectDebitReturnReversalAdmission", "reversalId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetDirectDebitReturnReversalAdmissionRequest struct {

	/*AdmissionID      Direct Debit Admission Id      */

	AdmissionID strfmt.UUID

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

func (o *GetDirectDebitReturnReversalAdmissionRequest) FromJson(j string) (*GetDirectDebitReturnReversalAdmissionRequest, error) {

	return o, nil
}

func (o *GetDirectDebitReturnReversalAdmissionRequest) WithAdmissionID(admissionID strfmt.UUID) *GetDirectDebitReturnReversalAdmissionRequest {

	o.AdmissionID = admissionID

	return o
}

func (o *GetDirectDebitReturnReversalAdmissionRequest) WithID(id strfmt.UUID) *GetDirectDebitReturnReversalAdmissionRequest {

	o.ID = id

	return o
}

func (o *GetDirectDebitReturnReversalAdmissionRequest) WithReturnID(returnID strfmt.UUID) *GetDirectDebitReturnReversalAdmissionRequest {

	o.ReturnID = returnID

	return o
}

func (o *GetDirectDebitReturnReversalAdmissionRequest) WithReversalID(reversalID strfmt.UUID) *GetDirectDebitReturnReversalAdmissionRequest {

	o.ReversalID = reversalID

	return o
}

// ////////////////
// WithContext adds the context to the get direct debit return reversal admission Request
func (o *GetDirectDebitReturnReversalAdmissionRequest) WithContext(ctx context.Context) *GetDirectDebitReturnReversalAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get direct debit return reversal admission Request
func (o *GetDirectDebitReturnReversalAdmissionRequest) WithHTTPClient(client *http.Client) *GetDirectDebitReturnReversalAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetDirectDebitReturnReversalAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param admissionId
	if err := r.SetPathParam("admissionId", o.AdmissionID.String()); err != nil {
		return err
	}

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
