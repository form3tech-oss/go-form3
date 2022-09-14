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
)

// SendersCorrespondentAccountHoldingEntity senders correspondent account holding entity
// swagger:model SendersCorrespondentAccountHoldingEntity
type SendersCorrespondentAccountHoldingEntity struct {

	// Sender's correspondent's address
	BankAddress []string `json:"bank_address,omitempty"`

	// SWIFT BIC for sender's correspondent
	BankID string `json:"bank_id,omitempty"`

	// bank id code
	BankIDCode BankIDCode `json:"bank_id_code,omitempty"`

	// Sender's correspondent's name
	BankName string `json:"bank_name,omitempty"`

	// Identifier of the financial institution which services the account
	BankPartyID string `json:"bank_party_id,omitempty"`
}

func SendersCorrespondentAccountHoldingEntityWithDefaults(defaults client.Defaults) *SendersCorrespondentAccountHoldingEntity {
	return &SendersCorrespondentAccountHoldingEntity{

		BankAddress: make([]string, 0),

		BankID: defaults.GetString("SendersCorrespondentAccountHoldingEntity", "bank_id"),

		// TODO BankIDCode: BankIDCode,

		BankName: defaults.GetString("SendersCorrespondentAccountHoldingEntity", "bank_name"),

		BankPartyID: defaults.GetString("SendersCorrespondentAccountHoldingEntity", "bank_party_id"),
	}
}

func (m *SendersCorrespondentAccountHoldingEntity) WithBankAddress(bankAddress []string) *SendersCorrespondentAccountHoldingEntity {

	m.BankAddress = bankAddress

	return m
}

func (m *SendersCorrespondentAccountHoldingEntity) WithBankID(bankID string) *SendersCorrespondentAccountHoldingEntity {

	m.BankID = bankID

	return m
}

func (m *SendersCorrespondentAccountHoldingEntity) WithBankIDCode(bankIDCode BankIDCode) *SendersCorrespondentAccountHoldingEntity {

	m.BankIDCode = bankIDCode

	return m
}

func (m *SendersCorrespondentAccountHoldingEntity) WithBankName(bankName string) *SendersCorrespondentAccountHoldingEntity {

	m.BankName = bankName

	return m
}

func (m *SendersCorrespondentAccountHoldingEntity) WithBankPartyID(bankPartyID string) *SendersCorrespondentAccountHoldingEntity {

	m.BankPartyID = bankPartyID

	return m
}

// Validate validates this senders correspondent account holding entity
func (m *SendersCorrespondentAccountHoldingEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendersCorrespondentAccountHoldingEntity) validateBankIDCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankIDCode) { // not required
		return nil
	}

	if err := m.BankIDCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bank_id_code")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendersCorrespondentAccountHoldingEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendersCorrespondentAccountHoldingEntity) UnmarshalBinary(b []byte) error {
	var res SendersCorrespondentAccountHoldingEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SendersCorrespondentAccountHoldingEntity) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
