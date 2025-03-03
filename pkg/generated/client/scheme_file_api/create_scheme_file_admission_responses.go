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

// CreateSchemeFileAdmissionReader is a Reader for the CreateSchemeFileAdmission structure.
type CreateSchemeFileAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSchemeFileAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSchemeFileAdmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSchemeFileAdmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateSchemeFileAdmissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateSchemeFileAdmissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateSchemeFileAdmissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateSchemeFileAdmissionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateSchemeFileAdmissionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSchemeFileAdmissionCreated creates a CreateSchemeFileAdmissionCreated with default headers values
func NewCreateSchemeFileAdmissionCreated() *CreateSchemeFileAdmissionCreated {
	return &CreateSchemeFileAdmissionCreated{}
}

/*
CreateSchemeFileAdmissionCreated handles this case with default header values.

Scheme File Admission Response
*/
type CreateSchemeFileAdmissionCreated struct {

	//Payload

	// isStream: false
	*models.SchemeFileAdmissionResponse
}

// IsSuccess returns true when this create scheme file admission created response has a 2xx status code
func (o *CreateSchemeFileAdmissionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create scheme file admission created response has a 3xx status code
func (o *CreateSchemeFileAdmissionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file admission created response has a 4xx status code
func (o *CreateSchemeFileAdmissionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create scheme file admission created response has a 5xx status code
func (o *CreateSchemeFileAdmissionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheme file admission created response a status code equal to that given
func (o *CreateSchemeFileAdmissionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create scheme file admission created response
func (o *CreateSchemeFileAdmissionCreated) Code() int {
	return 201
}

func (o *CreateSchemeFileAdmissionCreated) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles/{scheme_file_id}/admissions][%d] createSchemeFileAdmissionCreated", 201)
}

func (o *CreateSchemeFileAdmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SchemeFileAdmissionResponse = new(models.SchemeFileAdmissionResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SchemeFileAdmissionResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchemeFileAdmissionBadRequest creates a CreateSchemeFileAdmissionBadRequest with default headers values
func NewCreateSchemeFileAdmissionBadRequest() *CreateSchemeFileAdmissionBadRequest {
	return &CreateSchemeFileAdmissionBadRequest{}
}

/*
CreateSchemeFileAdmissionBadRequest handles this case with default header values.

Bad Request
*/
type CreateSchemeFileAdmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create scheme file admission bad request response has a 2xx status code
func (o *CreateSchemeFileAdmissionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheme file admission bad request response has a 3xx status code
func (o *CreateSchemeFileAdmissionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file admission bad request response has a 4xx status code
func (o *CreateSchemeFileAdmissionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheme file admission bad request response has a 5xx status code
func (o *CreateSchemeFileAdmissionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheme file admission bad request response a status code equal to that given
func (o *CreateSchemeFileAdmissionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create scheme file admission bad request response
func (o *CreateSchemeFileAdmissionBadRequest) Code() int {
	return 400
}

func (o *CreateSchemeFileAdmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles/{scheme_file_id}/admissions][%d] createSchemeFileAdmissionBadRequest", 400)
}

func (o *CreateSchemeFileAdmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchemeFileAdmissionUnauthorized creates a CreateSchemeFileAdmissionUnauthorized with default headers values
func NewCreateSchemeFileAdmissionUnauthorized() *CreateSchemeFileAdmissionUnauthorized {
	return &CreateSchemeFileAdmissionUnauthorized{}
}

/*
CreateSchemeFileAdmissionUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateSchemeFileAdmissionUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create scheme file admission unauthorized response has a 2xx status code
func (o *CreateSchemeFileAdmissionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheme file admission unauthorized response has a 3xx status code
func (o *CreateSchemeFileAdmissionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file admission unauthorized response has a 4xx status code
func (o *CreateSchemeFileAdmissionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheme file admission unauthorized response has a 5xx status code
func (o *CreateSchemeFileAdmissionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheme file admission unauthorized response a status code equal to that given
func (o *CreateSchemeFileAdmissionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create scheme file admission unauthorized response
func (o *CreateSchemeFileAdmissionUnauthorized) Code() int {
	return 401
}

func (o *CreateSchemeFileAdmissionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles/{scheme_file_id}/admissions][%d] createSchemeFileAdmissionUnauthorized", 401)
}

func (o *CreateSchemeFileAdmissionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchemeFileAdmissionForbidden creates a CreateSchemeFileAdmissionForbidden with default headers values
func NewCreateSchemeFileAdmissionForbidden() *CreateSchemeFileAdmissionForbidden {
	return &CreateSchemeFileAdmissionForbidden{}
}

/*
CreateSchemeFileAdmissionForbidden handles this case with default header values.

Forbidden
*/
type CreateSchemeFileAdmissionForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create scheme file admission forbidden response has a 2xx status code
func (o *CreateSchemeFileAdmissionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheme file admission forbidden response has a 3xx status code
func (o *CreateSchemeFileAdmissionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file admission forbidden response has a 4xx status code
func (o *CreateSchemeFileAdmissionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheme file admission forbidden response has a 5xx status code
func (o *CreateSchemeFileAdmissionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheme file admission forbidden response a status code equal to that given
func (o *CreateSchemeFileAdmissionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create scheme file admission forbidden response
func (o *CreateSchemeFileAdmissionForbidden) Code() int {
	return 403
}

func (o *CreateSchemeFileAdmissionForbidden) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles/{scheme_file_id}/admissions][%d] createSchemeFileAdmissionForbidden", 403)
}

func (o *CreateSchemeFileAdmissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchemeFileAdmissionNotFound creates a CreateSchemeFileAdmissionNotFound with default headers values
func NewCreateSchemeFileAdmissionNotFound() *CreateSchemeFileAdmissionNotFound {
	return &CreateSchemeFileAdmissionNotFound{}
}

/*
CreateSchemeFileAdmissionNotFound handles this case with default header values.

Not Found
*/
type CreateSchemeFileAdmissionNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create scheme file admission not found response has a 2xx status code
func (o *CreateSchemeFileAdmissionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheme file admission not found response has a 3xx status code
func (o *CreateSchemeFileAdmissionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file admission not found response has a 4xx status code
func (o *CreateSchemeFileAdmissionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheme file admission not found response has a 5xx status code
func (o *CreateSchemeFileAdmissionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheme file admission not found response a status code equal to that given
func (o *CreateSchemeFileAdmissionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create scheme file admission not found response
func (o *CreateSchemeFileAdmissionNotFound) Code() int {
	return 404
}

func (o *CreateSchemeFileAdmissionNotFound) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles/{scheme_file_id}/admissions][%d] createSchemeFileAdmissionNotFound", 404)
}

func (o *CreateSchemeFileAdmissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchemeFileAdmissionConflict creates a CreateSchemeFileAdmissionConflict with default headers values
func NewCreateSchemeFileAdmissionConflict() *CreateSchemeFileAdmissionConflict {
	return &CreateSchemeFileAdmissionConflict{}
}

/*
CreateSchemeFileAdmissionConflict handles this case with default header values.

Scheme File Admission Conflict
*/
type CreateSchemeFileAdmissionConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create scheme file admission conflict response has a 2xx status code
func (o *CreateSchemeFileAdmissionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheme file admission conflict response has a 3xx status code
func (o *CreateSchemeFileAdmissionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file admission conflict response has a 4xx status code
func (o *CreateSchemeFileAdmissionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create scheme file admission conflict response has a 5xx status code
func (o *CreateSchemeFileAdmissionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create scheme file admission conflict response a status code equal to that given
func (o *CreateSchemeFileAdmissionConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create scheme file admission conflict response
func (o *CreateSchemeFileAdmissionConflict) Code() int {
	return 409
}

func (o *CreateSchemeFileAdmissionConflict) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles/{scheme_file_id}/admissions][%d] createSchemeFileAdmissionConflict", 409)
}

func (o *CreateSchemeFileAdmissionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchemeFileAdmissionInternalServerError creates a CreateSchemeFileAdmissionInternalServerError with default headers values
func NewCreateSchemeFileAdmissionInternalServerError() *CreateSchemeFileAdmissionInternalServerError {
	return &CreateSchemeFileAdmissionInternalServerError{}
}

/*
CreateSchemeFileAdmissionInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateSchemeFileAdmissionInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create scheme file admission internal server error response has a 2xx status code
func (o *CreateSchemeFileAdmissionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create scheme file admission internal server error response has a 3xx status code
func (o *CreateSchemeFileAdmissionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create scheme file admission internal server error response has a 4xx status code
func (o *CreateSchemeFileAdmissionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create scheme file admission internal server error response has a 5xx status code
func (o *CreateSchemeFileAdmissionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create scheme file admission internal server error response a status code equal to that given
func (o *CreateSchemeFileAdmissionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create scheme file admission internal server error response
func (o *CreateSchemeFileAdmissionInternalServerError) Code() int {
	return 500
}

func (o *CreateSchemeFileAdmissionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /files/schemefiles/{scheme_file_id}/admissions][%d] createSchemeFileAdmissionInternalServerError", 500)
}

func (o *CreateSchemeFileAdmissionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
