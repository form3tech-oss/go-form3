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

// AccountNumberCode The type of identification given at `account_number` attribute
// swagger:model AccountNumberCode
type AccountNumberCode string

const (

	// AccountNumberCodeBBAN captures enum value "BBAN"
	AccountNumberCodeBBAN AccountNumberCode = "BBAN"

	// AccountNumberCodeIBAN captures enum value "IBAN"
	AccountNumberCodeIBAN AccountNumberCode = "IBAN"
)

// for schema
var accountNumberCodeEnum []interface{}

func init() {
	var res []AccountNumberCode
	if err := json.Unmarshal([]byte(`["BBAN","IBAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountNumberCodeEnum = append(accountNumberCodeEnum, v)
	}
}

func (m AccountNumberCode) validateAccountNumberCodeEnum(path, location string, value AccountNumberCode) error {
	if err := validate.Enum(path, location, value, accountNumberCodeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this account number code
func (m AccountNumberCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MaxLength("", "body", string(m), 4); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccountNumberCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountNumberCode) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
