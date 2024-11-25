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

// PaymentSubmissionTaskAssignee Helps to identify the owner of the task
// swagger:model PaymentSubmissionTaskAssignee
type PaymentSubmissionTaskAssignee string

const (

	// PaymentSubmissionTaskAssigneeCustomer captures enum value "customer"
	PaymentSubmissionTaskAssigneeCustomer PaymentSubmissionTaskAssignee = "customer"

	// PaymentSubmissionTaskAssigneeForm3 captures enum value "form3"
	PaymentSubmissionTaskAssigneeForm3 PaymentSubmissionTaskAssignee = "form3"
)

// for schema
var paymentSubmissionTaskAssigneeEnum []interface{}

func init() {
	var res []PaymentSubmissionTaskAssignee
	if err := json.Unmarshal([]byte(`["customer","form3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentSubmissionTaskAssigneeEnum = append(paymentSubmissionTaskAssigneeEnum, v)
	}
}

func (m PaymentSubmissionTaskAssignee) validatePaymentSubmissionTaskAssigneeEnum(path, location string, value PaymentSubmissionTaskAssignee) error {
	if err := validate.Enum(path, location, value, paymentSubmissionTaskAssigneeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment submission task assignee
func (m PaymentSubmissionTaskAssignee) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentSubmissionTaskAssigneeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentSubmissionTaskAssignee) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
