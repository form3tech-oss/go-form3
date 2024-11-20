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

// PaymentAdmission payment admission
// swagger:model PaymentAdmission
type PaymentAdmission struct {

	// attributes
	Attributes *PaymentAdmissionAttributes `json:"attributes,omitempty"`

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
	Relationships *PaymentAdmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func PaymentAdmissionWithDefaults(defaults client.Defaults) *PaymentAdmission {
	return &PaymentAdmission{

		Attributes: PaymentAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("PaymentAdmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("PaymentAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("PaymentAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("PaymentAdmission", "organisation_id"),

		Relationships: PaymentAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("PaymentAdmission", "type"),

		Version: defaults.GetInt64Ptr("PaymentAdmission", "version"),
	}
}

func (m *PaymentAdmission) WithAttributes(attributes PaymentAdmissionAttributes) *PaymentAdmission {

	m.Attributes = &attributes

	return m
}

func (m *PaymentAdmission) WithoutAttributes() *PaymentAdmission {
	m.Attributes = nil
	return m
}

func (m *PaymentAdmission) WithCreatedOn(createdOn strfmt.DateTime) *PaymentAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *PaymentAdmission) WithoutCreatedOn() *PaymentAdmission {
	m.CreatedOn = nil
	return m
}

func (m *PaymentAdmission) WithID(id strfmt.UUID) *PaymentAdmission {

	m.ID = &id

	return m
}

func (m *PaymentAdmission) WithoutID() *PaymentAdmission {
	m.ID = nil
	return m
}

func (m *PaymentAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *PaymentAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *PaymentAdmission) WithoutModifiedOn() *PaymentAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *PaymentAdmission) WithOrganisationID(organisationID strfmt.UUID) *PaymentAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *PaymentAdmission) WithoutOrganisationID() *PaymentAdmission {
	m.OrganisationID = nil
	return m
}

func (m *PaymentAdmission) WithRelationships(relationships PaymentAdmissionRelationships) *PaymentAdmission {

	m.Relationships = &relationships

	return m
}

func (m *PaymentAdmission) WithoutRelationships() *PaymentAdmission {
	m.Relationships = nil
	return m
}

func (m *PaymentAdmission) WithType(typeVar string) *PaymentAdmission {

	m.Type = typeVar

	return m
}

func (m *PaymentAdmission) WithVersion(version int64) *PaymentAdmission {

	m.Version = &version

	return m
}

func (m *PaymentAdmission) WithoutVersion() *PaymentAdmission {
	m.Version = nil
	return m
}

