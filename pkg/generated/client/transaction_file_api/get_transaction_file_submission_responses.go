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

// GetTransactionFileSubmissionReader is a Reader for the GetTransactionFileSubmission structure.
type GetTransactionFileSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionFileSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionFileSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetTransactionFileSubmissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTransactionFileSubmissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTransactionFileSubmissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetTransactionFileSubmissionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionFileSubmissionOK creates a GetTransactionFileSubmissionOK with default headers values
func NewGetTransactionFileSubmissionOK() *GetTransactionFileSubmissionOK {
	return &GetTransactionFileSubmissionOK{}
}

/*
GetTransactionFileSubmissionOK handles this case with default header values.

Transaction File Submission Response
*/
type GetTransactionFileSubmissionOK struct {

	//Payload

	// isStream: false
	*models.TransactionFileSubmissionResponse
}

func (o *GetTransactionFileSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /files/transactions/{transaction_file_id}/submissions/{transaction_file_submission_id}][%d] getTransactionFileSubmissionOK", 200)
}

func (o *GetTransactionFileSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.TransactionFileSubmissionResponse = new(models.TransactionFileSubmissionResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.TransactionFileSubmissionResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFileSubmissionUnauthorized creates a GetTransactionFileSubmissionUnauthorized with default headers values
func NewGetTransactionFileSubmissionUnauthorized() *GetTransactionFileSubmissionUnauthorized {
	return &GetTransactionFileSubmissionUnauthorized{}
}

/*
GetTransactionFileSubmissionUnauthorized handles this case with default header values.

Unauthorized
*/
type GetTransactionFileSubmissionUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFileSubmissionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /files/transactions/{transaction_file_id}/submissions/{transaction_file_submission_id}][%d] getTransactionFileSubmissionUnauthorized", 401)
}

func (o *GetTransactionFileSubmissionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFileSubmissionForbidden creates a GetTransactionFileSubmissionForbidden with default headers values
func NewGetTransactionFileSubmissionForbidden() *GetTransactionFileSubmissionForbidden {
	return &GetTransactionFileSubmissionForbidden{}
}

/*
GetTransactionFileSubmissionForbidden handles this case with default header values.

Forbidden
*/
type GetTransactionFileSubmissionForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFileSubmissionForbidden) Error() string {
	return fmt.Sprintf("[GET /files/transactions/{transaction_file_id}/submissions/{transaction_file_submission_id}][%d] getTransactionFileSubmissionForbidden", 403)
}

func (o *GetTransactionFileSubmissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFileSubmissionNotFound creates a GetTransactionFileSubmissionNotFound with default headers values
func NewGetTransactionFileSubmissionNotFound() *GetTransactionFileSubmissionNotFound {
	return &GetTransactionFileSubmissionNotFound{}
}

/*
GetTransactionFileSubmissionNotFound handles this case with default header values.

Not Found
*/
type GetTransactionFileSubmissionNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFileSubmissionNotFound) Error() string {
	return fmt.Sprintf("[GET /files/transactions/{transaction_file_id}/submissions/{transaction_file_submission_id}][%d] getTransactionFileSubmissionNotFound", 404)
}

func (o *GetTransactionFileSubmissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionFileSubmissionInternalServerError creates a GetTransactionFileSubmissionInternalServerError with default headers values
func NewGetTransactionFileSubmissionInternalServerError() *GetTransactionFileSubmissionInternalServerError {
	return &GetTransactionFileSubmissionInternalServerError{}
}

/*
GetTransactionFileSubmissionInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetTransactionFileSubmissionInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionFileSubmissionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /files/transactions/{transaction_file_id}/submissions/{transaction_file_submission_id}][%d] getTransactionFileSubmissionInternalServerError", 500)
}

func (o *GetTransactionFileSubmissionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
