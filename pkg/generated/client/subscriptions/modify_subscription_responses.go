// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// ModifySubscriptionReader is a Reader for the ModifySubscription structure.
type ModifySubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifySubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifySubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewModifySubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewModifySubscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewModifySubscriptionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewModifySubscriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifySubscriptionOK creates a ModifySubscriptionOK with default headers values
func NewModifySubscriptionOK() *ModifySubscriptionOK {
	return &ModifySubscriptionOK{}
}

/*
ModifySubscriptionOK handles this case with default header values.

Subscription details
*/
type ModifySubscriptionOK struct {

	//Payload

	// isStream: false
	*models.SubscriptionDetailsResponse
}

// IsSuccess returns true when this modify subscription o k response has a 2xx status code
func (o *ModifySubscriptionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modify subscription o k response has a 3xx status code
func (o *ModifySubscriptionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify subscription o k response has a 4xx status code
func (o *ModifySubscriptionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify subscription o k response has a 5xx status code
func (o *ModifySubscriptionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this modify subscription o k response a status code equal to that given
func (o *ModifySubscriptionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the modify subscription o k response
func (o *ModifySubscriptionOK) Code() int {
	return 200
}

func (o *ModifySubscriptionOK) Error() string {
	return fmt.Sprintf("[PATCH /notification/subscriptions/{id}][%d] modifySubscriptionOK", 200)
}

func (o *ModifySubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SubscriptionDetailsResponse = new(models.SubscriptionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SubscriptionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifySubscriptionBadRequest creates a ModifySubscriptionBadRequest with default headers values
func NewModifySubscriptionBadRequest() *ModifySubscriptionBadRequest {
	return &ModifySubscriptionBadRequest{}
}

/*
ModifySubscriptionBadRequest handles this case with default header values.

Bad Request
*/
type ModifySubscriptionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this modify subscription bad request response has a 2xx status code
func (o *ModifySubscriptionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify subscription bad request response has a 3xx status code
func (o *ModifySubscriptionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify subscription bad request response has a 4xx status code
func (o *ModifySubscriptionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify subscription bad request response has a 5xx status code
func (o *ModifySubscriptionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this modify subscription bad request response a status code equal to that given
func (o *ModifySubscriptionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the modify subscription bad request response
func (o *ModifySubscriptionBadRequest) Code() int {
	return 400
}

func (o *ModifySubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /notification/subscriptions/{id}][%d] modifySubscriptionBadRequest", 400)
}

func (o *ModifySubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifySubscriptionNotFound creates a ModifySubscriptionNotFound with default headers values
func NewModifySubscriptionNotFound() *ModifySubscriptionNotFound {
	return &ModifySubscriptionNotFound{}
}

/*
ModifySubscriptionNotFound handles this case with default header values.

Not Found
*/
type ModifySubscriptionNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this modify subscription not found response has a 2xx status code
func (o *ModifySubscriptionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify subscription not found response has a 3xx status code
func (o *ModifySubscriptionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify subscription not found response has a 4xx status code
func (o *ModifySubscriptionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify subscription not found response has a 5xx status code
func (o *ModifySubscriptionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this modify subscription not found response a status code equal to that given
func (o *ModifySubscriptionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the modify subscription not found response
func (o *ModifySubscriptionNotFound) Code() int {
	return 404
}

func (o *ModifySubscriptionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /notification/subscriptions/{id}][%d] modifySubscriptionNotFound", 404)
}

func (o *ModifySubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifySubscriptionConflict creates a ModifySubscriptionConflict with default headers values
func NewModifySubscriptionConflict() *ModifySubscriptionConflict {
	return &ModifySubscriptionConflict{}
}

/*
ModifySubscriptionConflict handles this case with default header values.

Conflict
*/
type ModifySubscriptionConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this modify subscription conflict response has a 2xx status code
func (o *ModifySubscriptionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify subscription conflict response has a 3xx status code
func (o *ModifySubscriptionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify subscription conflict response has a 4xx status code
func (o *ModifySubscriptionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify subscription conflict response has a 5xx status code
func (o *ModifySubscriptionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this modify subscription conflict response a status code equal to that given
func (o *ModifySubscriptionConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the modify subscription conflict response
func (o *ModifySubscriptionConflict) Code() int {
	return 409
}

func (o *ModifySubscriptionConflict) Error() string {
	return fmt.Sprintf("[PATCH /notification/subscriptions/{id}][%d] modifySubscriptionConflict", 409)
}

func (o *ModifySubscriptionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifySubscriptionInternalServerError creates a ModifySubscriptionInternalServerError with default headers values
func NewModifySubscriptionInternalServerError() *ModifySubscriptionInternalServerError {
	return &ModifySubscriptionInternalServerError{}
}

/*
ModifySubscriptionInternalServerError handles this case with default header values.

Internal Error
*/
type ModifySubscriptionInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this modify subscription internal server error response has a 2xx status code
func (o *ModifySubscriptionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify subscription internal server error response has a 3xx status code
func (o *ModifySubscriptionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify subscription internal server error response has a 4xx status code
func (o *ModifySubscriptionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify subscription internal server error response has a 5xx status code
func (o *ModifySubscriptionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this modify subscription internal server error response a status code equal to that given
func (o *ModifySubscriptionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the modify subscription internal server error response
func (o *ModifySubscriptionInternalServerError) Code() int {
	return 500
}

func (o *ModifySubscriptionInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /notification/subscriptions/{id}][%d] modifySubscriptionInternalServerError", 500)
}

func (o *ModifySubscriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
