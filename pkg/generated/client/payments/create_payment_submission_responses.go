// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// CreatePaymentSubmissionReader is a Reader for the CreatePaymentSubmission structure.
type CreatePaymentSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentSubmissionCreated creates a CreatePaymentSubmissionCreated with default headers values
func NewCreatePaymentSubmissionCreated() *CreatePaymentSubmissionCreated {
	return &CreatePaymentSubmissionCreated{}
}

/*CreatePaymentSubmissionCreated handles this case with default header values.

Submission creation response
*/
type CreatePaymentSubmissionCreated struct {

	//Payload

	// isStream: false
	*models.PaymentSubmissionCreationResponse
}

func (o *CreatePaymentSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/submissions][%d] createPaymentSubmissionCreated", 201)
}

func (o *CreatePaymentSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PaymentSubmissionCreationResponse = new(models.PaymentSubmissionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PaymentSubmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentSubmissionBadRequest creates a CreatePaymentSubmissionBadRequest with default headers values
func NewCreatePaymentSubmissionBadRequest() *CreatePaymentSubmissionBadRequest {
	return &CreatePaymentSubmissionBadRequest{}
}

/*CreatePaymentSubmissionBadRequest handles this case with default header values.

Submission creation error
*/
type CreatePaymentSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreatePaymentSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/submissions][%d] createPaymentSubmissionBadRequest", 400)
}

func (o *CreatePaymentSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
