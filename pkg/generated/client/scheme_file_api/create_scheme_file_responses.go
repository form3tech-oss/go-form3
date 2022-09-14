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

/*CreateSchemeFileCreated handles this case with default header values.

Scheme File Creation Response
*/
type CreateSchemeFileCreated struct {

	//Payload

	// isStream: false
	*models.SchemeFileResponse
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

/*CreateSchemeFileBadRequest handles this case with default header values.

Bad Request
*/
type CreateSchemeFileBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
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

/*CreateSchemeFileUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateSchemeFileUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
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

/*CreateSchemeFileForbidden handles this case with default header values.

Forbidden
*/
type CreateSchemeFileForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
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

/*CreateSchemeFileConflict handles this case with default header values.

Conflict
*/
type CreateSchemeFileConflict struct {

	//Payload

	// isStream: false
	*models.APIError
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

/*CreateSchemeFileInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateSchemeFileInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
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
