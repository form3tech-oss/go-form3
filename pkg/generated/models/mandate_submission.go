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

// MandateSubmission mandate submission
// swagger:model MandateSubmission
type MandateSubmission struct {

	// attributes
	Attributes *MandateSubmissionAttributes `json:"attributes,omitempty"`

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
	Relationships *MandateSubmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func MandateSubmissionWithDefaults(defaults client.Defaults) *MandateSubmission {
	return &MandateSubmission{

		Attributes: MandateSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("MandateSubmission", "created_on"),

		ID: defaults.GetStrfmtUUID("MandateSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("MandateSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("MandateSubmission", "organisation_id"),

		Relationships: MandateSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("MandateSubmission", "type"),

		Version: defaults.GetInt64Ptr("MandateSubmission", "version"),
	}
}

func (m *MandateSubmission) WithAttributes(attributes MandateSubmissionAttributes) *MandateSubmission {

	m.Attributes = &attributes

	return m
}

func (m *MandateSubmission) WithoutAttributes() *MandateSubmission {
	m.Attributes = nil
	return m
}

func (m *MandateSubmission) WithCreatedOn(createdOn strfmt.DateTime) *MandateSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *MandateSubmission) WithoutCreatedOn() *MandateSubmission {
	m.CreatedOn = nil
	return m
}

func (m *MandateSubmission) WithID(id strfmt.UUID) *MandateSubmission {

	m.ID = id

	return m
}

func (m *MandateSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *MandateSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *MandateSubmission) WithoutModifiedOn() *MandateSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *MandateSubmission) WithOrganisationID(organisationID strfmt.UUID) *MandateSubmission {

	m.OrganisationID = organisationID

	return m
}

func (m *MandateSubmission) WithRelationships(relationships MandateSubmissionRelationships) *MandateSubmission {

	m.Relationships = &relationships

	return m
}

func (m *MandateSubmission) WithoutRelationships() *MandateSubmission {
	m.Relationships = nil
	return m
}

func (m *MandateSubmission) WithType(typeVar string) *MandateSubmission {

	m.Type = typeVar

	return m
}

func (m *MandateSubmission) WithVersion(version int64) *MandateSubmission {

	m.Version = &version

	return m
}

func (m *MandateSubmission) WithoutVersion() *MandateSubmission {
	m.Version = nil
	return m
}

