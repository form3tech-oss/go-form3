// Code generated by go-swagger; DO NOT EDIT.

package mandates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// CreateMandateReader is a Reader for the CreateMandate structure.
type CreateMandateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMandateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateMandateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateMandateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateMandateCreated creates a CreateMandateCreated with default headers values
func NewCreateMandateCreated() *CreateMandateCreated {
	return &CreateMandateCreated{}
}

/*CreateMandateCreated handles this case with default header values.

Mandate creation response
*/
type CreateMandateCreated struct {

	//Payload

	// isStream: false
	*models.MandateCreationResponse
}

func (o *CreateMandateCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/mandates][%d] createMandateCreated", 201)
}

func (o *CreateMandateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MandateCreationResponse = new(models.MandateCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MandateCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMandateBadRequest creates a CreateMandateBadRequest with default headers values
func NewCreateMandateBadRequest() *CreateMandateBadRequest {
	return &CreateMandateBadRequest{}
}

/*CreateMandateBadRequest handles this case with default header values.

Mandate creation error
*/
type CreateMandateBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateMandateBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/mandates][%d] createMandateBadRequest", 400)
}

func (o *CreateMandateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
