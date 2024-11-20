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

// PostDirectdebitsIDReversalsReversalIDSubmissionsReader is a Reader for the PostDirectdebitsIDReversalsReversalIDSubmissions structure.
type PostDirectdebitsIDReversalsReversalIDSubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostDirectdebitsIDReversalsReversalIDSubmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDirectdebitsIDReversalsReversalIDSubmissionsCreated creates a PostDirectdebitsIDReversalsReversalIDSubmissionsCreated with default headers values
func NewPostDirectdebitsIDReversalsReversalIDSubmissionsCreated() *PostDirectdebitsIDReversalsReversalIDSubmissionsCreated {
	return &PostDirectdebitsIDReversalsReversalIDSubmissionsCreated{}
}

/*
PostDirectdebitsIDReversalsReversalIDSubmissionsCreated handles this case with default header values.

Reversal submission creation response
*/
type PostDirectdebitsIDReversalsReversalIDSubmissionsCreated struct {

	//Payload

	// isStream: false
	*models.DirectDebitReversalSubmissionCreationResponse
}

// IsSuccess returns true when this post directdebits Id reversals reversal Id submissions created response has a 2xx status code
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post directdebits Id reversals reversal Id submissions created response has a 3xx status code
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post directdebits Id reversals reversal Id submissions created response has a 4xx status code
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post directdebits Id reversals reversal Id submissions created response has a 5xx status code
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post directdebits Id reversals reversal Id submissions created response a status code equal to that given
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post directdebits Id reversals reversal Id submissions created response
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsCreated) Code() int {
	return 201
}

func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/reversals/{reversalId}/submissions][%d] postDirectdebitsIdReversalsReversalIdSubmissionsCreated", 201)
}

func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitReversalSubmissionCreationResponse = new(models.DirectDebitReversalSubmissionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitReversalSubmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest creates a PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest with default headers values
func NewPostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest() *PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest {
	return &PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest{}
}

/*
PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest handles this case with default header values.

Reversal submission creation error
*/
type PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this post directdebits Id reversals reversal Id submissions bad request response has a 2xx status code
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post directdebits Id reversals reversal Id submissions bad request response has a 3xx status code
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post directdebits Id reversals reversal Id submissions bad request response has a 4xx status code
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post directdebits Id reversals reversal Id submissions bad request response has a 5xx status code
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post directdebits Id reversals reversal Id submissions bad request response a status code equal to that given
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post directdebits Id reversals reversal Id submissions bad request response
func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest) Code() int {
	return 400
}

func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/reversals/{reversalId}/submissions][%d] postDirectdebitsIdReversalsReversalIdSubmissionsBadRequest", 400)
}

func (o *PostDirectdebitsIDReversalsReversalIDSubmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
