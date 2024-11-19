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

// ClaimReversal claim reversal
// swagger:model ClaimReversal
type ClaimReversal struct {

	// attributes
	// Required: true
	Attributes *ClaimReversalAttributes `json:"attributes"`

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
	Relationships *ClaimReversalRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ClaimReversalWithDefaults(defaults client.Defaults) *ClaimReversal {
	return &ClaimReversal{

		Attributes: ClaimReversalAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("ClaimReversal", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ClaimReversal", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("ClaimReversal", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ClaimReversal", "organisation_id"),

		Relationships: ClaimReversalRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ClaimReversal", "type"),

		Version: defaults.GetInt64Ptr("ClaimReversal", "version"),
	}
}

func (m *ClaimReversal) WithAttributes(attributes ClaimReversalAttributes) *ClaimReversal {

	m.Attributes = &attributes

	return m
}

func (m *ClaimReversal) WithoutAttributes() *ClaimReversal {
	m.Attributes = nil
	return m
}

func (m *ClaimReversal) WithCreatedOn(createdOn strfmt.DateTime) *ClaimReversal {

	m.CreatedOn = &createdOn

	return m
}

func (m *ClaimReversal) WithoutCreatedOn() *ClaimReversal {
	m.CreatedOn = nil
	return m
}

func (m *ClaimReversal) WithID(id strfmt.UUID) *ClaimReversal {

	m.ID = &id

	return m
}

func (m *ClaimReversal) WithoutID() *ClaimReversal {
	m.ID = nil
	return m
}

func (m *ClaimReversal) WithModifiedOn(modifiedOn strfmt.DateTime) *ClaimReversal {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *ClaimReversal) WithoutModifiedOn() *ClaimReversal {
	m.ModifiedOn = nil
	return m
}

func (m *ClaimReversal) WithOrganisationID(organisationID strfmt.UUID) *ClaimReversal {

	m.OrganisationID = &organisationID

	return m
}

func (m *ClaimReversal) WithoutOrganisationID() *ClaimReversal {
	m.OrganisationID = nil
	return m
}

func (m *ClaimReversal) WithRelationships(relationships ClaimReversalRelationships) *ClaimReversal {

	m.Relationships = &relationships

	return m
}

func (m *ClaimReversal) WithoutRelationships() *ClaimReversal {
	m.Relationships = nil
	return m
}

func (m *ClaimReversal) WithType(typeVar string) *ClaimReversal {

	m.Type = typeVar

	return m
}

func (m *ClaimReversal) WithVersion(version int64) *ClaimReversal {

	m.Version = &version

	return m
}

func (m *ClaimReversal) WithoutVersion() *ClaimReversal {
	m.Version = nil
	return m
}

// Validate validates this claim reversal
func (m *ClaimReversal) Validate(formats strfmt.Registry) error {
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

func (m *ClaimReversal) validateAttributes(formats strfmt.Registry) error {

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

func (m *ClaimReversal) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClaimReversal) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClaimReversal) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClaimReversal) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClaimReversal) validateRelationships(formats strfmt.Registry) error {

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

func (m *ClaimReversal) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ClaimReversal) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimReversal) UnmarshalBinary(b []byte) error {
	var res ClaimReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ClaimReversalAttributes claim reversal attributes
// swagger:model ClaimReversalAttributes
type ClaimReversalAttributes struct {

	// original instruction id
	// Required: true
	// Min Length: 1
	OriginalInstructionID *string `json:"original_instruction_id"`
}

func ClaimReversalAttributesWithDefaults(defaults client.Defaults) *ClaimReversalAttributes {
	return &ClaimReversalAttributes{

		OriginalInstructionID: defaults.GetStringPtr("ClaimReversalAttributes", "original_instruction_id"),
	}
}

func (m *ClaimReversalAttributes) WithOriginalInstructionID(originalInstructionID string) *ClaimReversalAttributes {

	m.OriginalInstructionID = &originalInstructionID

	return m
}

func (m *ClaimReversalAttributes) WithoutOriginalInstructionID() *ClaimReversalAttributes {
	m.OriginalInstructionID = nil
	return m
}

// Validate validates this claim reversal attributes
func (m *ClaimReversalAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOriginalInstructionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimReversalAttributes) validateOriginalInstructionID(formats strfmt.Registry) error {

	if err := validate.Required("attributes"+"."+"original_instruction_id", "body", m.OriginalInstructionID); err != nil {
		return err
	}

	if err := validate.MinLength("attributes"+"."+"original_instruction_id", "body", string(*m.OriginalInstructionID), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimReversalAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimReversalAttributes) UnmarshalBinary(b []byte) error {
	var res ClaimReversalAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimReversalAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
