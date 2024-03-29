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

// GetOrganisationNameverificationsIDReader is a Reader for the GetOrganisationNameverificationsID structure.
type GetOrganisationNameverificationsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganisationNameverificationsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganisationNameverificationsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOrganisationNameverificationsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetOrganisationNameverificationsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrganisationNameverificationsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetOrganisationNameverificationsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganisationNameverificationsIDOK creates a GetOrganisationNameverificationsIDOK with default headers values
func NewGetOrganisationNameverificationsIDOK() *GetOrganisationNameverificationsIDOK {
	return &GetOrganisationNameverificationsIDOK{}
}

/*GetOrganisationNameverificationsIDOK handles this case with default header values.

Name verification details
*/
type GetOrganisationNameverificationsIDOK struct {

	//Payload

	// isStream: false
	*models.NameVerificationDetailsResponse
}

func (o *GetOrganisationNameverificationsIDOK) Error() string {
	return fmt.Sprintf("[GET /organisation/nameverifications/{id}][%d] getOrganisationNameverificationsIdOK", 200)
}

func (o *GetOrganisationNameverificationsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.NameVerificationDetailsResponse = new(models.NameVerificationDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.NameVerificationDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationNameverificationsIDBadRequest creates a GetOrganisationNameverificationsIDBadRequest with default headers values
func NewGetOrganisationNameverificationsIDBadRequest() *GetOrganisationNameverificationsIDBadRequest {
	return &GetOrganisationNameverificationsIDBadRequest{}
}

/*GetOrganisationNameverificationsIDBadRequest handles this case with default header values.

Bad Request
*/
type GetOrganisationNameverificationsIDBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetOrganisationNameverificationsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /organisation/nameverifications/{id}][%d] getOrganisationNameverificationsIdBadRequest", 400)
}

func (o *GetOrganisationNameverificationsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationNameverificationsIDForbidden creates a GetOrganisationNameverificationsIDForbidden with default headers values
func NewGetOrganisationNameverificationsIDForbidden() *GetOrganisationNameverificationsIDForbidden {
	return &GetOrganisationNameverificationsIDForbidden{}
}

/*GetOrganisationNameverificationsIDForbidden handles this case with default header values.

Forbidden
*/
type GetOrganisationNameverificationsIDForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetOrganisationNameverificationsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /organisation/nameverifications/{id}][%d] getOrganisationNameverificationsIdForbidden", 403)
}

func (o *GetOrganisationNameverificationsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationNameverificationsIDNotFound creates a GetOrganisationNameverificationsIDNotFound with default headers values
func NewGetOrganisationNameverificationsIDNotFound() *GetOrganisationNameverificationsIDNotFound {
	return &GetOrganisationNameverificationsIDNotFound{}
}

/*GetOrganisationNameverificationsIDNotFound handles this case with default header values.

Not Found
*/
type GetOrganisationNameverificationsIDNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetOrganisationNameverificationsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /organisation/nameverifications/{id}][%d] getOrganisationNameverificationsIdNotFound", 404)
}

func (o *GetOrganisationNameverificationsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationNameverificationsIDInternalServerError creates a GetOrganisationNameverificationsIDInternalServerError with default headers values
func NewGetOrganisationNameverificationsIDInternalServerError() *GetOrganisationNameverificationsIDInternalServerError {
	return &GetOrganisationNameverificationsIDInternalServerError{}
}

/*GetOrganisationNameverificationsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetOrganisationNameverificationsIDInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetOrganisationNameverificationsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organisation/nameverifications/{id}][%d] getOrganisationNameverificationsIdInternalServerError", 500)
}

func (o *GetOrganisationNameverificationsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
