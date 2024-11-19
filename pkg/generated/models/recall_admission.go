// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RecallAdmission recall admission
// swagger:model RecallAdmission
type RecallAdmission struct {

	// attributes
	Attributes *RecallAdmissionAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *RecallAdmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func RecallAdmissionWithDefaults(defaults client.Defaults) *RecallAdmission {
	return &RecallAdmission{

		Attributes: RecallAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("RecallAdmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("RecallAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("RecallAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("RecallAdmission", "organisation_id"),

		Relationships: RecallAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("RecallAdmission", "type"),

		Version: defaults.GetInt64Ptr("RecallAdmission", "version"),
	}
}

func (m *RecallAdmission) WithAttributes(attributes RecallAdmissionAttributes) *RecallAdmission {

	m.Attributes = &attributes

	return m
}

func (m *RecallAdmission) WithoutAttributes() *RecallAdmission {
	m.Attributes = nil
	return m
}

func (m *RecallAdmission) WithCreatedOn(createdOn strfmt.DateTime) *RecallAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *RecallAdmission) WithoutCreatedOn() *RecallAdmission {
	m.CreatedOn = nil
	return m
}

func (m *RecallAdmission) WithID(id strfmt.UUID) *RecallAdmission {

	m.ID = &id

	return m
}

func (m *RecallAdmission) WithoutID() *RecallAdmission {
	m.ID = nil
	return m
}

func (m *RecallAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *RecallAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *RecallAdmission) WithoutModifiedOn() *RecallAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *RecallAdmission) WithOrganisationID(organisationID strfmt.UUID) *RecallAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *RecallAdmission) WithoutOrganisationID() *RecallAdmission {
	m.OrganisationID = nil
	return m
}

func (m *RecallAdmission) WithRelationships(relationships RecallAdmissionRelationships) *RecallAdmission {

	m.Relationships = &relationships

	return m
}

func (m *RecallAdmission) WithoutRelationships() *RecallAdmission {
	m.Relationships = nil
	return m
}

func (m *RecallAdmission) WithType(typeVar string) *RecallAdmission {

	m.Type = typeVar

	return m
}

func (m *RecallAdmission) WithVersion(version int64) *RecallAdmission {

	m.Version = &version

	return m
}

func (m *RecallAdmission) WithoutVersion() *RecallAdmission {
	m.Version = nil
	return m
}

// Validate validates this recall admission
func (m *RecallAdmission) Validate(formats strfmt.Registry) error {
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

func (m *RecallAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *RecallAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallAdmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *RecallAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *RecallAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallAdmission) UnmarshalBinary(b []byte) error {
	var res RecallAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
