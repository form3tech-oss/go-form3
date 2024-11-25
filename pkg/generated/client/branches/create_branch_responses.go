// Code generated by go-swagger; DO NOT EDIT.

package branches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// CreateBranchReader is a Reader for the CreateBranch structure.
type CreateBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateBranchCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewCreateBranchConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateBranchCreated creates a CreateBranchCreated with default headers values
func NewCreateBranchCreated() *CreateBranchCreated {
	return &CreateBranchCreated{}
}

/*
CreateBranchCreated handles this case with default header values.

Branch creation response
*/
type CreateBranchCreated struct {

	//Payload

	// isStream: false
	*models.BranchCreationResponse
}

// IsSuccess returns true when this create branch created response has a 2xx status code
func (o *CreateBranchCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create branch created response has a 3xx status code
func (o *CreateBranchCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create branch created response has a 4xx status code
func (o *CreateBranchCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create branch created response has a 5xx status code
func (o *CreateBranchCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create branch created response a status code equal to that given
func (o *CreateBranchCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create branch created response
func (o *CreateBranchCreated) Code() int {
	return 201
}

func (o *CreateBranchCreated) Error() string {
	return fmt.Sprintf("[POST /organisation/branches][%d] createBranchCreated", 201)
}

func (o *CreateBranchCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.BranchCreationResponse = new(models.BranchCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.BranchCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBranchConflict creates a CreateBranchConflict with default headers values
func NewCreateBranchConflict() *CreateBranchConflict {
	return &CreateBranchConflict{}
}

/*
CreateBranchConflict handles this case with default header values.

Branch creation error, constraint violation of organisation id and bank id
*/
type CreateBranchConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create branch conflict response has a 2xx status code
func (o *CreateBranchConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create branch conflict response has a 3xx status code
func (o *CreateBranchConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create branch conflict response has a 4xx status code
func (o *CreateBranchConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create branch conflict response has a 5xx status code
func (o *CreateBranchConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create branch conflict response a status code equal to that given
func (o *CreateBranchConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create branch conflict response
func (o *CreateBranchConflict) Code() int {
	return 409
}

func (o *CreateBranchConflict) Error() string {
	return fmt.Sprintf("[POST /organisation/branches][%d] createBranchConflict", 409)
}

func (o *CreateBranchConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
