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

// ListSubscriptionsReader is a Reader for the ListSubscriptions structure.
type ListSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSubscriptionsOK creates a ListSubscriptionsOK with default headers values
func NewListSubscriptionsOK() *ListSubscriptionsOK {
	return &ListSubscriptionsOK{}
}

/*
ListSubscriptionsOK handles this case with default header values.

List of subscription details
*/
type ListSubscriptionsOK struct {

	//Payload

	// isStream: false
	*models.SubscriptionDetailsListResponse
}

// IsSuccess returns true when this list subscriptions o k response has a 2xx status code
func (o *ListSubscriptionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list subscriptions o k response has a 3xx status code
func (o *ListSubscriptionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list subscriptions o k response has a 4xx status code
func (o *ListSubscriptionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list subscriptions o k response has a 5xx status code
func (o *ListSubscriptionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list subscriptions o k response a status code equal to that given
func (o *ListSubscriptionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list subscriptions o k response
func (o *ListSubscriptionsOK) Code() int {
	return 200
}

func (o *ListSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /notification/subscriptions][%d] listSubscriptionsOK", 200)
}

func (o *ListSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SubscriptionDetailsListResponse = new(models.SubscriptionDetailsListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SubscriptionDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
