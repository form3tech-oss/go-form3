// Code generated by go-swagger; DO NOT EDIT.

package mandates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// CreateMandateSubmissionReader is a Reader for the CreateMandateSubmission structure.
type CreateMandateSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMandateSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateMandateSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateMandateSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateMandateSubmissionCreated creates a CreateMandateSubmissionCreated with default headers values
func NewCreateMandateSubmissionCreated() *CreateMandateSubmissionCreated {
	return &CreateMandateSubmissionCreated{}
}

/*
CreateMandateSubmissionCreated handles this case with default header values.

Mandate Submission creation response
*/
type CreateMandateSubmissionCreated struct {

	//Payload

	// isStream: false
	*models.MandateSubmissionDetailsResponse
}

// IsSuccess returns true when this create mandate submission created response has a 2xx status code
func (o *CreateMandateSubmissionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create mandate submission created response has a 3xx status code
func (o *CreateMandateSubmissionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create mandate submission created response has a 4xx status code
func (o *CreateMandateSubmissionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create mandate submission created response has a 5xx status code
func (o *CreateMandateSubmissionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create mandate submission created response a status code equal to that given
func (o *CreateMandateSubmissionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create mandate submission created response
func (o *CreateMandateSubmissionCreated) Code() int {
	return 201
}

func (o *CreateMandateSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/mandates/{id}/submissions][%d] createMandateSubmissionCreated", 201)
}

func (o *CreateMandateSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MandateSubmissionDetailsResponse = new(models.MandateSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MandateSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMandateSubmissionBadRequest creates a CreateMandateSubmissionBadRequest with default headers values
func NewCreateMandateSubmissionBadRequest() *CreateMandateSubmissionBadRequest {
	return &CreateMandateSubmissionBadRequest{}
}

/*
CreateMandateSubmissionBadRequest handles this case with default header values.

Mandate Submission creation error
*/
type CreateMandateSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create mandate submission bad request response has a 2xx status code
func (o *CreateMandateSubmissionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create mandate submission bad request response has a 3xx status code
func (o *CreateMandateSubmissionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create mandate submission bad request response has a 4xx status code
func (o *CreateMandateSubmissionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create mandate submission bad request response has a 5xx status code
func (o *CreateMandateSubmissionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create mandate submission bad request response a status code equal to that given
func (o *CreateMandateSubmissionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create mandate submission bad request response
func (o *CreateMandateSubmissionBadRequest) Code() int {
	return 400
}

func (o *CreateMandateSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/mandates/{id}/submissions][%d] createMandateSubmissionBadRequest", 400)
}

func (o *CreateMandateSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
