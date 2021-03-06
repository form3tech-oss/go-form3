// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PartyAccountRecord party account record
// swagger:model PartyAccountRecord
type PartyAccountRecord struct {

	// data
	Data *PartyAccount `json:"data,omitempty"`
}

func PartyAccountRecordWithDefaults(defaults client.Defaults) *PartyAccountRecord {
	return &PartyAccountRecord{

		Data: PartyAccountWithDefaults(defaults),
	}
}

func (m *PartyAccountRecord) WithData(data PartyAccount) *PartyAccountRecord {

	m.Data = &data

	return m
}

func (m *PartyAccountRecord) WithoutData() *PartyAccountRecord {
	m.Data = nil
	return m
}

// Validate validates this party account record
func (m *PartyAccountRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyAccountRecord) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyAccountRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyAccountRecord) UnmarshalBinary(b []byte) error {
	var res PartyAccountRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyAccountRecord) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
