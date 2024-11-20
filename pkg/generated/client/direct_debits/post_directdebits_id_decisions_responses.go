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

// PostDirectdebitsIDDecisionsReader is a Reader for the PostDirectdebitsIDDecisions structure.
type PostDirectdebitsIDDecisionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDirectdebitsIDDecisionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostDirectdebitsIDDecisionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostDirectdebitsIDDecisionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostDirectdebitsIDDecisionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDirectdebitsIDDecisionsCreated creates a PostDirectdebitsIDDecisionsCreated with default headers values
func NewPostDirectdebitsIDDecisionsCreated() *PostDirectdebitsIDDecisionsCreated {
	return &PostDirectdebitsIDDecisionsCreated{}
}

/*
PostDirectdebitsIDDecisionsCreated handles this case with default header values.

Direct Debit decision creation response
*/
type PostDirectdebitsIDDecisionsCreated struct {

	//Payload

	// isStream: false
	*models.DirectDebitDecisionCreationResponse
}

// IsSuccess returns true when this post directdebits Id decisions created response has a 2xx status code
func (o *PostDirectdebitsIDDecisionsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post directdebits Id decisions created response has a 3xx status code
func (o *PostDirectdebitsIDDecisionsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post directdebits Id decisions created response has a 4xx status code
func (o *PostDirectdebitsIDDecisionsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post directdebits Id decisions created response has a 5xx status code
func (o *PostDirectdebitsIDDecisionsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post directdebits Id decisions created response a status code equal to that given
func (o *PostDirectdebitsIDDecisionsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post directdebits Id decisions created response
func (o *PostDirectdebitsIDDecisionsCreated) Code() int {
	return 201
}

func (o *PostDirectdebitsIDDecisionsCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions][%d] postDirectdebitsIdDecisionsCreated", 201)
}

func (o *PostDirectdebitsIDDecisionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitDecisionCreationResponse = new(models.DirectDebitDecisionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitDecisionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDirectdebitsIDDecisionsBadRequest creates a PostDirectdebitsIDDecisionsBadRequest with default headers values
func NewPostDirectdebitsIDDecisionsBadRequest() *PostDirectdebitsIDDecisionsBadRequest {
	return &PostDirectdebitsIDDecisionsBadRequest{}
}

/*
PostDirectdebitsIDDecisionsBadRequest handles this case with default header values.

Direct Debit decision creation error
*/
type PostDirectdebitsIDDecisionsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this post directdebits Id decisions bad request response has a 2xx status code
func (o *PostDirectdebitsIDDecisionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post directdebits Id decisions bad request response has a 3xx status code
func (o *PostDirectdebitsIDDecisionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post directdebits Id decisions bad request response has a 4xx status code
func (o *PostDirectdebitsIDDecisionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post directdebits Id decisions bad request response has a 5xx status code
func (o *PostDirectdebitsIDDecisionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post directdebits Id decisions bad request response a status code equal to that given
func (o *PostDirectdebitsIDDecisionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post directdebits Id decisions bad request response
func (o *PostDirectdebitsIDDecisionsBadRequest) Code() int {
	return 400
}

func (o *PostDirectdebitsIDDecisionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions][%d] postDirectdebitsIdDecisionsBadRequest", 400)
}

func (o *PostDirectdebitsIDDecisionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDirectdebitsIDDecisionsConflict creates a PostDirectdebitsIDDecisionsConflict with default headers values
func NewPostDirectdebitsIDDecisionsConflict() *PostDirectdebitsIDDecisionsConflict {
	return &PostDirectdebitsIDDecisionsConflict{}
}

/*
PostDirectdebitsIDDecisionsConflict handles this case with default header values.

Direct Debit decision creation conflict error
*/
type PostDirectdebitsIDDecisionsConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this post directdebits Id decisions conflict response has a 2xx status code
func (o *PostDirectdebitsIDDecisionsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post directdebits Id decisions conflict response has a 3xx status code
func (o *PostDirectdebitsIDDecisionsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post directdebits Id decisions conflict response has a 4xx status code
func (o *PostDirectdebitsIDDecisionsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post directdebits Id decisions conflict response has a 5xx status code
func (o *PostDirectdebitsIDDecisionsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post directdebits Id decisions conflict response a status code equal to that given
func (o *PostDirectdebitsIDDecisionsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the post directdebits Id decisions conflict response
func (o *PostDirectdebitsIDDecisionsConflict) Code() int {
	return 409
}

func (o *PostDirectdebitsIDDecisionsConflict) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions][%d] postDirectdebitsIdDecisionsConflict", 409)
}

func (o *PostDirectdebitsIDDecisionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
