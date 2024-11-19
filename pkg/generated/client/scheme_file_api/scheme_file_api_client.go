// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package scheme_file_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new scheme file api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for scheme file api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
create scheme file API
*/
func (a *CreateSchemeFileRequest) Do() (*CreateSchemeFileCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateSchemeFile",
		Method:             "POST",
		PathPattern:        "/files/schemefiles",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateSchemeFileReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSchemeFileCreated), nil

}

func (a *CreateSchemeFileRequest) MustDo() *CreateSchemeFileCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create scheme file admission API
*/
func (a *CreateSchemeFileAdmissionRequest) Do() (*CreateSchemeFileAdmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateSchemeFileAdmission",
		Method:             "POST",
		PathPattern:        "/files/schemefiles/{scheme_file_id}/admissions",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateSchemeFileAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSchemeFileAdmissionCreated), nil

}

func (a *CreateSchemeFileAdmissionRequest) MustDo() *CreateSchemeFileAdmissionCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create scheme file submission API
*/
func (a *CreateSchemeFileSubmissionRequest) Do() (*CreateSchemeFileSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateSchemeFileSubmission",
		Method:             "POST",
		PathPattern:        "/files/schemefiles/{scheme_file_id}/submissions",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateSchemeFileSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSchemeFileSubmissionCreated), nil

}

func (a *CreateSchemeFileSubmissionRequest) MustDo() *CreateSchemeFileSubmissionCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get scheme file API
*/
func (a *GetSchemeFileRequest) Do() (*GetSchemeFileOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSchemeFile",
		Method:             "GET",
		PathPattern:        "/files/schemefiles/{scheme_file_id}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetSchemeFileReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSchemeFileOK), nil

}

func (a *GetSchemeFileRequest) MustDo() *GetSchemeFileOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get scheme file admission API
*/
func (a *GetSchemeFileAdmissionRequest) Do() (*GetSchemeFileAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSchemeFileAdmission",
		Method:             "GET",
		PathPattern:        "/files/schemefiles/{scheme_file_id}/admissions/{scheme_file_admission_id}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetSchemeFileAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSchemeFileAdmissionOK), nil

}

func (a *GetSchemeFileAdmissionRequest) MustDo() *GetSchemeFileAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get scheme file submission API
*/
func (a *GetSchemeFileSubmissionRequest) Do() (*GetSchemeFileSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSchemeFileSubmission",
		Method:             "GET",
		PathPattern:        "/files/schemefiles/{scheme_file_id}/submissions/{scheme_file_submission_id}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetSchemeFileSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSchemeFileSubmissionOK), nil

}

func (a *GetSchemeFileSubmissionRequest) MustDo() *GetSchemeFileSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list scheme files API
*/
func (a *ListSchemeFilesRequest) Do() (*ListSchemeFilesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListSchemeFiles",
		Method:             "GET",
		PathPattern:        "/files/schemefiles",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListSchemeFilesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSchemeFilesOK), nil

}

func (a *ListSchemeFilesRequest) MustDo() *ListSchemeFilesOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
upload scheme file API
*/
func (a *UploadSchemeFileRequest) Do() (*UploadSchemeFileOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UploadSchemeFile",
		Method:             "PUT",
		PathPattern:        "/files/schemefiles/{scheme_file_id}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &UploadSchemeFileReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UploadSchemeFileOK), nil

}

func (a *UploadSchemeFileRequest) MustDo() *UploadSchemeFileOK {
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
