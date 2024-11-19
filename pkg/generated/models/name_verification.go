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

// NameVerification name verification
// swagger:model NameVerification
type NameVerification struct {

	// attributes
	// Required: true
	Attributes *NameVerificationAttributes `json:"attributes"`

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
	Relationships *NameVerificationRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func NameVerificationWithDefaults(defaults client.Defaults) *NameVerification {
	return &NameVerification{

		Attributes: NameVerificationAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("NameVerification", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("NameVerification", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("NameVerification", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NameVerification", "organisation_id"),

		Relationships: NameVerificationRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("NameVerification", "type"),

		Version: defaults.GetInt64Ptr("NameVerification", "version"),
	}
}

func (m *NameVerification) WithAttributes(attributes NameVerificationAttributes) *NameVerification {

	m.Attributes = &attributes

	return m
}

func (m *NameVerification) WithoutAttributes() *NameVerification {
	m.Attributes = nil
	return m
}

func (m *NameVerification) WithCreatedOn(createdOn strfmt.DateTime) *NameVerification {

	m.CreatedOn = &createdOn

	return m
}

func (m *NameVerification) WithoutCreatedOn() *NameVerification {
	m.CreatedOn = nil
	return m
}

func (m *NameVerification) WithID(id strfmt.UUID) *NameVerification {

	m.ID = &id

	return m
}

func (m *NameVerification) WithoutID() *NameVerification {
	m.ID = nil
	return m
}

func (m *NameVerification) WithModifiedOn(modifiedOn strfmt.DateTime) *NameVerification {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *NameVerification) WithoutModifiedOn() *NameVerification {
	m.ModifiedOn = nil
	return m
}

func (m *NameVerification) WithOrganisationID(organisationID strfmt.UUID) *NameVerification {

	m.OrganisationID = &organisationID

	return m
}

func (m *NameVerification) WithoutOrganisationID() *NameVerification {
	m.OrganisationID = nil
	return m
}

func (m *NameVerification) WithRelationships(relationships NameVerificationRelationships) *NameVerification {

	m.Relationships = &relationships

	return m
}

func (m *NameVerification) WithoutRelationships() *NameVerification {
	m.Relationships = nil
	return m
}

func (m *NameVerification) WithType(typeVar string) *NameVerification {

	m.Type = typeVar

	return m
}

func (m *NameVerification) WithVersion(version int64) *NameVerification {

	m.Version = &version

	return m
}

func (m *NameVerification) WithoutVersion() *NameVerification {
	m.Version = nil
	return m
}

// Validate validates this name verification
func (m *NameVerification) Validate(formats strfmt.Registry) error {
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

func (m *NameVerification) validateAttributes(formats strfmt.Registry) error {

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

func (m *NameVerification) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerification) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerification) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerification) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NameVerification) validateRelationships(formats strfmt.Registry) error {

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

func (m *NameVerification) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *NameVerification) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameVerification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameVerification) UnmarshalBinary(b []byte) error {
	var res NameVerification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NameVerification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
