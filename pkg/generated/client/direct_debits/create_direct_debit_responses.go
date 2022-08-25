// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v5/pkg/generated/models"
)

// CreateDirectDebitReader is a Reader for the CreateDirectDebit structure.
type CreateDirectDebitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDirectDebitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateDirectDebitCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateDirectDebitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDirectDebitCreated creates a CreateDirectDebitCreated with default headers values
func NewCreateDirectDebitCreated() *CreateDirectDebitCreated {
	return &CreateDirectDebitCreated{}
}

/*CreateDirectDebitCreated handles this case with default header values.

Direct Debit creation response
*/
type CreateDirectDebitCreated struct {

	//Payload

	// isStream: false
	*models.DirectDebitCreationResponse
}

func (o *CreateDirectDebitCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits][%d] createDirectDebitCreated", 201)
}

func (o *CreateDirectDebitCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitCreationResponse = new(models.DirectDebitCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDirectDebitBadRequest creates a CreateDirectDebitBadRequest with default headers values
func NewCreateDirectDebitBadRequest() *CreateDirectDebitBadRequest {
	return &CreateDirectDebitBadRequest{}
}

/*CreateDirectDebitBadRequest handles this case with default header values.

Direct Debit creation error
*/
type CreateDirectDebitBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateDirectDebitBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits][%d] createDirectDebitBadRequest", 400)
}

func (o *CreateDirectDebitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
