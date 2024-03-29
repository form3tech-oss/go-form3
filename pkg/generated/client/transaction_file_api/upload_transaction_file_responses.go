// Code generated by go-swagger; DO NOT EDIT.

package transaction_file_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// UploadTransactionFileReader is a Reader for the UploadTransactionFile structure.
type UploadTransactionFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadTransactionFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUploadTransactionFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUploadTransactionFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUploadTransactionFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUploadTransactionFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUploadTransactionFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUploadTransactionFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUploadTransactionFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadTransactionFileOK creates a UploadTransactionFileOK with default headers values
func NewUploadTransactionFileOK() *UploadTransactionFileOK {
	return &UploadTransactionFileOK{}
}

/*UploadTransactionFileOK handles this case with default header values.

Transaction File Response
*/
type UploadTransactionFileOK struct {

	//Payload

	// isStream: false
	*models.TransactionFileResponse
}

func (o *UploadTransactionFileOK) Error() string {
	return fmt.Sprintf("[PUT /files/transactions/{transaction_file_id}][%d] uploadTransactionFileOK", 200)
}

func (o *UploadTransactionFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.TransactionFileResponse = new(models.TransactionFileResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.TransactionFileResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadTransactionFileBadRequest creates a UploadTransactionFileBadRequest with default headers values
func NewUploadTransactionFileBadRequest() *UploadTransactionFileBadRequest {
	return &UploadTransactionFileBadRequest{}
}

/*UploadTransactionFileBadRequest handles this case with default header values.

Bad Request
*/
type UploadTransactionFileBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *UploadTransactionFileBadRequest) Error() string {
	return fmt.Sprintf("[PUT /files/transactions/{transaction_file_id}][%d] uploadTransactionFileBadRequest", 400)
}

func (o *UploadTransactionFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadTransactionFileUnauthorized creates a UploadTransactionFileUnauthorized with default headers values
func NewUploadTransactionFileUnauthorized() *UploadTransactionFileUnauthorized {
	return &UploadTransactionFileUnauthorized{}
}

/*UploadTransactionFileUnauthorized handles this case with default header values.

Unauthorized
*/
type UploadTransactionFileUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *UploadTransactionFileUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /files/transactions/{transaction_file_id}][%d] uploadTransactionFileUnauthorized", 401)
}

func (o *UploadTransactionFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadTransactionFileForbidden creates a UploadTransactionFileForbidden with default headers values
func NewUploadTransactionFileForbidden() *UploadTransactionFileForbidden {
	return &UploadTransactionFileForbidden{}
}

/*UploadTransactionFileForbidden handles this case with default header values.

Forbidden
*/
type UploadTransactionFileForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *UploadTransactionFileForbidden) Error() string {
	return fmt.Sprintf("[PUT /files/transactions/{transaction_file_id}][%d] uploadTransactionFileForbidden", 403)
}

func (o *UploadTransactionFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadTransactionFileNotFound creates a UploadTransactionFileNotFound with default headers values
func NewUploadTransactionFileNotFound() *UploadTransactionFileNotFound {
	return &UploadTransactionFileNotFound{}
}

/*UploadTransactionFileNotFound handles this case with default header values.

Transaction File Not Found
*/
type UploadTransactionFileNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *UploadTransactionFileNotFound) Error() string {
	return fmt.Sprintf("[PUT /files/transactions/{transaction_file_id}][%d] uploadTransactionFileNotFound", 404)
}

func (o *UploadTransactionFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadTransactionFileConflict creates a UploadTransactionFileConflict with default headers values
func NewUploadTransactionFileConflict() *UploadTransactionFileConflict {
	return &UploadTransactionFileConflict{}
}

/*UploadTransactionFileConflict handles this case with default header values.

Transaction File Conflict
*/
type UploadTransactionFileConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *UploadTransactionFileConflict) Error() string {
	return fmt.Sprintf("[PUT /files/transactions/{transaction_file_id}][%d] uploadTransactionFileConflict", 409)
}

func (o *UploadTransactionFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadTransactionFileInternalServerError creates a UploadTransactionFileInternalServerError with default headers values
func NewUploadTransactionFileInternalServerError() *UploadTransactionFileInternalServerError {
	return &UploadTransactionFileInternalServerError{}
}

/*UploadTransactionFileInternalServerError handles this case with default header values.

Internal Server Error
*/
type UploadTransactionFileInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *UploadTransactionFileInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /files/transactions/{transaction_file_id}][%d] uploadTransactionFileInternalServerError", 500)
}

func (o *UploadTransactionFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
