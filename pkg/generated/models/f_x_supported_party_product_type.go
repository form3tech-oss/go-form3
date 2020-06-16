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

// FXSupportedPartyProductType f x supported party product type
// swagger:model FXSupportedPartyProductType
type FXSupportedPartyProductType string

const (

	// FXSupportedPartyProductTypeInternationalPaymentsEbury captures enum value "international_payments_ebury"
	FXSupportedPartyProductTypeInternationalPaymentsEbury FXSupportedPartyProductType = "international_payments_ebury"
)

// for schema
var fXSupportedPartyProductTypeEnum []interface{}

func init() {
	var res []FXSupportedPartyProductType
	if err := json.Unmarshal([]byte(`["international_payments_ebury"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fXSupportedPartyProductTypeEnum = append(fXSupportedPartyProductTypeEnum, v)
	}
}

func (m FXSupportedPartyProductType) validateFXSupportedPartyProductTypeEnum(path, location string, value FXSupportedPartyProductType) error {
	if err := validate.Enum(path, location, value, fXSupportedPartyProductTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this f x supported party product type
func (m FXSupportedPartyProductType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFXSupportedPartyProductTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FXSupportedPartyProductType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}