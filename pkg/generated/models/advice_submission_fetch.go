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

// AdviceSubmissionFetch advice submission fetch
// swagger:model AdviceSubmissionFetch
type AdviceSubmissionFetch struct {

	// attributes
	Attributes *AdviceSubmissionAttributes `json:"attributes,omitempty"`

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
	Relationships *AdviceSubmissionFetchRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func AdviceSubmissionFetchWithDefaults(defaults client.Defaults) *AdviceSubmissionFetch {
	return &AdviceSubmissionFetch{

		Attributes: AdviceSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("AdviceSubmissionFetch", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("AdviceSubmissionFetch", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("AdviceSubmissionFetch", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("AdviceSubmissionFetch", "organisation_id"),

		Relationships: AdviceSubmissionFetchRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("AdviceSubmissionFetch", "type"),

		Version: defaults.GetInt64Ptr("AdviceSubmissionFetch", "version"),
	}
}

func (m *AdviceSubmissionFetch) WithAttributes(attributes AdviceSubmissionAttributes) *AdviceSubmissionFetch {

	m.Attributes = &attributes

	return m
}

func (m *AdviceSubmissionFetch) WithoutAttributes() *AdviceSubmissionFetch {
	m.Attributes = nil
	return m
}

func (m *AdviceSubmissionFetch) WithCreatedOn(createdOn strfmt.DateTime) *AdviceSubmissionFetch {

	m.CreatedOn = &createdOn

	return m
}

func (m *AdviceSubmissionFetch) WithoutCreatedOn() *AdviceSubmissionFetch {
	m.CreatedOn = nil
	return m
}

func (m *AdviceSubmissionFetch) WithID(id strfmt.UUID) *AdviceSubmissionFetch {

	m.ID = &id

	return m
}

func (m *AdviceSubmissionFetch) WithoutID() *AdviceSubmissionFetch {
	m.ID = nil
	return m
}

func (m *AdviceSubmissionFetch) WithModifiedOn(modifiedOn strfmt.DateTime) *AdviceSubmissionFetch {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *AdviceSubmissionFetch) WithoutModifiedOn() *AdviceSubmissionFetch {
	m.ModifiedOn = nil
	return m
}

func (m *AdviceSubmissionFetch) WithOrganisationID(organisationID strfmt.UUID) *AdviceSubmissionFetch {

	m.OrganisationID = &organisationID

	return m
}

func (m *AdviceSubmissionFetch) WithoutOrganisationID() *AdviceSubmissionFetch {
	m.OrganisationID = nil
	return m
}

func (m *AdviceSubmissionFetch) WithRelationships(relationships AdviceSubmissionFetchRelationships) *AdviceSubmissionFetch {

	m.Relationships = &relationships

	return m
}

func (m *AdviceSubmissionFetch) WithoutRelationships() *AdviceSubmissionFetch {
	m.Relationships = nil
	return m
}

func (m *AdviceSubmissionFetch) WithType(typeVar string) *AdviceSubmissionFetch {

	m.Type = typeVar

	return m
}

func (m *AdviceSubmissionFetch) WithVersion(version int64) *AdviceSubmissionFetch {

	m.Version = &version

	return m
}

func (m *AdviceSubmissionFetch) WithoutVersion() *AdviceSubmissionFetch {
	m.Version = nil
	return m
}

// Validate validates this advice submission fetch
func (m *AdviceSubmissionFetch) Validate(formats strfmt.Registry) error {
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

func (m *AdviceSubmissionFetch) validateAttributes(formats strfmt.Registry) error {

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

func (m *AdviceSubmissionFetch) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdviceSubmissionFetch) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdviceSubmissionFetch) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdviceSubmissionFetch) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdviceSubmissionFetch) validateRelationships(formats strfmt.Registry) error {

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

func (m *AdviceSubmissionFetch) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AdviceSubmissionFetch) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdviceSubmissionFetch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdviceSubmissionFetch) UnmarshalBinary(b []byte) error {
	var res AdviceSubmissionFetch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AdviceSubmissionFetch) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
