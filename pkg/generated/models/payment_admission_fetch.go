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

// PaymentAdmissionFetch payment admission fetch
// swagger:model PaymentAdmissionFetch
type PaymentAdmissionFetch struct {

	// attributes
	Attributes *PaymentAdmissionFetchAttributes `json:"attributes,omitempty"`

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
	Relationships *PaymentAdmissionFetchRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func PaymentAdmissionFetchWithDefaults(defaults client.Defaults) *PaymentAdmissionFetch {
	return &PaymentAdmissionFetch{

		Attributes: PaymentAdmissionFetchAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("PaymentAdmissionFetch", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("PaymentAdmissionFetch", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("PaymentAdmissionFetch", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("PaymentAdmissionFetch", "organisation_id"),

		Relationships: PaymentAdmissionFetchRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("PaymentAdmissionFetch", "type"),

		Version: defaults.GetInt64Ptr("PaymentAdmissionFetch", "version"),
	}
}

func (m *PaymentAdmissionFetch) WithAttributes(attributes PaymentAdmissionFetchAttributes) *PaymentAdmissionFetch {

	m.Attributes = &attributes

	return m
}

func (m *PaymentAdmissionFetch) WithoutAttributes() *PaymentAdmissionFetch {
	m.Attributes = nil
	return m
}

func (m *PaymentAdmissionFetch) WithCreatedOn(createdOn strfmt.DateTime) *PaymentAdmissionFetch {

	m.CreatedOn = &createdOn

	return m
}

func (m *PaymentAdmissionFetch) WithoutCreatedOn() *PaymentAdmissionFetch {
	m.CreatedOn = nil
	return m
}

func (m *PaymentAdmissionFetch) WithID(id strfmt.UUID) *PaymentAdmissionFetch {

	m.ID = &id

	return m
}

func (m *PaymentAdmissionFetch) WithoutID() *PaymentAdmissionFetch {
	m.ID = nil
	return m
}

func (m *PaymentAdmissionFetch) WithModifiedOn(modifiedOn strfmt.DateTime) *PaymentAdmissionFetch {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *PaymentAdmissionFetch) WithoutModifiedOn() *PaymentAdmissionFetch {
	m.ModifiedOn = nil
	return m
}

func (m *PaymentAdmissionFetch) WithOrganisationID(organisationID strfmt.UUID) *PaymentAdmissionFetch {

	m.OrganisationID = &organisationID

	return m
}

func (m *PaymentAdmissionFetch) WithoutOrganisationID() *PaymentAdmissionFetch {
	m.OrganisationID = nil
	return m
}

func (m *PaymentAdmissionFetch) WithRelationships(relationships PaymentAdmissionFetchRelationships) *PaymentAdmissionFetch {

	m.Relationships = &relationships

	return m
}

func (m *PaymentAdmissionFetch) WithoutRelationships() *PaymentAdmissionFetch {
	m.Relationships = nil
	return m
}

func (m *PaymentAdmissionFetch) WithType(typeVar string) *PaymentAdmissionFetch {

	m.Type = typeVar

	return m
}

func (m *PaymentAdmissionFetch) WithVersion(version int64) *PaymentAdmissionFetch {

	m.Version = &version

	return m
}

func (m *PaymentAdmissionFetch) WithoutVersion() *PaymentAdmissionFetch {
	m.Version = nil
	return m
}

// Validate validates this payment admission fetch
func (m *PaymentAdmissionFetch) Validate(formats strfmt.Registry) error {
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

func (m *PaymentAdmissionFetch) validateAttributes(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionFetch) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetch) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetch) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetch) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetch) validateRelationships(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionFetch) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetch) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdmissionFetch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionFetch) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionFetch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionFetch) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentAdmissionFetchAttributes payment admission fetch attributes
// swagger:model PaymentAdmissionFetchAttributes
type PaymentAdmissionFetchAttributes struct {

	// account validation outcome
	AccountValidationOutcome AccountValidationOutcome `json:"account_validation_outcome,omitempty"`

	// Date and time the payment admission was created
	// Read Only: true
	// Format: date-time
	AdmissionDatetime *strfmt.DateTime `json:"admission_datetime,omitempty"`

	// Clearing infrastructure through which the payment instruction was processed
	// Pattern: ^[0-9A-Za-z_]*$
	ClearingSystem string `json:"clearing_system,omitempty"`

	// posting status
	PostingStatus PostingStatus `json:"posting_status,omitempty"`

	// Additional payment reference assigned by the scheme
	ReferenceID string `json:"reference_id,omitempty"`

	// Route taken for an inbound payment
	// Enum: ["on_us","xp"]
	Route string `json:"route,omitempty"`

	// Date and time the inbound payment was received. Supports RFC3339 Nano
	// Pattern: ^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,9}Z$
	// Format: date-time
	SchemeReceivedDatetime *strfmt.DateTime `json:"scheme_received_datetime,omitempty"`

	// Refers to individual scheme where applicable
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// [Description](http://api-docs.form3.tech/api.html#enumerations-scheme-status-codes-for-bacs) of `scheme_status_code`
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// Cycle in which the payment will be settled
	// Minimum: 0
	SettlementCycle *int64 `json:"settlement_cycle,omitempty"`

	// Date on which the payment will be settled
	// Format: date
	SettlementDate *strfmt.Date `json:"settlement_date,omitempty"`

	// status
	Status PaymentAdmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason PaymentAdmissionStatusReason `json:"status_reason,omitempty"`

	// Date and time the inbound payment was successfully validated by Form3. Supports RFC3339 Nano
	// Pattern: ^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,9}Z$
	// Format: date-time
	ValidationCompletedDatetime *strfmt.DateTime `json:"validation_completed_datetime,omitempty"`
}

