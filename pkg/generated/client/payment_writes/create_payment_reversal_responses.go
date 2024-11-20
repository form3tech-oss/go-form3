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

// CreatePaymentReversalReader is a Reader for the CreatePaymentReversal structure.
type CreatePaymentReversalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentReversalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentReversalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentReversalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentReversalCreated creates a CreatePaymentReversalCreated with default headers values
func NewCreatePaymentReversalCreated() *CreatePaymentReversalCreated {
	return &CreatePaymentReversalCreated{}
}

/*
CreatePaymentReversalCreated handles this case with default header values.

Reversal creation response
*/
type CreatePaymentReversalCreated struct {

	//Payload

	// isStream: false
	*models.ReversalCreationResponse
}

func (o *CreatePaymentReversalCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/reversals][%d] createPaymentReversalCreated", 201)
}

func (o *CreatePaymentReversalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReversalCreationResponse = new(models.ReversalCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReversalCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReversalBadRequest creates a CreatePaymentReversalBadRequest with default headers values
func NewCreatePaymentReversalBadRequest() *CreatePaymentReversalBadRequest {
	return &CreatePaymentReversalBadRequest{}
}

/*
CreatePaymentReversalBadRequest handles this case with default header values.

Reversal creation error
*/
type CreatePaymentReversalBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreatePaymentReversalBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/reversals][%d] createPaymentReversalBadRequest", 400)
}

func (o *CreatePaymentReversalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
