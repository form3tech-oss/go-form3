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

// GetMandateSubmissionReader is a Reader for the GetMandateSubmission structure.
type GetMandateSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMandateSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMandateSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetMandateSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMandateSubmissionOK creates a GetMandateSubmissionOK with default headers values
func NewGetMandateSubmissionOK() *GetMandateSubmissionOK {
	return &GetMandateSubmissionOK{}
}

/*
GetMandateSubmissionOK handles this case with default header values.

Mandate Submission details
*/
type GetMandateSubmissionOK struct {

	//Payload

	// isStream: false
	*models.MandateSubmissionDetailsResponse
}

// IsSuccess returns true when this get mandate submission o k response has a 2xx status code
func (o *GetMandateSubmissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get mandate submission o k response has a 3xx status code
func (o *GetMandateSubmissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mandate submission o k response has a 4xx status code
func (o *GetMandateSubmissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mandate submission o k response has a 5xx status code
func (o *GetMandateSubmissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get mandate submission o k response a status code equal to that given
func (o *GetMandateSubmissionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get mandate submission o k response
func (o *GetMandateSubmissionOK) Code() int {
	return 200
}

func (o *GetMandateSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/mandates/{id}/submissions/{submissionId}][%d] getMandateSubmissionOK", 200)
}

func (o *GetMandateSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MandateSubmissionDetailsResponse = new(models.MandateSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MandateSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMandateSubmissionBadRequest creates a GetMandateSubmissionBadRequest with default headers values
func NewGetMandateSubmissionBadRequest() *GetMandateSubmissionBadRequest {
	return &GetMandateSubmissionBadRequest{}
}

/*
GetMandateSubmissionBadRequest handles this case with default header values.

Error
*/
type GetMandateSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this get mandate submission bad request response has a 2xx status code
func (o *GetMandateSubmissionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get mandate submission bad request response has a 3xx status code
func (o *GetMandateSubmissionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mandate submission bad request response has a 4xx status code
func (o *GetMandateSubmissionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get mandate submission bad request response has a 5xx status code
func (o *GetMandateSubmissionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get mandate submission bad request response a status code equal to that given
func (o *GetMandateSubmissionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get mandate submission bad request response
func (o *GetMandateSubmissionBadRequest) Code() int {
	return 400
}

func (o *GetMandateSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/mandates/{id}/submissions/{submissionId}][%d] getMandateSubmissionBadRequest", 400)
}

func (o *GetMandateSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
