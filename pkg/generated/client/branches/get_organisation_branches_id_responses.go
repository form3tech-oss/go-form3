// Code generated by go-swagger; DO NOT EDIT.

package branches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// GetOrganisationBranchesIDReader is a Reader for the GetOrganisationBranchesID structure.
type GetOrganisationBranchesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganisationBranchesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganisationBranchesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganisationBranchesIDOK creates a GetOrganisationBranchesIDOK with default headers values
func NewGetOrganisationBranchesIDOK() *GetOrganisationBranchesIDOK {
	return &GetOrganisationBranchesIDOK{}
}

/*GetOrganisationBranchesIDOK handles this case with default header values.

Branch details
*/
type GetOrganisationBranchesIDOK struct {

	//Payload

	// isStream: false
	*models.BranchDetailsResponse
}

func (o *GetOrganisationBranchesIDOK) Error() string {
	return fmt.Sprintf("[GET /organisation/branches/{id}][%d] getOrganisationBranchesIdOK", 200)
}

func (o *GetOrganisationBranchesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.BranchDetailsResponse = new(models.BranchDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.BranchDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
