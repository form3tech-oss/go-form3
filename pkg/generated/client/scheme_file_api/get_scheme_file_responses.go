// Code generated by go-swagger; DO NOT EDIT.

package scheme_file_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
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

// IsSuccess returns true when this get scheme file o k response has a 2xx status code
func (o *GetSchemeFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scheme file o k response has a 3xx status code
func (o *GetSchemeFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scheme file o k response has a 4xx status code
func (o *GetSchemeFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scheme file o k response has a 5xx status code
func (o *GetSchemeFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scheme file o k response a status code equal to that given
func (o *GetSchemeFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get scheme file o k response
func (o *GetSchemeFileOK) Code() int {
	return 200
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

// IsSuccess returns true when this get scheme file unauthorized response has a 2xx status code
func (o *GetSchemeFileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scheme file unauthorized response has a 3xx status code
func (o *GetSchemeFileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scheme file unauthorized response has a 4xx status code
func (o *GetSchemeFileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scheme file unauthorized response has a 5xx status code
func (o *GetSchemeFileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scheme file unauthorized response a status code equal to that given
func (o *GetSchemeFileUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get scheme file unauthorized response
func (o *GetSchemeFileUnauthorized) Code() int {
	return 401
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

// IsSuccess returns true when this get scheme file forbidden response has a 2xx status code
func (o *GetSchemeFileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scheme file forbidden response has a 3xx status code
func (o *GetSchemeFileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scheme file forbidden response has a 4xx status code
func (o *GetSchemeFileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scheme file forbidden response has a 5xx status code
func (o *GetSchemeFileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scheme file forbidden response a status code equal to that given
func (o *GetSchemeFileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get scheme file forbidden response
func (o *GetSchemeFileForbidden) Code() int {
	return 403
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

// IsSuccess returns true when this get scheme file not found response has a 2xx status code
func (o *GetSchemeFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scheme file not found response has a 3xx status code
func (o *GetSchemeFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scheme file not found response has a 4xx status code
func (o *GetSchemeFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scheme file not found response has a 5xx status code
func (o *GetSchemeFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get scheme file not found response a status code equal to that given
func (o *GetSchemeFileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get scheme file not found response
func (o *GetSchemeFileNotFound) Code() int {
	return 404
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

// IsSuccess returns true when this get scheme file internal server error response has a 2xx status code
func (o *GetSchemeFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scheme file internal server error response has a 3xx status code
func (o *GetSchemeFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scheme file internal server error response has a 4xx status code
func (o *GetSchemeFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scheme file internal server error response has a 5xx status code
func (o *GetSchemeFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scheme file internal server error response a status code equal to that given
func (o *GetSchemeFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get scheme file internal server error response
func (o *GetSchemeFileInternalServerError) Code() int {
	return 500
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
