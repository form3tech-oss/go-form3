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

// ReturnSubmission return submission
// swagger:model ReturnSubmission
type ReturnSubmission struct {

	// attributes
	Attributes *ReturnSubmissionAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// Unique resource ID
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
	Relationships *ReturnSubmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReturnSubmissionWithDefaults(defaults client.Defaults) *ReturnSubmission {
	return &ReturnSubmission{

		Attributes: ReturnSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("ReturnSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ReturnSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("ReturnSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ReturnSubmission", "organisation_id"),

		Relationships: ReturnSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReturnSubmission", "type"),

		Version: defaults.GetInt64Ptr("ReturnSubmission", "version"),
	}
}

func (m *ReturnSubmission) WithAttributes(attributes ReturnSubmissionAttributes) *ReturnSubmission {

	m.Attributes = &attributes

	return m
}

func (m *ReturnSubmission) WithoutAttributes() *ReturnSubmission {
	m.Attributes = nil
	return m
}

func (m *ReturnSubmission) WithCreatedOn(createdOn strfmt.DateTime) *ReturnSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *ReturnSubmission) WithoutCreatedOn() *ReturnSubmission {
	m.CreatedOn = nil
	return m
}

func (m *ReturnSubmission) WithID(id strfmt.UUID) *ReturnSubmission {

	m.ID = &id

	return m
}

func (m *ReturnSubmission) WithoutID() *ReturnSubmission {
	m.ID = nil
	return m
}

func (m *ReturnSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *ReturnSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *ReturnSubmission) WithoutModifiedOn() *ReturnSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *ReturnSubmission) WithOrganisationID(organisationID strfmt.UUID) *ReturnSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *ReturnSubmission) WithoutOrganisationID() *ReturnSubmission {
	m.OrganisationID = nil
	return m
}

func (m *ReturnSubmission) WithRelationships(relationships ReturnSubmissionRelationships) *ReturnSubmission {

	m.Relationships = &relationships

	return m
}

func (m *ReturnSubmission) WithoutRelationships() *ReturnSubmission {
	m.Relationships = nil
	return m
}

func (m *ReturnSubmission) WithType(typeVar string) *ReturnSubmission {

	m.Type = typeVar

	return m
}

func (m *ReturnSubmission) WithVersion(version int64) *ReturnSubmission {

	m.Version = &version

	return m
}

func (m *ReturnSubmission) WithoutVersion() *ReturnSubmission {
	m.Version = nil
	return m
}

