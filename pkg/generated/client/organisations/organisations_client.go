// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package organisations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organisations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for organisations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
create unit API
*/
func (a *CreateUnitRequest) Do() (*CreateUnitCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateUnit",
		Method:             "POST",
		PathPattern:        "/organisation/units",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateUnitReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUnitCreated), nil

}

func (a *CreateUnitRequest) MustDo() *CreateUnitCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get unit API
*/
func (a *GetUnitRequest) Do() (*GetUnitOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUnit",
		Method:             "GET",
		PathPattern:        "/organisation/units/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetUnitReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUnitOK), nil

}

func (a *GetUnitRequest) MustDo() *GetUnitOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list units API
*/
func (a *ListUnitsRequest) Do() (*ListUnitsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListUnits",
		Method:             "GET",
		PathPattern:        "/organisation/units",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListUnitsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListUnitsOK), nil

}

func (a *ListUnitsRequest) MustDo() *ListUnitsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
modify unit API
*/
func (a *ModifyUnitRequest) Do() (*ModifyUnitOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyUnit",
		Method:             "PATCH",
		PathPattern:        "/organisation/units/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ModifyUnitReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyUnitOK), nil

}

func (a *ModifyUnitRequest) MustDo() *ModifyUnitOK {
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
