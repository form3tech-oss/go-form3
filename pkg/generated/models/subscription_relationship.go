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

// SubscriptionRelationship subscription relationship
// swagger:model SubscriptionRelationship
type SubscriptionRelationship struct {

	// data
	// Read Only: true
	Data []SubscriptionRelationshipDataItems0 `json:"data"`
}

func SubscriptionRelationshipWithDefaults(defaults client.Defaults) *SubscriptionRelationship {
	return &SubscriptionRelationship{

		Data: make([]SubscriptionRelationshipDataItems0, 0),
	}
}

func (m *SubscriptionRelationship) WithData(data []SubscriptionRelationshipDataItems0) *SubscriptionRelationship {

	m.Data = data

	return m
}

// Validate validates this subscription relationship
func (m *SubscriptionRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionRelationship) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {

		if err := m.Data[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionRelationship) UnmarshalBinary(b []byte) error {
	var res SubscriptionRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SubscriptionRelationship) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// SubscriptionRelationshipDataItems0 subscription relationship data items0
// swagger:model SubscriptionRelationshipDataItems0
type SubscriptionRelationshipDataItems0 struct {

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// type
	// Read Only: true
	Type string `json:"type,omitempty"`
}

func SubscriptionRelationshipDataItems0WithDefaults(defaults client.Defaults) *SubscriptionRelationshipDataItems0 {
	return &SubscriptionRelationshipDataItems0{

		ID: defaults.GetStrfmtUUID("SubscriptionRelationshipDataItems0", "id"),

		Type: defaults.GetString("SubscriptionRelationshipDataItems0", "type"),
	}
}

func (m *SubscriptionRelationshipDataItems0) WithID(id strfmt.UUID) *SubscriptionRelationshipDataItems0 {

	m.ID = id

	return m
}

func (m *SubscriptionRelationshipDataItems0) WithType(typeVar string) *SubscriptionRelationshipDataItems0 {

	m.Type = typeVar

	return m
}

// Validate validates this subscription relationship data items0
func (m *SubscriptionRelationshipDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionRelationshipDataItems0) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionRelationshipDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionRelationshipDataItems0) UnmarshalBinary(b []byte) error {
	var res SubscriptionRelationshipDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SubscriptionRelationshipDataItems0) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
