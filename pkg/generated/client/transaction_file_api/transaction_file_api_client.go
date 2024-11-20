// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package transaction_file_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new transaction file api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for transaction file api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
create transaction file API
*/
func (a *CreateTransactionFileRequest) Do() (*CreateTransactionFileCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateTransactionFile",
		Method:             "POST",
		PathPattern:        "/files/transactions",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateTransactionFileReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateTransactionFileCreated), nil

}

func (a *CreateTransactionFileRequest) MustDo() *CreateTransactionFileCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create transaction file admission API
*/
func (a *CreateTransactionFileAdmissionRequest) Do() (*CreateTransactionFileAdmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateTransactionFileAdmission",
		Method:             "POST",
		PathPattern:        "/files/transactions/{transaction_file_id}/admissions",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateTransactionFileAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateTransactionFileAdmissionCreated), nil

}

func (a *CreateTransactionFileAdmissionRequest) MustDo() *CreateTransactionFileAdmissionCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create transaction file submission API
*/
func (a *CreateTransactionFileSubmissionRequest) Do() (*CreateTransactionFileSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateTransactionFileSubmission",
		Method:             "POST",
		PathPattern:        "/files/transactions/{transaction_file_id}/submissions",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateTransactionFileSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateTransactionFileSubmissionCreated), nil

}

func (a *CreateTransactionFileSubmissionRequest) MustDo() *CreateTransactionFileSubmissionCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction file API
*/
func (a *GetTransactionFileRequest) Do() (*GetTransactionFileOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionFile",
		Method:             "GET",
		PathPattern:        "/files/transactions/{transaction_file_id}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionFileReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionFileOK), nil

}

func (a *GetTransactionFileRequest) MustDo() *GetTransactionFileOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction file admission API
*/
func (a *GetTransactionFileAdmissionRequest) Do() (*GetTransactionFileAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionFileAdmission",
		Method:             "GET",
		PathPattern:        "/files/transactions/{transaction_file_id}/admissions/{transaction_file_admission_id}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionFileAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionFileAdmissionOK), nil

}

func (a *GetTransactionFileAdmissionRequest) MustDo() *GetTransactionFileAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction file submission API
*/
func (a *GetTransactionFileSubmissionRequest) Do() (*GetTransactionFileSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionFileSubmission",
		Method:             "GET",
		PathPattern:        "/files/transactions/{transaction_file_id}/submissions/{transaction_file_submission_id}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionFileSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionFileSubmissionOK), nil

}

func (a *GetTransactionFileSubmissionRequest) MustDo() *GetTransactionFileSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list transaction files API
*/
func (a *ListTransactionFilesRequest) Do() (*ListTransactionFilesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListTransactionFiles",
		Method:             "GET",
		PathPattern:        "/files/transactions",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListTransactionFilesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListTransactionFilesOK), nil

}

func (a *ListTransactionFilesRequest) MustDo() *ListTransactionFilesOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
upload transaction file API
*/
func (a *UploadTransactionFileRequest) Do() (*UploadTransactionFileOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UploadTransactionFile",
		Method:             "PUT",
		PathPattern:        "/files/transactions/{transaction_file_id}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &UploadTransactionFileReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UploadTransactionFileOK), nil

}

func (a *UploadTransactionFileRequest) MustDo() *UploadTransactionFileOK {
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
