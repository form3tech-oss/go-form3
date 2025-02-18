// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FinCrimePaymentCheckAgentAccountWithBankID fin crime payment check agent account with bank Id
// swagger:model FinCrimePaymentCheckAgentAccountWithBankId
type FinCrimePaymentCheckAgentAccountWithBankID struct {

	// Other ID of agent
	// Max Length: 30
	BankID string `json:"bank_id,omitempty"`

	// Code for the type of other ID for agent
	// Max Length: 5
	BankIDCode string `json:"bank_id_code,omitempty"`
}

func FinCrimePaymentCheckAgentAccountWithBankIDWithDefaults(defaults client.Defaults) *FinCrimePaymentCheckAgentAccountWithBankID {
	return &FinCrimePaymentCheckAgentAccountWithBankID{

		BankID: defaults.GetString("FinCrimePaymentCheckAgentAccountWithBankId", "bank_id"),

		BankIDCode: defaults.GetString("FinCrimePaymentCheckAgentAccountWithBankId", "bank_id_code"),
	}
}

func (m *FinCrimePaymentCheckAgentAccountWithBankID) WithBankID(bankID string) *FinCrimePaymentCheckAgentAccountWithBankID {

	m.BankID = bankID

	return m
}

func (m *FinCrimePaymentCheckAgentAccountWithBankID) WithBankIDCode(bankIDCode string) *FinCrimePaymentCheckAgentAccountWithBankID {

	m.BankIDCode = bankIDCode

	return m
}

// Validate validates this fin crime payment check agent account with bank Id
func (m *FinCrimePaymentCheckAgentAccountWithBankID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FinCrimePaymentCheckAgentAccountWithBankID) validateBankID(formats strfmt.Registry) error {

	if swag.IsZero(m.BankID) { // not required
		return nil
	}

	if err := validate.MaxLength("bank_id", "body", m.BankID, 30); err != nil {
		return err
	}

	return nil
}

func (m *FinCrimePaymentCheckAgentAccountWithBankID) validateBankIDCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankIDCode) { // not required
		return nil
	}

	if err := validate.MaxLength("bank_id_code", "body", m.BankIDCode, 5); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FinCrimePaymentCheckAgentAccountWithBankID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FinCrimePaymentCheckAgentAccountWithBankID) UnmarshalBinary(b []byte) error {
	var res FinCrimePaymentCheckAgentAccountWithBankID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FinCrimePaymentCheckAgentAccountWithBankID) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
