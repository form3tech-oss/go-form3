// Code generated by go-swagger; DO NOT EDIT.

package query_api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.GetQueryAdmission creates a new GetQueryAdmissionRequest object
// with the default values initialized.
func (c *Client) GetQueryAdmission() *GetQueryAdmissionRequest {
	var ()
	return &GetQueryAdmissionRequest{

		QueryAdmissionID: c.Defaults.GetStrfmtUUID("GetQueryAdmission", "query_admission_id"),

		QueryID: c.Defaults.GetStrfmtUUID("GetQueryAdmission", "query_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetQueryAdmissionRequest struct {

	/*QueryAdmissionID      Query Admission ID      */

	QueryAdmissionID strfmt.UUID

	/*QueryID      Query ID      */

	QueryID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetQueryAdmissionRequest) FromJson(j string) (*GetQueryAdmissionRequest, error) {

	return o, nil
}

func (o *GetQueryAdmissionRequest) WithQueryAdmissionID(queryAdmissionID strfmt.UUID) *GetQueryAdmissionRequest {

	o.QueryAdmissionID = queryAdmissionID

	return o
}

func (o *GetQueryAdmissionRequest) WithQueryID(queryID strfmt.UUID) *GetQueryAdmissionRequest {

	o.QueryID = queryID

	return o
}

// ////////////////
// WithContext adds the context to the get query admission Request
func (o *GetQueryAdmissionRequest) WithContext(ctx context.Context) *GetQueryAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get query admission Request
func (o *GetQueryAdmissionRequest) WithHTTPClient(client *http.Client) *GetQueryAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetQueryAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param query_admission_id
	if err := r.SetPathParam("query_admission_id", o.QueryAdmissionID.String()); err != nil {
		return err
	}

	// path param query_id
	if err := r.SetPathParam("query_id", o.QueryID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
