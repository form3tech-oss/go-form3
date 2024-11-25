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

// RecallAdmissionUpdate recall admission update
// swagger:model RecallAdmissionUpdate
type RecallAdmissionUpdate struct {

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

func RecallAdmissionUpdateWithDefaults(defaults client.Defaults) *RecallAdmissionUpdate {
	return &RecallAdmissionUpdate{

		Attributes: RecallAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("RecallAdmissionUpdate", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("RecallAdmissionUpdate", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("RecallAdmissionUpdate", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("RecallAdmissionUpdate", "organisation_id"),

		Relationships: RecallAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("RecallAdmissionUpdate", "type"),

		Version: defaults.GetInt64Ptr("RecallAdmissionUpdate", "version"),
	}
}

func (m *RecallAdmissionUpdate) WithAttributes(attributes RecallAdmissionAttributes) *RecallAdmissionUpdate {

	m.Attributes = &attributes

	return m
}

func (m *RecallAdmissionUpdate) WithoutAttributes() *RecallAdmissionUpdate {
	m.Attributes = nil
	return m
}

func (m *RecallAdmissionUpdate) WithCreatedOn(createdOn strfmt.DateTime) *RecallAdmissionUpdate {

	m.CreatedOn = &createdOn

	return m
}

func (m *RecallAdmissionUpdate) WithoutCreatedOn() *RecallAdmissionUpdate {
	m.CreatedOn = nil
	return m
}

func (m *RecallAdmissionUpdate) WithID(id strfmt.UUID) *RecallAdmissionUpdate {

	m.ID = &id

	return m
}

func (m *RecallAdmissionUpdate) WithoutID() *RecallAdmissionUpdate {
	m.ID = nil
	return m
}

func (m *RecallAdmissionUpdate) WithModifiedOn(modifiedOn strfmt.DateTime) *RecallAdmissionUpdate {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *RecallAdmissionUpdate) WithoutModifiedOn() *RecallAdmissionUpdate {
	m.ModifiedOn = nil
	return m
}

func (m *RecallAdmissionUpdate) WithOrganisationID(organisationID strfmt.UUID) *RecallAdmissionUpdate {

	m.OrganisationID = &organisationID

	return m
}

func (m *RecallAdmissionUpdate) WithoutOrganisationID() *RecallAdmissionUpdate {
	m.OrganisationID = nil
	return m
}

func (m *RecallAdmissionUpdate) WithRelationships(relationships RecallAdmissionRelationships) *RecallAdmissionUpdate {

	m.Relationships = &relationships

	return m
}

func (m *RecallAdmissionUpdate) WithoutRelationships() *RecallAdmissionUpdate {
	m.Relationships = nil
	return m
}

func (m *RecallAdmissionUpdate) WithType(typeVar string) *RecallAdmissionUpdate {

	m.Type = typeVar

	return m
}

func (m *RecallAdmissionUpdate) WithVersion(version int64) *RecallAdmissionUpdate {

	m.Version = &version

	return m
}

func (m *RecallAdmissionUpdate) WithoutVersion() *RecallAdmissionUpdate {
	m.Version = nil
	return m
}

// Validate validates this recall admission update
func (m *RecallAdmissionUpdate) Validate(formats strfmt.Registry) error {
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

func (m *RecallAdmissionUpdate) validateAttributes(formats strfmt.Registry) error {

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

func (m *RecallAdmissionUpdate) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallAdmissionUpdate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallAdmissionUpdate) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallAdmissionUpdate) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallAdmissionUpdate) validateRelationships(formats strfmt.Registry) error {

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

func (m *RecallAdmissionUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *RecallAdmissionUpdate) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallAdmissionUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallAdmissionUpdate) UnmarshalBinary(b []byte) error {
	var res RecallAdmissionUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallAdmissionUpdate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
