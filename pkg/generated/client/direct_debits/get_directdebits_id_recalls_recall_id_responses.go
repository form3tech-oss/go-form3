// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// GetDirectdebitsIDRecallsRecallIDReader is a Reader for the GetDirectdebitsIDRecallsRecallID structure.
type GetDirectdebitsIDRecallsRecallIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectdebitsIDRecallsRecallIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectdebitsIDRecallsRecallIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectdebitsIDRecallsRecallIDOK creates a GetDirectdebitsIDRecallsRecallIDOK with default headers values
func NewGetDirectdebitsIDRecallsRecallIDOK() *GetDirectdebitsIDRecallsRecallIDOK {
	return &GetDirectdebitsIDRecallsRecallIDOK{}
}

/*
GetDirectdebitsIDRecallsRecallIDOK handles this case with default header values.

Recall details
*/
type GetDirectdebitsIDRecallsRecallIDOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitRecallDetailsResponse
}

// IsSuccess returns true when this get directdebits Id recalls recall Id o k response has a 2xx status code
func (o *GetDirectdebitsIDRecallsRecallIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get directdebits Id recalls recall Id o k response has a 3xx status code
func (o *GetDirectdebitsIDRecallsRecallIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get directdebits Id recalls recall Id o k response has a 4xx status code
func (o *GetDirectdebitsIDRecallsRecallIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get directdebits Id recalls recall Id o k response has a 5xx status code
func (o *GetDirectdebitsIDRecallsRecallIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get directdebits Id recalls recall Id o k response a status code equal to that given
func (o *GetDirectdebitsIDRecallsRecallIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get directdebits Id recalls recall Id o k response
func (o *GetDirectdebitsIDRecallsRecallIDOK) Code() int {
	return 200
}

func (o *GetDirectdebitsIDRecallsRecallIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/recalls/{recallId}][%d] getDirectdebitsIdRecallsRecallIdOK", 200)
}

func (o *GetDirectdebitsIDRecallsRecallIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitRecallDetailsResponse = new(models.DirectDebitRecallDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitRecallDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
