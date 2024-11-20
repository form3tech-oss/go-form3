// Code generated by go-swagger; DO NOT EDIT.

package organisations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// ListUnitsReader is a Reader for the ListUnits structure.
type ListUnitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUnitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListUnitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListUnitsOK creates a ListUnitsOK with default headers values
func NewListUnitsOK() *ListUnitsOK {
	return &ListUnitsOK{}
}

/*
ListUnitsOK handles this case with default header values.

List of organisation details
*/
type ListUnitsOK struct {

	//Payload

	// isStream: false
	*models.OrganisationDetailsListResponse
}

func (o *ListUnitsOK) Error() string {
	return fmt.Sprintf("[GET /organisation/units][%d] listUnitsOK", 200)
}

func (o *ListUnitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.OrganisationDetailsListResponse = new(models.OrganisationDetailsListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.OrganisationDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
