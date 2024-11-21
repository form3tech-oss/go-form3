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

// CreatePaymentReader is a Reader for the CreatePayment structure.
type CreatePaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentCreated creates a CreatePaymentCreated with default headers values
func NewCreatePaymentCreated() *CreatePaymentCreated {
	return &CreatePaymentCreated{}
}

/*
CreatePaymentCreated handles this case with default header values.

Payment creation response
*/
type CreatePaymentCreated struct {

	//Payload

	// isStream: false
	*models.PaymentCreationResponse
}

// IsSuccess returns true when this create payment created response has a 2xx status code
func (o *CreatePaymentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create payment created response has a 3xx status code
func (o *CreatePaymentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment created response has a 4xx status code
func (o *CreatePaymentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create payment created response has a 5xx status code
func (o *CreatePaymentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment created response a status code equal to that given
func (o *CreatePaymentCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create payment created response
func (o *CreatePaymentCreated) Code() int {
	return 201
}

func (o *CreatePaymentCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments][%d] createPaymentCreated", 201)
}

func (o *CreatePaymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PaymentCreationResponse = new(models.PaymentCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PaymentCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentBadRequest creates a CreatePaymentBadRequest with default headers values
func NewCreatePaymentBadRequest() *CreatePaymentBadRequest {
	return &CreatePaymentBadRequest{}
}

/*
CreatePaymentBadRequest handles this case with default header values.

Payment creation error
*/
type CreatePaymentBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create payment bad request response has a 2xx status code
func (o *CreatePaymentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create payment bad request response has a 3xx status code
func (o *CreatePaymentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create payment bad request response has a 4xx status code
func (o *CreatePaymentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create payment bad request response has a 5xx status code
func (o *CreatePaymentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create payment bad request response a status code equal to that given
func (o *CreatePaymentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create payment bad request response
func (o *CreatePaymentBadRequest) Code() int {
	return 400
}

func (o *CreatePaymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments][%d] createPaymentBadRequest", 400)
}

func (o *CreatePaymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}