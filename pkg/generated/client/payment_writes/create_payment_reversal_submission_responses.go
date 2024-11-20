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

// CreatePaymentReversalSubmissionReader is a Reader for the CreatePaymentReversalSubmission structure.
type CreatePaymentReversalSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentReversalSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentReversalSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentReversalSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentReversalSubmissionCreated creates a CreatePaymentReversalSubmissionCreated with default headers values
func NewCreatePaymentReversalSubmissionCreated() *CreatePaymentReversalSubmissionCreated {
	return &CreatePaymentReversalSubmissionCreated{}
}

/*
CreatePaymentReversalSubmissionCreated handles this case with default header values.

Reversal submission creation response
*/
type CreatePaymentReversalSubmissionCreated struct {

	//Payload

	// isStream: false
	*models.ReversalSubmissionCreationResponse
}

func (o *CreatePaymentReversalSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/reversals/{reversalId}/submissions][%d] createPaymentReversalSubmissionCreated", 201)
}

func (o *CreatePaymentReversalSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReversalSubmissionCreationResponse = new(models.ReversalSubmissionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReversalSubmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReversalSubmissionBadRequest creates a CreatePaymentReversalSubmissionBadRequest with default headers values
func NewCreatePaymentReversalSubmissionBadRequest() *CreatePaymentReversalSubmissionBadRequest {
	return &CreatePaymentReversalSubmissionBadRequest{}
}

/*
CreatePaymentReversalSubmissionBadRequest handles this case with default header values.

Reversal submission creation error
*/
type CreatePaymentReversalSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreatePaymentReversalSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/reversals/{reversalId}/submissions][%d] createPaymentReversalSubmissionBadRequest", 400)
}

func (o *CreatePaymentReversalSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
