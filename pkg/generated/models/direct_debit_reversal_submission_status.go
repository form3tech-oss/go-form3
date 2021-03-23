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

// DirectDebitReversalSubmissionStatus direct debit reversal submission status
// swagger:model DirectDebitReversalSubmissionStatus
type DirectDebitReversalSubmissionStatus string

const (

	// DirectDebitReversalSubmissionStatusAccepted captures enum value "accepted"
	DirectDebitReversalSubmissionStatusAccepted DirectDebitReversalSubmissionStatus = "accepted"

	// DirectDebitReversalSubmissionStatusReleasedToGateway captures enum value "released_to_gateway"
	DirectDebitReversalSubmissionStatusReleasedToGateway DirectDebitReversalSubmissionStatus = "released_to_gateway"

	// DirectDebitReversalSubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	DirectDebitReversalSubmissionStatusDeliveryConfirmed DirectDebitReversalSubmissionStatus = "delivery_confirmed"

	// DirectDebitReversalSubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	DirectDebitReversalSubmissionStatusDeliveryFailed DirectDebitReversalSubmissionStatus = "delivery_failed"

	// DirectDebitReversalSubmissionStatusSubmitted captures enum value "submitted"
	DirectDebitReversalSubmissionStatusSubmitted DirectDebitReversalSubmissionStatus = "submitted"

	// DirectDebitReversalSubmissionStatusValidationPending captures enum value "validation_pending"
	DirectDebitReversalSubmissionStatusValidationPending DirectDebitReversalSubmissionStatus = "validation_pending"

	// DirectDebitReversalSubmissionStatusValidationPassed captures enum value "validation_passed"
	DirectDebitReversalSubmissionStatusValidationPassed DirectDebitReversalSubmissionStatus = "validation_passed"

	// DirectDebitReversalSubmissionStatusQueuedForDelivery captures enum value "queued_for_delivery"
	DirectDebitReversalSubmissionStatusQueuedForDelivery DirectDebitReversalSubmissionStatus = "queued_for_delivery"

	// DirectDebitReversalSubmissionStatusPending captures enum value "pending"
	DirectDebitReversalSubmissionStatusPending DirectDebitReversalSubmissionStatus = "pending"
)

// for schema
var directDebitReversalSubmissionStatusEnum []interface{}

func init() {
	var res []DirectDebitReversalSubmissionStatus
	if err := json.Unmarshal([]byte(`["accepted","released_to_gateway","delivery_confirmed","delivery_failed","submitted","validation_pending","validation_passed","queued_for_delivery","pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directDebitReversalSubmissionStatusEnum = append(directDebitReversalSubmissionStatusEnum, v)
	}
}

func (m DirectDebitReversalSubmissionStatus) validateDirectDebitReversalSubmissionStatusEnum(path, location string, value DirectDebitReversalSubmissionStatus) error {
	if err := validate.Enum(path, location, value, directDebitReversalSubmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this direct debit reversal submission status
func (m DirectDebitReversalSubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDirectDebitReversalSubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReversalSubmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}