// Code generated by go-swagger; DO NOT EDIT.

package query_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// GetQueryReader is a Reader for the GetQuery structure.
type GetQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetQueryOK creates a GetQueryOK with default headers values
func NewGetQueryOK() *GetQueryOK {
	return &GetQueryOK{}
}

/*
GetQueryOK handles this case with default header values.

query response
*/
type GetQueryOK struct {

	//Payload

	// isStream: false
	*models.QueryFetchResponse
}

func (o *GetQueryOK) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}][%d] getQueryOK", 200)
}

func (o *GetQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.QueryFetchResponse = new(models.QueryFetchResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.QueryFetchResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQueryBadRequest creates a GetQueryBadRequest with default headers values
func NewGetQueryBadRequest() *GetQueryBadRequest {
	return &GetQueryBadRequest{}
}

/*
GetQueryBadRequest handles this case with default header values.

bad request
*/
type GetQueryBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetQueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}][%d] getQueryBadRequest", 400)
}

func (o *GetQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQueryForbidden creates a GetQueryForbidden with default headers values
func NewGetQueryForbidden() *GetQueryForbidden {
	return &GetQueryForbidden{}
}

/*
GetQueryForbidden handles this case with default header values.

forbidden
*/
type GetQueryForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetQueryForbidden) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}][%d] getQueryForbidden", 403)
}

func (o *GetQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
