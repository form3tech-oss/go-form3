// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ReimbursementAccountHoldingEntity reimbursement account holding entity
// swagger:model ReimbursementAccountHoldingEntity
type ReimbursementAccountHoldingEntity struct {

	// Third party reimbursement institution address
	BankAddress []string `json:"bank_address,omitempty"`

	// Identification of third party reimbursement institution
	BankID string `json:"bank_id,omitempty"`

	// The type of identification provided at `bank_id` attribute. Must be ISO code as listed in the [External Code Sets spreadsheet](https://www.iso20022.org/external_code_list.page)
	BankIDCode string `json:"bank_id_code,omitempty"`

	// Third party reimbursement institution name
	BankName string `json:"bank_name,omitempty"`

	// Third party reimbursement institution identifier
	BankPartyID string `json:"bank_party_id,omitempty"`
}

func ReimbursementAccountHoldingEntityWithDefaults(defaults client.Defaults) *ReimbursementAccountHoldingEntity {
	return &ReimbursementAccountHoldingEntity{

		BankAddress: make([]string, 0),

		BankID: defaults.GetString("ReimbursementAccountHoldingEntity", "bank_id"),

		BankIDCode: defaults.GetString("ReimbursementAccountHoldingEntity", "bank_id_code"),

		BankName: defaults.GetString("ReimbursementAccountHoldingEntity", "bank_name"),

		BankPartyID: defaults.GetString("ReimbursementAccountHoldingEntity", "bank_party_id"),
	}
}

func (m *ReimbursementAccountHoldingEntity) WithBankAddress(bankAddress []string) *ReimbursementAccountHoldingEntity {

	m.BankAddress = bankAddress

	return m
}

func (m *ReimbursementAccountHoldingEntity) WithBankID(bankID string) *ReimbursementAccountHoldingEntity {

	m.BankID = bankID

	return m
}

func (m *ReimbursementAccountHoldingEntity) WithBankIDCode(bankIDCode string) *ReimbursementAccountHoldingEntity {

	m.BankIDCode = bankIDCode

	return m
}

func (m *ReimbursementAccountHoldingEntity) WithBankName(bankName string) *ReimbursementAccountHoldingEntity {

	m.BankName = bankName

	return m
}

func (m *ReimbursementAccountHoldingEntity) WithBankPartyID(bankPartyID string) *ReimbursementAccountHoldingEntity {

	m.BankPartyID = bankPartyID

	return m
}

// Validate validates this reimbursement account holding entity
func (m *ReimbursementAccountHoldingEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReimbursementAccountHoldingEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReimbursementAccountHoldingEntity) UnmarshalBinary(b []byte) error {
	var res ReimbursementAccountHoldingEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReimbursementAccountHoldingEntity) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
