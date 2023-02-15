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

// ReversalAdmissionStatus Status of the reversal admission
// swagger:model ReversalAdmissionStatus
type ReversalAdmissionStatus string

const (

	// ReversalAdmissionStatusPending captures enum value "pending"
	ReversalAdmissionStatusPending ReversalAdmissionStatus = "pending"

	// ReversalAdmissionStatusConfirmed captures enum value "confirmed"
	ReversalAdmissionStatusConfirmed ReversalAdmissionStatus = "confirmed"
)

// for schema
var reversalAdmissionStatusEnum []interface{}

func init() {
	var res []ReversalAdmissionStatus
	if err := json.Unmarshal([]byte(`["pending","confirmed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reversalAdmissionStatusEnum = append(reversalAdmissionStatusEnum, v)
	}
}

func (m ReversalAdmissionStatus) validateReversalAdmissionStatusEnum(path, location string, value ReversalAdmissionStatus) error {
	if err := validate.Enum(path, location, value, reversalAdmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this reversal admission status
func (m ReversalAdmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReversalAdmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalAdmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
