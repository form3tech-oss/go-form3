// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// Client.PostTransactionDirectdebitsIDReversalsReversalIDSubmissions creates a new PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest object
// with the default values initialized.
func (c *Client) PostTransactionDirectdebitsIDReversalsReversalIDSubmissions() *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest {
	var ()
	return &PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest{

		DirectDebitReversalSubmissionCreation: models.DirectDebitReversalSubmissionCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("PostTransactionDirectdebitsIDReversalsReversalIDSubmissions", "id"),

		ReversalID: c.Defaults.GetStrfmtUUID("PostTransactionDirectdebitsIDReversalsReversalIDSubmissions", "reversalId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest struct {

	/*ReversalSubmissionCreationRequest*/

	*models.DirectDebitReversalSubmissionCreation

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	/*ReversalID      Reversal Id      */

	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest) FromJson(j string) *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest {

	var m models.DirectDebitReversalSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.DirectDebitReversalSubmissionCreation = &m

	return o
}

func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest) WithReversalSubmissionCreationRequest(reversalSubmissionCreationRequest models.DirectDebitReversalSubmissionCreation) *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest {

	o.DirectDebitReversalSubmissionCreation = &reversalSubmissionCreationRequest

	return o
}

func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest) WithoutReversalSubmissionCreationRequest() *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest {

	o.DirectDebitReversalSubmissionCreation = &models.DirectDebitReversalSubmissionCreation{}

	return o
}

func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest) WithID(id strfmt.UUID) *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest {

	o.ID = id

	return o
}

func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest) WithReversalID(reversalID strfmt.UUID) *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest {

	o.ReversalID = reversalID

	return o
}

//////////////////
// WithContext adds the context to the post transaction directdebits ID reversals reversal ID submissions Request
func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest) WithContext(ctx context.Context) *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the post transaction directdebits ID reversals reversal ID submissions Request
func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest) WithHTTPClient(client *http.Client) *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.DirectDebitReversalSubmissionCreation != nil {
		if err := r.SetBodyParam(o.DirectDebitReversalSubmissionCreation); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
