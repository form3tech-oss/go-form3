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

// ListReturnSubmissionTasksReader is a Reader for the ListReturnSubmissionTasks structure.
type ListReturnSubmissionTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReturnSubmissionTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListReturnSubmissionTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListReturnSubmissionTasksOK creates a ListReturnSubmissionTasksOK with default headers values
func NewListReturnSubmissionTasksOK() *ListReturnSubmissionTasksOK {
	return &ListReturnSubmissionTasksOK{}
}

/*
ListReturnSubmissionTasksOK handles this case with default header values.

List of Task Details
*/
type ListReturnSubmissionTasksOK struct {

	//Payload

	// isStream: false
	*models.ReturnSubmissionTaskListResponse
}

// IsSuccess returns true when this list return submission tasks o k response has a 2xx status code
func (o *ListReturnSubmissionTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list return submission tasks o k response has a 3xx status code
func (o *ListReturnSubmissionTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list return submission tasks o k response has a 4xx status code
func (o *ListReturnSubmissionTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list return submission tasks o k response has a 5xx status code
func (o *ListReturnSubmissionTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list return submission tasks o k response a status code equal to that given
func (o *ListReturnSubmissionTasksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list return submission tasks o k response
func (o *ListReturnSubmissionTasksOK) Code() int {
	return 200
}

func (o *ListReturnSubmissionTasksOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{paymentId}/returns/{returnId}/submissions/{returnSubmissionId}/tasks][%d] listReturnSubmissionTasksOK", 200)
}

func (o *ListReturnSubmissionTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReturnSubmissionTaskListResponse = new(models.ReturnSubmissionTaskListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReturnSubmissionTaskListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}