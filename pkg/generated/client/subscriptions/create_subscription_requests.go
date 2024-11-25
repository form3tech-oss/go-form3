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

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// Client.CreateSubscription creates a new CreateSubscriptionRequest object
// with the default values initialized.
func (c *Client) CreateSubscription() *CreateSubscriptionRequest {
	var ()
	return &CreateSubscriptionRequest{

		SubscriptionCreation: models.SubscriptionCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateSubscriptionRequest struct {

	/*SubscriptionCreationRequest*/

	*models.SubscriptionCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateSubscriptionRequest) FromJson(j string) (*CreateSubscriptionRequest, error) {

	var m models.SubscriptionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.SubscriptionCreation = &m

	return o, nil
}

func (o *CreateSubscriptionRequest) WithSubscriptionCreationRequest(subscriptionCreationRequest models.SubscriptionCreation) *CreateSubscriptionRequest {

	o.SubscriptionCreation = &subscriptionCreationRequest

	return o
}

func (o *CreateSubscriptionRequest) WithoutSubscriptionCreationRequest() *CreateSubscriptionRequest {

	o.SubscriptionCreation = &models.SubscriptionCreation{}

	return o
}

// ////////////////
// WithContext adds the context to the create subscription Request
func (o *CreateSubscriptionRequest) WithContext(ctx context.Context) *CreateSubscriptionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create subscription Request
func (o *CreateSubscriptionRequest) WithHTTPClient(client *http.Client) *CreateSubscriptionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateSubscriptionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.SubscriptionCreation != nil {
		if err := r.SetBodyParam(o.SubscriptionCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