func PaymentAdmissionFetchAttributesWithDefaults(defaults client.Defaults) *PaymentAdmissionFetchAttributes {
	return &PaymentAdmissionFetchAttributes{

		// TODO AccountValidationOutcome: AccountValidationOutcome,

		AdmissionDatetime: defaults.GetStrfmtDateTimePtr("PaymentAdmissionFetchAttributes", "admission_datetime"),

		ClearingSystem: defaults.GetString("PaymentAdmissionFetchAttributes", "clearing_system"),

		// TODO PostingStatus: PostingStatus,

		ReferenceID: defaults.GetString("PaymentAdmissionFetchAttributes", "reference_id"),

		Route: defaults.GetString("PaymentAdmissionFetchAttributes", "route"),

		SchemeReceivedDatetime: defaults.GetStrfmtDateTimePtr("PaymentAdmissionFetchAttributes", "scheme_received_datetime"),

		SchemeStatusCode: defaults.GetString("PaymentAdmissionFetchAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("PaymentAdmissionFetchAttributes", "scheme_status_code_description"),

		SettlementCycle: defaults.GetInt64Ptr("PaymentAdmissionFetchAttributes", "settlement_cycle"),

		SettlementDate: defaults.GetStrfmtDatePtr("PaymentAdmissionFetchAttributes", "settlement_date"),

		// TODO Status: PaymentAdmissionStatus,

		// TODO StatusReason: PaymentAdmissionStatusReason,

		ValidationCompletedDatetime: defaults.GetStrfmtDateTimePtr("PaymentAdmissionFetchAttributes", "validation_completed_datetime"),
	}
}

func (m *PaymentAdmissionFetchAttributes) WithAccountValidationOutcome(accountValidationOutcome AccountValidationOutcome) *PaymentAdmissionFetchAttributes {

	m.AccountValidationOutcome = accountValidationOutcome

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithAdmissionDatetime(admissionDatetime strfmt.DateTime) *PaymentAdmissionFetchAttributes {

	m.AdmissionDatetime = &admissionDatetime

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithoutAdmissionDatetime() *PaymentAdmissionFetchAttributes {
	m.AdmissionDatetime = nil
	return m
}

func (m *PaymentAdmissionFetchAttributes) WithClearingSystem(clearingSystem string) *PaymentAdmissionFetchAttributes {

	m.ClearingSystem = clearingSystem

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithPostingStatus(postingStatus PostingStatus) *PaymentAdmissionFetchAttributes {

	m.PostingStatus = postingStatus

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithReferenceID(referenceID string) *PaymentAdmissionFetchAttributes {

	m.ReferenceID = referenceID

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithRoute(route string) *PaymentAdmissionFetchAttributes {

	m.Route = route

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithSchemeReceivedDatetime(schemeReceivedDatetime strfmt.DateTime) *PaymentAdmissionFetchAttributes {

	m.SchemeReceivedDatetime = &schemeReceivedDatetime

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithoutSchemeReceivedDatetime() *PaymentAdmissionFetchAttributes {
	m.SchemeReceivedDatetime = nil
	return m
}

func (m *PaymentAdmissionFetchAttributes) WithSchemeStatusCode(schemeStatusCode string) *PaymentAdmissionFetchAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *PaymentAdmissionFetchAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithSettlementCycle(settlementCycle int64) *PaymentAdmissionFetchAttributes {

	m.SettlementCycle = &settlementCycle

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithoutSettlementCycle() *PaymentAdmissionFetchAttributes {
	m.SettlementCycle = nil
	return m
}

func (m *PaymentAdmissionFetchAttributes) WithSettlementDate(settlementDate strfmt.Date) *PaymentAdmissionFetchAttributes {

	m.SettlementDate = &settlementDate

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithoutSettlementDate() *PaymentAdmissionFetchAttributes {
	m.SettlementDate = nil
	return m
}

func (m *PaymentAdmissionFetchAttributes) WithStatus(status PaymentAdmissionStatus) *PaymentAdmissionFetchAttributes {

	m.Status = status

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithStatusReason(statusReason PaymentAdmissionStatusReason) *PaymentAdmissionFetchAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithValidationCompletedDatetime(validationCompletedDatetime strfmt.DateTime) *PaymentAdmissionFetchAttributes {

	m.ValidationCompletedDatetime = &validationCompletedDatetime

	return m
}

func (m *PaymentAdmissionFetchAttributes) WithoutValidationCompletedDatetime() *PaymentAdmissionFetchAttributes {
	m.ValidationCompletedDatetime = nil
	return m
}

// Validate validates this payment admission fetch attributes
func (m *PaymentAdmissionFetchAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountValidationOutcome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClearingSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeReceivedDatetime(formats); err != nil {
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

	if err := m.validateStatusReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationCompletedDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdmissionFetchAttributes) validateAccountValidationOutcome(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountValidationOutcome) { // not required
		return nil
	}

	if err := m.AccountValidationOutcome.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "account_validation_outcome")
		}
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetchAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetchAttributes) validateClearingSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.ClearingSystem) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"clearing_system", "body", m.ClearingSystem, `^[0-9A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetchAttributes) validatePostingStatus(formats strfmt.Registry) error {

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

var paymentAdmissionFetchAttributesTypeRoutePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on_us","xp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentAdmissionFetchAttributesTypeRoutePropEnum = append(paymentAdmissionFetchAttributesTypeRoutePropEnum, v)
	}
}

const (

	// PaymentAdmissionFetchAttributesRouteOnUs captures enum value "on_us"
	PaymentAdmissionFetchAttributesRouteOnUs string = "on_us"

	// PaymentAdmissionFetchAttributesRouteXp captures enum value "xp"
	PaymentAdmissionFetchAttributesRouteXp string = "xp"
)

// prop value enum
func (m *PaymentAdmissionFetchAttributes) validateRouteEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentAdmissionFetchAttributesTypeRoutePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentAdmissionFetchAttributes) validateRoute(formats strfmt.Registry) error {

	if swag.IsZero(m.Route) { // not required
		return nil
	}

	// value enum
	if err := m.validateRouteEnum("attributes"+"."+"route", "body", m.Route); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetchAttributes) validateSchemeReceivedDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeReceivedDatetime) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"scheme_received_datetime", "body", m.SchemeReceivedDatetime.String(), `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,9}Z$`); err != nil {
		return err
	}

	if err := validate.FormatOf("attributes"+"."+"scheme_received_datetime", "body", "date-time", m.SchemeReceivedDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetchAttributes) validateSettlementCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementCycle) { // not required
		return nil
	}

	if err := validate.MinimumInt("attributes"+"."+"settlement_cycle", "body", int64(*m.SettlementCycle), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetchAttributes) validateSettlementDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"settlement_date", "body", "date", m.SettlementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionFetchAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionFetchAttributes) validateStatusReason(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionFetchAttributes) validateValidationCompletedDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationCompletedDatetime) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"validation_completed_datetime", "body", m.ValidationCompletedDatetime.String(), `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,9}Z$`); err != nil {
		return err
	}

	if err := validate.FormatOf("attributes"+"."+"validation_completed_datetime", "body", "date-time", m.ValidationCompletedDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdmissionFetchAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionFetchAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionFetchAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionFetchAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentAdmissionFetchRelationships payment admission fetch relationships
// swagger:model PaymentAdmissionFetchRelationships
type PaymentAdmissionFetchRelationships struct {

	// beneficiary account
	BeneficiaryAccount *RelationshipPaymentAdmissionBeneficiaryAccount `json:"beneficiary_account,omitempty"`

	// beneficiary branch
	BeneficiaryBranch *RelationshipPaymentAdmissionBeneficiaryBranch `json:"beneficiary_branch,omitempty"`

	// payment
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// payment admission task
	PaymentAdmissionTask *RelationshipPaymentAdmissionTasks `json:"payment_admission_task,omitempty"`
}

func PaymentAdmissionFetchRelationshipsWithDefaults(defaults client.Defaults) *PaymentAdmissionFetchRelationships {
	return &PaymentAdmissionFetchRelationships{

		BeneficiaryAccount: RelationshipPaymentAdmissionBeneficiaryAccountWithDefaults(defaults),

		BeneficiaryBranch: RelationshipPaymentAdmissionBeneficiaryBranchWithDefaults(defaults),

		Payment: RelationshipLinksWithDefaults(defaults),

		PaymentAdmissionTask: RelationshipPaymentAdmissionTasksWithDefaults(defaults),
	}
}

func (m *PaymentAdmissionFetchRelationships) WithBeneficiaryAccount(beneficiaryAccount RelationshipPaymentAdmissionBeneficiaryAccount) *PaymentAdmissionFetchRelationships {

	m.BeneficiaryAccount = &beneficiaryAccount

	return m
}

func (m *PaymentAdmissionFetchRelationships) WithoutBeneficiaryAccount() *PaymentAdmissionFetchRelationships {
	m.BeneficiaryAccount = nil
	return m
}

func (m *PaymentAdmissionFetchRelationships) WithBeneficiaryBranch(beneficiaryBranch RelationshipPaymentAdmissionBeneficiaryBranch) *PaymentAdmissionFetchRelationships {

	m.BeneficiaryBranch = &beneficiaryBranch

	return m
}

func (m *PaymentAdmissionFetchRelationships) WithoutBeneficiaryBranch() *PaymentAdmissionFetchRelationships {
	m.BeneficiaryBranch = nil
	return m
}

func (m *PaymentAdmissionFetchRelationships) WithPayment(payment RelationshipLinks) *PaymentAdmissionFetchRelationships {

	m.Payment = &payment

	return m
}

func (m *PaymentAdmissionFetchRelationships) WithoutPayment() *PaymentAdmissionFetchRelationships {
	m.Payment = nil
	return m
}

func (m *PaymentAdmissionFetchRelationships) WithPaymentAdmissionTask(paymentAdmissionTask RelationshipPaymentAdmissionTasks) *PaymentAdmissionFetchRelationships {

	m.PaymentAdmissionTask = &paymentAdmissionTask

	return m
}

func (m *PaymentAdmissionFetchRelationships) WithoutPaymentAdmissionTask() *PaymentAdmissionFetchRelationships {
	m.PaymentAdmissionTask = nil
	return m
}

// Validate validates this payment admission fetch relationships
func (m *PaymentAdmissionFetchRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBeneficiaryAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentAdmissionTask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdmissionFetchRelationships) validateBeneficiaryAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.BeneficiaryAccount) { // not required
		return nil
	}

	if m.BeneficiaryAccount != nil {
		if err := m.BeneficiaryAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "beneficiary_account")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAdmissionFetchRelationships) validateBeneficiaryBranch(formats strfmt.Registry) error {

	if swag.IsZero(m.BeneficiaryBranch) { // not required
		return nil
	}

	if m.BeneficiaryBranch != nil {
		if err := m.BeneficiaryBranch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "beneficiary_branch")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAdmissionFetchRelationships) validatePayment(formats strfmt.Registry) error {

	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "payment")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAdmissionFetchRelationships) validatePaymentAdmissionTask(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentAdmissionTask) { // not required
		return nil
	}

	if m.PaymentAdmissionTask != nil {
		if err := m.PaymentAdmissionTask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "payment_admission_task")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdmissionFetchRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionFetchRelationships) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionFetchRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionFetchRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
