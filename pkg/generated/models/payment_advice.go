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

// PaymentAdvice payment advice
// swagger:model PaymentAdvice
type PaymentAdvice struct {

	// attributes
	Attributes *PaymentAdviceAttributes `json:"attributes,omitempty"`

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
	Relationships *PaymentAdviceRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func PaymentAdviceWithDefaults(defaults client.Defaults) *PaymentAdvice {
	return &PaymentAdvice{

		Attributes: PaymentAdviceAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("PaymentAdvice", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("PaymentAdvice", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("PaymentAdvice", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("PaymentAdvice", "organisation_id"),

		Relationships: PaymentAdviceRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("PaymentAdvice", "type"),

		Version: defaults.GetInt64Ptr("PaymentAdvice", "version"),
	}
}

func (m *PaymentAdvice) WithAttributes(attributes PaymentAdviceAttributes) *PaymentAdvice {

	m.Attributes = &attributes

	return m
}

func (m *PaymentAdvice) WithoutAttributes() *PaymentAdvice {
	m.Attributes = nil
	return m
}

func (m *PaymentAdvice) WithCreatedOn(createdOn strfmt.DateTime) *PaymentAdvice {

	m.CreatedOn = &createdOn

	return m
}

func (m *PaymentAdvice) WithoutCreatedOn() *PaymentAdvice {
	m.CreatedOn = nil
	return m
}

func (m *PaymentAdvice) WithID(id strfmt.UUID) *PaymentAdvice {

	m.ID = &id

	return m
}

func (m *PaymentAdvice) WithoutID() *PaymentAdvice {
	m.ID = nil
	return m
}

func (m *PaymentAdvice) WithModifiedOn(modifiedOn strfmt.DateTime) *PaymentAdvice {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *PaymentAdvice) WithoutModifiedOn() *PaymentAdvice {
	m.ModifiedOn = nil
	return m
}

func (m *PaymentAdvice) WithOrganisationID(organisationID strfmt.UUID) *PaymentAdvice {

	m.OrganisationID = &organisationID

	return m
}

func (m *PaymentAdvice) WithoutOrganisationID() *PaymentAdvice {
	m.OrganisationID = nil
	return m
}

func (m *PaymentAdvice) WithRelationships(relationships PaymentAdviceRelationships) *PaymentAdvice {

	m.Relationships = &relationships

	return m
}

func (m *PaymentAdvice) WithoutRelationships() *PaymentAdvice {
	m.Relationships = nil
	return m
}

func (m *PaymentAdvice) WithType(typeVar string) *PaymentAdvice {

	m.Type = typeVar

	return m
}

func (m *PaymentAdvice) WithVersion(version int64) *PaymentAdvice {

	m.Version = &version

	return m
}

func (m *PaymentAdvice) WithoutVersion() *PaymentAdvice {
	m.Version = nil
	return m
}

// Validate validates this payment advice
func (m *PaymentAdvice) Validate(formats strfmt.Registry) error {
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

func (m *PaymentAdvice) validateAttributes(formats strfmt.Registry) error {

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

func (m *PaymentAdvice) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdvice) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdvice) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdvice) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdvice) validateRelationships(formats strfmt.Registry) error {

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

func (m *PaymentAdvice) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdvice) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdvice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdvice) UnmarshalBinary(b []byte) error {
	var res PaymentAdvice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdvice) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentAdviceAttributes payment advice attributes
// swagger:model PaymentAdviceAttributes
type PaymentAdviceAttributes struct {

	// beneficiary party
	BeneficiaryParty *PaymentAdviceAttributesBeneficiaryParty `json:"beneficiary_party,omitempty"`

	// reason code
	ReasonCode string `json:"reason_code,omitempty"`
}

func PaymentAdviceAttributesWithDefaults(defaults client.Defaults) *PaymentAdviceAttributes {
	return &PaymentAdviceAttributes{

		BeneficiaryParty: PaymentAdviceAttributesBeneficiaryPartyWithDefaults(defaults),

		ReasonCode: defaults.GetString("PaymentAdviceAttributes", "reason_code"),
	}
}

func (m *PaymentAdviceAttributes) WithBeneficiaryParty(beneficiaryParty PaymentAdviceAttributesBeneficiaryParty) *PaymentAdviceAttributes {

	m.BeneficiaryParty = &beneficiaryParty

	return m
}

func (m *PaymentAdviceAttributes) WithoutBeneficiaryParty() *PaymentAdviceAttributes {
	m.BeneficiaryParty = nil
	return m
}

func (m *PaymentAdviceAttributes) WithReasonCode(reasonCode string) *PaymentAdviceAttributes {

	m.ReasonCode = reasonCode

	return m
}

// Validate validates this payment advice attributes
func (m *PaymentAdviceAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBeneficiaryParty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdviceAttributes) validateBeneficiaryParty(formats strfmt.Registry) error {

	if swag.IsZero(m.BeneficiaryParty) { // not required
		return nil
	}

	if m.BeneficiaryParty != nil {
		if err := m.BeneficiaryParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "beneficiary_party")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdviceAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdviceAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentAdviceAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdviceAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentAdviceAttributesBeneficiaryParty payment advice attributes beneficiary party
// swagger:model PaymentAdviceAttributesBeneficiaryParty
type PaymentAdviceAttributesBeneficiaryParty struct {

	// new bank details
	NewBankDetails *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails `json:"new_bank_details,omitempty"`
}

func PaymentAdviceAttributesBeneficiaryPartyWithDefaults(defaults client.Defaults) *PaymentAdviceAttributesBeneficiaryParty {
	return &PaymentAdviceAttributesBeneficiaryParty{

		NewBankDetails: PaymentAdviceAttributesBeneficiaryPartyNewBankDetailsWithDefaults(defaults),
	}
}

func (m *PaymentAdviceAttributesBeneficiaryParty) WithNewBankDetails(newBankDetails PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) *PaymentAdviceAttributesBeneficiaryParty {

	m.NewBankDetails = &newBankDetails

	return m
}

func (m *PaymentAdviceAttributesBeneficiaryParty) WithoutNewBankDetails() *PaymentAdviceAttributesBeneficiaryParty {
	m.NewBankDetails = nil
	return m
}

// Validate validates this payment advice attributes beneficiary party
func (m *PaymentAdviceAttributesBeneficiaryParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewBankDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdviceAttributesBeneficiaryParty) validateNewBankDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.NewBankDetails) { // not required
		return nil
	}

	if m.NewBankDetails != nil {
		if err := m.NewBankDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "beneficiary_party" + "." + "new_bank_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdviceAttributesBeneficiaryParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdviceAttributesBeneficiaryParty) UnmarshalBinary(b []byte) error {
	var res PaymentAdviceAttributesBeneficiaryParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdviceAttributesBeneficiaryParty) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentAdviceAttributesBeneficiaryPartyNewBankDetails payment advice attributes beneficiary party new bank details
// swagger:model PaymentAdviceAttributesBeneficiaryPartyNewBankDetails
type PaymentAdviceAttributesBeneficiaryPartyNewBankDetails struct {

	// account number
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// account with
	AccountWith *AccountHoldingEntity `json:"account_with,omitempty"`

	// roll number
	RollNumber string `json:"roll_number,omitempty"`
}

func PaymentAdviceAttributesBeneficiaryPartyNewBankDetailsWithDefaults(defaults client.Defaults) *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails {
	return &PaymentAdviceAttributesBeneficiaryPartyNewBankDetails{

		AccountNumber: defaults.GetString("PaymentAdviceAttributesBeneficiaryPartyNewBankDetails", "account_number"),

		// TODO AccountNumberCode: AccountNumberCode,

		AccountWith: AccountHoldingEntityWithDefaults(defaults),

		RollNumber: defaults.GetString("PaymentAdviceAttributesBeneficiaryPartyNewBankDetails", "roll_number"),
	}
}

func (m *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) WithAccountNumber(accountNumber string) *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails {

	m.AccountNumber = accountNumber

	return m
}

func (m *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) WithAccountNumberCode(accountNumberCode AccountNumberCode) *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails {

	m.AccountNumberCode = accountNumberCode

	return m
}

func (m *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) WithAccountWith(accountWith AccountHoldingEntity) *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails {

	m.AccountWith = &accountWith

	return m
}

func (m *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) WithoutAccountWith() *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails {
	m.AccountWith = nil
	return m
}

func (m *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) WithRollNumber(rollNumber string) *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails {

	m.RollNumber = rollNumber

	return m
}

// Validate validates this payment advice attributes beneficiary party new bank details
func (m *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	if err := m.AccountNumberCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "beneficiary_party" + "." + "new_bank_details" + "." + "account_number_code")
		}
		return err
	}

	return nil
}

