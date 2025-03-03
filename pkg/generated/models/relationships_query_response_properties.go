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

// RelationshipsQueryResponseProperties relationships query response properties
// swagger:model RelationshipsQueryResponseProperties
type RelationshipsQueryResponseProperties struct {

	// The Query Response ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// type
	// Required: true
	Type *QueryResponseResourceType `json:"type"`
}

func RelationshipsQueryResponsePropertiesWithDefaults(defaults client.Defaults) *RelationshipsQueryResponseProperties {
	return &RelationshipsQueryResponseProperties{

		ID: defaults.GetStrfmtUUIDPtr("RelationshipsQueryResponseProperties", "id"),

		// TODO Type: QueryResponseResourceType,

	}
}

func (m *RelationshipsQueryResponseProperties) WithID(id strfmt.UUID) *RelationshipsQueryResponseProperties {

	m.ID = &id

	return m
}

func (m *RelationshipsQueryResponseProperties) WithoutID() *RelationshipsQueryResponseProperties {
	m.ID = nil
	return m
}

func (m *RelationshipsQueryResponseProperties) WithType(typeVar QueryResponseResourceType) *RelationshipsQueryResponseProperties {

	m.Type = &typeVar

	return m
}

func (m *RelationshipsQueryResponseProperties) WithoutType() *RelationshipsQueryResponseProperties {
	m.Type = nil
	return m
}

// Validate validates this relationships query response properties
func (m *RelationshipsQueryResponseProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
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

func (m *RelationshipsQueryResponseProperties) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RelationshipsQueryResponseProperties) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelationshipsQueryResponseProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipsQueryResponseProperties) UnmarshalBinary(b []byte) error {
	var res RelationshipsQueryResponseProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipsQueryResponseProperties) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
