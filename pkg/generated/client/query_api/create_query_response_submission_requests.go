// Code generated by go-swagger; DO NOT EDIT.

package query_api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// Client.CreateQueryResponseSubmission creates a new CreateQueryResponseSubmissionRequest object
// with the default values initialized.
func (c *Client) CreateQueryResponseSubmission() *CreateQueryResponseSubmissionRequest {
	var ()
	return &CreateQueryResponseSubmissionRequest{

		QueryResponseSubmissionCreation: models.QueryResponseSubmissionCreationWithDefaults(c.Defaults),

		QueryID: c.Defaults.GetStrfmtUUID("CreateQueryResponseSubmission", "query_id"),

		QueryResponseID: c.Defaults.GetStrfmtUUID("CreateQueryResponseSubmission", "query_response_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateQueryResponseSubmissionRequest struct {

	/*CreationRequest*/

	*models.QueryResponseSubmissionCreation

	/*QueryID      Query ID      */

	QueryID strfmt.UUID

	/*QueryResponseID      Query Response ID      */

	QueryResponseID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateQueryResponseSubmissionRequest) FromJson(j string) (*CreateQueryResponseSubmissionRequest, error) {

	var m models.QueryResponseSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.QueryResponseSubmissionCreation = &m

	return o, nil
}

func (o *CreateQueryResponseSubmissionRequest) WithCreationRequest(creationRequest models.QueryResponseSubmissionCreation) *CreateQueryResponseSubmissionRequest {

	o.QueryResponseSubmissionCreation = &creationRequest

	return o
}

func (o *CreateQueryResponseSubmissionRequest) WithoutCreationRequest() *CreateQueryResponseSubmissionRequest {

	o.QueryResponseSubmissionCreation = &models.QueryResponseSubmissionCreation{}

	return o
}

func (o *CreateQueryResponseSubmissionRequest) WithQueryID(queryID strfmt.UUID) *CreateQueryResponseSubmissionRequest {

	o.QueryID = queryID

	return o
}

func (o *CreateQueryResponseSubmissionRequest) WithQueryResponseID(queryResponseID strfmt.UUID) *CreateQueryResponseSubmissionRequest {

	o.QueryResponseID = queryResponseID

	return o
}

// ////////////////
// WithContext adds the context to the create query response submission Request
func (o *CreateQueryResponseSubmissionRequest) WithContext(ctx context.Context) *CreateQueryResponseSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create query response submission Request
func (o *CreateQueryResponseSubmissionRequest) WithHTTPClient(client *http.Client) *CreateQueryResponseSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateQueryResponseSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.QueryResponseSubmissionCreation != nil {
		if err := r.SetBodyParam(o.QueryResponseSubmissionCreation); err != nil {
			return err
		}
	}

	// path param query_id
	if err := r.SetPathParam("query_id", o.QueryID.String()); err != nil {
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