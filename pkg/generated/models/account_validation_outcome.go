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

// AccountValidationOutcome Outcome of the validation of the beneficiary's account number, which is sometimes performed on a customer's behalf
// swagger:model AccountValidationOutcome
type AccountValidationOutcome string

const (

	// AccountValidationOutcomePassed captures enum value "passed"
	AccountValidationOutcomePassed AccountValidationOutcome = "passed"

	// AccountValidationOutcomeFailed captures enum value "failed"
	AccountValidationOutcomeFailed AccountValidationOutcome = "failed"

	// AccountValidationOutcomeFailedAutoRejectDisabled captures enum value "failed_auto_reject_disabled"
	AccountValidationOutcomeFailedAutoRejectDisabled AccountValidationOutcome = "failed_auto_reject_disabled"

	// AccountValidationOutcomeFailedAutoRejectEnabled captures enum value "failed_auto_reject_enabled"
	AccountValidationOutcomeFailedAutoRejectEnabled AccountValidationOutcome = "failed_auto_reject_enabled"
)

// for schema
var accountValidationOutcomeEnum []interface{}

func init() {
	var res []AccountValidationOutcome
	if err := json.Unmarshal([]byte(`["passed","failed","failed_auto_reject_disabled","failed_auto_reject_enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountValidationOutcomeEnum = append(accountValidationOutcomeEnum, v)
	}
}

func (m AccountValidationOutcome) validateAccountValidationOutcomeEnum(path, location string, value AccountValidationOutcome) error {
	if err := validate.Enum(path, location, value, accountValidationOutcomeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this account validation outcome
func (m AccountValidationOutcome) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccountValidationOutcomeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountValidationOutcome) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}