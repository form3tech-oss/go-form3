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

// PaymentSubmissionTaskStatus status of the submission task
// swagger:model PaymentSubmissionTaskStatus
type PaymentSubmissionTaskStatus string

const (

	// PaymentSubmissionTaskStatusCompleted captures enum value "completed"
	PaymentSubmissionTaskStatusCompleted PaymentSubmissionTaskStatus = "completed"

	// PaymentSubmissionTaskStatusFailed captures enum value "failed"
	PaymentSubmissionTaskStatusFailed PaymentSubmissionTaskStatus = "failed"

	// PaymentSubmissionTaskStatusPending captures enum value "pending"
	PaymentSubmissionTaskStatusPending PaymentSubmissionTaskStatus = "pending"

	// PaymentSubmissionTaskStatusOnHold captures enum value "on_hold"
	PaymentSubmissionTaskStatusOnHold PaymentSubmissionTaskStatus = "on_hold"
)

// for schema
var paymentSubmissionTaskStatusEnum []interface{}

func init() {
	var res []PaymentSubmissionTaskStatus
	if err := json.Unmarshal([]byte(`["completed","failed","pending","on_hold"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentSubmissionTaskStatusEnum = append(paymentSubmissionTaskStatusEnum, v)
	}
}

func (m PaymentSubmissionTaskStatus) validatePaymentSubmissionTaskStatusEnum(path, location string, value PaymentSubmissionTaskStatus) error {
	if err := validate.Enum(path, location, value, paymentSubmissionTaskStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment submission task status
func (m PaymentSubmissionTaskStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentSubmissionTaskStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentSubmissionTaskStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
