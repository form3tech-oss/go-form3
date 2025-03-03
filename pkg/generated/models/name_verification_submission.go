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

// NameVerificationSubmission name verification submission
// swagger:model NameVerificationSubmission
type NameVerificationSubmission struct {

	// attributes
	Attributes *NameVerificationSubmissionAttributes `json:"attributes,omitempty"`

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
	Relationships *NameVerificationSubmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func NameVerificationSubmissionWithDefaults(defaults client.Defaults) *NameVerificationSubmission {
	return &NameVerificationSubmission{

		Attributes: NameVerificationSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("NameVerificationSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("NameVerificationSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("NameVerificationSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NameVerificationSubmission", "organisation_id"),

		Relationships: NameVerificationSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("NameVerificationSubmission", "type"),

		Version: defaults.GetInt64Ptr("NameVerificationSubmission", "version"),
	}
}

func (m *NameVerificationSubmission) WithAttributes(attributes NameVerificationSubmissionAttributes) *NameVerificationSubmission {

	m.Attributes = &attributes

	return m
}

func (m *NameVerificationSubmission) WithoutAttributes() *NameVerificationSubmission {
	m.Attributes = nil
	return m
}

func (m *NameVerificationSubmission) WithCreatedOn(createdOn strfmt.DateTime) *NameVerificationSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *NameVerificationSubmission) WithoutCreatedOn() *NameVerificationSubmission {
	m.CreatedOn = nil
	return m
}

func (m *NameVerificationSubmission) WithID(id strfmt.UUID) *NameVerificationSubmission {

	m.ID = &id

	return m
}

func (m *NameVerificationSubmission) WithoutID() *NameVerificationSubmission {
	m.ID = nil
	return m
}

func (m *NameVerificationSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *NameVerificationSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *NameVerificationSubmission) WithoutModifiedOn() *NameVerificationSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *NameVerificationSubmission) WithOrganisationID(organisationID strfmt.UUID) *NameVerificationSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *NameVerificationSubmission) WithoutOrganisationID() *NameVerificationSubmission {
	m.OrganisationID = nil
	return m
}

func (m *NameVerificationSubmission) WithRelationships(relationships NameVerificationSubmissionRelationships) *NameVerificationSubmission {

	m.Relationships = &relationships

	return m
}

func (m *NameVerificationSubmission) WithoutRelationships() *NameVerificationSubmission {
	m.Relationships = nil
	return m
}

func (m *NameVerificationSubmission) WithType(typeVar string) *NameVerificationSubmission {

	m.Type = typeVar

	return m
}

func (m *NameVerificationSubmission) WithVersion(version int64) *NameVerificationSubmission {

	m.Version = &version

	return m
}

func (m *NameVerificationSubmission) WithoutVersion() *NameVerificationSubmission {
	m.Version = nil
	return m
}

// Validate validates this name verification submission
func (m *NameVerificationSubmission) Validate(formats strfmt.Registry) error {
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

func (m *NameVerificationSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *NameVerificationSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerificationSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerificationSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerificationSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerificationSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *NameVerificationSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *NameVerificationSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameVerificationSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameVerificationSubmission) UnmarshalBinary(b []byte) error {
	var res NameVerificationSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NameVerificationSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
