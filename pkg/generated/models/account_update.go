// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountUpdate account update
// swagger:model AccountUpdate
type AccountUpdate struct {

	// attributes
	// Required: true
	Attributes *AccountUpdateAttributes `json:"attributes"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *AccountRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Required: true
	// Minimum: 0
	Version *int64 `json:"version"`
}

func AccountUpdateWithDefaults(defaults client.Defaults) *AccountUpdate {
	return &AccountUpdate{

		Attributes: AccountUpdateAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("AccountUpdate", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("AccountUpdate", "organisation_id"),

		Relationships: AccountRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("AccountUpdate", "type"),

		Version: defaults.GetInt64Ptr("AccountUpdate", "version"),
	}
}

func (m *AccountUpdate) WithAttributes(attributes AccountUpdateAttributes) *AccountUpdate {

	m.Attributes = &attributes

	return m
}

func (m *AccountUpdate) WithoutAttributes() *AccountUpdate {
	m.Attributes = nil
	return m
}

func (m *AccountUpdate) WithID(id strfmt.UUID) *AccountUpdate {

	m.ID = &id

	return m
}

func (m *AccountUpdate) WithoutID() *AccountUpdate {
	m.ID = nil
	return m
}

func (m *AccountUpdate) WithOrganisationID(organisationID strfmt.UUID) *AccountUpdate {

	m.OrganisationID = &organisationID

	return m
}

func (m *AccountUpdate) WithoutOrganisationID() *AccountUpdate {
	m.OrganisationID = nil
	return m
}

func (m *AccountUpdate) WithRelationships(relationships AccountRelationships) *AccountUpdate {

	m.Relationships = &relationships

	return m
}

func (m *AccountUpdate) WithoutRelationships() *AccountUpdate {
	m.Relationships = nil
	return m
}

func (m *AccountUpdate) WithType(typeVar string) *AccountUpdate {

	m.Type = typeVar

	return m
}

func (m *AccountUpdate) WithVersion(version int64) *AccountUpdate {

	m.Version = &version

	return m
}

func (m *AccountUpdate) WithoutVersion() *AccountUpdate {
	m.Version = nil
	return m
}

// Validate validates this account update
func (m *AccountUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

func (m *AccountUpdate) validateAttributes(formats strfmt.Registry) error {

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

func (m *AccountUpdate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdate) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdate) validateRelationships(formats strfmt.Registry) error {

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

func (m *AccountUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdate) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountUpdate) UnmarshalBinary(b []byte) error {
	var res AccountUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountUpdate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AccountUpdateAttributes account update attributes
// swagger:model AccountUpdateAttributes
type AccountUpdateAttributes struct {

	// acceptance qualifier
	AcceptanceQualifier AcceptanceQualifier `json:"acceptance_qualifier,omitempty"`

	// Is the account business or personal?
	// Enum: [Personal Business]
	AccountClassification string `json:"account_classification,omitempty"`

	// - deprecated - Is the account opted out of account matching, e.g. CoP?
	AccountMatchingOptOut bool `json:"account_matching_opt_out,omitempty"`

	// Account number of the account. A unique number will automatically be generated if not provided.
	// Pattern: ^[A-Z0-9]{0,64}$
	AccountNumber string `json:"account_number,omitempty"`

	// - deprecated - Alternative account names. Used for Confirmation of Payee matching.
	// Max Items: 3
	AlternativeBankAccountNames []string `json:"alternative_bank_account_names"`

	// Alternative names. Used for Confirmation of Payee matching.
	// Max Items: 3
	AlternativeNames []string `json:"alternative_names"`

	// - deprecated - Primary account name. Used for Confirmation of Payee matching. Required if confirmation_of_payee_enabled is true for the organisation.
	// Max Length: 140
	// Min Length: 1
	BankAccountName string `json:"bank_account_name,omitempty"`

	// Local country bank identifier. In the UK this is the sort code.
	// Pattern: ^[A-Z0-9]{0,16}$
	BankID string `json:"bank_id,omitempty"`

	// ISO 20022 code used to identify the type of bank ID being used
	// Pattern: ^[A-Z]{0,16}$
	BankIDCode string `json:"bank_id_code,omitempty"`

	// ISO 4217 code used to identify the base currency of the account
	// Pattern: ^[A-Z]{3}$
	BaseCurrency string `json:"base_currency,omitempty"`

	// SWIFT BIC in either 8 or 11 character format
	// Pattern: ^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$
	Bic string `json:"bic,omitempty"`

	// ISO 3166-1 code used to identify the domicile of the account
	// Pattern: ^[A-Z]{2}$
	Country string `json:"country,omitempty"`

	// A free-format reference that can be used to link this account to an external system
	// Pattern: ^[a-zA-Z0-9-$@., ]{0,256}$
	CustomerID string `json:"customer_id,omitempty"`

	// - deprecated - Customer first name.
	// Max Length: 40
	// Min Length: 1
	FirstName string `json:"first_name,omitempty"`

	// IBAN of the account. Will be calculated from other fields if not supplied.
	// Pattern: ^[A-Z]{2}[0-9]{2}[A-Z0-9]{0,64}$
	Iban string `json:"iban,omitempty"`

	// Is the account joint?
	JointAccount bool `json:"joint_account,omitempty"`

	// Account holder names (for example title, first name, last name). Used for Confirmation of Payee matching.
	// Max Items: 4
	Name []string `json:"name"`

	// Describes the status of the account for name matching via CoP. The value determines the code with which Form3 responds to matched CoP requests to this account.
	// Enum: [supported switched opted_out not_supported]
	NameMatchingStatus string `json:"name_matching_status,omitempty"`

	// organisation identification
	OrganisationIdentification *AccountAttributesOrganisationIdentification `json:"organisation_identification,omitempty"`

	// private identification
	PrivateIdentification *AccountAttributesPrivateIdentification `json:"private_identification,omitempty"`

	// - deprecated - Accounting system or service. It will be added to each payment received to an account.
	// Max Length: 35
	ProcessingService string `json:"processing_service,omitempty"`

	// When set will apply a validation mask on the payment reference to each payment received to an account.
	// Max Length: 35
	ReferenceMask string `json:"reference_mask,omitempty"`

	// Secondary identification, e.g. building society roll number. Used for Confirmation of Payee.
	// Max Length: 140
	// Min Length: 1
	SecondaryIdentification string `json:"secondary_identification,omitempty"`

	// Current status of the account
	// Enum: [pending failed confirmed closed]
	Status string `json:"status,omitempty"`

	// status reason
	StatusReason StatusReason `json:"status_reason,omitempty"`

	// - deprecated - Indicates whether the account has been switched using the Current Account Switch Service.
	Switched bool `json:"switched,omitempty"`

	// switched account details
	SwitchedAccountDetails *SwitchedAccountDetails `json:"switched_account_details,omitempty"`

	// - deprecated - Customer title.
	// Max Length: 40
	// Min Length: 1
	Title string `json:"title,omitempty"`

	// Account type
	// Max Length: 35
	Type string `json:"type,omitempty"`

	// All purpose list of key-value pairs to store specific data for the associated account. It will be added to each payment received to an account.
	// Max Items: 5
	UserDefinedData []*UserDefinedData `json:"user_defined_data,omitempty"`

	// - deprecated - All purpose field to store specific data for the associated account. It will be added to each payment received to an account.
	// Max Length: 35
	UserDefinedInformation string `json:"user_defined_information,omitempty"`

	// validation type
	ValidationType ValidationType `json:"validation_type,omitempty"`
}

func AccountUpdateAttributesWithDefaults(defaults client.Defaults) *AccountUpdateAttributes {
	return &AccountUpdateAttributes{

		// TODO AcceptanceQualifier: AcceptanceQualifier,

		AccountClassification: defaults.GetString("AccountUpdateAttributes", "account_classification"),

		AccountMatchingOptOut: defaults.GetBool("AccountUpdateAttributes", "account_matching_opt_out"),

		AccountNumber: defaults.GetString("AccountUpdateAttributes", "account_number"),

		AlternativeBankAccountNames: make([]string, 0),

		AlternativeNames: make([]string, 0),

		BankAccountName: defaults.GetString("AccountUpdateAttributes", "bank_account_name"),

		BankID: defaults.GetString("AccountUpdateAttributes", "bank_id"),

		BankIDCode: defaults.GetString("AccountUpdateAttributes", "bank_id_code"),

		BaseCurrency: defaults.GetString("AccountUpdateAttributes", "base_currency"),

		Bic: defaults.GetString("AccountUpdateAttributes", "bic"),

		Country: defaults.GetString("AccountUpdateAttributes", "country"),

		CustomerID: defaults.GetString("AccountUpdateAttributes", "customer_id"),

		FirstName: defaults.GetString("AccountUpdateAttributes", "first_name"),

		Iban: defaults.GetString("AccountUpdateAttributes", "iban"),

		JointAccount: defaults.GetBool("AccountUpdateAttributes", "joint_account"),

		Name: make([]string, 0),

		NameMatchingStatus: defaults.GetString("AccountUpdateAttributes", "name_matching_status"),

		OrganisationIdentification: AccountAttributesOrganisationIdentificationWithDefaults(defaults),

		PrivateIdentification: AccountAttributesPrivateIdentificationWithDefaults(defaults),

		ProcessingService: defaults.GetString("AccountUpdateAttributes", "processing_service"),

		ReferenceMask: defaults.GetString("AccountUpdateAttributes", "reference_mask"),

		SecondaryIdentification: defaults.GetString("AccountUpdateAttributes", "secondary_identification"),

		Status: defaults.GetString("AccountUpdateAttributes", "status"),

		// TODO StatusReason: StatusReason,

		Switched: defaults.GetBool("AccountUpdateAttributes", "switched"),

		SwitchedAccountDetails: SwitchedAccountDetailsWithDefaults(defaults),

		Title: defaults.GetString("AccountUpdateAttributes", "title"),

		Type: defaults.GetString("AccountUpdateAttributes", "type"),

		UserDefinedData: make([]*UserDefinedData, 0),

		UserDefinedInformation: defaults.GetString("AccountUpdateAttributes", "user_defined_information"),

		// TODO ValidationType: ValidationType,

	}
}

func (m *AccountUpdateAttributes) WithAcceptanceQualifier(acceptanceQualifier AcceptanceQualifier) *AccountUpdateAttributes {

	m.AcceptanceQualifier = acceptanceQualifier

	return m
}

func (m *AccountUpdateAttributes) WithAccountClassification(accountClassification string) *AccountUpdateAttributes {

	m.AccountClassification = accountClassification

	return m
}

func (m *AccountUpdateAttributes) WithAccountMatchingOptOut(accountMatchingOptOut bool) *AccountUpdateAttributes {

	m.AccountMatchingOptOut = accountMatchingOptOut

	return m
}

func (m *AccountUpdateAttributes) WithAccountNumber(accountNumber string) *AccountUpdateAttributes {

	m.AccountNumber = accountNumber

	return m
}

func (m *AccountUpdateAttributes) WithAlternativeBankAccountNames(alternativeBankAccountNames []string) *AccountUpdateAttributes {

	m.AlternativeBankAccountNames = alternativeBankAccountNames

	return m
}

func (m *AccountUpdateAttributes) WithAlternativeNames(alternativeNames []string) *AccountUpdateAttributes {

	m.AlternativeNames = alternativeNames

	return m
}

func (m *AccountUpdateAttributes) WithBankAccountName(bankAccountName string) *AccountUpdateAttributes {

	m.BankAccountName = bankAccountName

	return m
}

func (m *AccountUpdateAttributes) WithBankID(bankID string) *AccountUpdateAttributes {

	m.BankID = bankID

	return m
}

func (m *AccountUpdateAttributes) WithBankIDCode(bankIDCode string) *AccountUpdateAttributes {

	m.BankIDCode = bankIDCode

	return m
}

func (m *AccountUpdateAttributes) WithBaseCurrency(baseCurrency string) *AccountUpdateAttributes {

	m.BaseCurrency = baseCurrency

	return m
}

func (m *AccountUpdateAttributes) WithBic(bic string) *AccountUpdateAttributes {

	m.Bic = bic

	return m
}

func (m *AccountUpdateAttributes) WithCountry(country string) *AccountUpdateAttributes {

	m.Country = country

	return m
}

func (m *AccountUpdateAttributes) WithCustomerID(customerID string) *AccountUpdateAttributes {

	m.CustomerID = customerID

	return m
}

func (m *AccountUpdateAttributes) WithFirstName(firstName string) *AccountUpdateAttributes {

	m.FirstName = firstName

	return m
}

func (m *AccountUpdateAttributes) WithIban(iban string) *AccountUpdateAttributes {

	m.Iban = iban

	return m
}

func (m *AccountUpdateAttributes) WithJointAccount(jointAccount bool) *AccountUpdateAttributes {

	m.JointAccount = jointAccount

	return m
}

func (m *AccountUpdateAttributes) WithName(name []string) *AccountUpdateAttributes {

	m.Name = name

	return m
}

func (m *AccountUpdateAttributes) WithNameMatchingStatus(nameMatchingStatus string) *AccountUpdateAttributes {

	m.NameMatchingStatus = nameMatchingStatus

	return m
}

func (m *AccountUpdateAttributes) WithOrganisationIdentification(organisationIdentification AccountAttributesOrganisationIdentification) *AccountUpdateAttributes {

	m.OrganisationIdentification = &organisationIdentification

	return m
}

func (m *AccountUpdateAttributes) WithoutOrganisationIdentification() *AccountUpdateAttributes {
	m.OrganisationIdentification = nil
	return m
}

func (m *AccountUpdateAttributes) WithPrivateIdentification(privateIdentification AccountAttributesPrivateIdentification) *AccountUpdateAttributes {

	m.PrivateIdentification = &privateIdentification

	return m
}

func (m *AccountUpdateAttributes) WithoutPrivateIdentification() *AccountUpdateAttributes {
	m.PrivateIdentification = nil
	return m
}

func (m *AccountUpdateAttributes) WithProcessingService(processingService string) *AccountUpdateAttributes {

	m.ProcessingService = processingService

	return m
}

func (m *AccountUpdateAttributes) WithReferenceMask(referenceMask string) *AccountUpdateAttributes {

	m.ReferenceMask = referenceMask

	return m
}

func (m *AccountUpdateAttributes) WithSecondaryIdentification(secondaryIdentification string) *AccountUpdateAttributes {

	m.SecondaryIdentification = secondaryIdentification

	return m
}

func (m *AccountUpdateAttributes) WithStatus(status string) *AccountUpdateAttributes {

	m.Status = status

	return m
}

func (m *AccountUpdateAttributes) WithStatusReason(statusReason StatusReason) *AccountUpdateAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *AccountUpdateAttributes) WithSwitched(switched bool) *AccountUpdateAttributes {

	m.Switched = switched

	return m
}

func (m *AccountUpdateAttributes) WithSwitchedAccountDetails(switchedAccountDetails SwitchedAccountDetails) *AccountUpdateAttributes {

	m.SwitchedAccountDetails = &switchedAccountDetails

	return m
}

func (m *AccountUpdateAttributes) WithoutSwitchedAccountDetails() *AccountUpdateAttributes {
	m.SwitchedAccountDetails = nil
	return m
}

func (m *AccountUpdateAttributes) WithTitle(title string) *AccountUpdateAttributes {

	m.Title = title

	return m
}

func (m *AccountUpdateAttributes) WithType(typeVar string) *AccountUpdateAttributes {

	m.Type = typeVar

	return m
}

func (m *AccountUpdateAttributes) WithUserDefinedData(userDefinedData []*UserDefinedData) *AccountUpdateAttributes {

	m.UserDefinedData = userDefinedData

	return m
}

func (m *AccountUpdateAttributes) WithUserDefinedInformation(userDefinedInformation string) *AccountUpdateAttributes {

	m.UserDefinedInformation = userDefinedInformation

	return m
}

func (m *AccountUpdateAttributes) WithValidationType(validationType ValidationType) *AccountUpdateAttributes {

	m.ValidationType = validationType

	return m
}

// Validate validates this account update attributes
func (m *AccountUpdateAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptanceQualifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountClassification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlternativeBankAccountNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlternativeNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankAccountName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIban(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameMatchingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceMask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwitchedAccountDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDefinedData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDefinedInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountUpdateAttributes) validateAcceptanceQualifier(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptanceQualifier) { // not required
		return nil
	}

	if err := m.AcceptanceQualifier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "acceptance_qualifier")
		}
		return err
	}

	return nil
}

var accountUpdateAttributesTypeAccountClassificationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Personal","Business"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountUpdateAttributesTypeAccountClassificationPropEnum = append(accountUpdateAttributesTypeAccountClassificationPropEnum, v)
	}
}

const (

	// AccountUpdateAttributesAccountClassificationPersonal captures enum value "Personal"
	AccountUpdateAttributesAccountClassificationPersonal string = "Personal"

	// AccountUpdateAttributesAccountClassificationBusiness captures enum value "Business"
	AccountUpdateAttributesAccountClassificationBusiness string = "Business"
)

// prop value enum
func (m *AccountUpdateAttributes) validateAccountClassificationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountUpdateAttributesTypeAccountClassificationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountUpdateAttributes) validateAccountClassification(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountClassification) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccountClassificationEnum("attributes"+"."+"account_classification", "body", m.AccountClassification); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumber) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"account_number", "body", string(m.AccountNumber), `^[A-Z0-9]{0,64}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateAlternativeBankAccountNames(formats strfmt.Registry) error {

	if swag.IsZero(m.AlternativeBankAccountNames) { // not required
		return nil
	}

	iAlternativeBankAccountNamesSize := int64(len(m.AlternativeBankAccountNames))

	if err := validate.MaxItems("attributes"+"."+"alternative_bank_account_names", "body", iAlternativeBankAccountNamesSize, 3); err != nil {
		return err
	}

	for i := 0; i < len(m.AlternativeBankAccountNames); i++ {

		if err := validate.MinLength("attributes"+"."+"alternative_bank_account_names"+"."+strconv.Itoa(i), "body", string(m.AlternativeBankAccountNames[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("attributes"+"."+"alternative_bank_account_names"+"."+strconv.Itoa(i), "body", string(m.AlternativeBankAccountNames[i]), 140); err != nil {
			return err
		}

	}

	return nil
}

func (m *AccountUpdateAttributes) validateAlternativeNames(formats strfmt.Registry) error {

	if swag.IsZero(m.AlternativeNames) { // not required
		return nil
	}

	iAlternativeNamesSize := int64(len(m.AlternativeNames))

	if err := validate.MaxItems("attributes"+"."+"alternative_names", "body", iAlternativeNamesSize, 3); err != nil {
		return err
	}

	for i := 0; i < len(m.AlternativeNames); i++ {

		if err := validate.MinLength("attributes"+"."+"alternative_names"+"."+strconv.Itoa(i), "body", string(m.AlternativeNames[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("attributes"+"."+"alternative_names"+"."+strconv.Itoa(i), "body", string(m.AlternativeNames[i]), 140); err != nil {
			return err
		}

	}

	return nil
}

func (m *AccountUpdateAttributes) validateBankAccountName(formats strfmt.Registry) error {

	if swag.IsZero(m.BankAccountName) { // not required
		return nil
	}

	if err := validate.MinLength("attributes"+"."+"bank_account_name", "body", string(m.BankAccountName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("attributes"+"."+"bank_account_name", "body", string(m.BankAccountName), 140); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateBankID(formats strfmt.Registry) error {

	if swag.IsZero(m.BankID) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"bank_id", "body", string(m.BankID), `^[A-Z0-9]{0,16}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateBankIDCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankIDCode) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"bank_id_code", "body", string(m.BankIDCode), `^[A-Z]{0,16}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateBaseCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.BaseCurrency) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"base_currency", "body", string(m.BaseCurrency), `^[A-Z]{3}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateBic(formats strfmt.Registry) error {

	if swag.IsZero(m.Bic) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"bic", "body", string(m.Bic), `^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"country", "body", string(m.Country), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateCustomerID(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"customer_id", "body", string(m.CustomerID), `^[a-zA-Z0-9-$@., ]{0,256}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateFirstName(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstName) { // not required
		return nil
	}

	if err := validate.MinLength("attributes"+"."+"first_name", "body", string(m.FirstName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("attributes"+"."+"first_name", "body", string(m.FirstName), 40); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateIban(formats strfmt.Registry) error {

	if swag.IsZero(m.Iban) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"iban", "body", string(m.Iban), `^[A-Z]{2}[0-9]{2}[A-Z0-9]{0,64}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	iNameSize := int64(len(m.Name))

	if err := validate.MaxItems("attributes"+"."+"name", "body", iNameSize, 4); err != nil {
		return err
	}

	for i := 0; i < len(m.Name); i++ {

		if err := validate.MinLength("attributes"+"."+"name"+"."+strconv.Itoa(i), "body", string(m.Name[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("attributes"+"."+"name"+"."+strconv.Itoa(i), "body", string(m.Name[i]), 140); err != nil {
			return err
		}

	}

	return nil
}

var accountUpdateAttributesTypeNameMatchingStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["supported","switched","opted_out","not_supported"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountUpdateAttributesTypeNameMatchingStatusPropEnum = append(accountUpdateAttributesTypeNameMatchingStatusPropEnum, v)
	}
}

const (

	// AccountUpdateAttributesNameMatchingStatusSupported captures enum value "supported"
	AccountUpdateAttributesNameMatchingStatusSupported string = "supported"

	// AccountUpdateAttributesNameMatchingStatusSwitched captures enum value "switched"
	AccountUpdateAttributesNameMatchingStatusSwitched string = "switched"

	// AccountUpdateAttributesNameMatchingStatusOptedOut captures enum value "opted_out"
	AccountUpdateAttributesNameMatchingStatusOptedOut string = "opted_out"

	// AccountUpdateAttributesNameMatchingStatusNotSupported captures enum value "not_supported"
	AccountUpdateAttributesNameMatchingStatusNotSupported string = "not_supported"
)

// prop value enum
func (m *AccountUpdateAttributes) validateNameMatchingStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountUpdateAttributesTypeNameMatchingStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountUpdateAttributes) validateNameMatchingStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.NameMatchingStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameMatchingStatusEnum("attributes"+"."+"name_matching_status", "body", m.NameMatchingStatus); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateOrganisationIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationIdentification) { // not required
		return nil
	}

	if m.OrganisationIdentification != nil {
		if err := m.OrganisationIdentification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "organisation_identification")
			}
			return err
		}
	}

	return nil
}

func (m *AccountUpdateAttributes) validatePrivateIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivateIdentification) { // not required
		return nil
	}

	if m.PrivateIdentification != nil {
		if err := m.PrivateIdentification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "private_identification")
			}
			return err
		}
	}

	return nil
}

func (m *AccountUpdateAttributes) validateProcessingService(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessingService) { // not required
		return nil
	}

	if err := validate.MaxLength("attributes"+"."+"processing_service", "body", string(m.ProcessingService), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateReferenceMask(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceMask) { // not required
		return nil
	}

	if err := validate.MaxLength("attributes"+"."+"reference_mask", "body", string(m.ReferenceMask), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateSecondaryIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryIdentification) { // not required
		return nil
	}

	if err := validate.MinLength("attributes"+"."+"secondary_identification", "body", string(m.SecondaryIdentification), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("attributes"+"."+"secondary_identification", "body", string(m.SecondaryIdentification), 140); err != nil {
		return err
	}

	return nil
}

var accountUpdateAttributesTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","failed","confirmed","closed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountUpdateAttributesTypeStatusPropEnum = append(accountUpdateAttributesTypeStatusPropEnum, v)
	}
}

