// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package go_subscription_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new go subscription api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for go subscription api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get notification subscriptions ID errors API
*/
func (a *GetNotificationSubscriptionsIDErrorsRequest) Do() (*GetNotificationSubscriptionsIDErrorsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNotificationSubscriptionsIDErrors",
		Method:             "GET",
		PathPattern:        "/notification/subscriptions/{id}/errors",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetNotificationSubscriptionsIDErrorsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNotificationSubscriptionsIDErrorsOK), nil

}

func (a *GetNotificationSubscriptionsIDErrorsRequest) MustDo() *GetNotificationSubscriptionsIDErrorsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/////////

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
