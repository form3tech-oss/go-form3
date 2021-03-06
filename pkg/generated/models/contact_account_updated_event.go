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

// ContactAccountUpdatedEvent contact account updated event
// swagger:model ContactAccountUpdatedEvent
type ContactAccountUpdatedEvent struct {

	// data
	Data *ContactAccount `json:"data,omitempty"`
}

func ContactAccountUpdatedEventWithDefaults(defaults client.Defaults) *ContactAccountUpdatedEvent {
	return &ContactAccountUpdatedEvent{

		Data: ContactAccountWithDefaults(defaults),
	}
}

func (m *ContactAccountUpdatedEvent) WithData(data ContactAccount) *ContactAccountUpdatedEvent {

	m.Data = &data

	return m
}

func (m *ContactAccountUpdatedEvent) WithoutData() *ContactAccountUpdatedEvent {
	m.Data = nil
	return m
}

// Validate validates this contact account updated event
func (m *ContactAccountUpdatedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactAccountUpdatedEvent) validateData(formats strfmt.Registry) error {

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
func (m *ContactAccountUpdatedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactAccountUpdatedEvent) UnmarshalBinary(b []byte) error {
	var res ContactAccountUpdatedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ContactAccountUpdatedEvent) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
