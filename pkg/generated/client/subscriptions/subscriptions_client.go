// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new subscriptions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for subscriptions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
create subscription API
*/
func (a *CreateSubscriptionRequest) Do() (*CreateSubscriptionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateSubscription",
		Method:             "POST",
		PathPattern:        "/notification/subscriptions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateSubscriptionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSubscriptionCreated), nil

}

func (a *CreateSubscriptionRequest) MustDo() *CreateSubscriptionCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
delete subscription API
*/
func (a *DeleteSubscriptionRequest) Do() (*DeleteSubscriptionNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteSubscription",
		Method:             "DELETE",
		PathPattern:        "/notification/subscriptions/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &DeleteSubscriptionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSubscriptionNoContent), nil

}

func (a *DeleteSubscriptionRequest) MustDo() *DeleteSubscriptionNoContent {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get subscription API
*/
func (a *GetSubscriptionRequest) Do() (*GetSubscriptionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSubscription",
		Method:             "GET",
		PathPattern:        "/notification/subscriptions/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetSubscriptionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionOK), nil

}

func (a *GetSubscriptionRequest) MustDo() *GetSubscriptionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get subscriptions health API
*/
func (a *GetSubscriptionsHealthRequest) Do() (*GetSubscriptionsHealthOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSubscriptionsHealth",
		Method:             "GET",
		PathPattern:        "/notification/subscriptions/health",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetSubscriptionsHealthReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionsHealthOK), nil

}

func (a *GetSubscriptionsHealthRequest) MustDo() *GetSubscriptionsHealthOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
list subscriptions API
*/
func (a *ListSubscriptionsRequest) Do() (*ListSubscriptionsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListSubscriptions",
		Method:             "GET",
		PathPattern:        "/notification/subscriptions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListSubscriptionsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSubscriptionsOK), nil

}

func (a *ListSubscriptionsRequest) MustDo() *ListSubscriptionsOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
modify subscription API
*/
func (a *ModifySubscriptionRequest) Do() (*ModifySubscriptionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifySubscription",
		Method:             "PATCH",
		PathPattern:        "/notification/subscriptions/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ModifySubscriptionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifySubscriptionOK), nil

}

func (a *ModifySubscriptionRequest) MustDo() *ModifySubscriptionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/////////

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