// Validate validates this payment admission
func (m *PaymentAdmission) Validate(formats strfmt.Registry) error {
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

func (m *PaymentAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *PaymentAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *PaymentAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmission) UnmarshalBinary(b []byte) error {
	var res PaymentAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentAdmissionAttributes payment admission attributes
// swagger:model PaymentAdmissionAttributes
type PaymentAdmissionAttributes struct {

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
}

func PaymentAdmissionAttributesWithDefaults(defaults client.Defaults) *PaymentAdmissionAttributes {
	return &PaymentAdmissionAttributes{

		// TODO AccountValidationOutcome: AccountValidationOutcome,

		AdmissionDatetime: defaults.GetStrfmtDateTimePtr("PaymentAdmissionAttributes", "admission_datetime"),

		ClearingSystem: defaults.GetString("PaymentAdmissionAttributes", "clearing_system"),

		// TODO PostingStatus: PostingStatus,

		ReferenceID: defaults.GetString("PaymentAdmissionAttributes", "reference_id"),

		Route: defaults.GetString("PaymentAdmissionAttributes", "route"),

		SchemeStatusCode: defaults.GetString("PaymentAdmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("PaymentAdmissionAttributes", "scheme_status_code_description"),

		SettlementCycle: defaults.GetInt64Ptr("PaymentAdmissionAttributes", "settlement_cycle"),

		SettlementDate: defaults.GetStrfmtDatePtr("PaymentAdmissionAttributes", "settlement_date"),

		// TODO Status: PaymentAdmissionStatus,

		// TODO StatusReason: PaymentAdmissionStatusReason,

	}
}

func (m *PaymentAdmissionAttributes) WithAccountValidationOutcome(accountValidationOutcome AccountValidationOutcome) *PaymentAdmissionAttributes {

	m.AccountValidationOutcome = accountValidationOutcome

	return m
}

func (m *PaymentAdmissionAttributes) WithAdmissionDatetime(admissionDatetime strfmt.DateTime) *PaymentAdmissionAttributes {

	m.AdmissionDatetime = &admissionDatetime

	return m
}

func (m *PaymentAdmissionAttributes) WithoutAdmissionDatetime() *PaymentAdmissionAttributes {
	m.AdmissionDatetime = nil
	return m
}

func (m *PaymentAdmissionAttributes) WithClearingSystem(clearingSystem string) *PaymentAdmissionAttributes {

	m.ClearingSystem = clearingSystem

	return m
}

func (m *PaymentAdmissionAttributes) WithPostingStatus(postingStatus PostingStatus) *PaymentAdmissionAttributes {

	m.PostingStatus = postingStatus

	return m
}

func (m *PaymentAdmissionAttributes) WithReferenceID(referenceID string) *PaymentAdmissionAttributes {

	m.ReferenceID = referenceID

	return m
}

func (m *PaymentAdmissionAttributes) WithRoute(route string) *PaymentAdmissionAttributes {

	m.Route = route

	return m
}

func (m *PaymentAdmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *PaymentAdmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *PaymentAdmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *PaymentAdmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *PaymentAdmissionAttributes) WithSettlementCycle(settlementCycle int64) *PaymentAdmissionAttributes {

	m.SettlementCycle = &settlementCycle

	return m
}

func (m *PaymentAdmissionAttributes) WithoutSettlementCycle() *PaymentAdmissionAttributes {
	m.SettlementCycle = nil
	return m
}

func (m *PaymentAdmissionAttributes) WithSettlementDate(settlementDate strfmt.Date) *PaymentAdmissionAttributes {

	m.SettlementDate = &settlementDate

	return m
}

func (m *PaymentAdmissionAttributes) WithoutSettlementDate() *PaymentAdmissionAttributes {
	m.SettlementDate = nil
	return m
}

func (m *PaymentAdmissionAttributes) WithStatus(status PaymentAdmissionStatus) *PaymentAdmissionAttributes {

	m.Status = status

	return m
}

func (m *PaymentAdmissionAttributes) WithStatusReason(statusReason PaymentAdmissionStatusReason) *PaymentAdmissionAttributes {

	m.StatusReason = statusReason

	return m
}

// Validate validates this payment admission attributes
func (m *PaymentAdmissionAttributes) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdmissionAttributes) validateAccountValidationOutcome(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionAttributes) validateClearingSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.ClearingSystem) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"clearing_system", "body", string(m.ClearingSystem), `^[0-9A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionAttributes) validatePostingStatus(formats strfmt.Registry) error {

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

var paymentAdmissionAttributesTypeRoutePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on_us","xp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentAdmissionAttributesTypeRoutePropEnum = append(paymentAdmissionAttributesTypeRoutePropEnum, v)
	}
}

const (

	// PaymentAdmissionAttributesRouteOnUs captures enum value "on_us"
	PaymentAdmissionAttributesRouteOnUs string = "on_us"

	// PaymentAdmissionAttributesRouteXp captures enum value "xp"
	PaymentAdmissionAttributesRouteXp string = "xp"
)

// prop value enum
func (m *PaymentAdmissionAttributes) validateRouteEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentAdmissionAttributesTypeRoutePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentAdmissionAttributes) validateRoute(formats strfmt.Registry) error {

	if swag.IsZero(m.Route) { // not required
		return nil
	}

	// value enum
	if err := m.validateRouteEnum("attributes"+"."+"route", "body", m.Route); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionAttributes) validateSettlementCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementCycle) { // not required
		return nil
	}

	if err := validate.MinimumInt("attributes"+"."+"settlement_cycle", "body", int64(*m.SettlementCycle), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionAttributes) validateSettlementDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"settlement_date", "body", "date", m.SettlementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionAttributes) validateStatusReason(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PaymentAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentAdmissionRelationships payment admission relationships
// swagger:model PaymentAdmissionRelationships
type PaymentAdmissionRelationships struct {

	// beneficiary account
	BeneficiaryAccount *RelationshipPaymentAdmissionBeneficiaryAccount `json:"beneficiary_account,omitempty"`

	// beneficiary branch
	BeneficiaryBranch *RelationshipPaymentAdmissionBeneficiaryBranch `json:"beneficiary_branch,omitempty"`

	// payment
	Payment *RelationshipPayments `json:"payment,omitempty"`

	// payment admission task
	PaymentAdmissionTask *RelationshipPaymentAdmissionTasks `json:"payment_admission_task,omitempty"`
}

func PaymentAdmissionRelationshipsWithDefaults(defaults client.Defaults) *PaymentAdmissionRelationships {
	return &PaymentAdmissionRelationships{

		BeneficiaryAccount: RelationshipPaymentAdmissionBeneficiaryAccountWithDefaults(defaults),

		BeneficiaryBranch: RelationshipPaymentAdmissionBeneficiaryBranchWithDefaults(defaults),

		Payment: RelationshipPaymentsWithDefaults(defaults),

		PaymentAdmissionTask: RelationshipPaymentAdmissionTasksWithDefaults(defaults),
	}
}

func (m *PaymentAdmissionRelationships) WithBeneficiaryAccount(beneficiaryAccount RelationshipPaymentAdmissionBeneficiaryAccount) *PaymentAdmissionRelationships {

	m.BeneficiaryAccount = &beneficiaryAccount

	return m
}

func (m *PaymentAdmissionRelationships) WithoutBeneficiaryAccount() *PaymentAdmissionRelationships {
	m.BeneficiaryAccount = nil
	return m
}

func (m *PaymentAdmissionRelationships) WithBeneficiaryBranch(beneficiaryBranch RelationshipPaymentAdmissionBeneficiaryBranch) *PaymentAdmissionRelationships {

	m.BeneficiaryBranch = &beneficiaryBranch

	return m
}

func (m *PaymentAdmissionRelationships) WithoutBeneficiaryBranch() *PaymentAdmissionRelationships {
	m.BeneficiaryBranch = nil
	return m
}

func (m *PaymentAdmissionRelationships) WithPayment(payment RelationshipPayments) *PaymentAdmissionRelationships {

	m.Payment = &payment

	return m
}

func (m *PaymentAdmissionRelationships) WithoutPayment() *PaymentAdmissionRelationships {
	m.Payment = nil
	return m
}

func (m *PaymentAdmissionRelationships) WithPaymentAdmissionTask(paymentAdmissionTask RelationshipPaymentAdmissionTasks) *PaymentAdmissionRelationships {

	m.PaymentAdmissionTask = &paymentAdmissionTask

	return m
}

func (m *PaymentAdmissionRelationships) WithoutPaymentAdmissionTask() *PaymentAdmissionRelationships {
	m.PaymentAdmissionTask = nil
	return m
}

// Validate validates this payment admission relationships
func (m *PaymentAdmissionRelationships) Validate(formats strfmt.Registry) error {
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

func (m *PaymentAdmissionRelationships) validateBeneficiaryAccount(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionRelationships) validateBeneficiaryBranch(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionRelationships) validatePaymentAdmissionTask(formats strfmt.Registry) error {

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
func (m *PaymentAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
