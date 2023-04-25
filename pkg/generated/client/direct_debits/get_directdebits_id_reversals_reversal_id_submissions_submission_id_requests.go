// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionID creates a new GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest object
// with the default values initialized.
func (c *Client) GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionID() *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest {
	var ()
	return &GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest{

		ID: c.Defaults.GetStrfmtUUID("GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionID", "id"),

		ReversalID: c.Defaults.GetStrfmtUUID("GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionID", "reversalId"),

		SubmissionID: c.Defaults.GetStrfmtUUID("GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionID", "submissionId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest struct {

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	/*ReversalID      Reversal Id      */

	ReversalID strfmt.UUID

	/*SubmissionID      Direct Debit decision submission id      */

	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest) FromJson(j string) (*GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest, error) {

	return o, nil
}

func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest) WithID(id strfmt.UUID) *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest {

	o.ID = id

	return o
}

func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest) WithReversalID(reversalID strfmt.UUID) *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest {

	o.ReversalID = reversalID

	return o
}

func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest) WithSubmissionID(submissionID strfmt.UUID) *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest {

	o.SubmissionID = submissionID

	return o
}

//////////////////
// WithContext adds the context to the get directdebits ID reversals reversal ID submissions submission ID Request
func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest) WithContext(ctx context.Context) *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get directdebits ID reversals reversal ID submissions submission ID Request
func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest) WithHTTPClient(client *http.Client) *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	// path param submissionId
	if err := r.SetPathParam("submissionId", o.SubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
