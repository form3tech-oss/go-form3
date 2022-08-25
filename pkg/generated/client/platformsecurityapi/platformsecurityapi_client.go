// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package platformsecurityapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new platformsecurityapi API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for platformsecurityapi API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get platform security signing keys API
*/
func (a *GetPlatformSecuritySigningKeysRequest) Do() (*GetPlatformSecuritySigningKeysOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPlatformSecuritySigningKeys",
		Method:             "GET",
		PathPattern:        "/platform/security/signing_keys",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPlatformSecuritySigningKeysReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlatformSecuritySigningKeysOK), nil

}

func (a *GetPlatformSecuritySigningKeysRequest) MustDo() *GetPlatformSecuritySigningKeysOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get platform security signing keys signingkey ID API
*/
func (a *GetPlatformSecuritySigningKeysSigningkeyIDRequest) Do() (*GetPlatformSecuritySigningKeysSigningkeyIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPlatformSecuritySigningKeysSigningkeyID",
		Method:             "GET",
		PathPattern:        "/platform/security/signing_keys/{signingkey_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPlatformSecuritySigningKeysSigningkeyIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlatformSecuritySigningKeysSigningkeyIDOK), nil

}

func (a *GetPlatformSecuritySigningKeysSigningkeyIDRequest) MustDo() *GetPlatformSecuritySigningKeysSigningkeyIDOK {
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
