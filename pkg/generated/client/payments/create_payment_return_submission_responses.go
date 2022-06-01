// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v4/pkg/generated/models"
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

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentReturnSubmissionCreated creates a CreatePaymentReturnSubmissionCreated with default headers values
func NewCreatePaymentReturnSubmissionCreated() *CreatePaymentReturnSubmissionCreated {
	return &CreatePaymentReturnSubmissionCreated{}
}

/*CreatePaymentReturnSubmissionCreated handles this case with default header values.

Return submission creation response
*/
type CreatePaymentReturnSubmissionCreated struct {

	//Payload

	// isStream: false
	*models.ReturnSubmissionCreationResponse
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

/*CreatePaymentReturnSubmissionBadRequest handles this case with default header values.

Return submission creation error
*/
type CreatePaymentReturnSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
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
