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

// GetSchemeFileSubmissionReader is a Reader for the GetSchemeFileSubmission structure.
type GetSchemeFileSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchemeFileSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSchemeFileSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSchemeFileSubmissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSchemeFileSubmissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSchemeFileSubmissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSchemeFileSubmissionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSchemeFileSubmissionOK creates a GetSchemeFileSubmissionOK with default headers values
func NewGetSchemeFileSubmissionOK() *GetSchemeFileSubmissionOK {
	return &GetSchemeFileSubmissionOK{}
}

/*GetSchemeFileSubmissionOK handles this case with default header values.

Scheme File Submission Response
*/
type GetSchemeFileSubmissionOK struct {

	//Payload

	// isStream: false
	*models.SchemeFileSubmissionResponse
}

func (o *GetSchemeFileSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /files/schemefiles/{scheme_file_id}/submissions/{scheme_file_submission_id}][%d] getSchemeFileSubmissionOK", 200)
}

func (o *GetSchemeFileSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SchemeFileSubmissionResponse = new(models.SchemeFileSubmissionResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SchemeFileSubmissionResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemeFileSubmissionUnauthorized creates a GetSchemeFileSubmissionUnauthorized with default headers values
func NewGetSchemeFileSubmissionUnauthorized() *GetSchemeFileSubmissionUnauthorized {
	return &GetSchemeFileSubmissionUnauthorized{}
}

/*GetSchemeFileSubmissionUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSchemeFileSubmissionUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetSchemeFileSubmissionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /files/schemefiles/{scheme_file_id}/submissions/{scheme_file_submission_id}][%d] getSchemeFileSubmissionUnauthorized", 401)
}

func (o *GetSchemeFileSubmissionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemeFileSubmissionForbidden creates a GetSchemeFileSubmissionForbidden with default headers values
func NewGetSchemeFileSubmissionForbidden() *GetSchemeFileSubmissionForbidden {
	return &GetSchemeFileSubmissionForbidden{}
}

/*GetSchemeFileSubmissionForbidden handles this case with default header values.

Forbidden
*/
type GetSchemeFileSubmissionForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetSchemeFileSubmissionForbidden) Error() string {
	return fmt.Sprintf("[GET /files/schemefiles/{scheme_file_id}/submissions/{scheme_file_submission_id}][%d] getSchemeFileSubmissionForbidden", 403)
}

func (o *GetSchemeFileSubmissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemeFileSubmissionNotFound creates a GetSchemeFileSubmissionNotFound with default headers values
func NewGetSchemeFileSubmissionNotFound() *GetSchemeFileSubmissionNotFound {
	return &GetSchemeFileSubmissionNotFound{}
}

/*GetSchemeFileSubmissionNotFound handles this case with default header values.

Not Found
*/
type GetSchemeFileSubmissionNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetSchemeFileSubmissionNotFound) Error() string {
	return fmt.Sprintf("[GET /files/schemefiles/{scheme_file_id}/submissions/{scheme_file_submission_id}][%d] getSchemeFileSubmissionNotFound", 404)
}

func (o *GetSchemeFileSubmissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemeFileSubmissionInternalServerError creates a GetSchemeFileSubmissionInternalServerError with default headers values
func NewGetSchemeFileSubmissionInternalServerError() *GetSchemeFileSubmissionInternalServerError {
	return &GetSchemeFileSubmissionInternalServerError{}
}

/*GetSchemeFileSubmissionInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSchemeFileSubmissionInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetSchemeFileSubmissionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /files/schemefiles/{scheme_file_id}/submissions/{scheme_file_submission_id}][%d] getSchemeFileSubmissionInternalServerError", 500)
}

func (o *GetSchemeFileSubmissionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
