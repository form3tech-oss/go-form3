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

// ReturnSubmissionStatus [Status](http://draft-api-docs.form3.tech/api.html#enumerations-payment-submission-status) of the submission
// swagger:model ReturnSubmissionStatus
type ReturnSubmissionStatus string

const (

	// ReturnSubmissionStatusAccepted captures enum value "accepted"
	ReturnSubmissionStatusAccepted ReturnSubmissionStatus = "accepted"

	// ReturnSubmissionStatusLimitCheckPending captures enum value "limit_check_pending"
	ReturnSubmissionStatusLimitCheckPending ReturnSubmissionStatus = "limit_check_pending"

	// ReturnSubmissionStatusLimitCheckFailed captures enum value "limit_check_failed"
	ReturnSubmissionStatusLimitCheckFailed ReturnSubmissionStatus = "limit_check_failed"

	// ReturnSubmissionStatusLimitCheckPassed captures enum value "limit_check_passed"
	ReturnSubmissionStatusLimitCheckPassed ReturnSubmissionStatus = "limit_check_passed"

	// ReturnSubmissionStatusReleasedToGateway captures enum value "released_to_gateway"
	ReturnSubmissionStatusReleasedToGateway ReturnSubmissionStatus = "released_to_gateway"

	// ReturnSubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	ReturnSubmissionStatusDeliveryConfirmed ReturnSubmissionStatus = "delivery_confirmed"

	// ReturnSubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	ReturnSubmissionStatusDeliveryFailed ReturnSubmissionStatus = "delivery_failed"

	// ReturnSubmissionStatusSubmitted captures enum value "submitted"
	ReturnSubmissionStatusSubmitted ReturnSubmissionStatus = "submitted"

	// ReturnSubmissionStatusValidationPending captures enum value "validation_pending"
	ReturnSubmissionStatusValidationPending ReturnSubmissionStatus = "validation_pending"

	// ReturnSubmissionStatusQueuedForDelivery captures enum value "queued_for_delivery"
	ReturnSubmissionStatusQueuedForDelivery ReturnSubmissionStatus = "queued_for_delivery"
)

// for schema
var returnSubmissionStatusEnum []interface{}

func init() {
	var res []ReturnSubmissionStatus
	if err := json.Unmarshal([]byte(`["accepted","limit_check_pending","limit_check_failed","limit_check_passed","released_to_gateway","delivery_confirmed","delivery_failed","submitted","validation_pending","queued_for_delivery"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnSubmissionStatusEnum = append(returnSubmissionStatusEnum, v)
	}
}

func (m ReturnSubmissionStatus) validateReturnSubmissionStatusEnum(path, location string, value ReturnSubmissionStatus) error {
	if err := validate.Enum(path, location, value, returnSubmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this return submission status
func (m ReturnSubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReturnSubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnSubmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
