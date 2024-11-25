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

// ClaimSubmissionStatus claim submission status
// swagger:model ClaimSubmissionStatus
type ClaimSubmissionStatus string

const (

	// ClaimSubmissionStatusAccepted captures enum value "accepted"
	ClaimSubmissionStatusAccepted ClaimSubmissionStatus = "accepted"

	// ClaimSubmissionStatusValidationFailed captures enum value "validation_failed"
	ClaimSubmissionStatusValidationFailed ClaimSubmissionStatus = "validation_failed"

	// ClaimSubmissionStatusValidationPending captures enum value "validation_pending"
	ClaimSubmissionStatusValidationPending ClaimSubmissionStatus = "validation_pending"

	// ClaimSubmissionStatusValidationPassed captures enum value "validation_passed"
	ClaimSubmissionStatusValidationPassed ClaimSubmissionStatus = "validation_passed"

	// ClaimSubmissionStatusReleasedToGateway captures enum value "released_to_gateway"
	ClaimSubmissionStatusReleasedToGateway ClaimSubmissionStatus = "released_to_gateway"

	// ClaimSubmissionStatusQueuedForDelivery captures enum value "queued_for_delivery"
	ClaimSubmissionStatusQueuedForDelivery ClaimSubmissionStatus = "queued_for_delivery"

	// ClaimSubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	ClaimSubmissionStatusDeliveryConfirmed ClaimSubmissionStatus = "delivery_confirmed"

	// ClaimSubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	ClaimSubmissionStatusDeliveryFailed ClaimSubmissionStatus = "delivery_failed"
)

// for schema
var claimSubmissionStatusEnum []interface{}

func init() {
	var res []ClaimSubmissionStatus
	if err := json.Unmarshal([]byte(`["accepted","validation_failed","validation_pending","validation_passed","released_to_gateway","queued_for_delivery","delivery_confirmed","delivery_failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		claimSubmissionStatusEnum = append(claimSubmissionStatusEnum, v)
	}
}

func (m ClaimSubmissionStatus) validateClaimSubmissionStatusEnum(path, location string, value ClaimSubmissionStatus) error {
	if err := validate.Enum(path, location, value, claimSubmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this claim submission status
func (m ClaimSubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClaimSubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimSubmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
