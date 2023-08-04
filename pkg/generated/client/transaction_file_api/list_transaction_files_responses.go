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

// ListTransactionFilesReader is a Reader for the ListTransactionFiles structure.
type ListTransactionFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTransactionFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListTransactionFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListTransactionFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListTransactionFilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListTransactionFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListTransactionFilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListTransactionFilesOK creates a ListTransactionFilesOK with default headers values
func NewListTransactionFilesOK() *ListTransactionFilesOK {
	return &ListTransactionFilesOK{}
}

/*ListTransactionFilesOK handles this case with default header values.

List of transaction files
*/
type ListTransactionFilesOK struct {

	//Payload

	// isStream: false
	*models.ListTransactionFilesResponse
}

func (o *ListTransactionFilesOK) Error() string {
	return fmt.Sprintf("[GET /files/transactions][%d] listTransactionFilesOK", 200)
}

func (o *ListTransactionFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ListTransactionFilesResponse = new(models.ListTransactionFilesResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ListTransactionFilesResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransactionFilesBadRequest creates a ListTransactionFilesBadRequest with default headers values
func NewListTransactionFilesBadRequest() *ListTransactionFilesBadRequest {
	return &ListTransactionFilesBadRequest{}
}

/*ListTransactionFilesBadRequest handles this case with default header values.

Reports bad request
*/
type ListTransactionFilesBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListTransactionFilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /files/transactions][%d] listTransactionFilesBadRequest", 400)
}

func (o *ListTransactionFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransactionFilesUnauthorized creates a ListTransactionFilesUnauthorized with default headers values
func NewListTransactionFilesUnauthorized() *ListTransactionFilesUnauthorized {
	return &ListTransactionFilesUnauthorized{}
}

/*ListTransactionFilesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListTransactionFilesUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListTransactionFilesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /files/transactions][%d] listTransactionFilesUnauthorized", 401)
}

func (o *ListTransactionFilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransactionFilesForbidden creates a ListTransactionFilesForbidden with default headers values
func NewListTransactionFilesForbidden() *ListTransactionFilesForbidden {
	return &ListTransactionFilesForbidden{}
}

/*ListTransactionFilesForbidden handles this case with default header values.

Forbidden
*/
type ListTransactionFilesForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListTransactionFilesForbidden) Error() string {
	return fmt.Sprintf("[GET /files/transactions][%d] listTransactionFilesForbidden", 403)
}

func (o *ListTransactionFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransactionFilesInternalServerError creates a ListTransactionFilesInternalServerError with default headers values
func NewListTransactionFilesInternalServerError() *ListTransactionFilesInternalServerError {
	return &ListTransactionFilesInternalServerError{}
}

/*ListTransactionFilesInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListTransactionFilesInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListTransactionFilesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /files/transactions][%d] listTransactionFilesInternalServerError", 500)
}

func (o *ListTransactionFilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}