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

// BankIDCode bank Id code
// swagger:model BankIdCode
type BankIDCode string

const (

	// BankIDCodeSWBIC captures enum value "SWBIC"
	BankIDCodeSWBIC BankIDCode = "SWBIC"

	// BankIDCodeGBDSC captures enum value "GBDSC"
	BankIDCodeGBDSC BankIDCode = "GBDSC"

	// BankIDCodeBE captures enum value "BE"
	BankIDCodeBE BankIDCode = "BE"

	// BankIDCodeFR captures enum value "FR"
	BankIDCodeFR BankIDCode = "FR"

	// BankIDCodeDEBLZ captures enum value "DEBLZ"
	BankIDCodeDEBLZ BankIDCode = "DEBLZ"

	// BankIDCodeGRBIC captures enum value "GRBIC"
	BankIDCodeGRBIC BankIDCode = "GRBIC"

	// BankIDCodeITNCC captures enum value "ITNCC"
	BankIDCodeITNCC BankIDCode = "ITNCC"

	// BankIDCodePLKNR captures enum value "PLKNR"
	BankIDCodePLKNR BankIDCode = "PLKNR"

	// BankIDCodePTNCC captures enum value "PTNCC"
	BankIDCodePTNCC BankIDCode = "PTNCC"

	// BankIDCodeESNCC captures enum value "ESNCC"
	BankIDCodeESNCC BankIDCode = "ESNCC"

	// BankIDCodeCHBCC captures enum value "CHBCC"
	BankIDCodeCHBCC BankIDCode = "CHBCC"
)

// for schema
var bankIdCodeEnum []interface{}

func init() {
	var res []BankIDCode
	if err := json.Unmarshal([]byte(`["SWBIC","GBDSC","BE","FR","DEBLZ","GRBIC","ITNCC","PLKNR","PTNCC","ESNCC","CHBCC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bankIdCodeEnum = append(bankIdCodeEnum, v)
	}
}

func (m BankIDCode) validateBankIDCodeEnum(path, location string, value BankIDCode) error {
	if err := validate.Enum(path, location, value, bankIdCodeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this bank Id code
func (m BankIDCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBankIDCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankIDCode) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
