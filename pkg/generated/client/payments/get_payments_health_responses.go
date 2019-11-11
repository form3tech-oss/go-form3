// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetPaymentsHealthReader is a Reader for the GetPaymentsHealth structure.
type GetPaymentsHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.

func (o *GetPaymentsHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:

		result := NewGetPaymentsHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsHealthOK creates a GetPaymentsHealthOK with default headers values
func NewGetPaymentsHealthOK() *GetPaymentsHealthOK {
	return &GetPaymentsHealthOK{}
}

/*GetPaymentsHealthOK handles this case with default header values.

Payment service health
*/
type GetPaymentsHealthOK struct {

	//Payload

	// isStream: false
	*models.Health
}

func (o *GetPaymentsHealthOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/health][%d] getPaymentsHealthOK  %+v", 200, o)
}

func (o *GetPaymentsHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Health = new(models.Health)

	// response payload

	if err := consumer.Consume(response.Body(), o.Health); err != nil && err != io.EOF {
		return err
	}

	return nil
}
