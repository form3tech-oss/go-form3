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

// ReturnSubmissionTaskStatus status of the return submission task
// swagger:model ReturnSubmissionTaskStatus
type ReturnSubmissionTaskStatus string

const (

	// ReturnSubmissionTaskStatusCompleted captures enum value "completed"
	ReturnSubmissionTaskStatusCompleted ReturnSubmissionTaskStatus = "completed"

	// ReturnSubmissionTaskStatusFailed captures enum value "failed"
	ReturnSubmissionTaskStatusFailed ReturnSubmissionTaskStatus = "failed"

	// ReturnSubmissionTaskStatusPending captures enum value "pending"
	ReturnSubmissionTaskStatusPending ReturnSubmissionTaskStatus = "pending"

	// ReturnSubmissionTaskStatusOnHold captures enum value "on_hold"
	ReturnSubmissionTaskStatusOnHold ReturnSubmissionTaskStatus = "on_hold"
)

// for schema
var returnSubmissionTaskStatusEnum []interface{}

func init() {
	var res []ReturnSubmissionTaskStatus
	if err := json.Unmarshal([]byte(`["completed","failed","pending","on_hold"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnSubmissionTaskStatusEnum = append(returnSubmissionTaskStatusEnum, v)
	}
}

func (m ReturnSubmissionTaskStatus) validateReturnSubmissionTaskStatusEnum(path, location string, value ReturnSubmissionTaskStatus) error {
	if err := validate.Enum(path, location, value, returnSubmissionTaskStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this return submission task status
func (m ReturnSubmissionTaskStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReturnSubmissionTaskStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnSubmissionTaskStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
