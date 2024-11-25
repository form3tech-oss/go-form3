// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DirectDebitDecisionSubmissionStatus direct debit decision submission status
// swagger:model DirectDebitDecisionSubmissionStatus
type DirectDebitDecisionSubmissionStatus string

const (

	// DirectDebitDecisionSubmissionStatusAccepted captures enum value "accepted"
	DirectDebitDecisionSubmissionStatusAccepted DirectDebitDecisionSubmissionStatus = "accepted"

	// DirectDebitDecisionSubmissionStatusReleasedToGateway captures enum value "released_to_gateway"
	DirectDebitDecisionSubmissionStatusReleasedToGateway DirectDebitDecisionSubmissionStatus = "released_to_gateway"

	// DirectDebitDecisionSubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	DirectDebitDecisionSubmissionStatusDeliveryConfirmed DirectDebitDecisionSubmissionStatus = "delivery_confirmed"

	// DirectDebitDecisionSubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	DirectDebitDecisionSubmissionStatusDeliveryFailed DirectDebitDecisionSubmissionStatus = "delivery_failed"

	// DirectDebitDecisionSubmissionStatusSubmitted captures enum value "submitted"
	DirectDebitDecisionSubmissionStatusSubmitted DirectDebitDecisionSubmissionStatus = "submitted"

	// DirectDebitDecisionSubmissionStatusValidationPending captures enum value "validation_pending"
	DirectDebitDecisionSubmissionStatusValidationPending DirectDebitDecisionSubmissionStatus = "validation_pending"

	// DirectDebitDecisionSubmissionStatusValidationPassed captures enum value "validation_passed"
	DirectDebitDecisionSubmissionStatusValidationPassed DirectDebitDecisionSubmissionStatus = "validation_passed"

	// DirectDebitDecisionSubmissionStatusQueuedForDelivery captures enum value "queued_for_delivery"
	DirectDebitDecisionSubmissionStatusQueuedForDelivery DirectDebitDecisionSubmissionStatus = "queued_for_delivery"
)

// for schema
var directDebitDecisionSubmissionStatusEnum []interface{}

func init() {
	var res []DirectDebitDecisionSubmissionStatus
	if err := json.Unmarshal([]byte(`["accepted","released_to_gateway","delivery_confirmed","delivery_failed","submitted","validation_pending","validation_passed","queued_for_delivery"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directDebitDecisionSubmissionStatusEnum = append(directDebitDecisionSubmissionStatusEnum, v)
	}
}

func (m DirectDebitDecisionSubmissionStatus) validateDirectDebitDecisionSubmissionStatusEnum(path, location string, value DirectDebitDecisionSubmissionStatus) error {
	if err := validate.Enum(path, location, value, directDebitDecisionSubmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this direct debit decision submission status
func (m DirectDebitDecisionSubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDirectDebitDecisionSubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitDecisionSubmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
