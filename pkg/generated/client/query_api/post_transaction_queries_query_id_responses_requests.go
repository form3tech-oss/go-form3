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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v5/pkg/generated/models"
)

// Client.PostTransactionQueriesQueryIDResponses creates a new PostTransactionQueriesQueryIDResponsesRequest object
// with the default values initialized.
func (c *Client) PostTransactionQueriesQueryIDResponses() *PostTransactionQueriesQueryIDResponsesRequest {
	var ()
	return &PostTransactionQueriesQueryIDResponsesRequest{

		QueryResponseCreation: models.QueryResponseCreationWithDefaults(c.Defaults),

		QueryID: c.Defaults.GetStrfmtUUID("PostTransactionQueriesQueryIDResponses", "query_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type PostTransactionQueriesQueryIDResponsesRequest struct {

	/*CreationRequest*/

	*models.QueryResponseCreation

	/*QueryID      Query ID      */

	QueryID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *PostTransactionQueriesQueryIDResponsesRequest) FromJson(j string) (*PostTransactionQueriesQueryIDResponsesRequest, error) {

	var m models.QueryResponseCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.QueryResponseCreation = &m

	return o, nil
}

func (o *PostTransactionQueriesQueryIDResponsesRequest) WithCreationRequest(creationRequest models.QueryResponseCreation) *PostTransactionQueriesQueryIDResponsesRequest {

	o.QueryResponseCreation = &creationRequest

	return o
}

func (o *PostTransactionQueriesQueryIDResponsesRequest) WithoutCreationRequest() *PostTransactionQueriesQueryIDResponsesRequest {

	o.QueryResponseCreation = &models.QueryResponseCreation{}

	return o
}

func (o *PostTransactionQueriesQueryIDResponsesRequest) WithQueryID(queryID strfmt.UUID) *PostTransactionQueriesQueryIDResponsesRequest {

	o.QueryID = queryID

	return o
}

//////////////////
// WithContext adds the context to the post transaction queries query ID responses Request
func (o *PostTransactionQueriesQueryIDResponsesRequest) WithContext(ctx context.Context) *PostTransactionQueriesQueryIDResponsesRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the post transaction queries query ID responses Request
func (o *PostTransactionQueriesQueryIDResponsesRequest) WithHTTPClient(client *http.Client) *PostTransactionQueriesQueryIDResponsesRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *PostTransactionQueriesQueryIDResponsesRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.QueryResponseCreation != nil {
		if err := r.SetBodyParam(o.QueryResponseCreation); err != nil {
			return err
		}
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
