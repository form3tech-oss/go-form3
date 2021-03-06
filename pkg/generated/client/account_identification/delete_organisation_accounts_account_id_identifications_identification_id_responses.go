// Code generated by go-swagger; DO NOT EDIT.

package account_identification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDReader is a Reader for the DeleteOrganisationAccountsAccountIDIdentificationsIdentificationID structure.
type DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNoContent creates a DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNoContent with default headers values
func NewDeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNoContent() *DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNoContent {
	return &DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNoContent{}
}

/*DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNoContent handles this case with default header values.

Account Identification deleted
*/
type DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNoContent struct {
}

func (o *DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organisation/accounts/{account_id}/identifications/{identification_id}][%d] deleteOrganisationAccountsAccountIdIdentificationsIdentificationIdNoContent", 204)
}

func (o *DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNotFound creates a DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNotFound with default headers values
func NewDeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNotFound() *DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNotFound {
	return &DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNotFound{}
}

/*DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNotFound handles this case with default header values.

Account Identification not found
*/
type DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNotFound struct {
}

func (o *DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organisation/accounts/{account_id}/identifications/{identification_id}][%d] deleteOrganisationAccountsAccountIdIdentificationsIdentificationIdNotFound", 404)
}

func (o *DeleteOrganisationAccountsAccountIDIdentificationsIdentificationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
