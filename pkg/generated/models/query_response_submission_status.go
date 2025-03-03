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

// QueryResponseSubmissionStatus query response submission status
// swagger:model QueryResponseSubmissionStatus
type QueryResponseSubmissionStatus string

const (

	// QueryResponseSubmissionStatusAccepted captures enum value "accepted"
	QueryResponseSubmissionStatusAccepted QueryResponseSubmissionStatus = "accepted"

	// QueryResponseSubmissionStatusValidationPending captures enum value "validation_pending"
	QueryResponseSubmissionStatusValidationPending QueryResponseSubmissionStatus = "validation_pending"

	// QueryResponseSubmissionStatusReleasedToGateway captures enum value "released_to_gateway"
	QueryResponseSubmissionStatusReleasedToGateway QueryResponseSubmissionStatus = "released_to_gateway"

	// QueryResponseSubmissionStatusQueuedForDelivery captures enum value "queued_for_delivery"
	QueryResponseSubmissionStatusQueuedForDelivery QueryResponseSubmissionStatus = "queued_for_delivery"

	// QueryResponseSubmissionStatusSubmitted captures enum value "submitted"
	QueryResponseSubmissionStatusSubmitted QueryResponseSubmissionStatus = "submitted"

	// QueryResponseSubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	QueryResponseSubmissionStatusDeliveryConfirmed QueryResponseSubmissionStatus = "delivery_confirmed"

	// QueryResponseSubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	QueryResponseSubmissionStatusDeliveryFailed QueryResponseSubmissionStatus = "delivery_failed"
)

// for schema
var queryResponseSubmissionStatusEnum []interface{}

func init() {
	var res []QueryResponseSubmissionStatus
	if err := json.Unmarshal([]byte(`["accepted","validation_pending","released_to_gateway","queued_for_delivery","submitted","delivery_confirmed","delivery_failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryResponseSubmissionStatusEnum = append(queryResponseSubmissionStatusEnum, v)
	}
}

func (m QueryResponseSubmissionStatus) validateQueryResponseSubmissionStatusEnum(path, location string, value QueryResponseSubmissionStatus) error {
	if err := validate.Enum(path, location, value, queryResponseSubmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this query response submission status
func (m QueryResponseSubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateQueryResponseSubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponseSubmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
