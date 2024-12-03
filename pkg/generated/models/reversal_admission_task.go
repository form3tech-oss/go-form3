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

// ReversalAdmissionTask reversal admission task
// swagger:model ReversalAdmissionTask
type ReversalAdmissionTask struct {

	// attributes
	Attributes *ReversalAdmissionTaskAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// modified on
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// relationships
	Relationships *ReversalAdmissionTaskRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReversalAdmissionTaskWithDefaults(defaults client.Defaults) *ReversalAdmissionTask {
	return &ReversalAdmissionTask{

		Attributes: ReversalAdmissionTaskAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("ReversalAdmissionTask", "created_on"),

		ID: defaults.GetStrfmtUUID("ReversalAdmissionTask", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("ReversalAdmissionTask", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("ReversalAdmissionTask", "organisation_id"),

		Relationships: ReversalAdmissionTaskRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReversalAdmissionTask", "type"),

		Version: defaults.GetInt64Ptr("ReversalAdmissionTask", "version"),
	}
}

func (m *ReversalAdmissionTask) WithAttributes(attributes ReversalAdmissionTaskAttributes) *ReversalAdmissionTask {

	m.Attributes = &attributes

	return m
}

func (m *ReversalAdmissionTask) WithoutAttributes() *ReversalAdmissionTask {
	m.Attributes = nil
	return m
}

func (m *ReversalAdmissionTask) WithCreatedOn(createdOn strfmt.DateTime) *ReversalAdmissionTask {

	m.CreatedOn = createdOn

	return m
}

func (m *ReversalAdmissionTask) WithID(id strfmt.UUID) *ReversalAdmissionTask {

	m.ID = id

	return m
}

func (m *ReversalAdmissionTask) WithModifiedOn(modifiedOn strfmt.DateTime) *ReversalAdmissionTask {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *ReversalAdmissionTask) WithOrganisationID(organisationID strfmt.UUID) *ReversalAdmissionTask {

	m.OrganisationID = organisationID

	return m
}

func (m *ReversalAdmissionTask) WithRelationships(relationships ReversalAdmissionTaskRelationships) *ReversalAdmissionTask {

	m.Relationships = &relationships

	return m
}

func (m *ReversalAdmissionTask) WithoutRelationships() *ReversalAdmissionTask {
	m.Relationships = nil
	return m
}

func (m *ReversalAdmissionTask) WithType(typeVar string) *ReversalAdmissionTask {

	m.Type = typeVar

	return m
}

func (m *ReversalAdmissionTask) WithVersion(version int64) *ReversalAdmissionTask {

	m.Version = &version

	return m
}

func (m *ReversalAdmissionTask) WithoutVersion() *ReversalAdmissionTask {
	m.Version = nil
	return m
}

// Validate validates this reversal admission task
func (m *ReversalAdmissionTask) Validate(formats strfmt.Registry) error {
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

func (m *ReversalAdmissionTask) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReversalAdmissionTask) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionTask) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionTask) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionTask) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionTask) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReversalAdmissionTask) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionTask) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalAdmissionTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalAdmissionTask) UnmarshalBinary(b []byte) error {
	var res ReversalAdmissionTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalAdmissionTask) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
