// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Agent agent
// swagger:model Agent
type Agent struct {

	// Account number of the financial institution. Can be BBAN or IBAN.
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// address
	// Max Items: 4
	Address []string `json:"address,omitempty"`

	// identification
	Identification *AgentIdentification `json:"identification,omitempty"`

	// Name by which the financial institution is known
	Name string `json:"name,omitempty"`

	// role
	Role AgentRole `json:"role,omitempty"`
}

func AgentWithDefaults(defaults client.Defaults) *Agent {
	return &Agent{

		AccountNumber: defaults.GetString("Agent", "account_number"),

		// TODO AccountNumberCode: AccountNumberCode,

		Address: make([]string, 0),

		Identification: AgentIdentificationWithDefaults(defaults),

		Name: defaults.GetString("Agent", "name"),

		// TODO Role: AgentRole,

	}
}

func (m *Agent) WithAccountNumber(accountNumber string) *Agent {

	m.AccountNumber = accountNumber

	return m
}

func (m *Agent) WithAccountNumberCode(accountNumberCode AccountNumberCode) *Agent {

	m.AccountNumberCode = accountNumberCode

	return m
}

func (m *Agent) WithAddress(address []string) *Agent {

	m.Address = address

	return m
}

func (m *Agent) WithIdentification(identification AgentIdentification) *Agent {

	m.Identification = &identification

	return m
}

func (m *Agent) WithoutIdentification() *Agent {
	m.Identification = nil
	return m
}

func (m *Agent) WithName(name string) *Agent {

	m.Name = name

	return m
}

func (m *Agent) WithRole(role AgentRole) *Agent {

	m.Role = role

	return m
}

// Validate validates this agent
func (m *Agent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Agent) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	if err := m.AccountNumberCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("account_number_code")
		}
		return err
	}

	return nil
}

func (m *Agent) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	iAddressSize := int64(len(m.Address))

	if err := validate.MaxItems("address", "body", iAddressSize, 4); err != nil {
		return err
	}

	return nil
}

func (m *Agent) validateIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.Identification) { // not required
		return nil
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identification")
			}
			return err
		}
	}

	return nil
}

func (m *Agent) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if err := m.Role.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("role")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Agent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Agent) UnmarshalBinary(b []byte) error {
	var res Agent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Agent) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AgentIdentification agent identification
// swagger:model AgentIdentification
type AgentIdentification struct {

	// Identification code of the financial institution.
	// Max Length: 30
	BankID string `json:"bank_id,omitempty"`

	// Type of identification provided in bank_id field. Required when bank_id is provided, not used otherwise.
	// Max Length: 5
	BankIDCode string `json:"bank_id_code,omitempty"`

	// Array for additional ID(s) for agent
	BankIds []*AccountWithBankID `json:"bank_ids,omitempty"`
}

func AgentIdentificationWithDefaults(defaults client.Defaults) *AgentIdentification {
	return &AgentIdentification{

		BankID: defaults.GetString("AgentIdentification", "bank_id"),

		BankIDCode: defaults.GetString("AgentIdentification", "bank_id_code"),

		BankIds: make([]*AccountWithBankID, 0),
	}
}

func (m *AgentIdentification) WithBankID(bankID string) *AgentIdentification {

	m.BankID = bankID

	return m
}

func (m *AgentIdentification) WithBankIDCode(bankIDCode string) *AgentIdentification {

	m.BankIDCode = bankIDCode

	return m
}

func (m *AgentIdentification) WithBankIds(bankIds []*AccountWithBankID) *AgentIdentification {

	m.BankIds = bankIds

	return m
}

// Validate validates this agent identification
func (m *AgentIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentIdentification) validateBankID(formats strfmt.Registry) error {

	if swag.IsZero(m.BankID) { // not required
		return nil
	}

	if err := validate.MaxLength("identification"+"."+"bank_id", "body", m.BankID, 30); err != nil {
		return err
	}

	return nil
}

func (m *AgentIdentification) validateBankIDCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankIDCode) { // not required
		return nil
	}

	if err := validate.MaxLength("identification"+"."+"bank_id_code", "body", m.BankIDCode, 5); err != nil {
		return err
	}

	return nil
}

func (m *AgentIdentification) validateBankIds(formats strfmt.Registry) error {

	if swag.IsZero(m.BankIds) { // not required
		return nil
	}

	for i := 0; i < len(m.BankIds); i++ {
		if swag.IsZero(m.BankIds[i]) { // not required
			continue
		}

		if m.BankIds[i] != nil {
			if err := m.BankIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("identification" + "." + "bank_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgentIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentIdentification) UnmarshalBinary(b []byte) error {
	var res AgentIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AgentIdentification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
