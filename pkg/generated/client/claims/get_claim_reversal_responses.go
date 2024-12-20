// Code generated by go-swagger; DO NOT EDIT.

package claims

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// GetClaimReversalReader is a Reader for the GetClaimReversal structure.
type GetClaimReversalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClaimReversalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClaimReversalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetClaimReversalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetClaimReversalNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClaimReversalOK creates a GetClaimReversalOK with default headers values
func NewGetClaimReversalOK() *GetClaimReversalOK {
	return &GetClaimReversalOK{}
}

/*
GetClaimReversalOK handles this case with default header values.

Claim Reversal details
*/
type GetClaimReversalOK struct {

	//Payload

	// isStream: false
	*models.ClaimReversalDetailsResponse
}

// IsSuccess returns true when this get claim reversal o k response has a 2xx status code
func (o *GetClaimReversalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get claim reversal o k response has a 3xx status code
func (o *GetClaimReversalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get claim reversal o k response has a 4xx status code
func (o *GetClaimReversalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get claim reversal o k response has a 5xx status code
func (o *GetClaimReversalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get claim reversal o k response a status code equal to that given
func (o *GetClaimReversalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get claim reversal o k response
func (o *GetClaimReversalOK) Code() int {
	return 200
}

func (o *GetClaimReversalOK) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/reversals/{reversalId}][%d] getClaimReversalOK", 200)
}

func (o *GetClaimReversalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ClaimReversalDetailsResponse = new(models.ClaimReversalDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ClaimReversalDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClaimReversalBadRequest creates a GetClaimReversalBadRequest with default headers values
func NewGetClaimReversalBadRequest() *GetClaimReversalBadRequest {
	return &GetClaimReversalBadRequest{}
}

/*
GetClaimReversalBadRequest handles this case with default header values.

Bad Request
*/
type GetClaimReversalBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this get claim reversal bad request response has a 2xx status code
func (o *GetClaimReversalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get claim reversal bad request response has a 3xx status code
func (o *GetClaimReversalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get claim reversal bad request response has a 4xx status code
func (o *GetClaimReversalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get claim reversal bad request response has a 5xx status code
func (o *GetClaimReversalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get claim reversal bad request response a status code equal to that given
func (o *GetClaimReversalBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get claim reversal bad request response
func (o *GetClaimReversalBadRequest) Code() int {
	return 400
}

func (o *GetClaimReversalBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/reversals/{reversalId}][%d] getClaimReversalBadRequest", 400)
}

func (o *GetClaimReversalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClaimReversalNotFound creates a GetClaimReversalNotFound with default headers values
func NewGetClaimReversalNotFound() *GetClaimReversalNotFound {
	return &GetClaimReversalNotFound{}
}

/*
GetClaimReversalNotFound handles this case with default header values.

Not Found
*/
type GetClaimReversalNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this get claim reversal not found response has a 2xx status code
func (o *GetClaimReversalNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get claim reversal not found response has a 3xx status code
func (o *GetClaimReversalNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get claim reversal not found response has a 4xx status code
func (o *GetClaimReversalNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get claim reversal not found response has a 5xx status code
func (o *GetClaimReversalNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get claim reversal not found response a status code equal to that given
func (o *GetClaimReversalNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get claim reversal not found response
func (o *GetClaimReversalNotFound) Code() int {
	return 404
}

func (o *GetClaimReversalNotFound) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/reversals/{reversalId}][%d] getClaimReversalNotFound", 404)
}

func (o *GetClaimReversalNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
