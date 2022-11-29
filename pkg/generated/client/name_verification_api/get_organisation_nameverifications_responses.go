// Code generated by go-swagger; DO NOT EDIT.

package name_verification_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// GetOrganisationNameverificationsReader is a Reader for the GetOrganisationNameverifications structure.
type GetOrganisationNameverificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganisationNameverificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganisationNameverificationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOrganisationNameverificationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetOrganisationNameverificationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrganisationNameverificationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetOrganisationNameverificationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganisationNameverificationsOK creates a GetOrganisationNameverificationsOK with default headers values
func NewGetOrganisationNameverificationsOK() *GetOrganisationNameverificationsOK {
	return &GetOrganisationNameverificationsOK{}
}

/*GetOrganisationNameverificationsOK handles this case with default header values.

List of name verification details
*/
type GetOrganisationNameverificationsOK struct {

	//Payload

	// isStream: false
	*models.NameVerificationDetailsListResponse
}

func (o *GetOrganisationNameverificationsOK) Error() string {
	return fmt.Sprintf("[GET /organisation/nameverifications][%d] getOrganisationNameverificationsOK", 200)
}

func (o *GetOrganisationNameverificationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.NameVerificationDetailsListResponse = new(models.NameVerificationDetailsListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.NameVerificationDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationNameverificationsBadRequest creates a GetOrganisationNameverificationsBadRequest with default headers values
func NewGetOrganisationNameverificationsBadRequest() *GetOrganisationNameverificationsBadRequest {
	return &GetOrganisationNameverificationsBadRequest{}
}

/*GetOrganisationNameverificationsBadRequest handles this case with default header values.

Bad Request
*/
type GetOrganisationNameverificationsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetOrganisationNameverificationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /organisation/nameverifications][%d] getOrganisationNameverificationsBadRequest", 400)
}

func (o *GetOrganisationNameverificationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationNameverificationsForbidden creates a GetOrganisationNameverificationsForbidden with default headers values
func NewGetOrganisationNameverificationsForbidden() *GetOrganisationNameverificationsForbidden {
	return &GetOrganisationNameverificationsForbidden{}
}

/*GetOrganisationNameverificationsForbidden handles this case with default header values.

Forbidden
*/
type GetOrganisationNameverificationsForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetOrganisationNameverificationsForbidden) Error() string {
	return fmt.Sprintf("[GET /organisation/nameverifications][%d] getOrganisationNameverificationsForbidden", 403)
}

func (o *GetOrganisationNameverificationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationNameverificationsNotFound creates a GetOrganisationNameverificationsNotFound with default headers values
func NewGetOrganisationNameverificationsNotFound() *GetOrganisationNameverificationsNotFound {
	return &GetOrganisationNameverificationsNotFound{}
}

/*GetOrganisationNameverificationsNotFound handles this case with default header values.

Not Found
*/
type GetOrganisationNameverificationsNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetOrganisationNameverificationsNotFound) Error() string {
	return fmt.Sprintf("[GET /organisation/nameverifications][%d] getOrganisationNameverificationsNotFound", 404)
}

func (o *GetOrganisationNameverificationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationNameverificationsInternalServerError creates a GetOrganisationNameverificationsInternalServerError with default headers values
func NewGetOrganisationNameverificationsInternalServerError() *GetOrganisationNameverificationsInternalServerError {
	return &GetOrganisationNameverificationsInternalServerError{}
}

/*GetOrganisationNameverificationsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetOrganisationNameverificationsInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetOrganisationNameverificationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organisation/nameverifications][%d] getOrganisationNameverificationsInternalServerError", 500)
}

func (o *GetOrganisationNameverificationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