func (m *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) validateAccountWith(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountWith) { // not required
		return nil
	}

	if m.AccountWith != nil {
		if err := m.AccountWith.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "beneficiary_party" + "." + "new_bank_details" + "." + "account_with")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) UnmarshalBinary(b []byte) error {
	var res PaymentAdviceAttributesBeneficiaryPartyNewBankDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdviceAttributesBeneficiaryPartyNewBankDetails) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentAdviceRelationships payment advice relationships
// swagger:model PaymentAdviceRelationships
type PaymentAdviceRelationships struct {

	// advice submission
	AdviceSubmission *RelationshipLinks `json:"advice_submission,omitempty"`

	// payment
	Payment *RelationshipLinks `json:"payment,omitempty"`
}

func PaymentAdviceRelationshipsWithDefaults(defaults client.Defaults) *PaymentAdviceRelationships {
	return &PaymentAdviceRelationships{

		AdviceSubmission: RelationshipLinksWithDefaults(defaults),

		Payment: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *PaymentAdviceRelationships) WithAdviceSubmission(adviceSubmission RelationshipLinks) *PaymentAdviceRelationships {

	m.AdviceSubmission = &adviceSubmission

	return m
}

func (m *PaymentAdviceRelationships) WithoutAdviceSubmission() *PaymentAdviceRelationships {
	m.AdviceSubmission = nil
	return m
}

func (m *PaymentAdviceRelationships) WithPayment(payment RelationshipLinks) *PaymentAdviceRelationships {

	m.Payment = &payment

	return m
}

func (m *PaymentAdviceRelationships) WithoutPayment() *PaymentAdviceRelationships {
	m.Payment = nil
	return m
}

// Validate validates this payment advice relationships
func (m *PaymentAdviceRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdviceSubmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdviceRelationships) validateAdviceSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.AdviceSubmission) { // not required
		return nil
	}

	if m.AdviceSubmission != nil {
		if err := m.AdviceSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "advice_submission")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAdviceRelationships) validatePayment(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PaymentAdviceRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdviceRelationships) UnmarshalBinary(b []byte) error {
	var res PaymentAdviceRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdviceRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
