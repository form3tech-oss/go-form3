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

// AccountClassification account classification
// swagger:model AccountClassification
type AccountClassification string

const (

	// AccountClassificationPersonal captures enum value "personal"
	AccountClassificationPersonal AccountClassification = "personal"

	// AccountClassificationBusiness captures enum value "business"
	AccountClassificationBusiness AccountClassification = "business"
)

// for schema
var accountClassificationEnum []interface{}

func init() {
	var res []AccountClassification
	if err := json.Unmarshal([]byte(`["personal","business"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountClassificationEnum = append(accountClassificationEnum, v)
	}
}

func (m AccountClassification) validateAccountClassificationEnum(path, location string, value AccountClassification) error {
	if err := validate.Enum(path, location, value, accountClassificationEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this account classification
func (m AccountClassification) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccountClassificationEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountClassification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
