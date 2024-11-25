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

// Mandate mandate
// swagger:model Mandate
type Mandate struct {

	// attributes
	Attributes *MandateAttributes `json:"attributes,omitempty"`

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
	Relationships *MandateRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func MandateWithDefaults(defaults client.Defaults) *Mandate {
	return &Mandate{

		Attributes: MandateAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("Mandate", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("Mandate", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("Mandate", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("Mandate", "organisation_id"),

		Relationships: MandateRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("Mandate", "type"),

		Version: defaults.GetInt64Ptr("Mandate", "version"),
	}
}

func (m *Mandate) WithAttributes(attributes MandateAttributes) *Mandate {

	m.Attributes = &attributes

	return m
}

func (m *Mandate) WithoutAttributes() *Mandate {
	m.Attributes = nil
	return m
}

func (m *Mandate) WithCreatedOn(createdOn strfmt.DateTime) *Mandate {

	m.CreatedOn = &createdOn

	return m
}

func (m *Mandate) WithoutCreatedOn() *Mandate {
	m.CreatedOn = nil
	return m
}

func (m *Mandate) WithID(id strfmt.UUID) *Mandate {

	m.ID = &id

	return m
}

func (m *Mandate) WithoutID() *Mandate {
	m.ID = nil
	return m
}

func (m *Mandate) WithModifiedOn(modifiedOn strfmt.DateTime) *Mandate {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *Mandate) WithoutModifiedOn() *Mandate {
	m.ModifiedOn = nil
	return m
}

func (m *Mandate) WithOrganisationID(organisationID strfmt.UUID) *Mandate {

	m.OrganisationID = &organisationID

	return m
}

func (m *Mandate) WithoutOrganisationID() *Mandate {
	m.OrganisationID = nil
	return m
}

func (m *Mandate) WithRelationships(relationships MandateRelationships) *Mandate {

	m.Relationships = &relationships

	return m
}

func (m *Mandate) WithoutRelationships() *Mandate {
	m.Relationships = nil
	return m
}

func (m *Mandate) WithType(typeVar string) *Mandate {

	m.Type = typeVar

	return m
}

func (m *Mandate) WithVersion(version int64) *Mandate {

	m.Version = &version

	return m
}

func (m *Mandate) WithoutVersion() *Mandate {
	m.Version = nil
	return m
}

// Validate validates this mandate
func (m *Mandate) Validate(formats strfmt.Registry) error {
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

func (m *Mandate) validateAttributes(formats strfmt.Registry) error {

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

func (m *Mandate) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Mandate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Mandate) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Mandate) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Mandate) validateRelationships(formats strfmt.Registry) error {

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

func (m *Mandate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Mandate) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Mandate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Mandate) UnmarshalBinary(b []byte) error {
	var res Mandate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Mandate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
