// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SchemeMessageAdmission scheme message admission
// swagger:model SchemeMessageAdmission
type SchemeMessageAdmission struct {

	// attributes
	Attributes *SchemeMessageAdmissionAttributes `json:"attributes,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *SchemeMessageAdmissionRelationships `json:"relationships,omitempty"`

	// type
	// Enum: [scheme_message_admissions]
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func SchemeMessageAdmissionWithDefaults(defaults client.Defaults) *SchemeMessageAdmission {
	return &SchemeMessageAdmission{

		Attributes: SchemeMessageAdmissionAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("SchemeMessageAdmission", "id"),

		OrganisationID: defaults.GetStrfmtUUID("SchemeMessageAdmission", "organisation_id"),

		Relationships: SchemeMessageAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("SchemeMessageAdmission", "type"),

		Version: defaults.GetInt64Ptr("SchemeMessageAdmission", "version"),
	}
}

func (m *SchemeMessageAdmission) WithAttributes(attributes SchemeMessageAdmissionAttributes) *SchemeMessageAdmission {

	m.Attributes = &attributes

	return m
}

func (m *SchemeMessageAdmission) WithoutAttributes() *SchemeMessageAdmission {
	m.Attributes = nil
	return m
}

func (m *SchemeMessageAdmission) WithID(id strfmt.UUID) *SchemeMessageAdmission {

	m.ID = id

	return m
}

func (m *SchemeMessageAdmission) WithOrganisationID(organisationID strfmt.UUID) *SchemeMessageAdmission {

	m.OrganisationID = organisationID

	return m
}

func (m *SchemeMessageAdmission) WithRelationships(relationships SchemeMessageAdmissionRelationships) *SchemeMessageAdmission {

	m.Relationships = &relationships

	return m
}

func (m *SchemeMessageAdmission) WithoutRelationships() *SchemeMessageAdmission {
	m.Relationships = nil
	return m
}

func (m *SchemeMessageAdmission) WithType(typeVar string) *SchemeMessageAdmission {

	m.Type = typeVar

	return m
}

func (m *SchemeMessageAdmission) WithVersion(version int64) *SchemeMessageAdmission {

	m.Version = &version

	return m
}

func (m *SchemeMessageAdmission) WithoutVersion() *SchemeMessageAdmission {
	m.Version = nil
	return m
}

// Validate validates this scheme message admission
func (m *SchemeMessageAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeMessageAdmission) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *SchemeMessageAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeMessageAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeMessageAdmission) validateRelationships(formats strfmt.Registry) error {

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

var schemeMessageAdmissionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["scheme_message_admissions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemeMessageAdmissionTypeTypePropEnum = append(schemeMessageAdmissionTypeTypePropEnum, v)
	}
}

const (

	// SchemeMessageAdmissionTypeSchemeMessageAdmissions captures enum value "scheme_message_admissions"
	SchemeMessageAdmissionTypeSchemeMessageAdmissions string = "scheme_message_admissions"
)

// prop value enum
func (m *SchemeMessageAdmission) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, schemeMessageAdmissionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SchemeMessageAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *SchemeMessageAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeMessageAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeMessageAdmission) UnmarshalBinary(b []byte) error {
	var res SchemeMessageAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeMessageAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
