// Code generated by go-swagger; DO NOT EDIT.

package mandates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// ModifyMandateReader is a Reader for the ModifyMandate structure.
type ModifyMandateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyMandateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyMandateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewModifyMandateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyMandateOK creates a ModifyMandateOK with default headers values
func NewModifyMandateOK() *ModifyMandateOK {
	return &ModifyMandateOK{}
}

/*
ModifyMandateOK handles this case with default header values.

Mandate details
*/
type ModifyMandateOK struct {

	//Payload

	// isStream: false
	*models.MandateDetailsResponse
}

func (o *ModifyMandateOK) Error() string {
	return fmt.Sprintf("[PATCH /transaction/mandates/{id}][%d] modifyMandateOK", 200)
}

func (o *ModifyMandateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MandateDetailsResponse = new(models.MandateDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MandateDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyMandateBadRequest creates a ModifyMandateBadRequest with default headers values
func NewModifyMandateBadRequest() *ModifyMandateBadRequest {
	return &ModifyMandateBadRequest{}
}

/*
ModifyMandateBadRequest handles this case with default header values.

Mandate update error
*/
type ModifyMandateBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ModifyMandateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /transaction/mandates/{id}][%d] modifyMandateBadRequest", 400)
}

func (o *ModifyMandateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
