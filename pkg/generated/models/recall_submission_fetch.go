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

// RecallSubmissionFetch recall submission fetch
// swagger:model RecallSubmissionFetch
type RecallSubmissionFetch struct {

	// attributes
	Attributes *RecallSubmissionAttributes `json:"attributes,omitempty"`

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
	Relationships *RecallSubmissionFetchRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func RecallSubmissionFetchWithDefaults(defaults client.Defaults) *RecallSubmissionFetch {
	return &RecallSubmissionFetch{

		Attributes: RecallSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("RecallSubmissionFetch", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("RecallSubmissionFetch", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("RecallSubmissionFetch", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("RecallSubmissionFetch", "organisation_id"),

		Relationships: RecallSubmissionFetchRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("RecallSubmissionFetch", "type"),

		Version: defaults.GetInt64Ptr("RecallSubmissionFetch", "version"),
	}
}

func (m *RecallSubmissionFetch) WithAttributes(attributes RecallSubmissionAttributes) *RecallSubmissionFetch {

	m.Attributes = &attributes

	return m
}

func (m *RecallSubmissionFetch) WithoutAttributes() *RecallSubmissionFetch {
	m.Attributes = nil
	return m
}

func (m *RecallSubmissionFetch) WithCreatedOn(createdOn strfmt.DateTime) *RecallSubmissionFetch {

	m.CreatedOn = &createdOn

	return m
}

func (m *RecallSubmissionFetch) WithoutCreatedOn() *RecallSubmissionFetch {
	m.CreatedOn = nil
	return m
}

func (m *RecallSubmissionFetch) WithID(id strfmt.UUID) *RecallSubmissionFetch {

	m.ID = &id

	return m
}

func (m *RecallSubmissionFetch) WithoutID() *RecallSubmissionFetch {
	m.ID = nil
	return m
}

func (m *RecallSubmissionFetch) WithModifiedOn(modifiedOn strfmt.DateTime) *RecallSubmissionFetch {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *RecallSubmissionFetch) WithoutModifiedOn() *RecallSubmissionFetch {
	m.ModifiedOn = nil
	return m
}

func (m *RecallSubmissionFetch) WithOrganisationID(organisationID strfmt.UUID) *RecallSubmissionFetch {

	m.OrganisationID = &organisationID

	return m
}

func (m *RecallSubmissionFetch) WithoutOrganisationID() *RecallSubmissionFetch {
	m.OrganisationID = nil
	return m
}

func (m *RecallSubmissionFetch) WithRelationships(relationships RecallSubmissionFetchRelationships) *RecallSubmissionFetch {

	m.Relationships = &relationships

	return m
}

func (m *RecallSubmissionFetch) WithoutRelationships() *RecallSubmissionFetch {
	m.Relationships = nil
	return m
}

func (m *RecallSubmissionFetch) WithType(typeVar string) *RecallSubmissionFetch {

	m.Type = typeVar

	return m
}

func (m *RecallSubmissionFetch) WithVersion(version int64) *RecallSubmissionFetch {

	m.Version = &version

	return m
}

func (m *RecallSubmissionFetch) WithoutVersion() *RecallSubmissionFetch {
	m.Version = nil
	return m
}

// Validate validates this recall submission fetch
func (m *RecallSubmissionFetch) Validate(formats strfmt.Registry) error {
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

func (m *RecallSubmissionFetch) validateAttributes(formats strfmt.Registry) error {

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

func (m *RecallSubmissionFetch) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallSubmissionFetch) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallSubmissionFetch) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallSubmissionFetch) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallSubmissionFetch) validateRelationships(formats strfmt.Registry) error {

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

func (m *RecallSubmissionFetch) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *RecallSubmissionFetch) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallSubmissionFetch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallSubmissionFetch) UnmarshalBinary(b []byte) error {
	var res RecallSubmissionFetch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallSubmissionFetch) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
