// Code generated by go-swagger; DO NOT EDIT.

package account_validation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// GetValidationsGbdscSortcodesSortcodeReader is a Reader for the GetValidationsGbdscSortcodesSortcode structure.
type GetValidationsGbdscSortcodesSortcodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidationsGbdscSortcodesSortcodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetValidationsGbdscSortcodesSortcodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetValidationsGbdscSortcodesSortcodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetValidationsGbdscSortcodesSortcodeOK creates a GetValidationsGbdscSortcodesSortcodeOK with default headers values
func NewGetValidationsGbdscSortcodesSortcodeOK() *GetValidationsGbdscSortcodesSortcodeOK {
	return &GetValidationsGbdscSortcodesSortcodeOK{}
}

/*
GetValidationsGbdscSortcodesSortcodeOK handles this case with default header values.

Sort code details
*/
type GetValidationsGbdscSortcodesSortcodeOK struct {

	//Payload

	// isStream: false
	*models.SortCodeDetailsResponse
}

// IsSuccess returns true when this get validations gbdsc sortcodes sortcode o k response has a 2xx status code
func (o *GetValidationsGbdscSortcodesSortcodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get validations gbdsc sortcodes sortcode o k response has a 3xx status code
func (o *GetValidationsGbdscSortcodesSortcodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validations gbdsc sortcodes sortcode o k response has a 4xx status code
func (o *GetValidationsGbdscSortcodesSortcodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validations gbdsc sortcodes sortcode o k response has a 5xx status code
func (o *GetValidationsGbdscSortcodesSortcodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get validations gbdsc sortcodes sortcode o k response a status code equal to that given
func (o *GetValidationsGbdscSortcodesSortcodeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get validations gbdsc sortcodes sortcode o k response
func (o *GetValidationsGbdscSortcodesSortcodeOK) Code() int {
	return 200
}

func (o *GetValidationsGbdscSortcodesSortcodeOK) Error() string {
	return fmt.Sprintf("[GET /validations/gbdsc/sortcodes/{sortcode}][%d] getValidationsGbdscSortcodesSortcodeOK", 200)
}

func (o *GetValidationsGbdscSortcodesSortcodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SortCodeDetailsResponse = new(models.SortCodeDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SortCodeDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationsGbdscSortcodesSortcodeBadRequest creates a GetValidationsGbdscSortcodesSortcodeBadRequest with default headers values
func NewGetValidationsGbdscSortcodesSortcodeBadRequest() *GetValidationsGbdscSortcodesSortcodeBadRequest {
	return &GetValidationsGbdscSortcodesSortcodeBadRequest{}
}

/*
GetValidationsGbdscSortcodesSortcodeBadRequest handles this case with default header values.

Validation failed
*/
type GetValidationsGbdscSortcodesSortcodeBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this get validations gbdsc sortcodes sortcode bad request response has a 2xx status code
func (o *GetValidationsGbdscSortcodesSortcodeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validations gbdsc sortcodes sortcode bad request response has a 3xx status code
func (o *GetValidationsGbdscSortcodesSortcodeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validations gbdsc sortcodes sortcode bad request response has a 4xx status code
func (o *GetValidationsGbdscSortcodesSortcodeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get validations gbdsc sortcodes sortcode bad request response has a 5xx status code
func (o *GetValidationsGbdscSortcodesSortcodeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get validations gbdsc sortcodes sortcode bad request response a status code equal to that given
func (o *GetValidationsGbdscSortcodesSortcodeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get validations gbdsc sortcodes sortcode bad request response
func (o *GetValidationsGbdscSortcodesSortcodeBadRequest) Code() int {
	return 400
}

func (o *GetValidationsGbdscSortcodesSortcodeBadRequest) Error() string {
	return fmt.Sprintf("[GET /validations/gbdsc/sortcodes/{sortcode}][%d] getValidationsGbdscSortcodesSortcodeBadRequest", 400)
}

func (o *GetValidationsGbdscSortcodesSortcodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
