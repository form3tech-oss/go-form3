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

// RecallDecisionFetch recall decision fetch
// swagger:model RecallDecisionFetch
type RecallDecisionFetch struct {

	// attributes
	Attributes *RecallDecisionAttributes `json:"attributes,omitempty"`

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
	Relationships *RecallDecisionFetchRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func RecallDecisionFetchWithDefaults(defaults client.Defaults) *RecallDecisionFetch {
	return &RecallDecisionFetch{

		Attributes: RecallDecisionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("RecallDecisionFetch", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("RecallDecisionFetch", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("RecallDecisionFetch", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("RecallDecisionFetch", "organisation_id"),

		Relationships: RecallDecisionFetchRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("RecallDecisionFetch", "type"),

		Version: defaults.GetInt64Ptr("RecallDecisionFetch", "version"),
	}
}

func (m *RecallDecisionFetch) WithAttributes(attributes RecallDecisionAttributes) *RecallDecisionFetch {

	m.Attributes = &attributes

	return m
}

func (m *RecallDecisionFetch) WithoutAttributes() *RecallDecisionFetch {
	m.Attributes = nil
	return m
}

func (m *RecallDecisionFetch) WithCreatedOn(createdOn strfmt.DateTime) *RecallDecisionFetch {

	m.CreatedOn = &createdOn

	return m
}

func (m *RecallDecisionFetch) WithoutCreatedOn() *RecallDecisionFetch {
	m.CreatedOn = nil
	return m
}

func (m *RecallDecisionFetch) WithID(id strfmt.UUID) *RecallDecisionFetch {

	m.ID = &id

	return m
}

func (m *RecallDecisionFetch) WithoutID() *RecallDecisionFetch {
	m.ID = nil
	return m
}

func (m *RecallDecisionFetch) WithModifiedOn(modifiedOn strfmt.DateTime) *RecallDecisionFetch {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *RecallDecisionFetch) WithoutModifiedOn() *RecallDecisionFetch {
	m.ModifiedOn = nil
	return m
}

func (m *RecallDecisionFetch) WithOrganisationID(organisationID strfmt.UUID) *RecallDecisionFetch {

	m.OrganisationID = &organisationID

	return m
}

func (m *RecallDecisionFetch) WithoutOrganisationID() *RecallDecisionFetch {
	m.OrganisationID = nil
	return m
}

func (m *RecallDecisionFetch) WithRelationships(relationships RecallDecisionFetchRelationships) *RecallDecisionFetch {

	m.Relationships = &relationships

	return m
}

func (m *RecallDecisionFetch) WithoutRelationships() *RecallDecisionFetch {
	m.Relationships = nil
	return m
}

func (m *RecallDecisionFetch) WithType(typeVar string) *RecallDecisionFetch {

	m.Type = typeVar

	return m
}

func (m *RecallDecisionFetch) WithVersion(version int64) *RecallDecisionFetch {

	m.Version = &version

	return m
}

func (m *RecallDecisionFetch) WithoutVersion() *RecallDecisionFetch {
	m.Version = nil
	return m
}

// Validate validates this recall decision fetch
func (m *RecallDecisionFetch) Validate(formats strfmt.Registry) error {
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

func (m *RecallDecisionFetch) validateAttributes(formats strfmt.Registry) error {

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

func (m *RecallDecisionFetch) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionFetch) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionFetch) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionFetch) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionFetch) validateRelationships(formats strfmt.Registry) error {

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

func (m *RecallDecisionFetch) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionFetch) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecisionFetch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionFetch) UnmarshalBinary(b []byte) error {
	var res RecallDecisionFetch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionFetch) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}