// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package query_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new query api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for query api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get transaction queries API
*/
func (a *GetTransactionQueriesRequest) Do() (*GetTransactionQueriesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionQueries",
		Method:             "GET",
		PathPattern:        "/transaction/queries",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionQueriesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionQueriesOK), nil

}

func (a *GetTransactionQueriesRequest) MustDo() *GetTransactionQueriesOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction queries query ID API
*/
func (a *GetTransactionQueriesQueryIDRequest) Do() (*GetTransactionQueriesQueryIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionQueriesQueryID",
		Method:             "GET",
		PathPattern:        "/transaction/queries/{query_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionQueriesQueryIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionQueriesQueryIDOK), nil

}

func (a *GetTransactionQueriesQueryIDRequest) MustDo() *GetTransactionQueriesQueryIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction queries query ID admissions query admission ID API
*/
func (a *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDRequest) Do() (*GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionQueriesQueryIDAdmissionsQueryAdmissionID",
		Method:             "GET",
		PathPattern:        "/transaction/queries/{query_id}/admissions/{query_admission_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK), nil

}

func (a *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDRequest) MustDo() *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction queries query ID responses query response ID API
*/
func (a *GetTransactionQueriesQueryIDResponsesQueryResponseIDRequest) Do() (*GetTransactionQueriesQueryIDResponsesQueryResponseIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionQueriesQueryIDResponsesQueryResponseID",
		Method:             "GET",
		PathPattern:        "/transaction/queries/{query_id}/responses/{query_response_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionQueriesQueryIDResponsesQueryResponseIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionQueriesQueryIDResponsesQueryResponseIDOK), nil

}

func (a *GetTransactionQueriesQueryIDResponsesQueryResponseIDRequest) MustDo() *GetTransactionQueriesQueryIDResponsesQueryResponseIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction queries query ID responses query response ID submissions query response submission ID API
*/
func (a *GetTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsQueryResponseSubmissionIDRequest) Do() (*GetTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsQueryResponseSubmissionIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsQueryResponseSubmissionID",
		Method:             "GET",
		PathPattern:        "/transaction/queries/{query_id}/responses/{query_response_id}/submissions/{query_response_submission_id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsQueryResponseSubmissionIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsQueryResponseSubmissionIDOK), nil

}

func (a *GetTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsQueryResponseSubmissionIDRequest) MustDo() *GetTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsQueryResponseSubmissionIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post transaction queries query ID responses API
*/
func (a *PostTransactionQueriesQueryIDResponsesRequest) Do() (*PostTransactionQueriesQueryIDResponsesCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTransactionQueriesQueryIDResponses",
		Method:             "POST",
		PathPattern:        "/transaction/queries/{query_id}/responses",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostTransactionQueriesQueryIDResponsesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTransactionQueriesQueryIDResponsesCreated), nil

}

func (a *PostTransactionQueriesQueryIDResponsesRequest) MustDo() *PostTransactionQueriesQueryIDResponsesCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post transaction queries query ID responses query response ID submissions API
*/
func (a *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsRequest) Do() (*PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissions",
		Method:             "POST",
		PathPattern:        "/transaction/queries/{query_id}/responses/{query_response_id}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated), nil

}

func (a *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsRequest) MustDo() *PostTransactionQueriesQueryIDResponsesQueryResponseIDSubmissionsCreated {
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
