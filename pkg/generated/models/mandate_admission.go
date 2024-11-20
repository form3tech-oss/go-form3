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

// MandateAdmission mandate admission
// swagger:model MandateAdmission
type MandateAdmission struct {

	// attributes
	Attributes *MandateAdmissionAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// organisation id
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// relationships
	Relationships *MandateAdmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func MandateAdmissionWithDefaults(defaults client.Defaults) *MandateAdmission {
	return &MandateAdmission{

		Attributes: MandateAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("MandateAdmission", "created_on"),

		ID: defaults.GetStrfmtUUID("MandateAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("MandateAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("MandateAdmission", "organisation_id"),

		Relationships: MandateAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("MandateAdmission", "type"),

		Version: defaults.GetInt64Ptr("MandateAdmission", "version"),
	}
}

func (m *MandateAdmission) WithAttributes(attributes MandateAdmissionAttributes) *MandateAdmission {

	m.Attributes = &attributes

	return m
}

func (m *MandateAdmission) WithoutAttributes() *MandateAdmission {
	m.Attributes = nil
	return m
}

func (m *MandateAdmission) WithCreatedOn(createdOn strfmt.DateTime) *MandateAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *MandateAdmission) WithoutCreatedOn() *MandateAdmission {
	m.CreatedOn = nil
	return m
}

func (m *MandateAdmission) WithID(id strfmt.UUID) *MandateAdmission {

	m.ID = id

	return m
}

func (m *MandateAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *MandateAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *MandateAdmission) WithoutModifiedOn() *MandateAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *MandateAdmission) WithOrganisationID(organisationID strfmt.UUID) *MandateAdmission {

	m.OrganisationID = organisationID

	return m
}

func (m *MandateAdmission) WithRelationships(relationships MandateAdmissionRelationships) *MandateAdmission {

	m.Relationships = &relationships

	return m
}

func (m *MandateAdmission) WithoutRelationships() *MandateAdmission {
	m.Relationships = nil
	return m
}

func (m *MandateAdmission) WithType(typeVar string) *MandateAdmission {

	m.Type = typeVar

	return m
}

func (m *MandateAdmission) WithVersion(version int64) *MandateAdmission {

	m.Version = &version

	return m
}

func (m *MandateAdmission) WithoutVersion() *MandateAdmission {
	m.Version = nil
	return m
}

// Validate validates this mandate admission
func (m *MandateAdmission) Validate(formats strfmt.Registry) error {
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

func (m *MandateAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *MandateAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateAdmission) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateAdmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *MandateAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateAdmission) UnmarshalBinary(b []byte) error {
	var res MandateAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateAdmissionAttributes mandate admission attributes
// swagger:model MandateAdmissionAttributes
type MandateAdmissionAttributes struct {

	// admission datetime
	// Read Only: true
	// Format: date-time
	AdmissionDatetime strfmt.DateTime `json:"admission_datetime,omitempty"`

	// scheme status code
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// scheme status code description
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// status
	Status MandateAdmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason MandateAdmissionStatusReason `json:"status_reason,omitempty"`
}

func MandateAdmissionAttributesWithDefaults(defaults client.Defaults) *MandateAdmissionAttributes {
	return &MandateAdmissionAttributes{

		AdmissionDatetime: defaults.GetStrfmtDateTime("MandateAdmissionAttributes", "admission_datetime"),

		SchemeStatusCode: defaults.GetString("MandateAdmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("MandateAdmissionAttributes", "scheme_status_code_description"),

		// TODO Status: MandateAdmissionStatus,

		// TODO StatusReason: MandateAdmissionStatusReason,

	}
}

func (m *MandateAdmissionAttributes) WithAdmissionDatetime(admissionDatetime strfmt.DateTime) *MandateAdmissionAttributes {

	m.AdmissionDatetime = admissionDatetime

	return m
}

func (m *MandateAdmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *MandateAdmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *MandateAdmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *MandateAdmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *MandateAdmissionAttributes) WithStatus(status MandateAdmissionStatus) *MandateAdmissionAttributes {

	m.Status = status

	return m
}

func (m *MandateAdmissionAttributes) WithStatusReason(statusReason MandateAdmissionStatusReason) *MandateAdmissionAttributes {

	m.StatusReason = statusReason

	return m
}

// Validate validates this mandate admission attributes
func (m *MandateAdmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateAdmissionAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *MandateAdmissionAttributes) validateStatusReason(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusReason) { // not required
		return nil
	}

	if err := m.StatusReason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "status_reason")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res MandateAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
