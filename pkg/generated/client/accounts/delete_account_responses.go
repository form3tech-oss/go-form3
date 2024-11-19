// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAccountReader is a Reader for the DeleteAccount structure.
type DeleteAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAccountNoContent creates a DeleteAccountNoContent with default headers values
func NewDeleteAccountNoContent() *DeleteAccountNoContent {
	return &DeleteAccountNoContent{}
}

/*
DeleteAccountNoContent handles this case with default header values.

Account deleted
*/
type DeleteAccountNoContent struct {
}

func (o *DeleteAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organisation/accounts/{id}][%d] deleteAccountNoContent", 204)
}

func (o *DeleteAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
