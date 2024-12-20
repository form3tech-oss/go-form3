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

// Client.GetDirectdebitsIDDecisionsDecisionID creates a new GetDirectdebitsIDDecisionsDecisionIDRequest object
// with the default values initialized.
func (c *Client) GetDirectdebitsIDDecisionsDecisionID() *GetDirectdebitsIDDecisionsDecisionIDRequest {
	var ()
	return &GetDirectdebitsIDDecisionsDecisionIDRequest{

		DecisionID: c.Defaults.GetStrfmtUUID("GetDirectdebitsIDDecisionsDecisionID", "decisionId"),

		ID: c.Defaults.GetStrfmtUUID("GetDirectdebitsIDDecisionsDecisionID", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetDirectdebitsIDDecisionsDecisionIDRequest struct {

	/*DecisionID      Direct Debit decision id      */

	DecisionID strfmt.UUID

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetDirectdebitsIDDecisionsDecisionIDRequest) FromJson(j string) (*GetDirectdebitsIDDecisionsDecisionIDRequest, error) {

	return o, nil
}

func (o *GetDirectdebitsIDDecisionsDecisionIDRequest) WithDecisionID(decisionID strfmt.UUID) *GetDirectdebitsIDDecisionsDecisionIDRequest {

	o.DecisionID = decisionID

	return o
}

func (o *GetDirectdebitsIDDecisionsDecisionIDRequest) WithID(id strfmt.UUID) *GetDirectdebitsIDDecisionsDecisionIDRequest {

	o.ID = id

	return o
}

// ////////////////
// WithContext adds the context to the get directdebits ID decisions decision ID Request
func (o *GetDirectdebitsIDDecisionsDecisionIDRequest) WithContext(ctx context.Context) *GetDirectdebitsIDDecisionsDecisionIDRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get directdebits ID decisions decision ID Request
func (o *GetDirectdebitsIDDecisionsDecisionIDRequest) WithHTTPClient(client *http.Client) *GetDirectdebitsIDDecisionsDecisionIDRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetDirectdebitsIDDecisionsDecisionIDRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param decisionId
	if err := r.SetPathParam("decisionId", o.DecisionID.String()); err != nil {
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
