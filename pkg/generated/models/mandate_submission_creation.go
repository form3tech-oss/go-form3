// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MandateSubmissionCreation mandate submission creation
// swagger:model MandateSubmissionCreation
type MandateSubmissionCreation struct {

	// data
	Data *MandateSubmissionCreationData `json:"data,omitempty"`
}

func MandateSubmissionCreationWithDefaults(defaults client.Defaults) *MandateSubmissionCreation {
	return &MandateSubmissionCreation{

		Data: MandateSubmissionCreationDataWithDefaults(defaults),
	}
}

func (m *MandateSubmissionCreation) WithData(data MandateSubmissionCreationData) *MandateSubmissionCreation {

	m.Data = &data

	return m
}

func (m *MandateSubmissionCreation) WithoutData() *MandateSubmissionCreation {
	m.Data = nil
	return m
}

// Validate validates this mandate submission creation
func (m *MandateSubmissionCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateSubmissionCreation) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateSubmissionCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateSubmissionCreation) UnmarshalBinary(b []byte) error {
	var res MandateSubmissionCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateSubmissionCreation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateSubmissionCreationData mandate submission creation data
// swagger:model MandateSubmissionCreationData
type MandateSubmissionCreationData struct {

	// attributes
	Attributes *MandateSubmissionCreationDataAttributes `json:"attributes,omitempty"`

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

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func MandateSubmissionCreationDataWithDefaults(defaults client.Defaults) *MandateSubmissionCreationData {
	return &MandateSubmissionCreationData{

		Attributes: MandateSubmissionCreationDataAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("MandateSubmissionCreationData", "created_on"),

		ID: defaults.GetStrfmtUUID("MandateSubmissionCreationData", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("MandateSubmissionCreationData", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("MandateSubmissionCreationData", "organisation_id"),

		Type: defaults.GetString("MandateSubmissionCreationData", "type"),

		Version: defaults.GetInt64Ptr("MandateSubmissionCreationData", "version"),
	}
}

func (m *MandateSubmissionCreationData) WithAttributes(attributes MandateSubmissionCreationDataAttributes) *MandateSubmissionCreationData {

	m.Attributes = &attributes

	return m
}

func (m *MandateSubmissionCreationData) WithoutAttributes() *MandateSubmissionCreationData {
	m.Attributes = nil
	return m
}

func (m *MandateSubmissionCreationData) WithCreatedOn(createdOn strfmt.DateTime) *MandateSubmissionCreationData {

	m.CreatedOn = &createdOn

	return m
}

func (m *MandateSubmissionCreationData) WithoutCreatedOn() *MandateSubmissionCreationData {
	m.CreatedOn = nil
	return m
}

func (m *MandateSubmissionCreationData) WithID(id strfmt.UUID) *MandateSubmissionCreationData {

	m.ID = id

	return m
}

func (m *MandateSubmissionCreationData) WithModifiedOn(modifiedOn strfmt.DateTime) *MandateSubmissionCreationData {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *MandateSubmissionCreationData) WithoutModifiedOn() *MandateSubmissionCreationData {
	m.ModifiedOn = nil
	return m
}

func (m *MandateSubmissionCreationData) WithOrganisationID(organisationID strfmt.UUID) *MandateSubmissionCreationData {

	m.OrganisationID = organisationID

	return m
}

func (m *MandateSubmissionCreationData) WithType(typeVar string) *MandateSubmissionCreationData {

	m.Type = typeVar

	return m
}

func (m *MandateSubmissionCreationData) WithVersion(version int64) *MandateSubmissionCreationData {

	m.Version = &version

	return m
}

func (m *MandateSubmissionCreationData) WithoutVersion() *MandateSubmissionCreationData {
	m.Version = nil
	return m
}

// Validate validates this mandate submission creation data
func (m *MandateSubmissionCreationData) Validate(formats strfmt.Registry) error {
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

func (m *MandateSubmissionCreationData) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

func (m *MandateSubmissionCreationData) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionCreationData) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionCreationData) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionCreationData) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionCreationData) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("data"+"."+"type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionCreationData) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("data"+"."+"version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateSubmissionCreationData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateSubmissionCreationData) UnmarshalBinary(b []byte) error {
	var res MandateSubmissionCreationData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateSubmissionCreationData) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateSubmissionCreationDataAttributes mandate submission creation data attributes
// swagger:model MandateSubmissionCreationDataAttributes
type MandateSubmissionCreationDataAttributes struct {

	// last payment date
	// Format: date
	LastPaymentDate *strfmt.Date `json:"last_payment_date,omitempty"`

	// scheme file id
	// Max Length: 3
	// Min Length: 3
	// Pattern: ^[0-9]{3}$
	SchemeFileID *string `json:"scheme_file_id,omitempty"`

	// scheme submission id
	// Max Length: 6
	// Min Length: 6
	// Pattern: ^[0-9a-zA-Z]{6}$
	SchemeSubmissionID *string `json:"scheme_submission_id,omitempty"`

	// status
	Status MandateSubmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason *string `json:"status_reason,omitempty"`

	// submission datetime
	// Format: date-time
	SubmissionDatetime *strfmt.DateTime `json:"submission_datetime,omitempty"`

	// submission reason
	SubmissionReason string `json:"submission_reason,omitempty"`
}

func MandateSubmissionCreationDataAttributesWithDefaults(defaults client.Defaults) *MandateSubmissionCreationDataAttributes {
	return &MandateSubmissionCreationDataAttributes{

		LastPaymentDate: defaults.GetStrfmtDatePtr("MandateSubmissionCreationDataAttributes", "last_payment_date"),

		SchemeFileID: defaults.GetStringPtr("MandateSubmissionCreationDataAttributes", "scheme_file_id"),

		SchemeSubmissionID: defaults.GetStringPtr("MandateSubmissionCreationDataAttributes", "scheme_submission_id"),

		// TODO Status: MandateSubmissionStatus,

		StatusReason: defaults.GetStringPtr("MandateSubmissionCreationDataAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTimePtr("MandateSubmissionCreationDataAttributes", "submission_datetime"),

		SubmissionReason: defaults.GetString("MandateSubmissionCreationDataAttributes", "submission_reason"),
	}
}

func (m *MandateSubmissionCreationDataAttributes) WithLastPaymentDate(lastPaymentDate strfmt.Date) *MandateSubmissionCreationDataAttributes {

	m.LastPaymentDate = &lastPaymentDate

	return m
}

func (m *MandateSubmissionCreationDataAttributes) WithoutLastPaymentDate() *MandateSubmissionCreationDataAttributes {
	m.LastPaymentDate = nil
	return m
}

func (m *MandateSubmissionCreationDataAttributes) WithSchemeFileID(schemeFileID string) *MandateSubmissionCreationDataAttributes {

	m.SchemeFileID = &schemeFileID

	return m
}

func (m *MandateSubmissionCreationDataAttributes) WithoutSchemeFileID() *MandateSubmissionCreationDataAttributes {
	m.SchemeFileID = nil
	return m
}

func (m *MandateSubmissionCreationDataAttributes) WithSchemeSubmissionID(schemeSubmissionID string) *MandateSubmissionCreationDataAttributes {

	m.SchemeSubmissionID = &schemeSubmissionID

	return m
}

func (m *MandateSubmissionCreationDataAttributes) WithoutSchemeSubmissionID() *MandateSubmissionCreationDataAttributes {
	m.SchemeSubmissionID = nil
	return m
}

func (m *MandateSubmissionCreationDataAttributes) WithStatus(status MandateSubmissionStatus) *MandateSubmissionCreationDataAttributes {

	m.Status = status

	return m
}

func (m *MandateSubmissionCreationDataAttributes) WithStatusReason(statusReason string) *MandateSubmissionCreationDataAttributes {

	m.StatusReason = &statusReason

	return m
}

func (m *MandateSubmissionCreationDataAttributes) WithoutStatusReason() *MandateSubmissionCreationDataAttributes {
	m.StatusReason = nil
	return m
}

func (m *MandateSubmissionCreationDataAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *MandateSubmissionCreationDataAttributes {

	m.SubmissionDatetime = &submissionDatetime

	return m
}

func (m *MandateSubmissionCreationDataAttributes) WithoutSubmissionDatetime() *MandateSubmissionCreationDataAttributes {
	m.SubmissionDatetime = nil
	return m
}

func (m *MandateSubmissionCreationDataAttributes) WithSubmissionReason(submissionReason string) *MandateSubmissionCreationDataAttributes {

	m.SubmissionReason = submissionReason

	return m
}

// Validate validates this mandate submission creation data attributes
func (m *MandateSubmissionCreationDataAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastPaymentDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeFileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeSubmissionID(formats); err != nil {
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

func (m *MandateSubmissionCreationDataAttributes) validateLastPaymentDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastPaymentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"attributes"+"."+"last_payment_date", "body", "date", m.LastPaymentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionCreationDataAttributes) validateSchemeFileID(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeFileID) { // not required
		return nil
	}

	if err := validate.MinLength("data"+"."+"attributes"+"."+"scheme_file_id", "body", string(*m.SchemeFileID), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("data"+"."+"attributes"+"."+"scheme_file_id", "body", string(*m.SchemeFileID), 3); err != nil {
		return err
	}

	if err := validate.Pattern("data"+"."+"attributes"+"."+"scheme_file_id", "body", string(*m.SchemeFileID), `^[0-9]{3}$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionCreationDataAttributes) validateSchemeSubmissionID(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeSubmissionID) { // not required
		return nil
	}

	if err := validate.MinLength("data"+"."+"attributes"+"."+"scheme_submission_id", "body", string(*m.SchemeSubmissionID), 6); err != nil {
		return err
	}

	if err := validate.MaxLength("data"+"."+"attributes"+"."+"scheme_submission_id", "body", string(*m.SchemeSubmissionID), 6); err != nil {
		return err
	}

	if err := validate.Pattern("data"+"."+"attributes"+"."+"scheme_submission_id", "body", string(*m.SchemeSubmissionID), `^[0-9a-zA-Z]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionCreationDataAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "attributes" + "." + "status")
		}
		return err
	}

	return nil
}

func (m *MandateSubmissionCreationDataAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateSubmissionCreationDataAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateSubmissionCreationDataAttributes) UnmarshalBinary(b []byte) error {
	var res MandateSubmissionCreationDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateSubmissionCreationDataAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
