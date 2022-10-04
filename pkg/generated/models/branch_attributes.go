// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BranchAttributes branch attributes
// swagger:model BranchAttributes
type BranchAttributes struct {

	// acceptance qualifier
	AcceptanceQualifier AcceptanceQualifier `json:"acceptance_qualifier,omitempty"`

	// Local country bank identifier. In the UK this is the sort code.
	// Required: true
	// Pattern: ^[A-Z0-9]{1,11}$
	BankID *string `json:"bank_id"`

	// ISO 20022 code used to identify the type of bank ID being used
	// Required: true
	// Pattern: ^[A-Z]{0,16}$
	BankIDCode *string `json:"bank_id_code"`

	// if present – has effect of making secondary reference in payment mandatory
	ReferenceMask string `json:"reference_mask,omitempty"`

	// validation type
	ValidationType BranchValidationType `json:"validation_type,omitempty"`
}

func BranchAttributesWithDefaults(defaults client.Defaults) *BranchAttributes {
	return &BranchAttributes{

		// TODO AcceptanceQualifier: AcceptanceQualifier,

		BankID: defaults.GetStringPtr("BranchAttributes", "bank_id"),

		BankIDCode: defaults.GetStringPtr("BranchAttributes", "bank_id_code"),

		ReferenceMask: defaults.GetString("BranchAttributes", "reference_mask"),

		// TODO ValidationType: BranchValidationType,

	}
}

func (m *BranchAttributes) WithAcceptanceQualifier(acceptanceQualifier AcceptanceQualifier) *BranchAttributes {

	m.AcceptanceQualifier = acceptanceQualifier

	return m
}

func (m *BranchAttributes) WithBankID(bankID string) *BranchAttributes {

	m.BankID = &bankID

	return m
}

func (m *BranchAttributes) WithoutBankID() *BranchAttributes {
	m.BankID = nil
	return m
}

func (m *BranchAttributes) WithBankIDCode(bankIDCode string) *BranchAttributes {

	m.BankIDCode = &bankIDCode

	return m
}

func (m *BranchAttributes) WithoutBankIDCode() *BranchAttributes {
	m.BankIDCode = nil
	return m
}

func (m *BranchAttributes) WithReferenceMask(referenceMask string) *BranchAttributes {

	m.ReferenceMask = referenceMask

	return m
}

func (m *BranchAttributes) WithValidationType(validationType BranchValidationType) *BranchAttributes {

	m.ValidationType = validationType

	return m
}

// Validate validates this branch attributes
func (m *BranchAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptanceQualifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchAttributes) validateAcceptanceQualifier(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptanceQualifier) { // not required
		return nil
	}

	if err := m.AcceptanceQualifier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("acceptance_qualifier")
		}
		return err
	}

	return nil
}

func (m *BranchAttributes) validateBankID(formats strfmt.Registry) error {

	if err := validate.Required("bank_id", "body", m.BankID); err != nil {
		return err
	}

	if err := validate.Pattern("bank_id", "body", string(*m.BankID), `^[A-Z0-9]{1,11}$`); err != nil {
		return err
	}

	return nil
}

func (m *BranchAttributes) validateBankIDCode(formats strfmt.Registry) error {

	if err := validate.Required("bank_id_code", "body", m.BankIDCode); err != nil {
		return err
	}

	if err := validate.Pattern("bank_id_code", "body", string(*m.BankIDCode), `^[A-Z]{0,16}$`); err != nil {
		return err
	}

	return nil
}

func (m *BranchAttributes) validateValidationType(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationType) { // not required
		return nil
	}

	if err := m.ValidationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("validation_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchAttributes) UnmarshalBinary(b []byte) error {
	var res BranchAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *BranchAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
