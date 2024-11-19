// Code generated by go-swagger; DO NOT EDIT.

package scheme_file_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// GetSchemeFileReader is a Reader for the GetSchemeFile structure.
type GetSchemeFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchemeFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSchemeFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSchemeFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSchemeFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSchemeFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSchemeFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSchemeFileOK creates a GetSchemeFileOK with default headers values
func NewGetSchemeFileOK() *GetSchemeFileOK {
	return &GetSchemeFileOK{}
}

/*
GetSchemeFileOK handles this case with default header values.

Scheme File Response
*/
type GetSchemeFileOK struct {

	//Payload

	// isStream: false
	*models.SchemeFileResponse
}

func (o *GetSchemeFileOK) Error() string {
	return fmt.Sprintf("[GET /files/schemefiles/{scheme_file_id}][%d] getSchemeFileOK", 200)
}

func (o *GetSchemeFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SchemeFileResponse = new(models.SchemeFileResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SchemeFileResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemeFileUnauthorized creates a GetSchemeFileUnauthorized with default headers values
func NewGetSchemeFileUnauthorized() *GetSchemeFileUnauthorized {
	return &GetSchemeFileUnauthorized{}
}

/*
GetSchemeFileUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSchemeFileUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetSchemeFileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /files/schemefiles/{scheme_file_id}][%d] getSchemeFileUnauthorized", 401)
}

func (o *GetSchemeFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemeFileForbidden creates a GetSchemeFileForbidden with default headers values
func NewGetSchemeFileForbidden() *GetSchemeFileForbidden {
	return &GetSchemeFileForbidden{}
}

/*
GetSchemeFileForbidden handles this case with default header values.

Forbidden
*/
type GetSchemeFileForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetSchemeFileForbidden) Error() string {
	return fmt.Sprintf("[GET /files/schemefiles/{scheme_file_id}][%d] getSchemeFileForbidden", 403)
}

func (o *GetSchemeFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemeFileNotFound creates a GetSchemeFileNotFound with default headers values
func NewGetSchemeFileNotFound() *GetSchemeFileNotFound {
	return &GetSchemeFileNotFound{}
}

/*
GetSchemeFileNotFound handles this case with default header values.

Not Found
*/
type GetSchemeFileNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetSchemeFileNotFound) Error() string {
	return fmt.Sprintf("[GET /files/schemefiles/{scheme_file_id}][%d] getSchemeFileNotFound", 404)
}

func (o *GetSchemeFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemeFileInternalServerError creates a GetSchemeFileInternalServerError with default headers values
func NewGetSchemeFileInternalServerError() *GetSchemeFileInternalServerError {
	return &GetSchemeFileInternalServerError{}
}

/*
GetSchemeFileInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSchemeFileInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetSchemeFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /files/schemefiles/{scheme_file_id}][%d] getSchemeFileInternalServerError", 500)
}

func (o *GetSchemeFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
