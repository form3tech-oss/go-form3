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

// ReturnAdmissionTaskAssignee Helps to identify the owner of the task
// swagger:model ReturnAdmissionTaskAssignee
type ReturnAdmissionTaskAssignee string

const (

	// ReturnAdmissionTaskAssigneeCustomer captures enum value "customer"
	ReturnAdmissionTaskAssigneeCustomer ReturnAdmissionTaskAssignee = "customer"

	// ReturnAdmissionTaskAssigneeForm3 captures enum value "form3"
	ReturnAdmissionTaskAssigneeForm3 ReturnAdmissionTaskAssignee = "form3"
)

// for schema
var returnAdmissionTaskAssigneeEnum []interface{}

func init() {
	var res []ReturnAdmissionTaskAssignee
	if err := json.Unmarshal([]byte(`["customer","form3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnAdmissionTaskAssigneeEnum = append(returnAdmissionTaskAssigneeEnum, v)
	}
}

func (m ReturnAdmissionTaskAssignee) validateReturnAdmissionTaskAssigneeEnum(path, location string, value ReturnAdmissionTaskAssignee) error {
	if err := validate.Enum(path, location, value, returnAdmissionTaskAssigneeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this return admission task assignee
func (m ReturnAdmissionTaskAssignee) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReturnAdmissionTaskAssigneeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnAdmissionTaskAssignee) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
