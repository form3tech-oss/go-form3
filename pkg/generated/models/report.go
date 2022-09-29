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

// Report report
// swagger:model Report
type Report struct {

	// attributes
	// Required: true
	Attributes *ReportAttributes `json:"attributes"`

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
	Relationships *ReportRelationships `json:"relationships,omitempty"`

	// type
	// Enum: [reports]
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReportWithDefaults(defaults client.Defaults) *Report {
	return &Report{

		Attributes: ReportAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("Report", "created_on"),

		ID: defaults.GetStrfmtUUID("Report", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("Report", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("Report", "organisation_id"),

		Relationships: ReportRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("Report", "type"),

		Version: defaults.GetInt64Ptr("Report", "version"),
	}
}

func (m *Report) WithAttributes(attributes ReportAttributes) *Report {

	m.Attributes = &attributes

	return m
}

func (m *Report) WithoutAttributes() *Report {
	m.Attributes = nil
	return m
}

func (m *Report) WithCreatedOn(createdOn strfmt.DateTime) *Report {

	m.CreatedOn = createdOn

	return m
}

func (m *Report) WithID(id strfmt.UUID) *Report {

	m.ID = id

	return m
}

func (m *Report) WithModifiedOn(modifiedOn strfmt.DateTime) *Report {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *Report) WithOrganisationID(organisationID strfmt.UUID) *Report {

	m.OrganisationID = organisationID

	return m
}

func (m *Report) WithRelationships(relationships ReportRelationships) *Report {

	m.Relationships = &relationships

	return m
}

func (m *Report) WithoutRelationships() *Report {
	m.Relationships = nil
	return m
}

func (m *Report) WithType(typeVar string) *Report {

	m.Type = typeVar

	return m
}

func (m *Report) WithVersion(version int64) *Report {

	m.Version = &version

	return m
}

func (m *Report) WithoutVersion() *Report {
	m.Version = nil
	return m
}

// Validate validates this report
func (m *Report) Validate(formats strfmt.Registry) error {
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

func (m *Report) validateAttributes(formats strfmt.Registry) error {

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

func (m *Report) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateRelationships(formats strfmt.Registry) error {

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

var reportTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["reports"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportTypeTypePropEnum = append(reportTypeTypePropEnum, v)
	}
}

const (

	// ReportTypeReports captures enum value "reports"
	ReportTypeReports string = "reports"
)

// prop value enum
func (m *Report) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, reportTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Report) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Report) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Report) UnmarshalBinary(b []byte) error {
	var res Report
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Report) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
