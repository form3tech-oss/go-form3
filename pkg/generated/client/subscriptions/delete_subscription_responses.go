// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSubscriptionReader is a Reader for the DeleteSubscription structure.
type DeleteSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSubscriptionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSubscriptionNoContent creates a DeleteSubscriptionNoContent with default headers values
func NewDeleteSubscriptionNoContent() *DeleteSubscriptionNoContent {
	return &DeleteSubscriptionNoContent{}
}

/*DeleteSubscriptionNoContent handles this case with default header values.

Subscription deleted OK. No body content will be returned
*/
type DeleteSubscriptionNoContent struct {
}

func (o *DeleteSubscriptionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /notification/subscriptions/{id}][%d] deleteSubscriptionNoContent", 204)
}

func (o *DeleteSubscriptionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
