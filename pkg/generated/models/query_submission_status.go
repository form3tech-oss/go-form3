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

// QuerySubmissionStatus query submission status
// swagger:model QuerySubmissionStatus
type QuerySubmissionStatus string

const (

	// QuerySubmissionStatusAccepted captures enum value "accepted"
	QuerySubmissionStatusAccepted QuerySubmissionStatus = "accepted"

	// QuerySubmissionStatusValidationPending captures enum value "validation_pending"
	QuerySubmissionStatusValidationPending QuerySubmissionStatus = "validation_pending"

	// QuerySubmissionStatusReleasedToGateway captures enum value "released_to_gateway"
	QuerySubmissionStatusReleasedToGateway QuerySubmissionStatus = "released_to_gateway"

	// QuerySubmissionStatusQueuedForDelivery captures enum value "queued_for_delivery"
	QuerySubmissionStatusQueuedForDelivery QuerySubmissionStatus = "queued_for_delivery"

	// QuerySubmissionStatusSubmitted captures enum value "submitted"
	QuerySubmissionStatusSubmitted QuerySubmissionStatus = "submitted"

	// QuerySubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	QuerySubmissionStatusDeliveryConfirmed QuerySubmissionStatus = "delivery_confirmed"

	// QuerySubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	QuerySubmissionStatusDeliveryFailed QuerySubmissionStatus = "delivery_failed"
)

// for schema
var querySubmissionStatusEnum []interface{}

func init() {
	var res []QuerySubmissionStatus
	if err := json.Unmarshal([]byte(`["accepted","validation_pending","released_to_gateway","queued_for_delivery","submitted","delivery_confirmed","delivery_failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		querySubmissionStatusEnum = append(querySubmissionStatusEnum, v)
	}
}

func (m QuerySubmissionStatus) validateQuerySubmissionStatusEnum(path, location string, value QuerySubmissionStatus) error {
	if err := validate.Enum(path, location, value, querySubmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this query submission status
func (m QuerySubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateQuerySubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuerySubmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
