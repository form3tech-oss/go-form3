// Code generated by go-swagger; DO NOT EDIT.

package name_verification_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// PostOrganisationNameverificationsReader is a Reader for the PostOrganisationNameverifications structure.
type PostOrganisationNameverificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrganisationNameverificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostOrganisationNameverificationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostOrganisationNameverificationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostOrganisationNameverificationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostOrganisationNameverificationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostOrganisationNameverificationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostOrganisationNameverificationsCreated creates a PostOrganisationNameverificationsCreated with default headers values
func NewPostOrganisationNameverificationsCreated() *PostOrganisationNameverificationsCreated {
	return &PostOrganisationNameverificationsCreated{}
}

/*
PostOrganisationNameverificationsCreated handles this case with default header values.

created
*/
type PostOrganisationNameverificationsCreated struct {

	//Payload

	// isStream: false
	*models.NameVerificationCreationResponse
}

// IsSuccess returns true when this post organisation nameverifications created response has a 2xx status code
func (o *PostOrganisationNameverificationsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post organisation nameverifications created response has a 3xx status code
func (o *PostOrganisationNameverificationsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post organisation nameverifications created response has a 4xx status code
func (o *PostOrganisationNameverificationsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post organisation nameverifications created response has a 5xx status code
func (o *PostOrganisationNameverificationsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post organisation nameverifications created response a status code equal to that given
func (o *PostOrganisationNameverificationsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post organisation nameverifications created response
func (o *PostOrganisationNameverificationsCreated) Code() int {
	return 201
}

func (o *PostOrganisationNameverificationsCreated) Error() string {
	return fmt.Sprintf("[POST /organisation/nameverifications][%d] postOrganisationNameverificationsCreated", 201)
}

func (o *PostOrganisationNameverificationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.NameVerificationCreationResponse = new(models.NameVerificationCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.NameVerificationCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrganisationNameverificationsBadRequest creates a PostOrganisationNameverificationsBadRequest with default headers values
func NewPostOrganisationNameverificationsBadRequest() *PostOrganisationNameverificationsBadRequest {
	return &PostOrganisationNameverificationsBadRequest{}
}

/*
PostOrganisationNameverificationsBadRequest handles this case with default header values.

Bad Request
*/
type PostOrganisationNameverificationsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this post organisation nameverifications bad request response has a 2xx status code
func (o *PostOrganisationNameverificationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post organisation nameverifications bad request response has a 3xx status code
func (o *PostOrganisationNameverificationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post organisation nameverifications bad request response has a 4xx status code
func (o *PostOrganisationNameverificationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post organisation nameverifications bad request response has a 5xx status code
func (o *PostOrganisationNameverificationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post organisation nameverifications bad request response a status code equal to that given
func (o *PostOrganisationNameverificationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post organisation nameverifications bad request response
func (o *PostOrganisationNameverificationsBadRequest) Code() int {
	return 400
}

func (o *PostOrganisationNameverificationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /organisation/nameverifications][%d] postOrganisationNameverificationsBadRequest", 400)
}

func (o *PostOrganisationNameverificationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrganisationNameverificationsForbidden creates a PostOrganisationNameverificationsForbidden with default headers values
func NewPostOrganisationNameverificationsForbidden() *PostOrganisationNameverificationsForbidden {
	return &PostOrganisationNameverificationsForbidden{}
}

/*
PostOrganisationNameverificationsForbidden handles this case with default header values.

Forbidden
*/
type PostOrganisationNameverificationsForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this post organisation nameverifications forbidden response has a 2xx status code
func (o *PostOrganisationNameverificationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post organisation nameverifications forbidden response has a 3xx status code
func (o *PostOrganisationNameverificationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post organisation nameverifications forbidden response has a 4xx status code
func (o *PostOrganisationNameverificationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post organisation nameverifications forbidden response has a 5xx status code
func (o *PostOrganisationNameverificationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post organisation nameverifications forbidden response a status code equal to that given
func (o *PostOrganisationNameverificationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post organisation nameverifications forbidden response
func (o *PostOrganisationNameverificationsForbidden) Code() int {
	return 403
}

func (o *PostOrganisationNameverificationsForbidden) Error() string {
	return fmt.Sprintf("[POST /organisation/nameverifications][%d] postOrganisationNameverificationsForbidden", 403)
}

func (o *PostOrganisationNameverificationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrganisationNameverificationsConflict creates a PostOrganisationNameverificationsConflict with default headers values
func NewPostOrganisationNameverificationsConflict() *PostOrganisationNameverificationsConflict {
	return &PostOrganisationNameverificationsConflict{}
}

/*
PostOrganisationNameverificationsConflict handles this case with default header values.

Conflict
*/
type PostOrganisationNameverificationsConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this post organisation nameverifications conflict response has a 2xx status code
func (o *PostOrganisationNameverificationsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post organisation nameverifications conflict response has a 3xx status code
func (o *PostOrganisationNameverificationsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post organisation nameverifications conflict response has a 4xx status code
func (o *PostOrganisationNameverificationsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post organisation nameverifications conflict response has a 5xx status code
func (o *PostOrganisationNameverificationsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post organisation nameverifications conflict response a status code equal to that given
func (o *PostOrganisationNameverificationsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the post organisation nameverifications conflict response
func (o *PostOrganisationNameverificationsConflict) Code() int {
	return 409
}

func (o *PostOrganisationNameverificationsConflict) Error() string {
	return fmt.Sprintf("[POST /organisation/nameverifications][%d] postOrganisationNameverificationsConflict", 409)
}

func (o *PostOrganisationNameverificationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrganisationNameverificationsInternalServerError creates a PostOrganisationNameverificationsInternalServerError with default headers values
func NewPostOrganisationNameverificationsInternalServerError() *PostOrganisationNameverificationsInternalServerError {
	return &PostOrganisationNameverificationsInternalServerError{}
}

/*
PostOrganisationNameverificationsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostOrganisationNameverificationsInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this post organisation nameverifications internal server error response has a 2xx status code
func (o *PostOrganisationNameverificationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post organisation nameverifications internal server error response has a 3xx status code
func (o *PostOrganisationNameverificationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post organisation nameverifications internal server error response has a 4xx status code
func (o *PostOrganisationNameverificationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post organisation nameverifications internal server error response has a 5xx status code
func (o *PostOrganisationNameverificationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post organisation nameverifications internal server error response a status code equal to that given
func (o *PostOrganisationNameverificationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post organisation nameverifications internal server error response
func (o *PostOrganisationNameverificationsInternalServerError) Code() int {
	return 500
}

func (o *PostOrganisationNameverificationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /organisation/nameverifications][%d] postOrganisationNameverificationsInternalServerError", 500)
}

func (o *PostOrganisationNameverificationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
