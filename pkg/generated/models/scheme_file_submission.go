// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SchemeFileSubmission scheme file submission
// swagger:model SchemeFileSubmission
type SchemeFileSubmission struct {

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// type
	// Enum: [scheme_file_submissions]
	Type string `json:"type,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`

	// created on
	// Read Only: true
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// modified on
	// Read Only: true
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// attributes
	// Required: true
	Attributes *SchemeFileSubmissionAttributes `json:"attributes"`

	// relationships
	// Read Only: true
	Relationships *SchemeFileSubmissionRelationships `json:"relationships,omitempty"`
}

func SchemeFileSubmissionWithDefaults(defaults client.Defaults) *SchemeFileSubmission {
	return &SchemeFileSubmission{

		ID: defaults.GetStrfmtUUIDPtr("SchemeFileSubmission", "id"),

		Type: defaults.GetString("SchemeFileSubmission", "type"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("SchemeFileSubmission", "organisation_id"),

		Version: defaults.GetInt64Ptr("SchemeFileSubmission", "version"),

		CreatedOn: defaults.GetStrfmtDateTime("SchemeFileSubmission", "created_on"),

		ModifiedOn: defaults.GetStrfmtDateTime("SchemeFileSubmission", "modified_on"),

		Attributes: SchemeFileSubmissionAttributesWithDefaults(defaults),

		Relationships: SchemeFileSubmissionRelationshipsWithDefaults(defaults),
	}
}

func (m *SchemeFileSubmission) WithID(id strfmt.UUID) *SchemeFileSubmission {

	m.ID = &id

	return m
}

func (m *SchemeFileSubmission) WithoutID() *SchemeFileSubmission {
	m.ID = nil
	return m
}

func (m *SchemeFileSubmission) WithType(typeVar string) *SchemeFileSubmission {

	m.Type = typeVar

	return m
}

func (m *SchemeFileSubmission) WithOrganisationID(organisationID strfmt.UUID) *SchemeFileSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *SchemeFileSubmission) WithoutOrganisationID() *SchemeFileSubmission {
	m.OrganisationID = nil
	return m
}

func (m *SchemeFileSubmission) WithVersion(version int64) *SchemeFileSubmission {

	m.Version = &version

	return m
}

func (m *SchemeFileSubmission) WithoutVersion() *SchemeFileSubmission {
	m.Version = nil
	return m
}

func (m *SchemeFileSubmission) WithCreatedOn(createdOn strfmt.DateTime) *SchemeFileSubmission {

	m.CreatedOn = createdOn

	return m
}

func (m *SchemeFileSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *SchemeFileSubmission {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *SchemeFileSubmission) WithAttributes(attributes SchemeFileSubmissionAttributes) *SchemeFileSubmission {

	m.Attributes = &attributes

	return m
}

func (m *SchemeFileSubmission) WithoutAttributes() *SchemeFileSubmission {
	m.Attributes = nil
	return m
}

func (m *SchemeFileSubmission) WithRelationships(relationships SchemeFileSubmissionRelationships) *SchemeFileSubmission {

	m.Relationships = &relationships

	return m
}

func (m *SchemeFileSubmission) WithoutRelationships() *SchemeFileSubmission {
	m.Relationships = nil
	return m
}

// Validate validates this scheme file submission
func (m *SchemeFileSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeFileSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var schemeFileSubmissionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["scheme_file_submissions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemeFileSubmissionTypeTypePropEnum = append(schemeFileSubmissionTypeTypePropEnum, v)
	}
}

const (

	// SchemeFileSubmissionTypeSchemeFileSubmissions captures enum value "scheme_file_submissions"
	SchemeFileSubmissionTypeSchemeFileSubmissions string = "scheme_file_submissions"
)

// prop value enum
func (m *SchemeFileSubmission) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, schemeFileSubmissionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SchemeFileSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *SchemeFileSubmission) validateRelationships(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *SchemeFileSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeFileSubmission) UnmarshalBinary(b []byte) error {
	var res SchemeFileSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeFileSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// SchemeFileSubmissionAttributes scheme file submission attributes
// swagger:model SchemeFileSubmissionAttributes
type SchemeFileSubmissionAttributes struct {

	// Time the submission request was received by Form3. Used to compute the total processing time
	// Read Only: true
	// Format: date-time
	StartDatetime strfmt.DateTime `json:"start_datetime,omitempty"`

	// status
	Status SchemeFileSubmissionStatus `json:"status,omitempty"`

	// Plain-text description of the status attribute
	// Read Only: true
	StatusReason string `json:"status_reason,omitempty"`

	// Time when the Form3 system begins processing of the submission
	// Read Only: true
	// Format: date-time
	SubmissionDatetime strfmt.DateTime `json:"submission_datetime,omitempty"`
}

func SchemeFileSubmissionAttributesWithDefaults(defaults client.Defaults) *SchemeFileSubmissionAttributes {
	return &SchemeFileSubmissionAttributes{

		StartDatetime: defaults.GetStrfmtDateTime("SchemeFileSubmissionAttributes", "start_datetime"),

		// TODO Status: SchemeFileSubmissionStatus,

		StatusReason: defaults.GetString("SchemeFileSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTime("SchemeFileSubmissionAttributes", "submission_datetime"),
	}
}

func (m *SchemeFileSubmissionAttributes) WithStartDatetime(startDatetime strfmt.DateTime) *SchemeFileSubmissionAttributes {

	m.StartDatetime = startDatetime

	return m
}

func (m *SchemeFileSubmissionAttributes) WithStatus(status SchemeFileSubmissionStatus) *SchemeFileSubmissionAttributes {

	m.Status = status

	return m
}

func (m *SchemeFileSubmissionAttributes) WithStatusReason(statusReason string) *SchemeFileSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *SchemeFileSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *SchemeFileSubmissionAttributes {

	m.SubmissionDatetime = submissionDatetime

	return m
}

// Validate validates this scheme file submission attributes
func (m *SchemeFileSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeFileSubmissionAttributes) validateStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"start_datetime", "body", "date-time", m.StartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *SchemeFileSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeFileSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeFileSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res SchemeFileSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeFileSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
