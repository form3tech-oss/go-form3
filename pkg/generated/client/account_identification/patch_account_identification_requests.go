// Code generated by go-swagger; DO NOT EDIT.

package account_identification

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

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// Client.PatchAccountIdentification creates a new PatchAccountIdentificationRequest object
// with the default values initialized.
func (c *Client) PatchAccountIdentification() *PatchAccountIdentificationRequest {
	var ()
	return &PatchAccountIdentificationRequest{

		AccountIdentificationRequest: models.AccountIdentificationRequestWithDefaults(c.Defaults),

		AccountID: c.Defaults.GetStrfmtUUID("PatchAccountIdentification", "account_id"),

		IdentificationID: c.Defaults.GetStrfmtUUID("PatchAccountIdentification", "identification_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type PatchAccountIdentificationRequest struct {

	/*AccountIdentificationAmendRequest*/

	*models.AccountIdentificationRequest

	/*AccountID      Account Id      */

	AccountID strfmt.UUID

	/*IdentificationID      Account Identification Id; Must match id in the payload      */

	IdentificationID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *PatchAccountIdentificationRequest) FromJson(j string) (*PatchAccountIdentificationRequest, error) {

	var m models.AccountIdentificationRequest
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.AccountIdentificationRequest = &m

	return o, nil
}

func (o *PatchAccountIdentificationRequest) WithAccountIdentificationAmendRequest(accountIdentificationAmendRequest models.AccountIdentificationRequest) *PatchAccountIdentificationRequest {

	o.AccountIdentificationRequest = &accountIdentificationAmendRequest

	return o
}

func (o *PatchAccountIdentificationRequest) WithoutAccountIdentificationAmendRequest() *PatchAccountIdentificationRequest {

	o.AccountIdentificationRequest = &models.AccountIdentificationRequest{}

	return o
}

func (o *PatchAccountIdentificationRequest) WithAccountID(accountID strfmt.UUID) *PatchAccountIdentificationRequest {

	o.AccountID = accountID

	return o
}

func (o *PatchAccountIdentificationRequest) WithIdentificationID(identificationID strfmt.UUID) *PatchAccountIdentificationRequest {

	o.IdentificationID = identificationID

	return o
}

// ////////////////
// WithContext adds the context to the patch account identification Request
func (o *PatchAccountIdentificationRequest) WithContext(ctx context.Context) *PatchAccountIdentificationRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the patch account identification Request
func (o *PatchAccountIdentificationRequest) WithHTTPClient(client *http.Client) *PatchAccountIdentificationRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *PatchAccountIdentificationRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.AccountIdentificationRequest != nil {
		if err := r.SetBodyParam(o.AccountIdentificationRequest); err != nil {
			return err
		}
	}

	// path param account_id
	if err := r.SetPathParam("account_id", o.AccountID.String()); err != nil {
		return err
	}

	// path param identification_id
	if err := r.SetPathParam("identification_id", o.IdentificationID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
