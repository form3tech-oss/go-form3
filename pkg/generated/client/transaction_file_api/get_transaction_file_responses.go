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

// GetTransactionFileReader is a Reader for the GetTransactionFile structure.
type GetTransactionFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetTransactionFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTransactionFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTransactionFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetTransactionFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionFileOK creates a GetTransactionFileOK with default headers values
func NewGetTransactionFileOK() *GetTransactionFileOK {
	return &GetTransactionFileOK{}
}

/*GetTransactionFileOK handles this case with default header values.

Transaction File Response
*/
type GetTransactionFileOK struct {

	//Payload

	// isStream: false
	*models.TransactionFileResponse
}

func (o *GetTransactionFileOK) Error() string {
	return fmt.Sprintf("[GET /files/transactions/{transaction_file_id}][%d] getTransactionFileOK", 200)
}

func (o *GetTransactionFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.TransactionFileResponse = new(models.TransactionFileResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.TransactionFileResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFileUnauthorized creates a GetTransactionFileUnauthorized with default headers values
func NewGetTransactionFileUnauthorized() *GetTransactionFileUnauthorized {
	return &GetTransactionFileUnauthorized{}
}

/*GetTransactionFileUnauthorized handles this case with default header values.

Unauthorized
*/
type GetTransactionFileUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /files/transactions/{transaction_file_id}][%d] getTransactionFileUnauthorized", 401)
}

func (o *GetTransactionFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFileForbidden creates a GetTransactionFileForbidden with default headers values
func NewGetTransactionFileForbidden() *GetTransactionFileForbidden {
	return &GetTransactionFileForbidden{}
}

/*GetTransactionFileForbidden handles this case with default header values.

Forbidden
*/
type GetTransactionFileForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFileForbidden) Error() string {
	return fmt.Sprintf("[GET /files/transactions/{transaction_file_id}][%d] getTransactionFileForbidden", 403)
}

func (o *GetTransactionFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFileNotFound creates a GetTransactionFileNotFound with default headers values
func NewGetTransactionFileNotFound() *GetTransactionFileNotFound {
	return &GetTransactionFileNotFound{}
}

/*GetTransactionFileNotFound handles this case with default header values.

Not Found
*/
type GetTransactionFileNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFileNotFound) Error() string {
	return fmt.Sprintf("[GET /files/transactions/{transaction_file_id}][%d] getTransactionFileNotFound", 404)
}

func (o *GetTransactionFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFileInternalServerError creates a GetTransactionFileInternalServerError with default headers values
func NewGetTransactionFileInternalServerError() *GetTransactionFileInternalServerError {
	return &GetTransactionFileInternalServerError{}
}

/*GetTransactionFileInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetTransactionFileInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /files/transactions/{transaction_file_id}][%d] getTransactionFileInternalServerError", 500)
}

func (o *GetTransactionFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
