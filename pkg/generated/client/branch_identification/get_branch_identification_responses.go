// Code generated by go-swagger; DO NOT EDIT.

package branch_identification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// GetBranchIdentificationReader is a Reader for the GetBranchIdentification structure.
type GetBranchIdentificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBranchIdentificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBranchIdentificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBranchIdentificationOK creates a GetBranchIdentificationOK with default headers values
func NewGetBranchIdentificationOK() *GetBranchIdentificationOK {
	return &GetBranchIdentificationOK{}
}

/*
GetBranchIdentificationOK handles this case with default header values.

Branch Identification response
*/
type GetBranchIdentificationOK struct {

	//Payload

	// isStream: false
	*models.BranchIdentificationResponse
}

func (o *GetBranchIdentificationOK) Error() string {
	return fmt.Sprintf("[GET /organisation/branches/{branch_id}/identifications/{identification_id}][%d] getBranchIdentificationOK", 200)
}

func (o *GetBranchIdentificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.BranchIdentificationResponse = new(models.BranchIdentificationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.BranchIdentificationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
