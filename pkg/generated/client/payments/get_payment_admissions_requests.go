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

// Client.GetPaymentAdmissions creates a new GetPaymentAdmissionsRequest object
// with the default values initialized.
func (c *Client) GetPaymentAdmissions() *GetPaymentAdmissionsRequest {
	var ()
	return &GetPaymentAdmissionsRequest{

		AdmissionID: c.Defaults.GetStrfmtUUID("GetPaymentAdmissions", "admissionId"),

		ID: c.Defaults.GetStrfmtUUID("GetPaymentAdmissions", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPaymentAdmissionsRequest struct {

	/*AdmissionID      Admission Id      */

	AdmissionID strfmt.UUID

	/*ID      Payment Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPaymentAdmissionsRequest) FromJson(j string) *GetPaymentAdmissionsRequest {

	return o
}

func (o *GetPaymentAdmissionsRequest) WithAdmissionID(admissionID strfmt.UUID) *GetPaymentAdmissionsRequest {

	o.AdmissionID = admissionID

	return o
}

func (o *GetPaymentAdmissionsRequest) WithID(id strfmt.UUID) *GetPaymentAdmissionsRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the get payment admissions Request
func (o *GetPaymentAdmissionsRequest) WithContext(ctx context.Context) *GetPaymentAdmissionsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get payment admissions Request
func (o *GetPaymentAdmissionsRequest) WithHTTPClient(client *http.Client) *GetPaymentAdmissionsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPaymentAdmissionsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