// Validate validates this mandate submission
func (m *MandateSubmission) Validate(formats strfmt.Registry) error {
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

func (m *MandateSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *MandateSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmission) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *MandateSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateSubmission) UnmarshalBinary(b []byte) error {
	var res MandateSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateSubmissionAttributes mandate submission attributes
// swagger:model MandateSubmissionAttributes
type MandateSubmissionAttributes struct {

	// file identifier
	// Pattern: ^[0-9a-zA-Z]+$
	FileIdentifier *string `json:"file_identifier,omitempty"`

	// file number
	// Pattern: ^[0-9]+$
	FileNumber *string `json:"file_number,omitempty"`

	// last payment date
	// Format: date
	LastPaymentDate *strfmt.Date `json:"last_payment_date,omitempty"`

	// original mandate
	OriginalMandate *MandateAttributes `json:"original_mandate,omitempty"`

	// status
	Status MandateSubmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// submission datetime
	// Format: date-time
	SubmissionDatetime strfmt.DateTime `json:"submission_datetime,omitempty"`

	// submission reason
	SubmissionReason string `json:"submission_reason,omitempty"`

	// submitted mandate
	SubmittedMandate *MandateAttributes `json:"submitted_mandate,omitempty"`
}

func MandateSubmissionAttributesWithDefaults(defaults client.Defaults) *MandateSubmissionAttributes {
	return &MandateSubmissionAttributes{

		FileIdentifier: defaults.GetStringPtr("MandateSubmissionAttributes", "file_identifier"),

		FileNumber: defaults.GetStringPtr("MandateSubmissionAttributes", "file_number"),

		LastPaymentDate: defaults.GetStrfmtDatePtr("MandateSubmissionAttributes", "last_payment_date"),

		OriginalMandate: MandateAttributesWithDefaults(defaults),

		// TODO Status: MandateSubmissionStatus,

		StatusReason: defaults.GetString("MandateSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTime("MandateSubmissionAttributes", "submission_datetime"),

		SubmissionReason: defaults.GetString("MandateSubmissionAttributes", "submission_reason"),

		SubmittedMandate: MandateAttributesWithDefaults(defaults),
	}
}

func (m *MandateSubmissionAttributes) WithFileIdentifier(fileIdentifier string) *MandateSubmissionAttributes {

	m.FileIdentifier = &fileIdentifier

	return m
}

func (m *MandateSubmissionAttributes) WithoutFileIdentifier() *MandateSubmissionAttributes {
	m.FileIdentifier = nil
	return m
}

func (m *MandateSubmissionAttributes) WithFileNumber(fileNumber string) *MandateSubmissionAttributes {

	m.FileNumber = &fileNumber

	return m
}

func (m *MandateSubmissionAttributes) WithoutFileNumber() *MandateSubmissionAttributes {
	m.FileNumber = nil
	return m
}

func (m *MandateSubmissionAttributes) WithLastPaymentDate(lastPaymentDate strfmt.Date) *MandateSubmissionAttributes {

	m.LastPaymentDate = &lastPaymentDate

	return m
}

func (m *MandateSubmissionAttributes) WithoutLastPaymentDate() *MandateSubmissionAttributes {
	m.LastPaymentDate = nil
	return m
}

func (m *MandateSubmissionAttributes) WithOriginalMandate(originalMandate MandateAttributes) *MandateSubmissionAttributes {

	m.OriginalMandate = &originalMandate

	return m
}

func (m *MandateSubmissionAttributes) WithoutOriginalMandate() *MandateSubmissionAttributes {
	m.OriginalMandate = nil
	return m
}

func (m *MandateSubmissionAttributes) WithStatus(status MandateSubmissionStatus) *MandateSubmissionAttributes {

	m.Status = status

	return m
}

func (m *MandateSubmissionAttributes) WithStatusReason(statusReason string) *MandateSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *MandateSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *MandateSubmissionAttributes {

	m.SubmissionDatetime = submissionDatetime

	return m
}

func (m *MandateSubmissionAttributes) WithSubmissionReason(submissionReason string) *MandateSubmissionAttributes {

	m.SubmissionReason = submissionReason

	return m
}

func (m *MandateSubmissionAttributes) WithSubmittedMandate(submittedMandate MandateAttributes) *MandateSubmissionAttributes {

	m.SubmittedMandate = &submittedMandate

	return m
}

func (m *MandateSubmissionAttributes) WithoutSubmittedMandate() *MandateSubmissionAttributes {
	m.SubmittedMandate = nil
	return m
}

// Validate validates this mandate submission attributes
func (m *MandateSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastPaymentDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalMandate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmittedMandate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateSubmissionAttributes) validateFileIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.FileIdentifier) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"file_identifier", "body", string(*m.FileIdentifier), `^[0-9a-zA-Z]+$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionAttributes) validateFileNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.FileNumber) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"file_number", "body", string(*m.FileNumber), `^[0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionAttributes) validateLastPaymentDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastPaymentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"last_payment_date", "body", "date", m.LastPaymentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionAttributes) validateOriginalMandate(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginalMandate) { // not required
		return nil
	}

	if m.OriginalMandate != nil {
		if err := m.OriginalMandate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "original_mandate")
			}
			return err
		}
	}

	return nil
}

func (m *MandateSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *MandateSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionAttributes) validateSubmittedMandate(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmittedMandate) { // not required
		return nil
	}

	if m.SubmittedMandate != nil {
		if err := m.SubmittedMandate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "submitted_mandate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res MandateSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