const (

	// AccountUpdateAttributesStatusPending captures enum value "pending"
	AccountUpdateAttributesStatusPending string = "pending"

	// AccountUpdateAttributesStatusFailed captures enum value "failed"
	AccountUpdateAttributesStatusFailed string = "failed"

	// AccountUpdateAttributesStatusConfirmed captures enum value "confirmed"
	AccountUpdateAttributesStatusConfirmed string = "confirmed"

	// AccountUpdateAttributesStatusClosed captures enum value "closed"
	AccountUpdateAttributesStatusClosed string = "closed"
)

// prop value enum
func (m *AccountUpdateAttributes) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountUpdateAttributesTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountUpdateAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("attributes"+"."+"status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateStatusReason(formats strfmt.Registry) error {

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

func (m *AccountUpdateAttributes) validateSwitchedAccountDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.SwitchedAccountDetails) { // not required
		return nil
	}

	if m.SwitchedAccountDetails != nil {
		if err := m.SwitchedAccountDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "switched_account_details")
			}
			return err
		}
	}

	return nil
}

func (m *AccountUpdateAttributes) validateTitle(formats strfmt.Registry) error {

	if swag.IsZero(m.Title) { // not required
		return nil
	}

	if err := validate.MinLength("attributes"+"."+"title", "body", string(m.Title), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("attributes"+"."+"title", "body", string(m.Title), 40); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.MaxLength("attributes"+"."+"type", "body", string(m.Type), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateUserDefinedData(formats strfmt.Registry) error {

	if swag.IsZero(m.UserDefinedData) { // not required
		return nil
	}

	iUserDefinedDataSize := int64(len(m.UserDefinedData))

	if err := validate.MaxItems("attributes"+"."+"user_defined_data", "body", iUserDefinedDataSize, 5); err != nil {
		return err
	}

	for i := 0; i < len(m.UserDefinedData); i++ {
		if swag.IsZero(m.UserDefinedData[i]) { // not required
			continue
		}

		if m.UserDefinedData[i] != nil {
			if err := m.UserDefinedData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + "user_defined_data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountUpdateAttributes) validateUserDefinedInformation(formats strfmt.Registry) error {

	if swag.IsZero(m.UserDefinedInformation) { // not required
		return nil
	}

	if err := validate.MaxLength("attributes"+"."+"user_defined_information", "body", string(m.UserDefinedInformation), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountUpdateAttributes) validateValidationType(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationType) { // not required
		return nil
	}

	if err := m.ValidationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "validation_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountUpdateAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountUpdateAttributes) UnmarshalBinary(b []byte) error {
	var res AccountUpdateAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountUpdateAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
