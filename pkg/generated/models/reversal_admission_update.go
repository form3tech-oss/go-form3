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

// ReversalAdmissionUpdate reversal admission update
// swagger:model ReversalAdmissionUpdate
type ReversalAdmissionUpdate struct {

	// attributes
	Attributes *ReversalAdmissionUpdateAttributes `json:"attributes,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *ReversalAdmissionUpdateRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Required: true
	// Minimum: 0
	Version *int64 `json:"version"`
}

func ReversalAdmissionUpdateWithDefaults(defaults client.Defaults) *ReversalAdmissionUpdate {
	return &ReversalAdmissionUpdate{

		Attributes: ReversalAdmissionUpdateAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("ReversalAdmissionUpdate", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ReversalAdmissionUpdate", "organisation_id"),

		Relationships: ReversalAdmissionUpdateRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReversalAdmissionUpdate", "type"),

		Version: defaults.GetInt64Ptr("ReversalAdmissionUpdate", "version"),
	}
}

func (m *ReversalAdmissionUpdate) WithAttributes(attributes ReversalAdmissionUpdateAttributes) *ReversalAdmissionUpdate {

	m.Attributes = &attributes

	return m
}

func (m *ReversalAdmissionUpdate) WithoutAttributes() *ReversalAdmissionUpdate {
	m.Attributes = nil
	return m
}

func (m *ReversalAdmissionUpdate) WithID(id strfmt.UUID) *ReversalAdmissionUpdate {

	m.ID = &id

	return m
}

func (m *ReversalAdmissionUpdate) WithoutID() *ReversalAdmissionUpdate {
	m.ID = nil
	return m
}

func (m *ReversalAdmissionUpdate) WithOrganisationID(organisationID strfmt.UUID) *ReversalAdmissionUpdate {

	m.OrganisationID = &organisationID

	return m
}

func (m *ReversalAdmissionUpdate) WithoutOrganisationID() *ReversalAdmissionUpdate {
	m.OrganisationID = nil
	return m
}

func (m *ReversalAdmissionUpdate) WithRelationships(relationships ReversalAdmissionUpdateRelationships) *ReversalAdmissionUpdate {

	m.Relationships = &relationships

	return m
}

func (m *ReversalAdmissionUpdate) WithoutRelationships() *ReversalAdmissionUpdate {
	m.Relationships = nil
	return m
}

func (m *ReversalAdmissionUpdate) WithType(typeVar string) *ReversalAdmissionUpdate {

	m.Type = typeVar

	return m
}

func (m *ReversalAdmissionUpdate) WithVersion(version int64) *ReversalAdmissionUpdate {

	m.Version = &version

	return m
}

func (m *ReversalAdmissionUpdate) WithoutVersion() *ReversalAdmissionUpdate {
	m.Version = nil
	return m
}

// Validate validates this reversal admission update
func (m *ReversalAdmissionUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

func (m *ReversalAdmissionUpdate) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReversalAdmissionUpdate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionUpdate) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionUpdate) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReversalAdmissionUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionUpdate) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalAdmissionUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalAdmissionUpdate) UnmarshalBinary(b []byte) error {
	var res ReversalAdmissionUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalAdmissionUpdate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReversalAdmissionUpdateAttributes reversal admission update attributes
// swagger:model ReversalAdmissionUpdateAttributes
type ReversalAdmissionUpdateAttributes struct {

	// Scheme-specific status code. Refer to scheme documentation where available.
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// Description of `scheme_status_code`
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// source gateway
	SourceGateway string `json:"source_gateway,omitempty"`

	// status
	Status ReversalAdmissionStatus `json:"status,omitempty"`
}

func ReversalAdmissionUpdateAttributesWithDefaults(defaults client.Defaults) *ReversalAdmissionUpdateAttributes {
	return &ReversalAdmissionUpdateAttributes{

		SchemeStatusCode: defaults.GetString("ReversalAdmissionUpdateAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("ReversalAdmissionUpdateAttributes", "scheme_status_code_description"),

		SourceGateway: defaults.GetString("ReversalAdmissionUpdateAttributes", "source_gateway"),

		// TODO Status: ReversalAdmissionStatus,

	}
}

func (m *ReversalAdmissionUpdateAttributes) WithSchemeStatusCode(schemeStatusCode string) *ReversalAdmissionUpdateAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *ReversalAdmissionUpdateAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *ReversalAdmissionUpdateAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *ReversalAdmissionUpdateAttributes) WithSourceGateway(sourceGateway string) *ReversalAdmissionUpdateAttributes {

	m.SourceGateway = sourceGateway

	return m
}

func (m *ReversalAdmissionUpdateAttributes) WithStatus(status ReversalAdmissionStatus) *ReversalAdmissionUpdateAttributes {

	m.Status = status

	return m
}

// Validate validates this reversal admission update attributes
func (m *ReversalAdmissionUpdateAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalAdmissionUpdateAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalAdmissionUpdateAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalAdmissionUpdateAttributes) UnmarshalBinary(b []byte) error {
	var res ReversalAdmissionUpdateAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalAdmissionUpdateAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
