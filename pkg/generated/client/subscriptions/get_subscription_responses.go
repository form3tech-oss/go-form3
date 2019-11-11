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

// GetSubscriptionReader is a Reader for the GetSubscription structure.
type GetSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.

func (o *GetSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:

		result := NewGetSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSubscriptionOK creates a GetSubscriptionOK with default headers values
func NewGetSubscriptionOK() *GetSubscriptionOK {
	return &GetSubscriptionOK{}
}

/*GetSubscriptionOK handles this case with default header values.

Subscription details
*/
type GetSubscriptionOK struct {

	//Payload

	// isStream: false
	*models.SubscriptionDetailsResponse
}

func (o *GetSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /notification/subscriptions/{id}][%d] getSubscriptionOK  %+v", 200, o)
}

func (o *GetSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SubscriptionDetailsResponse = new(models.SubscriptionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SubscriptionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
