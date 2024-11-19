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

// ReturnAdmissionTaskStatus status of the return admission task
// swagger:model ReturnAdmissionTaskStatus
type ReturnAdmissionTaskStatus string

const (

	// ReturnAdmissionTaskStatusCompleted captures enum value "completed"
	ReturnAdmissionTaskStatusCompleted ReturnAdmissionTaskStatus = "completed"

	// ReturnAdmissionTaskStatusFailed captures enum value "failed"
	ReturnAdmissionTaskStatusFailed ReturnAdmissionTaskStatus = "failed"

	// ReturnAdmissionTaskStatusPending captures enum value "pending"
	ReturnAdmissionTaskStatusPending ReturnAdmissionTaskStatus = "pending"

	// ReturnAdmissionTaskStatusOnHold captures enum value "on_hold"
	ReturnAdmissionTaskStatusOnHold ReturnAdmissionTaskStatus = "on_hold"
)

// for schema
var returnAdmissionTaskStatusEnum []interface{}

func init() {
	var res []ReturnAdmissionTaskStatus
	if err := json.Unmarshal([]byte(`["completed","failed","pending","on_hold"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnAdmissionTaskStatusEnum = append(returnAdmissionTaskStatusEnum, v)
	}
}

func (m ReturnAdmissionTaskStatus) validateReturnAdmissionTaskStatusEnum(path, location string, value ReturnAdmissionTaskStatus) error {
	if err := validate.Enum(path, location, value, returnAdmissionTaskStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this return admission task status
func (m ReturnAdmissionTaskStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReturnAdmissionTaskStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnAdmissionTaskStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
