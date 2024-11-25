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

// CreatePaymentReturnSubmissionReader is a Reader for the CreatePaymentReturnSubmission structure.
type CreatePaymentReturnSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentReturnSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentReturnSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentReturnSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreatePaymentReturnSubmissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreatePaymentReturnSubmissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreatePaymentReturnSubmissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreatePaymentReturnSubmissionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreatePaymentReturnSubmissionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentReturnSubmissionCreated creates a CreatePaymentReturnSubmissionCreated with default headers values
func NewCreatePaymentReturnSubmissionCreated() *CreatePaymentReturnSubmissionCreated {
	return &CreatePaymentReturnSubmissionCreated{}
}

/*
CreatePaymentReturnSubmissionCreated handles this case with default header values.

Return submission creation response
*/
type CreatePaymentReturnSubmissionCreated struct {

	//Payload

	// isStream: false
	*models.ReturnSubmissionCreationResponse
}

// IsSuccess returns true when this create payment return submission created response has a 2xx status code
func (o *CreatePaymentReturnSubmissionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create payment return submission created response has a 3xx status code
func (o *CreatePaymentReturnSubmissionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return submission created response has a 4xx status code
func (o *CreatePaymentReturnSubmissionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create payment return submission created response has a 5xx status code
func (o *CreatePaymentReturnSubmissionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return submission created response a status code equal to that given
func (o *CreatePaymentReturnSubmissionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create payment return submission created response
func (o *CreatePaymentReturnSubmissionCreated) Code() int {
	return 201
}

func (o *CreatePaymentReturnSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/submissions][%d] createPaymentReturnSubmissionCreated", 201)
}

func (o *CreatePaymentReturnSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReturnSubmissionCreationResponse = new(models.ReturnSubmissionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReturnSubmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnSubmissionBadRequest creates a CreatePaymentReturnSubmissionBadRequest with default headers values
func NewCreatePaymentReturnSubmissionBadRequest() *CreatePaymentReturnSubmissionBadRequest {
	return &CreatePaymentReturnSubmissionBadRequest{}
}

/*
CreatePaymentReturnSubmissionBadRequest handles this case with default header values.

Return submission creation error
*/
type CreatePaymentReturnSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return submission bad request response has a 2xx status code
func (o *CreatePaymentReturnSubmissionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return submission bad request response has a 3xx status code
func (o *CreatePaymentReturnSubmissionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return submission bad request response has a 4xx status code
func (o *CreatePaymentReturnSubmissionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create payment return submission bad request response has a 5xx status code
func (o *CreatePaymentReturnSubmissionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return submission bad request response a status code equal to that given
func (o *CreatePaymentReturnSubmissionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create payment return submission bad request response
func (o *CreatePaymentReturnSubmissionBadRequest) Code() int {
	return 400
}

func (o *CreatePaymentReturnSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/submissions][%d] createPaymentReturnSubmissionBadRequest", 400)
}

func (o *CreatePaymentReturnSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnSubmissionUnauthorized creates a CreatePaymentReturnSubmissionUnauthorized with default headers values
func NewCreatePaymentReturnSubmissionUnauthorized() *CreatePaymentReturnSubmissionUnauthorized {
	return &CreatePaymentReturnSubmissionUnauthorized{}
}

/*
CreatePaymentReturnSubmissionUnauthorized handles this case with default header values.

Return submission creation error
*/
type CreatePaymentReturnSubmissionUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return submission unauthorized response has a 2xx status code
func (o *CreatePaymentReturnSubmissionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return submission unauthorized response has a 3xx status code
func (o *CreatePaymentReturnSubmissionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return submission unauthorized response has a 4xx status code
func (o *CreatePaymentReturnSubmissionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create payment return submission unauthorized response has a 5xx status code
func (o *CreatePaymentReturnSubmissionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return submission unauthorized response a status code equal to that given
func (o *CreatePaymentReturnSubmissionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create payment return submission unauthorized response
func (o *CreatePaymentReturnSubmissionUnauthorized) Code() int {
	return 401
}

func (o *CreatePaymentReturnSubmissionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/submissions][%d] createPaymentReturnSubmissionUnauthorized", 401)
}

func (o *CreatePaymentReturnSubmissionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnSubmissionForbidden creates a CreatePaymentReturnSubmissionForbidden with default headers values
func NewCreatePaymentReturnSubmissionForbidden() *CreatePaymentReturnSubmissionForbidden {
	return &CreatePaymentReturnSubmissionForbidden{}
}

/*
CreatePaymentReturnSubmissionForbidden handles this case with default header values.

Return submission creation error
*/
type CreatePaymentReturnSubmissionForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return submission forbidden response has a 2xx status code
func (o *CreatePaymentReturnSubmissionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return submission forbidden response has a 3xx status code
func (o *CreatePaymentReturnSubmissionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return submission forbidden response has a 4xx status code
func (o *CreatePaymentReturnSubmissionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create payment return submission forbidden response has a 5xx status code
func (o *CreatePaymentReturnSubmissionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return submission forbidden response a status code equal to that given
func (o *CreatePaymentReturnSubmissionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create payment return submission forbidden response
func (o *CreatePaymentReturnSubmissionForbidden) Code() int {
	return 403
}

func (o *CreatePaymentReturnSubmissionForbidden) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/submissions][%d] createPaymentReturnSubmissionForbidden", 403)
}

func (o *CreatePaymentReturnSubmissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnSubmissionNotFound creates a CreatePaymentReturnSubmissionNotFound with default headers values
func NewCreatePaymentReturnSubmissionNotFound() *CreatePaymentReturnSubmissionNotFound {
	return &CreatePaymentReturnSubmissionNotFound{}
}

/*
CreatePaymentReturnSubmissionNotFound handles this case with default header values.

Return submission creation error
*/
type CreatePaymentReturnSubmissionNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return submission not found response has a 2xx status code
func (o *CreatePaymentReturnSubmissionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return submission not found response has a 3xx status code
func (o *CreatePaymentReturnSubmissionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return submission not found response has a 4xx status code
func (o *CreatePaymentReturnSubmissionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create payment return submission not found response has a 5xx status code
func (o *CreatePaymentReturnSubmissionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return submission not found response a status code equal to that given
func (o *CreatePaymentReturnSubmissionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create payment return submission not found response
func (o *CreatePaymentReturnSubmissionNotFound) Code() int {
	return 404
}

func (o *CreatePaymentReturnSubmissionNotFound) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/submissions][%d] createPaymentReturnSubmissionNotFound", 404)
}

func (o *CreatePaymentReturnSubmissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnSubmissionConflict creates a CreatePaymentReturnSubmissionConflict with default headers values
func NewCreatePaymentReturnSubmissionConflict() *CreatePaymentReturnSubmissionConflict {
	return &CreatePaymentReturnSubmissionConflict{}
}

/*
CreatePaymentReturnSubmissionConflict handles this case with default header values.

Return submission creation error
*/
type CreatePaymentReturnSubmissionConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return submission conflict response has a 2xx status code
func (o *CreatePaymentReturnSubmissionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return submission conflict response has a 3xx status code
func (o *CreatePaymentReturnSubmissionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return submission conflict response has a 4xx status code
func (o *CreatePaymentReturnSubmissionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create payment return submission conflict response has a 5xx status code
func (o *CreatePaymentReturnSubmissionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment return submission conflict response a status code equal to that given
func (o *CreatePaymentReturnSubmissionConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create payment return submission conflict response
func (o *CreatePaymentReturnSubmissionConflict) Code() int {
	return 409
}

func (o *CreatePaymentReturnSubmissionConflict) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/submissions][%d] createPaymentReturnSubmissionConflict", 409)
}

func (o *CreatePaymentReturnSubmissionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnSubmissionInternalServerError creates a CreatePaymentReturnSubmissionInternalServerError with default headers values
func NewCreatePaymentReturnSubmissionInternalServerError() *CreatePaymentReturnSubmissionInternalServerError {
	return &CreatePaymentReturnSubmissionInternalServerError{}
}

/*
CreatePaymentReturnSubmissionInternalServerError handles this case with default header values.

Return submission creation error
*/
type CreatePaymentReturnSubmissionInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment return submission internal server error response has a 2xx status code
func (o *CreatePaymentReturnSubmissionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment return submission internal server error response has a 3xx status code
func (o *CreatePaymentReturnSubmissionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment return submission internal server error response has a 4xx status code
func (o *CreatePaymentReturnSubmissionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create payment return submission internal server error response has a 5xx status code
func (o *CreatePaymentReturnSubmissionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create payment return submission internal server error response a status code equal to that given
func (o *CreatePaymentReturnSubmissionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create payment return submission internal server error response
func (o *CreatePaymentReturnSubmissionInternalServerError) Code() int {
	return 500
}

func (o *CreatePaymentReturnSubmissionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/submissions][%d] createPaymentReturnSubmissionInternalServerError", 500)
}

func (o *CreatePaymentReturnSubmissionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