// Validate validates this return submission
func (m *ReturnSubmission) Validate(formats strfmt.Registry) error {
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

func (m *ReturnSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReturnSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReturnSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmission) UnmarshalBinary(b []byte) error {
	var res ReturnSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReturnSubmissionAttributes return submission attributes
// swagger:model ReturnSubmissionAttributes
type ReturnSubmissionAttributes struct {

	// Indicates if the submission was created automatically by the system (true) or manually (false)
	Auto *bool `json:"auto,omitempty"`

	// Identification code of the file sent to scheme.
	// Pattern: ^[0-9a-zA-Z]+$
	FileIdentifier *string `json:"file_identifier,omitempty"`

	// Number of the file sent to scheme.
	// Pattern: ^[0-9]+$
	FileNumber *string `json:"file_number,omitempty"`

	// Time a payment was released from being held due to a limit breach
	// Read Only: true
	// Format: date-time
	LimitBreachEndDatetime *strfmt.DateTime `json:"limit_breach_end_datetime,omitempty"`

	// Start time a payment was held due to a limit breach
	// Read Only: true
	// Format: date-time
	LimitBreachStartDatetime *strfmt.DateTime `json:"limit_breach_start_datetime,omitempty"`

	// posting status
	PostingStatus PostingStatus `json:"posting_status,omitempty"`

	// Details of the account to which funds are redirected (if applicable)
	RedirectedAccountNumber string `json:"redirected_account_number,omitempty"`

	// Details of the bank to which funds are redirected (if applicable)
	RedirectedBankID string `json:"redirected_bank_id,omitempty"`

	// Additional payment reference assigned by the scheme
	ReferenceID string `json:"reference_id,omitempty"`

	// Route taken for a return
	// Enum: ["on_us","xp"]
	Route string `json:"route,omitempty"`

	// Scheme-specific status (if submission has been submitted to a scheme)
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// [Description](http://api-docs.form3.tech/api.html#enumerations-scheme-status-codes-for-bacs) of `scheme_status_code`
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// Cycle in which the payment will be settled
	// Minimum: 0
	SettlementCycle *int64 `json:"settlement_cycle,omitempty"`

	// Date that the payment will be settled
	// Format: date
	SettlementDate *strfmt.Date `json:"settlement_date,omitempty"`

	// status
	Status ReturnSubmissionStatus `json:"status,omitempty"`

	// Description of the submission status
	StatusReason string `json:"status_reason,omitempty"`

	// Time of the submission
	// Read Only: true
	// Format: date-time
	SubmissionDatetime *strfmt.DateTime `json:"submission_datetime,omitempty"`

	// Time the request was received by Form3. Used to compute the total transaction time of a payment.
	// Read Only: true
	// Format: date-time
	TransactionStartDatetime *strfmt.DateTime `json:"transaction_start_datetime,omitempty"`
}

func ReturnSubmissionAttributesWithDefaults(defaults client.Defaults) *ReturnSubmissionAttributes {
	return &ReturnSubmissionAttributes{

		Auto: defaults.GetBoolPtr("ReturnSubmissionAttributes", "auto"),

		FileIdentifier: defaults.GetStringPtr("ReturnSubmissionAttributes", "file_identifier"),

		FileNumber: defaults.GetStringPtr("ReturnSubmissionAttributes", "file_number"),

		LimitBreachEndDatetime: defaults.GetStrfmtDateTimePtr("ReturnSubmissionAttributes", "limit_breach_end_datetime"),

		LimitBreachStartDatetime: defaults.GetStrfmtDateTimePtr("ReturnSubmissionAttributes", "limit_breach_start_datetime"),

		// TODO PostingStatus: PostingStatus,

		RedirectedAccountNumber: defaults.GetString("ReturnSubmissionAttributes", "redirected_account_number"),

		RedirectedBankID: defaults.GetString("ReturnSubmissionAttributes", "redirected_bank_id"),

		ReferenceID: defaults.GetString("ReturnSubmissionAttributes", "reference_id"),

		Route: defaults.GetString("ReturnSubmissionAttributes", "route"),

		SchemeStatusCode: defaults.GetString("ReturnSubmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("ReturnSubmissionAttributes", "scheme_status_code_description"),

		SettlementCycle: defaults.GetInt64Ptr("ReturnSubmissionAttributes", "settlement_cycle"),

		SettlementDate: defaults.GetStrfmtDatePtr("ReturnSubmissionAttributes", "settlement_date"),

		// TODO Status: ReturnSubmissionStatus,

		StatusReason: defaults.GetString("ReturnSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTimePtr("ReturnSubmissionAttributes", "submission_datetime"),

		TransactionStartDatetime: defaults.GetStrfmtDateTimePtr("ReturnSubmissionAttributes", "transaction_start_datetime"),
	}
}

func (m *ReturnSubmissionAttributes) WithAuto(auto bool) *ReturnSubmissionAttributes {

	m.Auto = &auto

	return m
}

func (m *ReturnSubmissionAttributes) WithoutAuto() *ReturnSubmissionAttributes {
	m.Auto = nil
	return m
}

func (m *ReturnSubmissionAttributes) WithFileIdentifier(fileIdentifier string) *ReturnSubmissionAttributes {

	m.FileIdentifier = &fileIdentifier

	return m
}

func (m *ReturnSubmissionAttributes) WithoutFileIdentifier() *ReturnSubmissionAttributes {
	m.FileIdentifier = nil
	return m
}

func (m *ReturnSubmissionAttributes) WithFileNumber(fileNumber string) *ReturnSubmissionAttributes {

	m.FileNumber = &fileNumber

	return m
}

func (m *ReturnSubmissionAttributes) WithoutFileNumber() *ReturnSubmissionAttributes {
	m.FileNumber = nil
	return m
}

func (m *ReturnSubmissionAttributes) WithLimitBreachEndDatetime(limitBreachEndDatetime strfmt.DateTime) *ReturnSubmissionAttributes {

	m.LimitBreachEndDatetime = &limitBreachEndDatetime

	return m
}

func (m *ReturnSubmissionAttributes) WithoutLimitBreachEndDatetime() *ReturnSubmissionAttributes {
	m.LimitBreachEndDatetime = nil
	return m
}

func (m *ReturnSubmissionAttributes) WithLimitBreachStartDatetime(limitBreachStartDatetime strfmt.DateTime) *ReturnSubmissionAttributes {

	m.LimitBreachStartDatetime = &limitBreachStartDatetime

	return m
}

func (m *ReturnSubmissionAttributes) WithoutLimitBreachStartDatetime() *ReturnSubmissionAttributes {
	m.LimitBreachStartDatetime = nil
	return m
}

func (m *ReturnSubmissionAttributes) WithPostingStatus(postingStatus PostingStatus) *ReturnSubmissionAttributes {

	m.PostingStatus = postingStatus

	return m
}

func (m *ReturnSubmissionAttributes) WithRedirectedAccountNumber(redirectedAccountNumber string) *ReturnSubmissionAttributes {

	m.RedirectedAccountNumber = redirectedAccountNumber

	return m
}

func (m *ReturnSubmissionAttributes) WithRedirectedBankID(redirectedBankID string) *ReturnSubmissionAttributes {

	m.RedirectedBankID = redirectedBankID

	return m
}

func (m *ReturnSubmissionAttributes) WithReferenceID(referenceID string) *ReturnSubmissionAttributes {

	m.ReferenceID = referenceID

	return m
}

func (m *ReturnSubmissionAttributes) WithRoute(route string) *ReturnSubmissionAttributes {

	m.Route = route

	return m
}

func (m *ReturnSubmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *ReturnSubmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *ReturnSubmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *ReturnSubmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *ReturnSubmissionAttributes) WithSettlementCycle(settlementCycle int64) *ReturnSubmissionAttributes {

	m.SettlementCycle = &settlementCycle

	return m
}

func (m *ReturnSubmissionAttributes) WithoutSettlementCycle() *ReturnSubmissionAttributes {
	m.SettlementCycle = nil
	return m
}

func (m *ReturnSubmissionAttributes) WithSettlementDate(settlementDate strfmt.Date) *ReturnSubmissionAttributes {

	m.SettlementDate = &settlementDate

	return m
}

func (m *ReturnSubmissionAttributes) WithoutSettlementDate() *ReturnSubmissionAttributes {
	m.SettlementDate = nil
	return m
}

func (m *ReturnSubmissionAttributes) WithStatus(status ReturnSubmissionStatus) *ReturnSubmissionAttributes {

	m.Status = status

	return m
}

func (m *ReturnSubmissionAttributes) WithStatusReason(statusReason string) *ReturnSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *ReturnSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *ReturnSubmissionAttributes {

	m.SubmissionDatetime = &submissionDatetime

	return m
}

func (m *ReturnSubmissionAttributes) WithoutSubmissionDatetime() *ReturnSubmissionAttributes {
	m.SubmissionDatetime = nil
	return m
}

func (m *ReturnSubmissionAttributes) WithTransactionStartDatetime(transactionStartDatetime strfmt.DateTime) *ReturnSubmissionAttributes {

	m.TransactionStartDatetime = &transactionStartDatetime

	return m
}

func (m *ReturnSubmissionAttributes) WithoutTransactionStartDatetime() *ReturnSubmissionAttributes {
	m.TransactionStartDatetime = nil
	return m
}

// Validate validates this return submission attributes
func (m *ReturnSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitBreachEndDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitBreachStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlementCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlementDate(formats); err != nil {
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

func (m *ReturnSubmissionAttributes) validateFileIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.FileIdentifier) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"file_identifier", "body", *m.FileIdentifier, `^[0-9a-zA-Z]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateFileNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.FileNumber) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"file_number", "body", *m.FileNumber, `^[0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateLimitBreachEndDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitBreachEndDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"limit_breach_end_datetime", "body", "date-time", m.LimitBreachEndDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateLimitBreachStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitBreachStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"limit_breach_start_datetime", "body", "date-time", m.LimitBreachStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validatePostingStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.PostingStatus) { // not required
		return nil
	}

	if err := m.PostingStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "posting_status")
		}
		return err
	}

	return nil
}

var returnSubmissionAttributesTypeRoutePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on_us","xp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnSubmissionAttributesTypeRoutePropEnum = append(returnSubmissionAttributesTypeRoutePropEnum, v)
	}
}

const (

	// ReturnSubmissionAttributesRouteOnUs captures enum value "on_us"
	ReturnSubmissionAttributesRouteOnUs string = "on_us"

	// ReturnSubmissionAttributesRouteXp captures enum value "xp"
	ReturnSubmissionAttributesRouteXp string = "xp"
)

// prop value enum
func (m *ReturnSubmissionAttributes) validateRouteEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, returnSubmissionAttributesTypeRoutePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReturnSubmissionAttributes) validateRoute(formats strfmt.Registry) error {

	if swag.IsZero(m.Route) { // not required
		return nil
	}

	// value enum
	if err := m.validateRouteEnum("attributes"+"."+"route", "body", m.Route); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateSettlementCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementCycle) { // not required
		return nil
	}

	if err := validate.MinimumInt("attributes"+"."+"settlement_cycle", "body", int64(*m.SettlementCycle), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateSettlementDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"settlement_date", "body", "date", m.SettlementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *ReturnSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionAttributes) validateTransactionStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"transaction_start_datetime", "body", "date-time", m.TransactionStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
