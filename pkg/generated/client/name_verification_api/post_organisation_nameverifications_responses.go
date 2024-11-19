// Code generated by go-swagger; DO NOT EDIT.

package name_verification_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
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
