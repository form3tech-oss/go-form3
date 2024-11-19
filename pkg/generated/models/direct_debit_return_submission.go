// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectDebitReturnSubmission direct debit return submission
// swagger:model DirectDebitReturnSubmission
type DirectDebitReturnSubmission struct {

	// attributes
	Attributes *DirectDebitReturnSubmissionAttributes `json:"attributes,omitempty"`

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
	Relationships *DirectDebitReturnSubmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitReturnSubmissionWithDefaults(defaults client.Defaults) *DirectDebitReturnSubmission {
	return &DirectDebitReturnSubmission{

		Attributes: DirectDebitReturnSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturnSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitReturnSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturnSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebitReturnSubmission", "organisation_id"),

		Relationships: DirectDebitReturnSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitReturnSubmission", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitReturnSubmission", "version"),
	}
}

func (m *DirectDebitReturnSubmission) WithAttributes(attributes DirectDebitReturnSubmissionAttributes) *DirectDebitReturnSubmission {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitReturnSubmission) WithoutAttributes() *DirectDebitReturnSubmission {
	m.Attributes = nil
	return m
}

func (m *DirectDebitReturnSubmission) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitReturnSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitReturnSubmission) WithoutCreatedOn() *DirectDebitReturnSubmission {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitReturnSubmission) WithID(id strfmt.UUID) *DirectDebitReturnSubmission {

	m.ID = &id

	return m
}

func (m *DirectDebitReturnSubmission) WithoutID() *DirectDebitReturnSubmission {
	m.ID = nil
	return m
}

func (m *DirectDebitReturnSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitReturnSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitReturnSubmission) WithoutModifiedOn() *DirectDebitReturnSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitReturnSubmission) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitReturnSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebitReturnSubmission) WithoutOrganisationID() *DirectDebitReturnSubmission {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebitReturnSubmission) WithRelationships(relationships DirectDebitReturnSubmissionRelationships) *DirectDebitReturnSubmission {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitReturnSubmission) WithoutRelationships() *DirectDebitReturnSubmission {
	m.Relationships = nil
	return m
}

func (m *DirectDebitReturnSubmission) WithType(typeVar string) *DirectDebitReturnSubmission {

	m.Type = typeVar

	return m
}

func (m *DirectDebitReturnSubmission) WithVersion(version int64) *DirectDebitReturnSubmission {

	m.Version = &version

	return m
}

func (m *DirectDebitReturnSubmission) WithoutVersion() *DirectDebitReturnSubmission {
	m.Version = nil
	return m
}

// Validate validates this direct debit return submission
func (m *DirectDebitReturnSubmission) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitReturnSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitReturnSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitReturnSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnSubmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnSubmissionAttributes direct debit return submission attributes
// swagger:model DirectDebitReturnSubmissionAttributes
type DirectDebitReturnSubmissionAttributes struct {

	// file identifier
	// Pattern: ^[0-9a-zA-Z]+$
	FileIdentifier *string `json:"file_identifier,omitempty"`

	// file number
	// Pattern: ^[0-9]+$
	FileNumber *string `json:"file_number,omitempty"`

	// scheme status code
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// scheme status code description
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// status
	Status DirectDebitReturnSubmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// submission datetime
	// Read Only: true
	// Format: date-time
	SubmissionDatetime strfmt.DateTime `json:"submission_datetime,omitempty"`

	// transaction start datetime
	// Read Only: true
	// Format: date-time
	TransactionStartDatetime strfmt.DateTime `json:"transaction_start_datetime,omitempty"`
}

func DirectDebitReturnSubmissionAttributesWithDefaults(defaults client.Defaults) *DirectDebitReturnSubmissionAttributes {
	return &DirectDebitReturnSubmissionAttributes{

		FileIdentifier: defaults.GetStringPtr("DirectDebitReturnSubmissionAttributes", "file_identifier"),

		FileNumber: defaults.GetStringPtr("DirectDebitReturnSubmissionAttributes", "file_number"),

		SchemeStatusCode: defaults.GetString("DirectDebitReturnSubmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("DirectDebitReturnSubmissionAttributes", "scheme_status_code_description"),

		// TODO Status: DirectDebitReturnSubmissionStatus,

		StatusReason: defaults.GetString("DirectDebitReturnSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTime("DirectDebitReturnSubmissionAttributes", "submission_datetime"),

		TransactionStartDatetime: defaults.GetStrfmtDateTime("DirectDebitReturnSubmissionAttributes", "transaction_start_datetime"),
	}
}

func (m *DirectDebitReturnSubmissionAttributes) WithFileIdentifier(fileIdentifier string) *DirectDebitReturnSubmissionAttributes {

	m.FileIdentifier = &fileIdentifier

	return m
}

func (m *DirectDebitReturnSubmissionAttributes) WithoutFileIdentifier() *DirectDebitReturnSubmissionAttributes {
	m.FileIdentifier = nil
	return m
}

func (m *DirectDebitReturnSubmissionAttributes) WithFileNumber(fileNumber string) *DirectDebitReturnSubmissionAttributes {

	m.FileNumber = &fileNumber

	return m
}

func (m *DirectDebitReturnSubmissionAttributes) WithoutFileNumber() *DirectDebitReturnSubmissionAttributes {
	m.FileNumber = nil
	return m
}

func (m *DirectDebitReturnSubmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *DirectDebitReturnSubmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *DirectDebitReturnSubmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *DirectDebitReturnSubmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *DirectDebitReturnSubmissionAttributes) WithStatus(status DirectDebitReturnSubmissionStatus) *DirectDebitReturnSubmissionAttributes {

	m.Status = status

	return m
}

func (m *DirectDebitReturnSubmissionAttributes) WithStatusReason(statusReason string) *DirectDebitReturnSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *DirectDebitReturnSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *DirectDebitReturnSubmissionAttributes {

	m.SubmissionDatetime = submissionDatetime

	return m
}

func (m *DirectDebitReturnSubmissionAttributes) WithTransactionStartDatetime(transactionStartDatetime strfmt.DateTime) *DirectDebitReturnSubmissionAttributes {

	m.TransactionStartDatetime = transactionStartDatetime

	return m
}

// Validate validates this direct debit return submission attributes
func (m *DirectDebitReturnSubmissionAttributes) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitReturnSubmissionAttributes) validateFileIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.FileIdentifier) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"file_identifier", "body", string(*m.FileIdentifier), `^[0-9a-zA-Z]+$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnSubmissionAttributes) validateFileNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.FileNumber) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"file_number", "body", string(*m.FileNumber), `^[0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *DirectDebitReturnSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnSubmissionAttributes) validateTransactionStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"transaction_start_datetime", "body", "date-time", m.TransactionStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnSubmissionRelationships direct debit return submission relationships
// swagger:model DirectDebitReturnSubmissionRelationships
type DirectDebitReturnSubmissionRelationships struct {

	// direct debit
	DirectDebit *DirectDebitReturnSubmissionRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit return
	DirectDebitReturn *DirectDebitReturnSubmissionRelationshipsDirectDebitReturn `json:"direct_debit_return,omitempty"`
}

func DirectDebitReturnSubmissionRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitReturnSubmissionRelationships {
	return &DirectDebitReturnSubmissionRelationships{

		DirectDebit: DirectDebitReturnSubmissionRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitReturn: DirectDebitReturnSubmissionRelationshipsDirectDebitReturnWithDefaults(defaults),
	}
}

func (m *DirectDebitReturnSubmissionRelationships) WithDirectDebit(directDebit DirectDebitReturnSubmissionRelationshipsDirectDebit) *DirectDebitReturnSubmissionRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitReturnSubmissionRelationships) WithoutDirectDebit() *DirectDebitReturnSubmissionRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitReturnSubmissionRelationships) WithDirectDebitReturn(directDebitReturn DirectDebitReturnSubmissionRelationshipsDirectDebitReturn) *DirectDebitReturnSubmissionRelationships {

	m.DirectDebitReturn = &directDebitReturn

	return m
}

func (m *DirectDebitReturnSubmissionRelationships) WithoutDirectDebitReturn() *DirectDebitReturnSubmissionRelationships {
	m.DirectDebitReturn = nil
	return m
}

// Validate validates this direct debit return submission relationships
func (m *DirectDebitReturnSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReturn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnSubmissionRelationships) validateDirectDebit(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebit) { // not required
		return nil
	}

	if m.DirectDebit != nil {
		if err := m.DirectDebit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitReturnSubmissionRelationships) validateDirectDebitReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReturn) { // not required
		return nil
	}

	if m.DirectDebitReturn != nil {
		if err := m.DirectDebitReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_return")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnSubmissionRelationshipsDirectDebit direct debit return submission relationships direct debit
// swagger:model DirectDebitReturnSubmissionRelationshipsDirectDebit
type DirectDebitReturnSubmissionRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitReturnSubmissionRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitReturnSubmissionRelationshipsDirectDebit {
	return &DirectDebitReturnSubmissionRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitReturnSubmissionRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitReturnSubmissionRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit return submission relationships direct debit
func (m *DirectDebitReturnSubmissionRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnSubmissionRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationships" + "." + "direct_debit" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnSubmissionRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnSubmissionRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnSubmissionRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnSubmissionRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnSubmissionRelationshipsDirectDebitReturn direct debit return submission relationships direct debit return
// swagger:model DirectDebitReturnSubmissionRelationshipsDirectDebitReturn
type DirectDebitReturnSubmissionRelationshipsDirectDebitReturn struct {

	// data
	Data []*DirectDebitReturn `json:"data"`
}

func DirectDebitReturnSubmissionRelationshipsDirectDebitReturnWithDefaults(defaults client.Defaults) *DirectDebitReturnSubmissionRelationshipsDirectDebitReturn {
	return &DirectDebitReturnSubmissionRelationshipsDirectDebitReturn{

		Data: make([]*DirectDebitReturn, 0),
	}
}

func (m *DirectDebitReturnSubmissionRelationshipsDirectDebitReturn) WithData(data []*DirectDebitReturn) *DirectDebitReturnSubmissionRelationshipsDirectDebitReturn {

	m.Data = data

	return m
}

// Validate validates this direct debit return submission relationships direct debit return
func (m *DirectDebitReturnSubmissionRelationshipsDirectDebitReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnSubmissionRelationshipsDirectDebitReturn) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationships" + "." + "direct_debit_return" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnSubmissionRelationshipsDirectDebitReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnSubmissionRelationshipsDirectDebitReturn) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnSubmissionRelationshipsDirectDebitReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnSubmissionRelationshipsDirectDebitReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
