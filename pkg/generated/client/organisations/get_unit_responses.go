// Code generated by go-swagger; DO NOT EDIT.

package organisations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// GetUnitReader is a Reader for the GetUnit structure.
type GetUnitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUnitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUnitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUnitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUnitOK creates a GetUnitOK with default headers values
func NewGetUnitOK() *GetUnitOK {
	return &GetUnitOK{}
}

/*
GetUnitOK handles this case with default header values.

Organisation details
*/
type GetUnitOK struct {

	//Payload

	// isStream: false
	*models.OrganisationDetailsResponse
}

func (o *GetUnitOK) Error() string {
	return fmt.Sprintf("[GET /organisation/units/{id}][%d] getUnitOK", 200)
}

func (o *GetUnitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.OrganisationDetailsResponse = new(models.OrganisationDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.OrganisationDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUnitNotFound creates a GetUnitNotFound with default headers values
func NewGetUnitNotFound() *GetUnitNotFound {
	return &GetUnitNotFound{}
}

/*
GetUnitNotFound handles this case with default header values.

Not Found
*/
type GetUnitNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetUnitNotFound) Error() string {
	return fmt.Sprintf("[GET /organisation/units/{id}][%d] getUnitNotFound", 404)
}

func (o *GetUnitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
