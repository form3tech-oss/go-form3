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

// Client.GetDirectDebitAdmission creates a new GetDirectDebitAdmissionRequest object
// with the default values initialized.
func (c *Client) GetDirectDebitAdmission() *GetDirectDebitAdmissionRequest {
	var ()
	return &GetDirectDebitAdmissionRequest{

		AdmissionID: c.Defaults.GetStrfmtUUID("GetDirectDebitAdmission", "admissionId"),

		ID: c.Defaults.GetStrfmtUUID("GetDirectDebitAdmission", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetDirectDebitAdmissionRequest struct {

	/*AdmissionID      Direct Debit Admission Id      */

	AdmissionID strfmt.UUID

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetDirectDebitAdmissionRequest) FromJson(j string) (*GetDirectDebitAdmissionRequest, error) {

	return o, nil
}

func (o *GetDirectDebitAdmissionRequest) WithAdmissionID(admissionID strfmt.UUID) *GetDirectDebitAdmissionRequest {

	o.AdmissionID = admissionID

	return o
}

func (o *GetDirectDebitAdmissionRequest) WithID(id strfmt.UUID) *GetDirectDebitAdmissionRequest {

	o.ID = id

	return o
}

// ////////////////
// WithContext adds the context to the get direct debit admission Request
func (o *GetDirectDebitAdmissionRequest) WithContext(ctx context.Context) *GetDirectDebitAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get direct debit admission Request
func (o *GetDirectDebitAdmissionRequest) WithHTTPClient(client *http.Client) *GetDirectDebitAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetDirectDebitAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
