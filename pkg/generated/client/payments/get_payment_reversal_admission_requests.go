// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package payments

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetPaymentReversalAdmission creates a new GetPaymentReversalAdmissionRequest object
// with the default values initialized.
func (c *Client) GetPaymentReversalAdmission() *GetPaymentReversalAdmissionRequest {
	var ()
	return &GetPaymentReversalAdmissionRequest{

		AdmissionID: c.Defaults.GetStrfmtUUID("GetPaymentReversalAdmission", "admissionId"),

		ID: c.Defaults.GetStrfmtUUID("GetPaymentReversalAdmission", "id"),

		ReversalID: c.Defaults.GetStrfmtUUID("GetPaymentReversalAdmission", "reversalId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPaymentReversalAdmissionRequest struct {

	/*AdmissionID      Admission Id      */

	AdmissionID strfmt.UUID

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*ReversalID      Reversal Id      */

	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPaymentReversalAdmissionRequest) FromJson(j string) *GetPaymentReversalAdmissionRequest {

	return o
}

func (o *GetPaymentReversalAdmissionRequest) WithAdmissionID(admissionID strfmt.UUID) *GetPaymentReversalAdmissionRequest {

	o.AdmissionID = admissionID

	return o
}

func (o *GetPaymentReversalAdmissionRequest) WithID(id strfmt.UUID) *GetPaymentReversalAdmissionRequest {

	o.ID = id

	return o
}

func (o *GetPaymentReversalAdmissionRequest) WithReversalID(reversalID strfmt.UUID) *GetPaymentReversalAdmissionRequest {

	o.ReversalID = reversalID

	return o
}

//////////////////
// WithContext adds the context to the get payment reversal admission Request
func (o *GetPaymentReversalAdmissionRequest) WithContext(ctx context.Context) *GetPaymentReversalAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get payment reversal admission Request
func (o *GetPaymentReversalAdmissionRequest) WithHTTPClient(client *http.Client) *GetPaymentReversalAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPaymentReversalAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
