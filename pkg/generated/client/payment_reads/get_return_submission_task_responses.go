// Code generated by go-swagger; DO NOT EDIT.

package payment_reads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// GetReturnSubmissionTaskReader is a Reader for the GetReturnSubmissionTask structure.
type GetReturnSubmissionTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReturnSubmissionTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReturnSubmissionTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReturnSubmissionTaskOK creates a GetReturnSubmissionTaskOK with default headers values
func NewGetReturnSubmissionTaskOK() *GetReturnSubmissionTaskOK {
	return &GetReturnSubmissionTaskOK{}
}

/*
GetReturnSubmissionTaskOK handles this case with default header values.

Return Submission Task Details
*/
type GetReturnSubmissionTaskOK struct {

	//Payload

	// isStream: false
	*models.ReturnSubmissionTaskFetchResponse
}

// IsSuccess returns true when this get return submission task o k response has a 2xx status code
func (o *GetReturnSubmissionTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get return submission task o k response has a 3xx status code
func (o *GetReturnSubmissionTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get return submission task o k response has a 4xx status code
func (o *GetReturnSubmissionTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get return submission task o k response has a 5xx status code
func (o *GetReturnSubmissionTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get return submission task o k response a status code equal to that given
func (o *GetReturnSubmissionTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get return submission task o k response
func (o *GetReturnSubmissionTaskOK) Code() int {
	return 200
}

func (o *GetReturnSubmissionTaskOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{paymentId}/returns/{returnId}/submissions/{returnSubmissionId}/tasks/{taskId}][%d] getReturnSubmissionTaskOK", 200)
}

func (o *GetReturnSubmissionTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReturnSubmissionTaskFetchResponse = new(models.ReturnSubmissionTaskFetchResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReturnSubmissionTaskFetchResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}