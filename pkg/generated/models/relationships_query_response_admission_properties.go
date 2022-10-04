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

// RelationshipsQueryResponseAdmissionProperties relationships query response admission properties
// swagger:model RelationshipsQueryResponseAdmissionProperties
type RelationshipsQueryResponseAdmissionProperties struct {

	// The response admission for this query
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// type
	// Required: true
	Type *QueryResponseAdmissionResourceType `json:"type"`
}

func RelationshipsQueryResponseAdmissionPropertiesWithDefaults(defaults client.Defaults) *RelationshipsQueryResponseAdmissionProperties {
	return &RelationshipsQueryResponseAdmissionProperties{

		ID: defaults.GetStrfmtUUIDPtr("RelationshipsQueryResponseAdmissionProperties", "id"),

		// TODO Type: QueryResponseAdmissionResourceType,

	}
}

func (m *RelationshipsQueryResponseAdmissionProperties) WithID(id strfmt.UUID) *RelationshipsQueryResponseAdmissionProperties {

	m.ID = &id

	return m
}

func (m *RelationshipsQueryResponseAdmissionProperties) WithoutID() *RelationshipsQueryResponseAdmissionProperties {
	m.ID = nil
	return m
}

func (m *RelationshipsQueryResponseAdmissionProperties) WithType(typeVar QueryResponseAdmissionResourceType) *RelationshipsQueryResponseAdmissionProperties {

	m.Type = &typeVar

	return m
}

func (m *RelationshipsQueryResponseAdmissionProperties) WithoutType() *RelationshipsQueryResponseAdmissionProperties {
	m.Type = nil
	return m
}

// Validate validates this relationships query response admission properties
func (m *RelationshipsQueryResponseAdmissionProperties) Validate(formats strfmt.Registry) error {
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

func (m *RelationshipsQueryResponseAdmissionProperties) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RelationshipsQueryResponseAdmissionProperties) validateType(formats strfmt.Registry) error {

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
func (m *RelationshipsQueryResponseAdmissionProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipsQueryResponseAdmissionProperties) UnmarshalBinary(b []byte) error {
	var res RelationshipsQueryResponseAdmissionProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipsQueryResponseAdmissionProperties) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
