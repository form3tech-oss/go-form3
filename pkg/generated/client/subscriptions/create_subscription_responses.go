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

// CreateSubscriptionReader is a Reader for the CreateSubscription structure.
type CreateSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSubscriptionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSubscriptionCreated creates a CreateSubscriptionCreated with default headers values
func NewCreateSubscriptionCreated() *CreateSubscriptionCreated {
	return &CreateSubscriptionCreated{}
}

/*CreateSubscriptionCreated handles this case with default header values.

Subscription creation response
*/
type CreateSubscriptionCreated struct {

	//Payload

	// isStream: false
	*models.SubscriptionCreationResponse
}

func (o *CreateSubscriptionCreated) Error() string {
	return fmt.Sprintf("[POST /notification/subscriptions][%d] createSubscriptionCreated", 201)
}

func (o *CreateSubscriptionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SubscriptionCreationResponse = new(models.SubscriptionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SubscriptionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
