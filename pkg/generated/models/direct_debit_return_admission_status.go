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

// DirectDebitReturnAdmissionStatus direct debit return admission status
// swagger:model DirectDebitReturnAdmissionStatus
type DirectDebitReturnAdmissionStatus string

const (

	// DirectDebitReturnAdmissionStatusConfirmed captures enum value "confirmed"
	DirectDebitReturnAdmissionStatusConfirmed DirectDebitReturnAdmissionStatus = "confirmed"

	// DirectDebitReturnAdmissionStatusPending captures enum value "pending"
	DirectDebitReturnAdmissionStatusPending DirectDebitReturnAdmissionStatus = "pending"

	// DirectDebitReturnAdmissionStatusFailed captures enum value "failed"
	DirectDebitReturnAdmissionStatusFailed DirectDebitReturnAdmissionStatus = "failed"

	// DirectDebitReturnAdmissionStatusLimitCheckPending captures enum value "limit_check_pending"
	DirectDebitReturnAdmissionStatusLimitCheckPending DirectDebitReturnAdmissionStatus = "limit_check_pending"

	// DirectDebitReturnAdmissionStatusLimitCheckPassed captures enum value "limit_check_passed"
	DirectDebitReturnAdmissionStatusLimitCheckPassed DirectDebitReturnAdmissionStatus = "limit_check_passed"
)

// for schema
var directDebitReturnAdmissionStatusEnum []interface{}

func init() {
	var res []DirectDebitReturnAdmissionStatus
	if err := json.Unmarshal([]byte(`["confirmed","pending","failed","limit_check_pending","limit_check_passed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directDebitReturnAdmissionStatusEnum = append(directDebitReturnAdmissionStatusEnum, v)
	}
}

func (m DirectDebitReturnAdmissionStatus) validateDirectDebitReturnAdmissionStatusEnum(path, location string, value DirectDebitReturnAdmissionStatus) error {
	if err := validate.Enum(path, location, value, directDebitReturnAdmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this direct debit return admission status
func (m DirectDebitReturnAdmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDirectDebitReturnAdmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnAdmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
