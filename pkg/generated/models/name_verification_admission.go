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

// NameVerificationAdmission name verification admission
// swagger:model NameVerificationAdmission
type NameVerificationAdmission struct {

	// attributes
	// Required: true
	Attributes *NameVerificationAdmissionAttributes `json:"attributes"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *NameVerificationAdmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func NameVerificationAdmissionWithDefaults(defaults client.Defaults) *NameVerificationAdmission {
	return &NameVerificationAdmission{

		Attributes: NameVerificationAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("NameVerificationAdmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("NameVerificationAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("NameVerificationAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NameVerificationAdmission", "organisation_id"),

		Relationships: NameVerificationAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("NameVerificationAdmission", "type"),

		Version: defaults.GetInt64Ptr("NameVerificationAdmission", "version"),
	}
}

func (m *NameVerificationAdmission) WithAttributes(attributes NameVerificationAdmissionAttributes) *NameVerificationAdmission {

	m.Attributes = &attributes

	return m
}

func (m *NameVerificationAdmission) WithoutAttributes() *NameVerificationAdmission {
	m.Attributes = nil
	return m
}

func (m *NameVerificationAdmission) WithCreatedOn(createdOn strfmt.DateTime) *NameVerificationAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *NameVerificationAdmission) WithoutCreatedOn() *NameVerificationAdmission {
	m.CreatedOn = nil
	return m
}

func (m *NameVerificationAdmission) WithID(id strfmt.UUID) *NameVerificationAdmission {

	m.ID = &id

	return m
}

func (m *NameVerificationAdmission) WithoutID() *NameVerificationAdmission {
	m.ID = nil
	return m
}

func (m *NameVerificationAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *NameVerificationAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *NameVerificationAdmission) WithoutModifiedOn() *NameVerificationAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *NameVerificationAdmission) WithOrganisationID(organisationID strfmt.UUID) *NameVerificationAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *NameVerificationAdmission) WithoutOrganisationID() *NameVerificationAdmission {
	m.OrganisationID = nil
	return m
}

func (m *NameVerificationAdmission) WithRelationships(relationships NameVerificationAdmissionRelationships) *NameVerificationAdmission {

	m.Relationships = &relationships

	return m
}

func (m *NameVerificationAdmission) WithoutRelationships() *NameVerificationAdmission {
	m.Relationships = nil
	return m
}

func (m *NameVerificationAdmission) WithType(typeVar string) *NameVerificationAdmission {

	m.Type = typeVar

	return m
}

func (m *NameVerificationAdmission) WithVersion(version int64) *NameVerificationAdmission {

	m.Version = &version

	return m
}

func (m *NameVerificationAdmission) WithoutVersion() *NameVerificationAdmission {
	m.Version = nil
	return m
}

// Validate validates this name verification admission
func (m *NameVerificationAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
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

func (m *NameVerificationAdmission) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
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

func (m *NameVerificationAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerificationAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerificationAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerificationAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerificationAdmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *NameVerificationAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *NameVerificationAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameVerificationAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameVerificationAdmission) UnmarshalBinary(b []byte) error {
	var res NameVerificationAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NameVerificationAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
