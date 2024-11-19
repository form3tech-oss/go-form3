// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

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

// Client.ModifySubscription creates a new ModifySubscriptionRequest object
// with the default values initialized.
func (c *Client) ModifySubscription() *ModifySubscriptionRequest {
	var ()
	return &ModifySubscriptionRequest{

		SubscriptionAmendment: models.SubscriptionAmendmentWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("ModifySubscription", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ModifySubscriptionRequest struct {

	/*SubscriptionUpdateRequest*/

	*models.SubscriptionAmendment

	/*ID      Subscription Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ModifySubscriptionRequest) FromJson(j string) (*ModifySubscriptionRequest, error) {

	var m models.SubscriptionAmendment
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.SubscriptionAmendment = &m

	return o, nil
}

func (o *ModifySubscriptionRequest) WithSubscriptionUpdateRequest(subscriptionUpdateRequest models.SubscriptionAmendment) *ModifySubscriptionRequest {

	o.SubscriptionAmendment = &subscriptionUpdateRequest

	return o
}

func (o *ModifySubscriptionRequest) WithoutSubscriptionUpdateRequest() *ModifySubscriptionRequest {

	o.SubscriptionAmendment = &models.SubscriptionAmendment{}

	return o
}

func (o *ModifySubscriptionRequest) WithID(id strfmt.UUID) *ModifySubscriptionRequest {

	o.ID = id

	return o
}

// ////////////////
// WithContext adds the context to the modify subscription Request
func (o *ModifySubscriptionRequest) WithContext(ctx context.Context) *ModifySubscriptionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the modify subscription Request
func (o *ModifySubscriptionRequest) WithHTTPClient(client *http.Client) *ModifySubscriptionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ModifySubscriptionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.SubscriptionAmendment != nil {
		if err := r.SetBodyParam(o.SubscriptionAmendment); err != nil {
			return err
		}
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
