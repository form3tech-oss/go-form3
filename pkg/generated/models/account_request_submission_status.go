// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// AccountRequestSubmissionStatus account request submission status
// swagger:model AccountRequestSubmissionStatus
type AccountRequestSubmissionStatus string

const (

	// AccountRequestSubmissionStatusAccepted captures enum value "accepted"
	AccountRequestSubmissionStatusAccepted AccountRequestSubmissionStatus = "accepted"

	// AccountRequestSubmissionStatusValidationPending captures enum value "validation_pending"
	AccountRequestSubmissionStatusValidationPending AccountRequestSubmissionStatus = "validation_pending"

	// AccountRequestSubmissionStatusValidationPassed captures enum value "validation_passed"
	AccountRequestSubmissionStatusValidationPassed AccountRequestSubmissionStatus = "validation_passed"

	// AccountRequestSubmissionStatusReleasedToGateway captures enum value "released_to_gateway"
	AccountRequestSubmissionStatusReleasedToGateway AccountRequestSubmissionStatus = "released_to_gateway"

	// AccountRequestSubmissionStatusSubmitted captures enum value "submitted"
	AccountRequestSubmissionStatusSubmitted AccountRequestSubmissionStatus = "submitted"

	// AccountRequestSubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	AccountRequestSubmissionStatusDeliveryConfirmed AccountRequestSubmissionStatus = "delivery_confirmed"

	// AccountRequestSubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	AccountRequestSubmissionStatusDeliveryFailed AccountRequestSubmissionStatus = "delivery_failed"
)

// for schema
var accountRequestSubmissionStatusEnum []interface{}

func init() {
	var res []AccountRequestSubmissionStatus
	if err := json.Unmarshal([]byte(`["accepted","validation_pending","validation_passed","released_to_gateway","submitted","delivery_confirmed","delivery_failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountRequestSubmissionStatusEnum = append(accountRequestSubmissionStatusEnum, v)
	}
}

func (m AccountRequestSubmissionStatus) validateAccountRequestSubmissionStatusEnum(path, location string, value AccountRequestSubmissionStatus) error {
	if err := validate.Enum(path, location, value, accountRequestSubmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this account request submission status
func (m AccountRequestSubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccountRequestSubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountRequestSubmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
