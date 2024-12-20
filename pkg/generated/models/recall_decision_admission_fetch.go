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

// RecallDecisionAdmissionFetch recall decision admission fetch
// swagger:model RecallDecisionAdmissionFetch
type RecallDecisionAdmissionFetch struct {

	// attributes
	Attributes *RecallDecisionAdmissionAttributes `json:"attributes,omitempty"`

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
	Relationships *RecallDecisionAdmissionFetchRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func RecallDecisionAdmissionFetchWithDefaults(defaults client.Defaults) *RecallDecisionAdmissionFetch {
	return &RecallDecisionAdmissionFetch{

		Attributes: RecallDecisionAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("RecallDecisionAdmissionFetch", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("RecallDecisionAdmissionFetch", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("RecallDecisionAdmissionFetch", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("RecallDecisionAdmissionFetch", "organisation_id"),

		Relationships: RecallDecisionAdmissionFetchRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("RecallDecisionAdmissionFetch", "type"),

		Version: defaults.GetInt64Ptr("RecallDecisionAdmissionFetch", "version"),
	}
}

func (m *RecallDecisionAdmissionFetch) WithAttributes(attributes RecallDecisionAdmissionAttributes) *RecallDecisionAdmissionFetch {

	m.Attributes = &attributes

	return m
}

func (m *RecallDecisionAdmissionFetch) WithoutAttributes() *RecallDecisionAdmissionFetch {
	m.Attributes = nil
	return m
}

func (m *RecallDecisionAdmissionFetch) WithCreatedOn(createdOn strfmt.DateTime) *RecallDecisionAdmissionFetch {

	m.CreatedOn = &createdOn

	return m
}

func (m *RecallDecisionAdmissionFetch) WithoutCreatedOn() *RecallDecisionAdmissionFetch {
	m.CreatedOn = nil
	return m
}

func (m *RecallDecisionAdmissionFetch) WithID(id strfmt.UUID) *RecallDecisionAdmissionFetch {

	m.ID = &id

	return m
}

func (m *RecallDecisionAdmissionFetch) WithoutID() *RecallDecisionAdmissionFetch {
	m.ID = nil
	return m
}

func (m *RecallDecisionAdmissionFetch) WithModifiedOn(modifiedOn strfmt.DateTime) *RecallDecisionAdmissionFetch {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *RecallDecisionAdmissionFetch) WithoutModifiedOn() *RecallDecisionAdmissionFetch {
	m.ModifiedOn = nil
	return m
}

func (m *RecallDecisionAdmissionFetch) WithOrganisationID(organisationID strfmt.UUID) *RecallDecisionAdmissionFetch {

	m.OrganisationID = &organisationID

	return m
}

func (m *RecallDecisionAdmissionFetch) WithoutOrganisationID() *RecallDecisionAdmissionFetch {
	m.OrganisationID = nil
	return m
}

func (m *RecallDecisionAdmissionFetch) WithRelationships(relationships RecallDecisionAdmissionFetchRelationships) *RecallDecisionAdmissionFetch {

	m.Relationships = &relationships

	return m
}

func (m *RecallDecisionAdmissionFetch) WithoutRelationships() *RecallDecisionAdmissionFetch {
	m.Relationships = nil
	return m
}

func (m *RecallDecisionAdmissionFetch) WithType(typeVar string) *RecallDecisionAdmissionFetch {

	m.Type = typeVar

	return m
}

func (m *RecallDecisionAdmissionFetch) WithVersion(version int64) *RecallDecisionAdmissionFetch {

	m.Version = &version

	return m
}

func (m *RecallDecisionAdmissionFetch) WithoutVersion() *RecallDecisionAdmissionFetch {
	m.Version = nil
	return m
}

// Validate validates this recall decision admission fetch
func (m *RecallDecisionAdmissionFetch) Validate(formats strfmt.Registry) error {
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

func (m *RecallDecisionAdmissionFetch) validateAttributes(formats strfmt.Registry) error {

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

func (m *RecallDecisionAdmissionFetch) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionAdmissionFetch) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionAdmissionFetch) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionAdmissionFetch) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionAdmissionFetch) validateRelationships(formats strfmt.Registry) error {

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

func (m *RecallDecisionAdmissionFetch) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionAdmissionFetch) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecisionAdmissionFetch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionAdmissionFetch) UnmarshalBinary(b []byte) error {
	var res RecallDecisionAdmissionFetch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionAdmissionFetch) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
