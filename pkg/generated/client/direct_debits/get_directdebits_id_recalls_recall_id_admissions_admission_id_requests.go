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

// Client.GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID creates a new GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest object
// with the default values initialized.
func (c *Client) GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID() *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {
	var ()
	return &GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest{

		AdmissionID: c.Defaults.GetStrfmtUUID("GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID", "admissionId"),

		ID: c.Defaults.GetStrfmtUUID("GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID", "id"),

		RecallID: c.Defaults.GetStrfmtUUID("GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID", "recallId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest struct {

	/*AdmissionID      Direct Debit Admission Id      */

	AdmissionID strfmt.UUID

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	/*RecallID      Recall Id      */

	RecallID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) FromJson(j string) (*GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest, error) {

	return o, nil
}

func (o *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WithAdmissionID(admissionID strfmt.UUID) *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {

	o.AdmissionID = admissionID

	return o
}

func (o *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WithID(id strfmt.UUID) *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {

	o.ID = id

	return o
}

func (o *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WithRecallID(recallID strfmt.UUID) *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {

	o.RecallID = recallID

	return o
}

// ////////////////
// WithContext adds the context to the get directdebits ID recalls recall ID admissions admission ID Request
func (o *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WithContext(ctx context.Context) *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get directdebits ID recalls recall ID admissions admission ID Request
func (o *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WithHTTPClient(client *http.Client) *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param recallId
	if err := r.SetPathParam("recallId", o.RecallID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
