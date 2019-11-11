// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportAdmission report admission
// swagger:model ReportAdmission
type ReportAdmission struct {

	// attributes
	Attributes *ReportAdmissionAttributes `json:"attributes,omitempty"`

	// created on
	// Read Only: true
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// modified on
	// Read Only: true
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *ReportAdmissionRelationships `json:"relationships,omitempty"`

	// type
	// Enum: [report_admissions]
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

// line 140

func ReportAdmissionWithDefaults(defaults client.Defaults) *ReportAdmission {
	return &ReportAdmission{

		Attributes: ReportAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("ReportAdmission", "created_on"),

		ID: defaults.GetStrfmtUUID("ReportAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("ReportAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("ReportAdmission", "organisation_id"),

		Relationships: ReportAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReportAdmission", "type"),

		Version: defaults.GetInt64Ptr("ReportAdmission", "version"),
	}
}

func (m *ReportAdmission) WithAttributes(attributes ReportAdmissionAttributes) *ReportAdmission {

	m.Attributes = &attributes

	return m
}

func (m *ReportAdmission) WithoutAttributes() *ReportAdmission {
	m.Attributes = nil
	return m
}

func (m *ReportAdmission) WithCreatedOn(createdOn strfmt.DateTime) *ReportAdmission {

	m.CreatedOn = createdOn

	return m
}

func (m *ReportAdmission) WithID(id strfmt.UUID) *ReportAdmission {

	m.ID = id

	return m
}

func (m *ReportAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *ReportAdmission {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *ReportAdmission) WithOrganisationID(organisationID strfmt.UUID) *ReportAdmission {

	m.OrganisationID = organisationID

	return m
}

func (m *ReportAdmission) WithRelationships(relationships ReportAdmissionRelationships) *ReportAdmission {

	m.Relationships = &relationships

	return m
}

func (m *ReportAdmission) WithoutRelationships() *ReportAdmission {
	m.Relationships = nil
	return m
}

func (m *ReportAdmission) WithType(typeVar string) *ReportAdmission {

	m.Type = typeVar

	return m
}

func (m *ReportAdmission) WithVersion(version int64) *ReportAdmission {

	m.Version = &version

	return m
}

func (m *ReportAdmission) WithoutVersion() *ReportAdmission {
	m.Version = nil
	return m
}

// Validate validates this report admission
func (m *ReportAdmission) Validate(formats strfmt.Registry) error {
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

func (m *ReportAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReportAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportAdmission) validateRelationships(formats strfmt.Registry) error {

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

var reportAdmissionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["report_admissions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportAdmissionTypeTypePropEnum = append(reportAdmissionTypeTypePropEnum, v)
	}
}

const (

	// ReportAdmissionTypeReportAdmissions captures enum value "report_admissions"
	ReportAdmissionTypeReportAdmissions string = "report_admissions"
)

// prop value enum
func (m *ReportAdmission) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, reportAdmissionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReportAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ReportAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportAdmission) UnmarshalBinary(b []byte) error {
	var res ReportAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
