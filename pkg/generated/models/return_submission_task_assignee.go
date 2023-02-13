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

// ReturnSubmissionTaskAssignee Helps to identify the owner of the task
// swagger:model ReturnSubmissionTaskAssignee
type ReturnSubmissionTaskAssignee string

const (

	// ReturnSubmissionTaskAssigneeCustomer captures enum value "customer"
	ReturnSubmissionTaskAssigneeCustomer ReturnSubmissionTaskAssignee = "customer"

	// ReturnSubmissionTaskAssigneeForm3 captures enum value "form3"
	ReturnSubmissionTaskAssigneeForm3 ReturnSubmissionTaskAssignee = "form3"
)

// for schema
var returnSubmissionTaskAssigneeEnum []interface{}

func init() {
	var res []ReturnSubmissionTaskAssignee
	if err := json.Unmarshal([]byte(`["customer","form3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnSubmissionTaskAssigneeEnum = append(returnSubmissionTaskAssigneeEnum, v)
	}
}

func (m ReturnSubmissionTaskAssignee) validateReturnSubmissionTaskAssigneeEnum(path, location string, value ReturnSubmissionTaskAssignee) error {
	if err := validate.Enum(path, location, value, returnSubmissionTaskAssigneeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this return submission task assignee
func (m ReturnSubmissionTaskAssignee) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReturnSubmissionTaskAssigneeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnSubmissionTaskAssignee) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
