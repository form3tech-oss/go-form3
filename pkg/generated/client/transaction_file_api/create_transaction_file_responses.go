// Code generated by go-swagger; DO NOT EDIT.

package transaction_file_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// CreateTransactionFileReader is a Reader for the CreateTransactionFile structure.
type CreateTransactionFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTransactionFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateTransactionFileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateTransactionFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateTransactionFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateTransactionFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateTransactionFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateTransactionFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateTransactionFileCreated creates a CreateTransactionFileCreated with default headers values
func NewCreateTransactionFileCreated() *CreateTransactionFileCreated {
	return &CreateTransactionFileCreated{}
}

/*
CreateTransactionFileCreated handles this case with default header values.

Transaction File Creation Response
*/
type CreateTransactionFileCreated struct {

	//Payload

	// isStream: false
	*models.TransactionFileResponse
}

func (o *CreateTransactionFileCreated) Error() string {
	return fmt.Sprintf("[POST /files/transactions][%d] createTransactionFileCreated", 201)
}

func (o *CreateTransactionFileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.TransactionFileResponse = new(models.TransactionFileResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.TransactionFileResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionFileBadRequest creates a CreateTransactionFileBadRequest with default headers values
func NewCreateTransactionFileBadRequest() *CreateTransactionFileBadRequest {
	return &CreateTransactionFileBadRequest{}
}

/*
CreateTransactionFileBadRequest handles this case with default header values.

Bad Request
*/
type CreateTransactionFileBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateTransactionFileBadRequest) Error() string {
	return fmt.Sprintf("[POST /files/transactions][%d] createTransactionFileBadRequest", 400)
}

func (o *CreateTransactionFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionFileUnauthorized creates a CreateTransactionFileUnauthorized with default headers values
func NewCreateTransactionFileUnauthorized() *CreateTransactionFileUnauthorized {
	return &CreateTransactionFileUnauthorized{}
}

/*
CreateTransactionFileUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateTransactionFileUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateTransactionFileUnauthorized) Error() string {
	return fmt.Sprintf("[POST /files/transactions][%d] createTransactionFileUnauthorized", 401)
}

func (o *CreateTransactionFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionFileForbidden creates a CreateTransactionFileForbidden with default headers values
func NewCreateTransactionFileForbidden() *CreateTransactionFileForbidden {
	return &CreateTransactionFileForbidden{}
}

/*
CreateTransactionFileForbidden handles this case with default header values.

Forbidden
*/
type CreateTransactionFileForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateTransactionFileForbidden) Error() string {
	return fmt.Sprintf("[POST /files/transactions][%d] createTransactionFileForbidden", 403)
}

func (o *CreateTransactionFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionFileConflict creates a CreateTransactionFileConflict with default headers values
func NewCreateTransactionFileConflict() *CreateTransactionFileConflict {
	return &CreateTransactionFileConflict{}
}

/*
CreateTransactionFileConflict handles this case with default header values.

Conflict
*/
type CreateTransactionFileConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateTransactionFileConflict) Error() string {
	return fmt.Sprintf("[POST /files/transactions][%d] createTransactionFileConflict", 409)
}

func (o *CreateTransactionFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionFileInternalServerError creates a CreateTransactionFileInternalServerError with default headers values
func NewCreateTransactionFileInternalServerError() *CreateTransactionFileInternalServerError {
	return &CreateTransactionFileInternalServerError{}
}

/*
CreateTransactionFileInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateTransactionFileInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateTransactionFileInternalServerError) Error() string {
	return fmt.Sprintf("[POST /files/transactions][%d] createTransactionFileInternalServerError", 500)
}

func (o *CreateTransactionFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
