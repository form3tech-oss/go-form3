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

// DirectDebitDecisionAdmission direct debit decision admission
// swagger:model DirectDebitDecisionAdmission
type DirectDebitDecisionAdmission struct {

	// attributes
	// Required: true
	Attributes *DirectDebitDecisionAdmissionAttributes `json:"attributes"`

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
	Relationships *DirectDebitDecisionAdmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitDecisionAdmissionWithDefaults(defaults client.Defaults) *DirectDebitDecisionAdmission {
	return &DirectDebitDecisionAdmission{

		Attributes: DirectDebitDecisionAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitDecisionAdmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitDecisionAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitDecisionAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebitDecisionAdmission", "organisation_id"),

		Relationships: DirectDebitDecisionAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitDecisionAdmission", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitDecisionAdmission", "version"),
	}
}

func (m *DirectDebitDecisionAdmission) WithAttributes(attributes DirectDebitDecisionAdmissionAttributes) *DirectDebitDecisionAdmission {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitDecisionAdmission) WithoutAttributes() *DirectDebitDecisionAdmission {
	m.Attributes = nil
	return m
}

func (m *DirectDebitDecisionAdmission) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitDecisionAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitDecisionAdmission) WithoutCreatedOn() *DirectDebitDecisionAdmission {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitDecisionAdmission) WithID(id strfmt.UUID) *DirectDebitDecisionAdmission {

	m.ID = &id

	return m
}

func (m *DirectDebitDecisionAdmission) WithoutID() *DirectDebitDecisionAdmission {
	m.ID = nil
	return m
}

func (m *DirectDebitDecisionAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitDecisionAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitDecisionAdmission) WithoutModifiedOn() *DirectDebitDecisionAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitDecisionAdmission) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitDecisionAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebitDecisionAdmission) WithoutOrganisationID() *DirectDebitDecisionAdmission {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebitDecisionAdmission) WithRelationships(relationships DirectDebitDecisionAdmissionRelationships) *DirectDebitDecisionAdmission {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitDecisionAdmission) WithoutRelationships() *DirectDebitDecisionAdmission {
	m.Relationships = nil
	return m
}

func (m *DirectDebitDecisionAdmission) WithType(typeVar string) *DirectDebitDecisionAdmission {

	m.Type = typeVar

	return m
}

func (m *DirectDebitDecisionAdmission) WithVersion(version int64) *DirectDebitDecisionAdmission {

	m.Version = &version

	return m
}

func (m *DirectDebitDecisionAdmission) WithoutVersion() *DirectDebitDecisionAdmission {
	m.Version = nil
	return m
}

// Validate validates this direct debit decision admission
func (m *DirectDebitDecisionAdmission) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitDecisionAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitDecisionAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecisionAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecisionAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecisionAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecisionAdmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitDecisionAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecisionAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecisionAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionAdmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitDecisionAdmissionAttributes direct debit decision admission attributes
// swagger:model DirectDebitDecisionAdmissionAttributes
type DirectDebitDecisionAdmissionAttributes struct {

	// admission datetime
	// Read Only: true
	// Format: date-time
	AdmissionDatetime strfmt.DateTime `json:"admission_datetime,omitempty"`

	// source gateway
	SourceGateway string `json:"source_gateway,omitempty"`

	// status
	Status DirectDebitDecisionAdmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason DirectDebitDecisionAdmissionStatusReason `json:"status_reason,omitempty"`
}

func DirectDebitDecisionAdmissionAttributesWithDefaults(defaults client.Defaults) *DirectDebitDecisionAdmissionAttributes {
	return &DirectDebitDecisionAdmissionAttributes{

		AdmissionDatetime: defaults.GetStrfmtDateTime("DirectDebitDecisionAdmissionAttributes", "admission_datetime"),

		SourceGateway: defaults.GetString("DirectDebitDecisionAdmissionAttributes", "source_gateway"),

		// TODO Status: DirectDebitDecisionAdmissionStatus,

		// TODO StatusReason: DirectDebitDecisionAdmissionStatusReason,

	}
}

func (m *DirectDebitDecisionAdmissionAttributes) WithAdmissionDatetime(admissionDatetime strfmt.DateTime) *DirectDebitDecisionAdmissionAttributes {

	m.AdmissionDatetime = admissionDatetime

	return m
}

func (m *DirectDebitDecisionAdmissionAttributes) WithSourceGateway(sourceGateway string) *DirectDebitDecisionAdmissionAttributes {

	m.SourceGateway = sourceGateway

	return m
}

func (m *DirectDebitDecisionAdmissionAttributes) WithStatus(status DirectDebitDecisionAdmissionStatus) *DirectDebitDecisionAdmissionAttributes {

	m.Status = status

	return m
}

func (m *DirectDebitDecisionAdmissionAttributes) WithStatusReason(statusReason DirectDebitDecisionAdmissionStatusReason) *DirectDebitDecisionAdmissionAttributes {

	m.StatusReason = statusReason

	return m
}

// Validate validates this direct debit decision admission attributes
func (m *DirectDebitDecisionAdmissionAttributes) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitDecisionAdmissionAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecisionAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *DirectDebitDecisionAdmissionAttributes) validateStatusReason(formats strfmt.Registry) error {

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
func (m *DirectDebitDecisionAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
