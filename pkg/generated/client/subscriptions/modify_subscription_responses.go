// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
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

/*ModifySubscriptionOK handles this case with default header values.

Subscription details
*/
type ModifySubscriptionOK struct {

	//Payload

	// isStream: false
	*models.SubscriptionDetailsResponse
}

func (o *ModifySubscriptionOK) Error() string {
	return fmt.Sprintf("[PATCH /notification/subscriptions/{id}][%d] modifySubscriptionOK  %+v", 200, o)
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

/*ModifySubscriptionBadRequest handles this case with default header values.

Bad Request
*/
type ModifySubscriptionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ModifySubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /notification/subscriptions/{id}][%d] modifySubscriptionBadRequest  %+v", 400, o)
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

/*ModifySubscriptionNotFound handles this case with default header values.

Not Found
*/
type ModifySubscriptionNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ModifySubscriptionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /notification/subscriptions/{id}][%d] modifySubscriptionNotFound  %+v", 404, o)
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

/*ModifySubscriptionConflict handles this case with default header values.

Conflict
*/
type ModifySubscriptionConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ModifySubscriptionConflict) Error() string {
	return fmt.Sprintf("[PATCH /notification/subscriptions/{id}][%d] modifySubscriptionConflict  %+v", 409, o)
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

/*ModifySubscriptionInternalServerError handles this case with default header values.

Internal Error
*/
type ModifySubscriptionInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ModifySubscriptionInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /notification/subscriptions/{id}][%d] modifySubscriptionInternalServerError  %+v", 500, o)
}

func (o *ModifySubscriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
