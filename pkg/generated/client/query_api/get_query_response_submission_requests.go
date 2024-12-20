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

// Client.GetQueryResponseSubmission creates a new GetQueryResponseSubmissionRequest object
// with the default values initialized.
func (c *Client) GetQueryResponseSubmission() *GetQueryResponseSubmissionRequest {
	var ()
	return &GetQueryResponseSubmissionRequest{

		QueryID: c.Defaults.GetStrfmtUUID("GetQueryResponseSubmission", "query_id"),

		QueryResponseID: c.Defaults.GetStrfmtUUID("GetQueryResponseSubmission", "query_response_id"),

		QueryResponseSubmissionID: c.Defaults.GetStrfmtUUID("GetQueryResponseSubmission", "query_response_submission_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetQueryResponseSubmissionRequest struct {

	/*QueryID      Query ID      */

	QueryID strfmt.UUID

	/*QueryResponseID      Query Response ID      */

	QueryResponseID strfmt.UUID

	/*QueryResponseSubmissionID      Query Response Submission ID      */

	QueryResponseSubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetQueryResponseSubmissionRequest) FromJson(j string) (*GetQueryResponseSubmissionRequest, error) {

	return o, nil
}

func (o *GetQueryResponseSubmissionRequest) WithQueryID(queryID strfmt.UUID) *GetQueryResponseSubmissionRequest {

	o.QueryID = queryID

	return o
}

func (o *GetQueryResponseSubmissionRequest) WithQueryResponseID(queryResponseID strfmt.UUID) *GetQueryResponseSubmissionRequest {

	o.QueryResponseID = queryResponseID

	return o
}

func (o *GetQueryResponseSubmissionRequest) WithQueryResponseSubmissionID(queryResponseSubmissionID strfmt.UUID) *GetQueryResponseSubmissionRequest {

	o.QueryResponseSubmissionID = queryResponseSubmissionID

	return o
}

// ////////////////
// WithContext adds the context to the get query response submission Request
func (o *GetQueryResponseSubmissionRequest) WithContext(ctx context.Context) *GetQueryResponseSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get query response submission Request
func (o *GetQueryResponseSubmissionRequest) WithHTTPClient(client *http.Client) *GetQueryResponseSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetQueryResponseSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param query_id
	if err := r.SetPathParam("query_id", o.QueryID.String()); err != nil {
		return err
	}

	// path param query_response_id
	if err := r.SetPathParam("query_response_id", o.QueryResponseID.String()); err != nil {
		return err
	}

	// path param query_response_submission_id
	if err := r.SetPathParam("query_response_submission_id", o.QueryResponseSubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
