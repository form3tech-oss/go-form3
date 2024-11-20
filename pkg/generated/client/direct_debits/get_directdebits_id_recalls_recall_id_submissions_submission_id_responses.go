// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDReader is a Reader for the GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionID structure.
type GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK creates a GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK with default headers values
func NewGetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK() *GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK {
	return &GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK{}
}

/*
GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK handles this case with default header values.

Recall submission details
*/
type GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitRecallSubmissionDetailsResponse
}

// IsSuccess returns true when this get directdebits Id recalls recall Id submissions submission Id o k response has a 2xx status code
func (o *GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get directdebits Id recalls recall Id submissions submission Id o k response has a 3xx status code
func (o *GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get directdebits Id recalls recall Id submissions submission Id o k response has a 4xx status code
func (o *GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get directdebits Id recalls recall Id submissions submission Id o k response has a 5xx status code
func (o *GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get directdebits Id recalls recall Id submissions submission Id o k response a status code equal to that given
func (o *GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get directdebits Id recalls recall Id submissions submission Id o k response
func (o *GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK) Code() int {
	return 200
}

func (o *GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/recalls/{recallId}/submissions/{submissionId}][%d] getDirectdebitsIdRecallsRecallIdSubmissionsSubmissionIdOK", 200)
}

func (o *GetDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitRecallSubmissionDetailsResponse = new(models.DirectDebitRecallSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitRecallSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
