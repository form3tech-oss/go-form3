// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// ModifyAccountReader is a Reader for the ModifyAccount structure.
type ModifyAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyAccountOK creates a ModifyAccountOK with default headers values
func NewModifyAccountOK() *ModifyAccountOK {
	return &ModifyAccountOK{}
}

/*
ModifyAccountOK handles this case with default header values.

Account updated
*/
type ModifyAccountOK struct {

	//Payload

	// isStream: false
	*models.AccountDetailsResponse
}

// IsSuccess returns true when this modify account o k response has a 2xx status code
func (o *ModifyAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modify account o k response has a 3xx status code
func (o *ModifyAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify account o k response has a 4xx status code
func (o *ModifyAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify account o k response has a 5xx status code
func (o *ModifyAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this modify account o k response a status code equal to that given
func (o *ModifyAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the modify account o k response
func (o *ModifyAccountOK) Code() int {
	return 200
}

func (o *ModifyAccountOK) Error() string {
	return fmt.Sprintf("[PATCH /organisation/accounts/{id}][%d] modifyAccountOK", 200)
}

func (o *ModifyAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountDetailsResponse = new(models.AccountDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
