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

// CreateSchemeFileReader is a Reader for the CreateSchemeFile structure.
type CreateSchemeFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSchemeFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSchemeFileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSchemeFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateSchemeFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateSchemeFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateSchemeFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateSchemeFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSchemeFileCreated creates a CreateSchemeFileCreated with default headers values
func NewCreateSchemeFileCreated() *CreateSchemeFileCreated {
	return &CreateSchemeFileCreated{}
}

/*
CreateSchemeFileCreated handles this case with default header values.

Scheme File Creation Response
*/
type CreateSchemeFileCreated struct {

	//Payload

	// isStream: false
	*models.SchemeFileResponse
}

// IsSuccess returns true when this create scheme file created response has a 2xx status code
func (o *CreateSchemeFileCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create scheme file created response has a 3xx status code
func (o *CreateSchemeFileCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file created response has a 4xx status code
func (o *CreateSchemeFileCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create scheme file created response has a 5xx status code
func (o *CreateSchemeFileCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheme file created response a status code equal to that given
func (o *CreateSchemeFileCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create scheme file created response
func (o *CreateSchemeFileCreated) Code() int {
	return 201
}

func (o *CreateSchemeFileCreated) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles][%d] createSchemeFileCreated", 201)
}

func (o *CreateSchemeFileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SchemeFileResponse = new(models.SchemeFileResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SchemeFileResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchemeFileBadRequest creates a CreateSchemeFileBadRequest with default headers values
func NewCreateSchemeFileBadRequest() *CreateSchemeFileBadRequest {
	return &CreateSchemeFileBadRequest{}
}

/*
CreateSchemeFileBadRequest handles this case with default header values.

Bad Request
*/
type CreateSchemeFileBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create scheme file bad request response has a 2xx status code
func (o *CreateSchemeFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheme file bad request response has a 3xx status code
func (o *CreateSchemeFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file bad request response has a 4xx status code
func (o *CreateSchemeFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheme file bad request response has a 5xx status code
func (o *CreateSchemeFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheme file bad request response a status code equal to that given
func (o *CreateSchemeFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create scheme file bad request response
func (o *CreateSchemeFileBadRequest) Code() int {
	return 400
}

func (o *CreateSchemeFileBadRequest) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles][%d] createSchemeFileBadRequest", 400)
}

func (o *CreateSchemeFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchemeFileUnauthorized creates a CreateSchemeFileUnauthorized with default headers values
func NewCreateSchemeFileUnauthorized() *CreateSchemeFileUnauthorized {
	return &CreateSchemeFileUnauthorized{}
}

/*
CreateSchemeFileUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateSchemeFileUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create scheme file unauthorized response has a 2xx status code
func (o *CreateSchemeFileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheme file unauthorized response has a 3xx status code
func (o *CreateSchemeFileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file unauthorized response has a 4xx status code
func (o *CreateSchemeFileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheme file unauthorized response has a 5xx status code
func (o *CreateSchemeFileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheme file unauthorized response a status code equal to that given
func (o *CreateSchemeFileUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create scheme file unauthorized response
func (o *CreateSchemeFileUnauthorized) Code() int {
	return 401
}

func (o *CreateSchemeFileUnauthorized) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles][%d] createSchemeFileUnauthorized", 401)
}

func (o *CreateSchemeFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchemeFileForbidden creates a CreateSchemeFileForbidden with default headers values
func NewCreateSchemeFileForbidden() *CreateSchemeFileForbidden {
	return &CreateSchemeFileForbidden{}
}

/*
CreateSchemeFileForbidden handles this case with default header values.

Forbidden
*/
type CreateSchemeFileForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create scheme file forbidden response has a 2xx status code
func (o *CreateSchemeFileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheme file forbidden response has a 3xx status code
func (o *CreateSchemeFileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file forbidden response has a 4xx status code
func (o *CreateSchemeFileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheme file forbidden response has a 5xx status code
func (o *CreateSchemeFileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheme file forbidden response a status code equal to that given
func (o *CreateSchemeFileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create scheme file forbidden response
func (o *CreateSchemeFileForbidden) Code() int {
	return 403
}

func (o *CreateSchemeFileForbidden) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles][%d] createSchemeFileForbidden", 403)
}

func (o *CreateSchemeFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchemeFileConflict creates a CreateSchemeFileConflict with default headers values
func NewCreateSchemeFileConflict() *CreateSchemeFileConflict {
	return &CreateSchemeFileConflict{}
}

/*
CreateSchemeFileConflict handles this case with default header values.

Conflict
*/
type CreateSchemeFileConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create scheme file conflict response has a 2xx status code
func (o *CreateSchemeFileConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheme file conflict response has a 3xx status code
func (o *CreateSchemeFileConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file conflict response has a 4xx status code
func (o *CreateSchemeFileConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheme file conflict response has a 5xx status code
func (o *CreateSchemeFileConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheme file conflict response a status code equal to that given
func (o *CreateSchemeFileConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create scheme file conflict response
func (o *CreateSchemeFileConflict) Code() int {
	return 409
}

func (o *CreateSchemeFileConflict) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles][%d] createSchemeFileConflict", 409)
}

func (o *CreateSchemeFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchemeFileInternalServerError creates a CreateSchemeFileInternalServerError with default headers values
func NewCreateSchemeFileInternalServerError() *CreateSchemeFileInternalServerError {
	return &CreateSchemeFileInternalServerError{}
}

/*
CreateSchemeFileInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateSchemeFileInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create scheme file internal server error response has a 2xx status code
func (o *CreateSchemeFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheme file internal server error response has a 3xx status code
func (o *CreateSchemeFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file internal server error response has a 4xx status code
func (o *CreateSchemeFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create scheme file internal server error response has a 5xx status code
func (o *CreateSchemeFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create scheme file internal server error response a status code equal to that given
func (o *CreateSchemeFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create scheme file internal server error response
func (o *CreateSchemeFileInternalServerError) Code() int {
	return 500
}

func (o *CreateSchemeFileInternalServerError) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles][%d] createSchemeFileInternalServerError", 500)
}

func (o *CreateSchemeFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
