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

// DirectDebitSubmission direct debit submission
// swagger:model DirectDebitSubmission
type DirectDebitSubmission struct {

	// attributes
	Attributes *DirectDebitSubmissionAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// Unique ID
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
	Relationships *DirectDebitSubmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version of the resource
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitSubmissionWithDefaults(defaults client.Defaults) *DirectDebitSubmission {
	return &DirectDebitSubmission{

		Attributes: DirectDebitSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebitSubmission", "organisation_id"),

		Relationships: DirectDebitSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitSubmission", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitSubmission", "version"),
	}
}

func (m *DirectDebitSubmission) WithAttributes(attributes DirectDebitSubmissionAttributes) *DirectDebitSubmission {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitSubmission) WithoutAttributes() *DirectDebitSubmission {
	m.Attributes = nil
	return m
}

func (m *DirectDebitSubmission) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitSubmission) WithoutCreatedOn() *DirectDebitSubmission {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitSubmission) WithID(id strfmt.UUID) *DirectDebitSubmission {

	m.ID = &id

	return m
}

func (m *DirectDebitSubmission) WithoutID() *DirectDebitSubmission {
	m.ID = nil
	return m
}

func (m *DirectDebitSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitSubmission) WithoutModifiedOn() *DirectDebitSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitSubmission) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebitSubmission) WithoutOrganisationID() *DirectDebitSubmission {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebitSubmission) WithRelationships(relationships DirectDebitSubmissionRelationships) *DirectDebitSubmission {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitSubmission) WithoutRelationships() *DirectDebitSubmission {
	m.Relationships = nil
	return m
}

func (m *DirectDebitSubmission) WithType(typeVar string) *DirectDebitSubmission {

	m.Type = typeVar

	return m
}

func (m *DirectDebitSubmission) WithVersion(version int64) *DirectDebitSubmission {

	m.Version = &version

	return m
}

func (m *DirectDebitSubmission) WithoutVersion() *DirectDebitSubmission {
	m.Version = nil
	return m
}

// Validate validates this direct debit submission
func (m *DirectDebitSubmission) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitSubmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitSubmissionAttributes direct debit submission attributes
// swagger:model DirectDebitSubmissionAttributes
type DirectDebitSubmissionAttributes struct {

	// file identifier
	// Pattern: ^[0-9a-zA-Z]+$
	FileIdentifier *string `json:"file_identifier,omitempty"`

	// file number
	// Pattern: ^[0-9]+$
	FileNumber *string `json:"file_number,omitempty"`

	// Scheme-specific status (if submission has been submitted to a scheme)
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// [Description](http://api-docs.form3.tech/api.html#enumerations-scheme-status-codes-for-bacs) of `scheme_status_code`
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// Status of the submission
	Status DirectDebitSubmissionStatus `json:"status,omitempty"`

	// Description of the submission status
	StatusReason string `json:"status_reason,omitempty"`

	// Date of the submission
	// Read Only: true
	// Format: date-time
	SubmissionDatetime strfmt.DateTime `json:"submission_datetime,omitempty"`

	// Time the request was received by Form3. Used to compute the total transaction time of an operation.
	// Read Only: true
	// Format: date-time
	TransactionStartDatetime strfmt.DateTime `json:"transaction_start_datetime,omitempty"`
}

func DirectDebitSubmissionAttributesWithDefaults(defaults client.Defaults) *DirectDebitSubmissionAttributes {
	return &DirectDebitSubmissionAttributes{

		FileIdentifier: defaults.GetStringPtr("DirectDebitSubmissionAttributes", "file_identifier"),

		FileNumber: defaults.GetStringPtr("DirectDebitSubmissionAttributes", "file_number"),

		SchemeStatusCode: defaults.GetString("DirectDebitSubmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("DirectDebitSubmissionAttributes", "scheme_status_code_description"),

		// TODO Status: DirectDebitSubmissionStatus,

		StatusReason: defaults.GetString("DirectDebitSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTime("DirectDebitSubmissionAttributes", "submission_datetime"),

		TransactionStartDatetime: defaults.GetStrfmtDateTime("DirectDebitSubmissionAttributes", "transaction_start_datetime"),
	}
}

func (m *DirectDebitSubmissionAttributes) WithFileIdentifier(fileIdentifier string) *DirectDebitSubmissionAttributes {

	m.FileIdentifier = &fileIdentifier

	return m
}

func (m *DirectDebitSubmissionAttributes) WithoutFileIdentifier() *DirectDebitSubmissionAttributes {
	m.FileIdentifier = nil
	return m
}

func (m *DirectDebitSubmissionAttributes) WithFileNumber(fileNumber string) *DirectDebitSubmissionAttributes {

	m.FileNumber = &fileNumber

	return m
}

func (m *DirectDebitSubmissionAttributes) WithoutFileNumber() *DirectDebitSubmissionAttributes {
	m.FileNumber = nil
	return m
}

func (m *DirectDebitSubmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *DirectDebitSubmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *DirectDebitSubmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *DirectDebitSubmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *DirectDebitSubmissionAttributes) WithStatus(status DirectDebitSubmissionStatus) *DirectDebitSubmissionAttributes {

	m.Status = status

	return m
}

func (m *DirectDebitSubmissionAttributes) WithStatusReason(statusReason string) *DirectDebitSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *DirectDebitSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *DirectDebitSubmissionAttributes {

	m.SubmissionDatetime = submissionDatetime

	return m
}

func (m *DirectDebitSubmissionAttributes) WithTransactionStartDatetime(transactionStartDatetime strfmt.DateTime) *DirectDebitSubmissionAttributes {

	m.TransactionStartDatetime = transactionStartDatetime

	return m
}

// Validate validates this direct debit submission attributes
func (m *DirectDebitSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitSubmissionAttributes) validateFileIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.FileIdentifier) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"file_identifier", "body", string(*m.FileIdentifier), `^[0-9a-zA-Z]+$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitSubmissionAttributes) validateFileNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.FileNumber) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"file_number", "body", string(*m.FileNumber), `^[0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *DirectDebitSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitSubmissionAttributes) validateTransactionStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"transaction_start_datetime", "body", "date-time", m.TransactionStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
