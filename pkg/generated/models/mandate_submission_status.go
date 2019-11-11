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

// MandateSubmissionStatus mandate submission status
// swagger:model MandateSubmissionStatus
type MandateSubmissionStatus string

const (

	// MandateSubmissionStatusAccepted captures enum value "accepted"
	MandateSubmissionStatusAccepted MandateSubmissionStatus = "accepted"

	// MandateSubmissionStatusValidationFailed captures enum value "validation_failed"
	MandateSubmissionStatusValidationFailed MandateSubmissionStatus = "validation_failed"

	// MandateSubmissionStatusValidationPending captures enum value "validation_pending"
	MandateSubmissionStatusValidationPending MandateSubmissionStatus = "validation_pending"

	// MandateSubmissionStatusValidationPassed captures enum value "validation_passed"
	MandateSubmissionStatusValidationPassed MandateSubmissionStatus = "validation_passed"

	// MandateSubmissionStatusReleasedToGateway captures enum value "released_to_gateway"
	MandateSubmissionStatusReleasedToGateway MandateSubmissionStatus = "released_to_gateway"

	// MandateSubmissionStatusQueuedForDelivery captures enum value "queued_for_delivery"
	MandateSubmissionStatusQueuedForDelivery MandateSubmissionStatus = "queued_for_delivery"

	// MandateSubmissionStatusSubmitted captures enum value "submitted"
	MandateSubmissionStatusSubmitted MandateSubmissionStatus = "submitted"

	// MandateSubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	MandateSubmissionStatusDeliveryConfirmed MandateSubmissionStatus = "delivery_confirmed"

	// MandateSubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	MandateSubmissionStatusDeliveryFailed MandateSubmissionStatus = "delivery_failed"
)

// for schema
var mandateSubmissionStatusEnum []interface{}

func init() {
	var res []MandateSubmissionStatus
	if err := json.Unmarshal([]byte(`["accepted","validation_failed","validation_pending","validation_passed","released_to_gateway","queued_for_delivery","submitted","delivery_confirmed","delivery_failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mandateSubmissionStatusEnum = append(mandateSubmissionStatusEnum, v)
	}
}

func (m MandateSubmissionStatus) validateMandateSubmissionStatusEnum(path, location string, value MandateSubmissionStatus) error {
	if err := validate.Enum(path, location, value, mandateSubmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this mandate submission status
func (m MandateSubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMandateSubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateSubmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
