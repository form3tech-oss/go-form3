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

// QueryAccountNumberCode The type of identification given at `account_number` attribute
// swagger:model QueryAccountNumberCode
type QueryAccountNumberCode string

const (

	// QueryAccountNumberCodeIBAN captures enum value "IBAN"
	QueryAccountNumberCodeIBAN QueryAccountNumberCode = "IBAN"

	// QueryAccountNumberCodeBBAN captures enum value "BBAN"
	QueryAccountNumberCodeBBAN QueryAccountNumberCode = "BBAN"
)

// for schema
var queryAccountNumberCodeEnum []interface{}

func init() {
	var res []QueryAccountNumberCode
	if err := json.Unmarshal([]byte(`["IBAN","BBAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryAccountNumberCodeEnum = append(queryAccountNumberCodeEnum, v)
	}
}

func (m QueryAccountNumberCode) validateQueryAccountNumberCodeEnum(path, location string, value QueryAccountNumberCode) error {
	if err := validate.Enum(path, location, value, queryAccountNumberCodeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this query account number code
func (m QueryAccountNumberCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MaxLength("", "body", string(m), 4); err != nil {
		return err
	}

	// value enum
	if err := m.validateQueryAccountNumberCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryAccountNumberCode) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
