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

// AccountNumber account number
// swagger:model AccountNumber
type AccountNumber struct {

	// id
	ID string `json:"id,omitempty"`

	// relationships
	Relationships *AccountNumberRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`
}

func AccountNumberWithDefaults(defaults client.Defaults) *AccountNumber {
	return &AccountNumber{

		ID: defaults.GetString("AccountNumber", "id"),

		Relationships: AccountNumberRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("AccountNumber", "type"),
	}
}

func (m *AccountNumber) WithID(id string) *AccountNumber {

	m.ID = id

	return m
}

func (m *AccountNumber) WithRelationships(relationships AccountNumberRelationships) *AccountNumber {

	m.Relationships = &relationships

	return m
}

func (m *AccountNumber) WithoutRelationships() *AccountNumber {
	m.Relationships = nil
	return m
}

func (m *AccountNumber) WithType(typeVar string) *AccountNumber {

	m.Type = typeVar

	return m
}

// Validate validates this account number
func (m *AccountNumber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountNumber) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

func (m *AccountNumber) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountNumber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountNumber) UnmarshalBinary(b []byte) error {
	var res AccountNumber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountNumber) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AccountNumberRelationships account number relationships
// swagger:model AccountNumberRelationships
type AccountNumberRelationships struct {

	// sort code
	SortCode *SortCode `json:"sort_code,omitempty"`
}

func AccountNumberRelationshipsWithDefaults(defaults client.Defaults) *AccountNumberRelationships {
	return &AccountNumberRelationships{

		SortCode: SortCodeWithDefaults(defaults),
	}
}

func (m *AccountNumberRelationships) WithSortCode(sortCode SortCode) *AccountNumberRelationships {

	m.SortCode = &sortCode

	return m
}

func (m *AccountNumberRelationships) WithoutSortCode() *AccountNumberRelationships {
	m.SortCode = nil
	return m
}

// Validate validates this account number relationships
func (m *AccountNumberRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSortCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountNumberRelationships) validateSortCode(formats strfmt.Registry) error {

	if swag.IsZero(m.SortCode) { // not required
		return nil
	}

	if m.SortCode != nil {
		if err := m.SortCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "sort_code")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountNumberRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountNumberRelationships) UnmarshalBinary(b []byte) error {
	var res AccountNumberRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountNumberRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
