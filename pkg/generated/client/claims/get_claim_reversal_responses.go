// Code generated by go-swagger; DO NOT EDIT.

package claims

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
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

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClaimReversalOK creates a GetClaimReversalOK with default headers values
func NewGetClaimReversalOK() *GetClaimReversalOK {
	return &GetClaimReversalOK{}
}

/*GetClaimReversalOK handles this case with default header values.

Claim Reversal details
*/
type GetClaimReversalOK struct {

	//Payload

	// isStream: false
	*models.ClaimReversalDetailsResponse
}

func (o *GetClaimReversalOK) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/reversals/{reversalId}][%d] getClaimReversalOK  %+v", 200, o)
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

/*GetClaimReversalBadRequest handles this case with default header values.

Error
*/
type GetClaimReversalBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetClaimReversalBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/reversals/{reversalId}][%d] getClaimReversalBadRequest  %+v", 400, o)
}

func (o *GetClaimReversalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
