// Code generated by go-swagger; DO NOT EDIT.

package claims

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// GetClaimReversalSubmissionReader is a Reader for the GetClaimReversalSubmission structure.
type GetClaimReversalSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClaimReversalSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClaimReversalSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetClaimReversalSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetClaimReversalSubmissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClaimReversalSubmissionOK creates a GetClaimReversalSubmissionOK with default headers values
func NewGetClaimReversalSubmissionOK() *GetClaimReversalSubmissionOK {
	return &GetClaimReversalSubmissionOK{}
}

/*
GetClaimReversalSubmissionOK handles this case with default header values.

Claim Reversal Submission details
*/
type GetClaimReversalSubmissionOK struct {

	//Payload

	// isStream: false
	*models.ClaimReversalSubmissionDetailsResponse
}

// IsSuccess returns true when this get claim reversal submission o k response has a 2xx status code
func (o *GetClaimReversalSubmissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get claim reversal submission o k response has a 3xx status code
func (o *GetClaimReversalSubmissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get claim reversal submission o k response has a 4xx status code
func (o *GetClaimReversalSubmissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get claim reversal submission o k response has a 5xx status code
func (o *GetClaimReversalSubmissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get claim reversal submission o k response a status code equal to that given
func (o *GetClaimReversalSubmissionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get claim reversal submission o k response
func (o *GetClaimReversalSubmissionOK) Code() int {
	return 200
}

func (o *GetClaimReversalSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/reversals/{reversalId}/submissions/{submissionId}][%d] getClaimReversalSubmissionOK", 200)
}

func (o *GetClaimReversalSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ClaimReversalSubmissionDetailsResponse = new(models.ClaimReversalSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ClaimReversalSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClaimReversalSubmissionBadRequest creates a GetClaimReversalSubmissionBadRequest with default headers values
func NewGetClaimReversalSubmissionBadRequest() *GetClaimReversalSubmissionBadRequest {
	return &GetClaimReversalSubmissionBadRequest{}
}

/*
GetClaimReversalSubmissionBadRequest handles this case with default header values.

Bad Request
*/
type GetClaimReversalSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this get claim reversal submission bad request response has a 2xx status code
func (o *GetClaimReversalSubmissionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get claim reversal submission bad request response has a 3xx status code
func (o *GetClaimReversalSubmissionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get claim reversal submission bad request response has a 4xx status code
func (o *GetClaimReversalSubmissionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get claim reversal submission bad request response has a 5xx status code
func (o *GetClaimReversalSubmissionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get claim reversal submission bad request response a status code equal to that given
func (o *GetClaimReversalSubmissionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get claim reversal submission bad request response
func (o *GetClaimReversalSubmissionBadRequest) Code() int {
	return 400
}

func (o *GetClaimReversalSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/reversals/{reversalId}/submissions/{submissionId}][%d] getClaimReversalSubmissionBadRequest", 400)
}

func (o *GetClaimReversalSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClaimReversalSubmissionNotFound creates a GetClaimReversalSubmissionNotFound with default headers values
func NewGetClaimReversalSubmissionNotFound() *GetClaimReversalSubmissionNotFound {
	return &GetClaimReversalSubmissionNotFound{}
}

/*
GetClaimReversalSubmissionNotFound handles this case with default header values.

Not Found
*/
type GetClaimReversalSubmissionNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this get claim reversal submission not found response has a 2xx status code
func (o *GetClaimReversalSubmissionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get claim reversal submission not found response has a 3xx status code
func (o *GetClaimReversalSubmissionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get claim reversal submission not found response has a 4xx status code
func (o *GetClaimReversalSubmissionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get claim reversal submission not found response has a 5xx status code
func (o *GetClaimReversalSubmissionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get claim reversal submission not found response a status code equal to that given
func (o *GetClaimReversalSubmissionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get claim reversal submission not found response
func (o *GetClaimReversalSubmissionNotFound) Code() int {
	return 404
}

func (o *GetClaimReversalSubmissionNotFound) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/reversals/{reversalId}/submissions/{submissionId}][%d] getClaimReversalSubmissionNotFound", 404)
}

func (o *GetClaimReversalSubmissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
