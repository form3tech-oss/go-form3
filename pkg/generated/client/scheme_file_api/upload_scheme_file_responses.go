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

// UploadSchemeFileReader is a Reader for the UploadSchemeFile structure.
type UploadSchemeFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadSchemeFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUploadSchemeFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUploadSchemeFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUploadSchemeFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUploadSchemeFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUploadSchemeFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUploadSchemeFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUploadSchemeFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadSchemeFileOK creates a UploadSchemeFileOK with default headers values
func NewUploadSchemeFileOK() *UploadSchemeFileOK {
	return &UploadSchemeFileOK{}
}

/*
UploadSchemeFileOK handles this case with default header values.

Scheme File Response
*/
type UploadSchemeFileOK struct {

	//Payload

	// isStream: false
	*models.SchemeFileResponse
}

// IsSuccess returns true when this upload scheme file o k response has a 2xx status code
func (o *UploadSchemeFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload scheme file o k response has a 3xx status code
func (o *UploadSchemeFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload scheme file o k response has a 4xx status code
func (o *UploadSchemeFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload scheme file o k response has a 5xx status code
func (o *UploadSchemeFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload scheme file o k response a status code equal to that given
func (o *UploadSchemeFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upload scheme file o k response
func (o *UploadSchemeFileOK) Code() int {
	return 200
}

func (o *UploadSchemeFileOK) Error() string {
	return fmt.Sprintf("[PUT /files/schemefiles/{scheme_file_id}][%d] uploadSchemeFileOK", 200)
}

func (o *UploadSchemeFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SchemeFileResponse = new(models.SchemeFileResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SchemeFileResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSchemeFileBadRequest creates a UploadSchemeFileBadRequest with default headers values
func NewUploadSchemeFileBadRequest() *UploadSchemeFileBadRequest {
	return &UploadSchemeFileBadRequest{}
}

/*
UploadSchemeFileBadRequest handles this case with default header values.

Bad Request
*/
type UploadSchemeFileBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this upload scheme file bad request response has a 2xx status code
func (o *UploadSchemeFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload scheme file bad request response has a 3xx status code
func (o *UploadSchemeFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload scheme file bad request response has a 4xx status code
func (o *UploadSchemeFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload scheme file bad request response has a 5xx status code
func (o *UploadSchemeFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload scheme file bad request response a status code equal to that given
func (o *UploadSchemeFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the upload scheme file bad request response
func (o *UploadSchemeFileBadRequest) Code() int {
	return 400
}

func (o *UploadSchemeFileBadRequest) Error() string {
	return fmt.Sprintf("[PUT /files/schemefiles/{scheme_file_id}][%d] uploadSchemeFileBadRequest", 400)
}

func (o *UploadSchemeFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSchemeFileUnauthorized creates a UploadSchemeFileUnauthorized with default headers values
func NewUploadSchemeFileUnauthorized() *UploadSchemeFileUnauthorized {
	return &UploadSchemeFileUnauthorized{}
}

/*
UploadSchemeFileUnauthorized handles this case with default header values.

Unauthorized
*/
type UploadSchemeFileUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this upload scheme file unauthorized response has a 2xx status code
func (o *UploadSchemeFileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload scheme file unauthorized response has a 3xx status code
func (o *UploadSchemeFileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload scheme file unauthorized response has a 4xx status code
func (o *UploadSchemeFileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload scheme file unauthorized response has a 5xx status code
func (o *UploadSchemeFileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this upload scheme file unauthorized response a status code equal to that given
func (o *UploadSchemeFileUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the upload scheme file unauthorized response
func (o *UploadSchemeFileUnauthorized) Code() int {
	return 401
}

func (o *UploadSchemeFileUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /files/schemefiles/{scheme_file_id}][%d] uploadSchemeFileUnauthorized", 401)
}

func (o *UploadSchemeFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSchemeFileForbidden creates a UploadSchemeFileForbidden with default headers values
func NewUploadSchemeFileForbidden() *UploadSchemeFileForbidden {
	return &UploadSchemeFileForbidden{}
}

/*
UploadSchemeFileForbidden handles this case with default header values.

Forbidden
*/
type UploadSchemeFileForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this upload scheme file forbidden response has a 2xx status code
func (o *UploadSchemeFileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload scheme file forbidden response has a 3xx status code
func (o *UploadSchemeFileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload scheme file forbidden response has a 4xx status code
func (o *UploadSchemeFileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload scheme file forbidden response has a 5xx status code
func (o *UploadSchemeFileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this upload scheme file forbidden response a status code equal to that given
func (o *UploadSchemeFileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the upload scheme file forbidden response
func (o *UploadSchemeFileForbidden) Code() int {
	return 403
}

func (o *UploadSchemeFileForbidden) Error() string {
	return fmt.Sprintf("[PUT /files/schemefiles/{scheme_file_id}][%d] uploadSchemeFileForbidden", 403)
}

func (o *UploadSchemeFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSchemeFileNotFound creates a UploadSchemeFileNotFound with default headers values
func NewUploadSchemeFileNotFound() *UploadSchemeFileNotFound {
	return &UploadSchemeFileNotFound{}
}

/*
UploadSchemeFileNotFound handles this case with default header values.

Scheme File Not Found
*/
type UploadSchemeFileNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this upload scheme file not found response has a 2xx status code
func (o *UploadSchemeFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload scheme file not found response has a 3xx status code
func (o *UploadSchemeFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload scheme file not found response has a 4xx status code
func (o *UploadSchemeFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload scheme file not found response has a 5xx status code
func (o *UploadSchemeFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this upload scheme file not found response a status code equal to that given
func (o *UploadSchemeFileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the upload scheme file not found response
func (o *UploadSchemeFileNotFound) Code() int {
	return 404
}

func (o *UploadSchemeFileNotFound) Error() string {
	return fmt.Sprintf("[PUT /files/schemefiles/{scheme_file_id}][%d] uploadSchemeFileNotFound", 404)
}

func (o *UploadSchemeFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSchemeFileConflict creates a UploadSchemeFileConflict with default headers values
func NewUploadSchemeFileConflict() *UploadSchemeFileConflict {
	return &UploadSchemeFileConflict{}
}

/*
UploadSchemeFileConflict handles this case with default header values.

Scheme File Conflict
*/
type UploadSchemeFileConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this upload scheme file conflict response has a 2xx status code
func (o *UploadSchemeFileConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload scheme file conflict response has a 3xx status code
func (o *UploadSchemeFileConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload scheme file conflict response has a 4xx status code
func (o *UploadSchemeFileConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload scheme file conflict response has a 5xx status code
func (o *UploadSchemeFileConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this upload scheme file conflict response a status code equal to that given
func (o *UploadSchemeFileConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the upload scheme file conflict response
func (o *UploadSchemeFileConflict) Code() int {
	return 409
}

func (o *UploadSchemeFileConflict) Error() string {
	return fmt.Sprintf("[PUT /files/schemefiles/{scheme_file_id}][%d] uploadSchemeFileConflict", 409)
}

func (o *UploadSchemeFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSchemeFileInternalServerError creates a UploadSchemeFileInternalServerError with default headers values
func NewUploadSchemeFileInternalServerError() *UploadSchemeFileInternalServerError {
	return &UploadSchemeFileInternalServerError{}
}

/*
UploadSchemeFileInternalServerError handles this case with default header values.

Internal Server Error
*/
type UploadSchemeFileInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this upload scheme file internal server error response has a 2xx status code
func (o *UploadSchemeFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload scheme file internal server error response has a 3xx status code
func (o *UploadSchemeFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload scheme file internal server error response has a 4xx status code
func (o *UploadSchemeFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload scheme file internal server error response has a 5xx status code
func (o *UploadSchemeFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this upload scheme file internal server error response a status code equal to that given
func (o *UploadSchemeFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the upload scheme file internal server error response
func (o *UploadSchemeFileInternalServerError) Code() int {
	return 500
}

func (o *UploadSchemeFileInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /files/schemefiles/{scheme_file_id}][%d] uploadSchemeFileInternalServerError", 500)
}

func (o *UploadSchemeFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
