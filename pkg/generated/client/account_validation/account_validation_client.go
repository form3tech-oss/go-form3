// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package account_validation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new account validation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for account validation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get validations gbdsc sortcodes sortcode API
*/
func (a *GetValidationsGbdscSortcodesSortcodeRequest) Do() (*GetValidationsGbdscSortcodesSortcodeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetValidationsGbdscSortcodesSortcode",
		Method:             "GET",
		PathPattern:        "/validations/gbdsc/sortcodes/{sortcode}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetValidationsGbdscSortcodesSortcodeReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetValidationsGbdscSortcodesSortcodeOK), nil

}

func (a *GetValidationsGbdscSortcodesSortcodeRequest) MustDo() *GetValidationsGbdscSortcodesSortcodeOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get validations gbdsc sortcodes sortcode accountnumbers accountnumber API
*/
func (a *GetValidationsGbdscSortcodesSortcodeAccountnumbersAccountnumberRequest) Do() (*GetValidationsGbdscSortcodesSortcodeAccountnumbersAccountnumberOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetValidationsGbdscSortcodesSortcodeAccountnumbersAccountnumber",
		Method:             "GET",
		PathPattern:        "/validations/gbdsc/sortcodes/{sortcode}/accountnumbers/{accountnumber}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetValidationsGbdscSortcodesSortcodeAccountnumbersAccountnumberReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetValidationsGbdscSortcodesSortcodeAccountnumbersAccountnumberOK), nil

}

func (a *GetValidationsGbdscSortcodesSortcodeAccountnumbersAccountnumberRequest) MustDo() *GetValidationsGbdscSortcodesSortcodeAccountnumbersAccountnumberOK {
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
