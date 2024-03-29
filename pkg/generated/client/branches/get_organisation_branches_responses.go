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

// GetOrganisationBranchesReader is a Reader for the GetOrganisationBranches structure.
type GetOrganisationBranchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganisationBranchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganisationBranchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganisationBranchesOK creates a GetOrganisationBranchesOK with default headers values
func NewGetOrganisationBranchesOK() *GetOrganisationBranchesOK {
	return &GetOrganisationBranchesOK{}
}

/*GetOrganisationBranchesOK handles this case with default header values.

List of branch details
*/
type GetOrganisationBranchesOK struct {

	//Payload

	// isStream: false
	*models.BranchDetailsListResponse
}

func (o *GetOrganisationBranchesOK) Error() string {
	return fmt.Sprintf("[GET /organisation/branches][%d] getOrganisationBranchesOK", 200)
}

func (o *GetOrganisationBranchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.BranchDetailsListResponse = new(models.BranchDetailsListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.BranchDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
