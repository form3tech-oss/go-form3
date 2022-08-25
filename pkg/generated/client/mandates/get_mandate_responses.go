// Code generated by go-swagger; DO NOT EDIT.

package mandates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v5/pkg/generated/models"
)

// GetMandateReader is a Reader for the GetMandate structure.
type GetMandateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMandateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMandateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMandateOK creates a GetMandateOK with default headers values
func NewGetMandateOK() *GetMandateOK {
	return &GetMandateOK{}
}

/*GetMandateOK handles this case with default header values.

Mandate details
*/
type GetMandateOK struct {

	//Payload

	// isStream: false
	*models.MandateDetailsResponse
}

func (o *GetMandateOK) Error() string {
	return fmt.Sprintf("[GET /transaction/mandates/{id}][%d] getMandateOK", 200)
}

func (o *GetMandateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MandateDetailsResponse = new(models.MandateDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MandateDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
