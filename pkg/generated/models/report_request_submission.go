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

// ReportRequestSubmission report request submission
// swagger:model ReportRequestSubmission
type ReportRequestSubmission struct {

	// attributes
	// Required: true
	Attributes *ReportRequestSubmissionAttributes `json:"attributes"`

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
	// Required: true
	Relationships *ReportRequestSubmissionRelationships `json:"relationships"`

	// type
	// Required: true
	// Enum: ["report_admissions"]
	Type string `json:"type"`

	// version
	// Required: true
	// Minimum: 0
	Version *int64 `json:"version"`
}

func ReportRequestSubmissionWithDefaults(defaults client.Defaults) *ReportRequestSubmission {
	return &ReportRequestSubmission{

		Attributes: ReportRequestSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("ReportRequestSubmission", "created_on"),

		ID: defaults.GetStrfmtUUID("ReportRequestSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("ReportRequestSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("ReportRequestSubmission", "organisation_id"),

		Relationships: ReportRequestSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReportRequestSubmission", "type"),

		Version: defaults.GetInt64Ptr("ReportRequestSubmission", "version"),
	}
}

func (m *ReportRequestSubmission) WithAttributes(attributes ReportRequestSubmissionAttributes) *ReportRequestSubmission {

	m.Attributes = &attributes

	return m
}

func (m *ReportRequestSubmission) WithoutAttributes() *ReportRequestSubmission {
	m.Attributes = nil
	return m
}

func (m *ReportRequestSubmission) WithCreatedOn(createdOn strfmt.DateTime) *ReportRequestSubmission {

	m.CreatedOn = createdOn

	return m
}

func (m *ReportRequestSubmission) WithID(id strfmt.UUID) *ReportRequestSubmission {

	m.ID = id

	return m
}

func (m *ReportRequestSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *ReportRequestSubmission {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *ReportRequestSubmission) WithOrganisationID(organisationID strfmt.UUID) *ReportRequestSubmission {

	m.OrganisationID = organisationID

	return m
}

func (m *ReportRequestSubmission) WithRelationships(relationships ReportRequestSubmissionRelationships) *ReportRequestSubmission {

	m.Relationships = &relationships

	return m
}

func (m *ReportRequestSubmission) WithoutRelationships() *ReportRequestSubmission {
	m.Relationships = nil
	return m
}

func (m *ReportRequestSubmission) WithType(typeVar string) *ReportRequestSubmission {

	m.Type = typeVar

	return m
}

func (m *ReportRequestSubmission) WithVersion(version int64) *ReportRequestSubmission {

	m.Version = &version

	return m
}

func (m *ReportRequestSubmission) WithoutVersion() *ReportRequestSubmission {
	m.Version = nil
	return m
}

// Validate validates this report request submission
func (m *ReportRequestSubmission) Validate(formats strfmt.Registry) error {
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

func (m *ReportRequestSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReportRequestSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportRequestSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportRequestSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportRequestSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportRequestSubmission) validateRelationships(formats strfmt.Registry) error {

	if err := validate.Required("relationships", "body", m.Relationships); err != nil {
		return err
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

var reportRequestSubmissionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["report_admissions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportRequestSubmissionTypeTypePropEnum = append(reportRequestSubmissionTypeTypePropEnum, v)
	}
}

const (

	// ReportRequestSubmissionTypeReportAdmissions captures enum value "report_admissions"
	ReportRequestSubmissionTypeReportAdmissions string = "report_admissions"
)

// prop value enum
func (m *ReportRequestSubmission) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, reportRequestSubmissionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReportRequestSubmission) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ReportRequestSubmission) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportRequestSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportRequestSubmission) UnmarshalBinary(b []byte) error {
	var res ReportRequestSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportRequestSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
