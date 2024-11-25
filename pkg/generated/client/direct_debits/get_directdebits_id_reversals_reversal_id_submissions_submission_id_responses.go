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

// GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDReader is a Reader for the GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionID structure.
type GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK creates a GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK with default headers values
func NewGetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK() *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK {
	return &GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK{}
}

/*
GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK handles this case with default header values.

Reversal submission details
*/
type GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitReversalSubmissionDetailsResponse
}

// IsSuccess returns true when this get directdebits Id reversals reversal Id submissions submission Id o k response has a 2xx status code
func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get directdebits Id reversals reversal Id submissions submission Id o k response has a 3xx status code
func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get directdebits Id reversals reversal Id submissions submission Id o k response has a 4xx status code
func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get directdebits Id reversals reversal Id submissions submission Id o k response has a 5xx status code
func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get directdebits Id reversals reversal Id submissions submission Id o k response a status code equal to that given
func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get directdebits Id reversals reversal Id submissions submission Id o k response
func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK) Code() int {
	return 200
}

func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/reversals/{reversalId}/submissions/{submissionId}][%d] getDirectdebitsIdReversalsReversalIdSubmissionsSubmissionIdOK", 200)
}

func (o *GetDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitReversalSubmissionDetailsResponse = new(models.DirectDebitReversalSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitReversalSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
