// Code generated by go-swagger; DO NOT EDIT.

package claims

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// GetClaimReader is a Reader for the GetClaim structure.
type GetClaimReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClaimReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClaimOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetClaimBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClaimOK creates a GetClaimOK with default headers values
func NewGetClaimOK() *GetClaimOK {
	return &GetClaimOK{}
}

/*GetClaimOK handles this case with default header values.

Claim details
*/
type GetClaimOK struct {

	//Payload

	// isStream: false
	*models.ClaimDetailsResponse
}

func (o *GetClaimOK) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}][%d] getClaimOK", 200)
}

func (o *GetClaimOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ClaimDetailsResponse = new(models.ClaimDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ClaimDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClaimBadRequest creates a GetClaimBadRequest with default headers values
func NewGetClaimBadRequest() *GetClaimBadRequest {
	return &GetClaimBadRequest{}
}

/*GetClaimBadRequest handles this case with default header values.

Error
*/
type GetClaimBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetClaimBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}][%d] getClaimBadRequest", 400)
}

func (o *GetClaimBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
