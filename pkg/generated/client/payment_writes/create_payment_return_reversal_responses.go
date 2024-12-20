// Code generated by go-swagger; DO NOT EDIT.

package payment_writes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// CreatePaymentReturnReversalReader is a Reader for the CreatePaymentReturnReversal structure.
type CreatePaymentReturnReversalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentReturnReversalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentReturnReversalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentReturnReversalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreatePaymentReturnReversalUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreatePaymentReturnReversalForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreatePaymentReturnReversalNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreatePaymentReturnReversalConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreatePaymentReturnReversalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentReturnReversalCreated creates a CreatePaymentReturnReversalCreated with default headers values
func NewCreatePaymentReturnReversalCreated() *CreatePaymentReturnReversalCreated {
	return &CreatePaymentReturnReversalCreated{}
}

/*
CreatePaymentReturnReversalCreated handles this case with default header values.

Reversal creation response
*/
type CreatePaymentReturnReversalCreated struct {

	//Payload

	// isStream: false
	*models.ReturnReversalCreationResponse
}

// IsSuccess returns true when this create payment return reversal created response has a 2xx status code
func (o *CreatePaymentReturnReversalCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create payment return reversal created response has a 3xx status code
func (o *CreatePaymentReturnReversalCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return reversal created response has a 4xx status code
func (o *CreatePaymentReturnReversalCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create payment return reversal created response has a 5xx status code
func (o *CreatePaymentReturnReversalCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return reversal created response a status code equal to that given
func (o *CreatePaymentReturnReversalCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create payment return reversal created response
func (o *CreatePaymentReturnReversalCreated) Code() int {
	return 201
}

func (o *CreatePaymentReturnReversalCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/reversals][%d] createPaymentReturnReversalCreated", 201)
}

func (o *CreatePaymentReturnReversalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReturnReversalCreationResponse = new(models.ReturnReversalCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReturnReversalCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnReversalBadRequest creates a CreatePaymentReturnReversalBadRequest with default headers values
func NewCreatePaymentReturnReversalBadRequest() *CreatePaymentReturnReversalBadRequest {
	return &CreatePaymentReturnReversalBadRequest{}
}

/*
CreatePaymentReturnReversalBadRequest handles this case with default header values.

Reversal creation error
*/
type CreatePaymentReturnReversalBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return reversal bad request response has a 2xx status code
func (o *CreatePaymentReturnReversalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return reversal bad request response has a 3xx status code
func (o *CreatePaymentReturnReversalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return reversal bad request response has a 4xx status code
func (o *CreatePaymentReturnReversalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create payment return reversal bad request response has a 5xx status code
func (o *CreatePaymentReturnReversalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return reversal bad request response a status code equal to that given
func (o *CreatePaymentReturnReversalBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create payment return reversal bad request response
func (o *CreatePaymentReturnReversalBadRequest) Code() int {
	return 400
}

func (o *CreatePaymentReturnReversalBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/reversals][%d] createPaymentReturnReversalBadRequest", 400)
}

func (o *CreatePaymentReturnReversalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnReversalUnauthorized creates a CreatePaymentReturnReversalUnauthorized with default headers values
func NewCreatePaymentReturnReversalUnauthorized() *CreatePaymentReturnReversalUnauthorized {
	return &CreatePaymentReturnReversalUnauthorized{}
}

/*
CreatePaymentReturnReversalUnauthorized handles this case with default header values.

Reversal creation error
*/
type CreatePaymentReturnReversalUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return reversal unauthorized response has a 2xx status code
func (o *CreatePaymentReturnReversalUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return reversal unauthorized response has a 3xx status code
func (o *CreatePaymentReturnReversalUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return reversal unauthorized response has a 4xx status code
func (o *CreatePaymentReturnReversalUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create payment return reversal unauthorized response has a 5xx status code
func (o *CreatePaymentReturnReversalUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return reversal unauthorized response a status code equal to that given
func (o *CreatePaymentReturnReversalUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create payment return reversal unauthorized response
func (o *CreatePaymentReturnReversalUnauthorized) Code() int {
	return 401
}

func (o *CreatePaymentReturnReversalUnauthorized) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/reversals][%d] createPaymentReturnReversalUnauthorized", 401)
}

func (o *CreatePaymentReturnReversalUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnReversalForbidden creates a CreatePaymentReturnReversalForbidden with default headers values
func NewCreatePaymentReturnReversalForbidden() *CreatePaymentReturnReversalForbidden {
	return &CreatePaymentReturnReversalForbidden{}
}

/*
CreatePaymentReturnReversalForbidden handles this case with default header values.

Reversal creation error
*/
type CreatePaymentReturnReversalForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return reversal forbidden response has a 2xx status code
func (o *CreatePaymentReturnReversalForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return reversal forbidden response has a 3xx status code
func (o *CreatePaymentReturnReversalForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return reversal forbidden response has a 4xx status code
func (o *CreatePaymentReturnReversalForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create payment return reversal forbidden response has a 5xx status code
func (o *CreatePaymentReturnReversalForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return reversal forbidden response a status code equal to that given
func (o *CreatePaymentReturnReversalForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create payment return reversal forbidden response
func (o *CreatePaymentReturnReversalForbidden) Code() int {
	return 403
}

func (o *CreatePaymentReturnReversalForbidden) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/reversals][%d] createPaymentReturnReversalForbidden", 403)
}

func (o *CreatePaymentReturnReversalForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnReversalNotFound creates a CreatePaymentReturnReversalNotFound with default headers values
func NewCreatePaymentReturnReversalNotFound() *CreatePaymentReturnReversalNotFound {
	return &CreatePaymentReturnReversalNotFound{}
}

/*
CreatePaymentReturnReversalNotFound handles this case with default header values.

Reversal creation error
*/
type CreatePaymentReturnReversalNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return reversal not found response has a 2xx status code
func (o *CreatePaymentReturnReversalNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return reversal not found response has a 3xx status code
func (o *CreatePaymentReturnReversalNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return reversal not found response has a 4xx status code
func (o *CreatePaymentReturnReversalNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create payment return reversal not found response has a 5xx status code
func (o *CreatePaymentReturnReversalNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return reversal not found response a status code equal to that given
func (o *CreatePaymentReturnReversalNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create payment return reversal not found response
func (o *CreatePaymentReturnReversalNotFound) Code() int {
	return 404
}

func (o *CreatePaymentReturnReversalNotFound) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/reversals][%d] createPaymentReturnReversalNotFound", 404)
}

func (o *CreatePaymentReturnReversalNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnReversalConflict creates a CreatePaymentReturnReversalConflict with default headers values
func NewCreatePaymentReturnReversalConflict() *CreatePaymentReturnReversalConflict {
	return &CreatePaymentReturnReversalConflict{}
}

/*
CreatePaymentReturnReversalConflict handles this case with default header values.

Reversal creation error
*/
type CreatePaymentReturnReversalConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return reversal conflict response has a 2xx status code
func (o *CreatePaymentReturnReversalConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return reversal conflict response has a 3xx status code
func (o *CreatePaymentReturnReversalConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return reversal conflict response has a 4xx status code
func (o *CreatePaymentReturnReversalConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create payment return reversal conflict response has a 5xx status code
func (o *CreatePaymentReturnReversalConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return reversal conflict response a status code equal to that given
func (o *CreatePaymentReturnReversalConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create payment return reversal conflict response
func (o *CreatePaymentReturnReversalConflict) Code() int {
	return 409
}

func (o *CreatePaymentReturnReversalConflict) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/reversals][%d] createPaymentReturnReversalConflict", 409)
}

func (o *CreatePaymentReturnReversalConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnReversalInternalServerError creates a CreatePaymentReturnReversalInternalServerError with default headers values
func NewCreatePaymentReturnReversalInternalServerError() *CreatePaymentReturnReversalInternalServerError {
	return &CreatePaymentReturnReversalInternalServerError{}
}

/*
CreatePaymentReturnReversalInternalServerError handles this case with default header values.

Reversal creation error
*/
type CreatePaymentReturnReversalInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return reversal internal server error response has a 2xx status code
func (o *CreatePaymentReturnReversalInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return reversal internal server error response has a 3xx status code
func (o *CreatePaymentReturnReversalInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return reversal internal server error response has a 4xx status code
func (o *CreatePaymentReturnReversalInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create payment return reversal internal server error response has a 5xx status code
func (o *CreatePaymentReturnReversalInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create payment return reversal internal server error response a status code equal to that given
func (o *CreatePaymentReturnReversalInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create payment return reversal internal server error response
func (o *CreatePaymentReturnReversalInternalServerError) Code() int {
	return 500
}

func (o *CreatePaymentReturnReversalInternalServerError) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/reversals][%d] createPaymentReturnReversalInternalServerError", 500)
}

func (o *CreatePaymentReturnReversalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
