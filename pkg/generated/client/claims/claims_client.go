// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package claims

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new claims API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for claims API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
create claim API
*/
func (a *CreateClaimRequest) Do() (*CreateClaimCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateClaim",
		Method:             "POST",
		PathPattern:        "/transaction/claims",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateClaimReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClaimCreated), nil

}

func (a *CreateClaimRequest) MustDo() *CreateClaimCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create claim reversal API
*/
func (a *CreateClaimReversalRequest) Do() (*CreateClaimReversalCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateClaimReversal",
		Method:             "POST",
		PathPattern:        "/transaction/claims/{id}/reversals",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateClaimReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClaimReversalCreated), nil

}

func (a *CreateClaimReversalRequest) MustDo() *CreateClaimReversalCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create claim reversal submission API
*/
func (a *CreateClaimReversalSubmissionRequest) Do() (*CreateClaimReversalSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateClaimReversalSubmission",
		Method:             "POST",
		PathPattern:        "/transaction/claims/{id}/reversals/{reversalId}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateClaimReversalSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClaimReversalSubmissionCreated), nil

}

func (a *CreateClaimReversalSubmissionRequest) MustDo() *CreateClaimReversalSubmissionCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create claim submission API
*/
func (a *CreateClaimSubmissionRequest) Do() (*CreateClaimSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateClaimSubmission",
		Method:             "POST",
		PathPattern:        "/transaction/claims/{id}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateClaimSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClaimSubmissionCreated), nil

}

func (a *CreateClaimSubmissionRequest) MustDo() *CreateClaimSubmissionCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get claim API
*/
func (a *GetClaimRequest) Do() (*GetClaimOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClaim",
		Method:             "GET",
		PathPattern:        "/transaction/claims/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetClaimReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClaimOK), nil

}

func (a *GetClaimRequest) MustDo() *GetClaimOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get claim reversal API
*/
func (a *GetClaimReversalRequest) Do() (*GetClaimReversalOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClaimReversal",
		Method:             "GET",
		PathPattern:        "/transaction/claims/{id}/reversals/{reversalId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetClaimReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClaimReversalOK), nil

}

func (a *GetClaimReversalRequest) MustDo() *GetClaimReversalOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get claim reversal submission API
*/
func (a *GetClaimReversalSubmissionRequest) Do() (*GetClaimReversalSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClaimReversalSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/claims/{id}/reversals/{reversalId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetClaimReversalSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClaimReversalSubmissionOK), nil

}

func (a *GetClaimReversalSubmissionRequest) MustDo() *GetClaimReversalSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get claim submission API
*/
func (a *GetClaimSubmissionRequest) Do() (*GetClaimSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClaimSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/claims/{id}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetClaimSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClaimSubmissionOK), nil

}

func (a *GetClaimSubmissionRequest) MustDo() *GetClaimSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list claims API
*/
func (a *ListClaimsRequest) Do() (*ListClaimsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListClaims",
		Method:             "GET",
		PathPattern:        "/transaction/claims",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListClaimsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListClaimsOK), nil

}

func (a *ListClaimsRequest) MustDo() *ListClaimsOK {
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
