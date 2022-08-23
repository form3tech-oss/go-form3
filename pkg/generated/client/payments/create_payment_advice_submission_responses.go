// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v5/pkg/generated/models"
)

// CreatePaymentAdviceSubmissionReader is a Reader for the CreatePaymentAdviceSubmission structure.
type CreatePaymentAdviceSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentAdviceSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentAdviceSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentAdviceSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentAdviceSubmissionCreated creates a CreatePaymentAdviceSubmissionCreated with default headers values
func NewCreatePaymentAdviceSubmissionCreated() *CreatePaymentAdviceSubmissionCreated {
	return &CreatePaymentAdviceSubmissionCreated{}
}

/*CreatePaymentAdviceSubmissionCreated handles this case with default header values.

Advice submission creation response
*/
type CreatePaymentAdviceSubmissionCreated struct {

	//Payload

	// isStream: false
	*models.AdviceSubmissionCreationResponse
}

func (o *CreatePaymentAdviceSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/advices/{adviceId}/submissions][%d] createPaymentAdviceSubmissionCreated", 201)
}

func (o *CreatePaymentAdviceSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AdviceSubmissionCreationResponse = new(models.AdviceSubmissionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AdviceSubmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentAdviceSubmissionBadRequest creates a CreatePaymentAdviceSubmissionBadRequest with default headers values
func NewCreatePaymentAdviceSubmissionBadRequest() *CreatePaymentAdviceSubmissionBadRequest {
	return &CreatePaymentAdviceSubmissionBadRequest{}
}

/*CreatePaymentAdviceSubmissionBadRequest handles this case with default header values.

Advice submission creation error
*/
type CreatePaymentAdviceSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreatePaymentAdviceSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/advices/{adviceId}/submissions][%d] createPaymentAdviceSubmissionBadRequest", 400)
}

func (o *CreatePaymentAdviceSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
