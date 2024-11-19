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

// ReturnReversalAdmission return reversal admission
// swagger:model ReturnReversalAdmission
type ReturnReversalAdmission struct {

	// attributes
	Attributes *ReturnReversalAdmissionAttributes `json:"attributes,omitempty"`

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
	Relationships *ReturnReversalAdmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReturnReversalAdmissionWithDefaults(defaults client.Defaults) *ReturnReversalAdmission {
	return &ReturnReversalAdmission{

		Attributes: ReturnReversalAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("ReturnReversalAdmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ReturnReversalAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("ReturnReversalAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ReturnReversalAdmission", "organisation_id"),

		Relationships: ReturnReversalAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReturnReversalAdmission", "type"),

		Version: defaults.GetInt64Ptr("ReturnReversalAdmission", "version"),
	}
}

func (m *ReturnReversalAdmission) WithAttributes(attributes ReturnReversalAdmissionAttributes) *ReturnReversalAdmission {

	m.Attributes = &attributes

	return m
}

func (m *ReturnReversalAdmission) WithoutAttributes() *ReturnReversalAdmission {
	m.Attributes = nil
	return m
}

func (m *ReturnReversalAdmission) WithCreatedOn(createdOn strfmt.DateTime) *ReturnReversalAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *ReturnReversalAdmission) WithoutCreatedOn() *ReturnReversalAdmission {
	m.CreatedOn = nil
	return m
}

func (m *ReturnReversalAdmission) WithID(id strfmt.UUID) *ReturnReversalAdmission {

	m.ID = &id

	return m
}

func (m *ReturnReversalAdmission) WithoutID() *ReturnReversalAdmission {
	m.ID = nil
	return m
}

func (m *ReturnReversalAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *ReturnReversalAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *ReturnReversalAdmission) WithoutModifiedOn() *ReturnReversalAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *ReturnReversalAdmission) WithOrganisationID(organisationID strfmt.UUID) *ReturnReversalAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *ReturnReversalAdmission) WithoutOrganisationID() *ReturnReversalAdmission {
	m.OrganisationID = nil
	return m
}

func (m *ReturnReversalAdmission) WithRelationships(relationships ReturnReversalAdmissionRelationships) *ReturnReversalAdmission {

	m.Relationships = &relationships

	return m
}

func (m *ReturnReversalAdmission) WithoutRelationships() *ReturnReversalAdmission {
	m.Relationships = nil
	return m
}

func (m *ReturnReversalAdmission) WithType(typeVar string) *ReturnReversalAdmission {

	m.Type = typeVar

	return m
}

func (m *ReturnReversalAdmission) WithVersion(version int64) *ReturnReversalAdmission {

	m.Version = &version

	return m
}

func (m *ReturnReversalAdmission) WithoutVersion() *ReturnReversalAdmission {
	m.Version = nil
	return m
}

// Validate validates this return reversal admission
func (m *ReturnReversalAdmission) Validate(formats strfmt.Registry) error {
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

func (m *ReturnReversalAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReturnReversalAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnReversalAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnReversalAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnReversalAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnReversalAdmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReturnReversalAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnReversalAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnReversalAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnReversalAdmission) UnmarshalBinary(b []byte) error {
	var res ReturnReversalAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnReversalAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReturnReversalAdmissionAttributes return reversal admission attributes
// swagger:model ReturnReversalAdmissionAttributes
type ReturnReversalAdmissionAttributes struct {

	// Scheme-specific status code. Refer to individual scheme where applicable
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// [Description](http://draft-api-docs.form3.tech/api.html#enumerations-scheme-status-codes-for-bacs) of scheme_status_code
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`
}

func ReturnReversalAdmissionAttributesWithDefaults(defaults client.Defaults) *ReturnReversalAdmissionAttributes {
	return &ReturnReversalAdmissionAttributes{

		SchemeStatusCode: defaults.GetString("ReturnReversalAdmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("ReturnReversalAdmissionAttributes", "scheme_status_code_description"),
	}
}

func (m *ReturnReversalAdmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *ReturnReversalAdmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *ReturnReversalAdmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *ReturnReversalAdmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

// Validate validates this return reversal admission attributes
func (m *ReturnReversalAdmissionAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReturnReversalAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnReversalAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnReversalAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnReversalAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
