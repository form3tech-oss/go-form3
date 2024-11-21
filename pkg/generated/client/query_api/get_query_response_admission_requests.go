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

// Client.GetQueryResponseAdmission creates a new GetQueryResponseAdmissionRequest object
// with the default values initialized.
func (c *Client) GetQueryResponseAdmission() *GetQueryResponseAdmissionRequest {
	var ()
	return &GetQueryResponseAdmissionRequest{

		QueryID: c.Defaults.GetStrfmtUUID("GetQueryResponseAdmission", "query_id"),

		QueryResponseAdmissionID: c.Defaults.GetStrfmtUUID("GetQueryResponseAdmission", "query_response_admission_id"),

		QueryResponseID: c.Defaults.GetStrfmtUUID("GetQueryResponseAdmission", "query_response_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetQueryResponseAdmissionRequest struct {

	/*QueryID      Query ID      */

	QueryID strfmt.UUID

	/*QueryResponseAdmissionID      Query Response Admission ID      */

	QueryResponseAdmissionID strfmt.UUID

	/*QueryResponseID      Query Response ID      */

	QueryResponseID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetQueryResponseAdmissionRequest) FromJson(j string) (*GetQueryResponseAdmissionRequest, error) {

	return o, nil
}

func (o *GetQueryResponseAdmissionRequest) WithQueryID(queryID strfmt.UUID) *GetQueryResponseAdmissionRequest {

	o.QueryID = queryID

	return o
}

func (o *GetQueryResponseAdmissionRequest) WithQueryResponseAdmissionID(queryResponseAdmissionID strfmt.UUID) *GetQueryResponseAdmissionRequest {

	o.QueryResponseAdmissionID = queryResponseAdmissionID

	return o
}

func (o *GetQueryResponseAdmissionRequest) WithQueryResponseID(queryResponseID strfmt.UUID) *GetQueryResponseAdmissionRequest {

	o.QueryResponseID = queryResponseID

	return o
}

// ////////////////
// WithContext adds the context to the get query response admission Request
func (o *GetQueryResponseAdmissionRequest) WithContext(ctx context.Context) *GetQueryResponseAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get query response admission Request
func (o *GetQueryResponseAdmissionRequest) WithHTTPClient(client *http.Client) *GetQueryResponseAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetQueryResponseAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param query_id
	if err := r.SetPathParam("query_id", o.QueryID.String()); err != nil {
		return err
	}

	// path param query_response_admission_id
	if err := r.SetPathParam("query_response_admission_id", o.QueryResponseAdmissionID.String()); err != nil {
		return err
	}

	// path param query_response_id
	if err := r.SetPathParam("query_response_id", o.QueryResponseID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}