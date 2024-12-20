// Code generated by go-swagger; DO NOT EDIT.

package scheme_file_api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.GetSchemeFileAdmission creates a new GetSchemeFileAdmissionRequest object
// with the default values initialized.
func (c *Client) GetSchemeFileAdmission() *GetSchemeFileAdmissionRequest {
	var ()
	return &GetSchemeFileAdmissionRequest{

		SchemeFileAdmissionID: c.Defaults.GetStrfmtUUID("GetSchemeFileAdmission", "scheme_file_admission_id"),

		SchemeFileID: c.Defaults.GetStrfmtUUID("GetSchemeFileAdmission", "scheme_file_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetSchemeFileAdmissionRequest struct {

	/*SchemeFileAdmissionID      Scheme File Admission Id      */

	SchemeFileAdmissionID strfmt.UUID

	/*SchemeFileID      Scheme File Id      */

	SchemeFileID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetSchemeFileAdmissionRequest) FromJson(j string) (*GetSchemeFileAdmissionRequest, error) {

	return o, nil
}

func (o *GetSchemeFileAdmissionRequest) WithSchemeFileAdmissionID(schemeFileAdmissionID strfmt.UUID) *GetSchemeFileAdmissionRequest {

	o.SchemeFileAdmissionID = schemeFileAdmissionID

	return o
}

func (o *GetSchemeFileAdmissionRequest) WithSchemeFileID(schemeFileID strfmt.UUID) *GetSchemeFileAdmissionRequest {

	o.SchemeFileID = schemeFileID

	return o
}

// ////////////////
// WithContext adds the context to the get scheme file admission Request
func (o *GetSchemeFileAdmissionRequest) WithContext(ctx context.Context) *GetSchemeFileAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get scheme file admission Request
func (o *GetSchemeFileAdmissionRequest) WithHTTPClient(client *http.Client) *GetSchemeFileAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetSchemeFileAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param scheme_file_admission_id
	if err := r.SetPathParam("scheme_file_admission_id", o.SchemeFileAdmissionID.String()); err != nil {
		return err
	}

	// path param scheme_file_id
	if err := r.SetPathParam("scheme_file_id", o.SchemeFileID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
