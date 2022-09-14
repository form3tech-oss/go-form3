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

// AccountAttributes account attributes
// swagger:model AccountAttributes
type AccountAttributes struct {

	// acceptance qualifier
	AcceptanceQualifier AcceptanceQualifier `json:"acceptance_qualifier,omitempty"`

	// Is the account business or personal?
	// Enum: [Personal Business]
	AccountClassification *string `json:"account_classification,omitempty"`

	// - deprecated - Is the account opted out of account matching, e.g. CoP?
	AccountMatchingOptOut *bool `json:"account_matching_opt_out,omitempty"`

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
	// Required: true
	// Pattern: ^[A-Z]{2}$
	Country *string `json:"country"`

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
	JointAccount *bool `json:"joint_account,omitempty"`

	// Account holder names (for example title, first name, last name). Used for Confirmation of Payee matching.
	// Max Items: 4
	Name []string `json:"name"`

	// Describes the status of the account for name matching via CoP. The value determines the code with which Form3 responds to matched CoP requests to this account.
	// Enum: [supported switched opted_out not_supported]
	NameMatchingStatus *string `json:"name_matching_status,omitempty"`

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
	Switched *bool `json:"switched,omitempty"`

	// - deprecated - Customer title.
	// Max Length: 40
	// Min Length: 1
	Title string `json:"title,omitempty"`

	// All purpose list of key-value pairs to store specific data for the associated account. It will be added to each payment received to an account.
	// Max Items: 5
	UserDefinedData []*UserDefinedData `json:"user_defined_data,omitempty"`

	// - deprecated - All purpose field to store specific data for the associated account. It will be added to each payment received to an account.
	// Max Length: 35
	UserDefinedInformation string `json:"user_defined_information,omitempty"`

	// validation type
	ValidationType ValidationType `json:"validation_type,omitempty"`
}

func AccountAttributesWithDefaults(defaults client.Defaults) *AccountAttributes {
	return &AccountAttributes{

		// TODO AcceptanceQualifier: AcceptanceQualifier,

		AccountClassification: defaults.GetStringPtr("AccountAttributes", "account_classification"),

		AccountMatchingOptOut: defaults.GetBoolPtr("AccountAttributes", "account_matching_opt_out"),

		AccountNumber: defaults.GetString("AccountAttributes", "account_number"),

		AlternativeBankAccountNames: make([]string, 0),

		AlternativeNames: make([]string, 0),

		BankAccountName: defaults.GetString("AccountAttributes", "bank_account_name"),

		BankID: defaults.GetString("AccountAttributes", "bank_id"),

		BankIDCode: defaults.GetString("AccountAttributes", "bank_id_code"),

		BaseCurrency: defaults.GetString("AccountAttributes", "base_currency"),

		Bic: defaults.GetString("AccountAttributes", "bic"),

		Country: defaults.GetStringPtr("AccountAttributes", "country"),

		CustomerID: defaults.GetString("AccountAttributes", "customer_id"),

		FirstName: defaults.GetString("AccountAttributes", "first_name"),

		Iban: defaults.GetString("AccountAttributes", "iban"),

		JointAccount: defaults.GetBoolPtr("AccountAttributes", "joint_account"),

		Name: make([]string, 0),

		NameMatchingStatus: defaults.GetStringPtr("AccountAttributes", "name_matching_status"),

		OrganisationIdentification: AccountAttributesOrganisationIdentificationWithDefaults(defaults),

		PrivateIdentification: AccountAttributesPrivateIdentificationWithDefaults(defaults),

		ProcessingService: defaults.GetString("AccountAttributes", "processing_service"),

		ReferenceMask: defaults.GetString("AccountAttributes", "reference_mask"),

		SecondaryIdentification: defaults.GetString("AccountAttributes", "secondary_identification"),

		Status: defaults.GetString("AccountAttributes", "status"),

		// TODO StatusReason: StatusReason,

		Switched: defaults.GetBoolPtr("AccountAttributes", "switched"),

		Title: defaults.GetString("AccountAttributes", "title"),

		UserDefinedData: make([]*UserDefinedData, 0),

		UserDefinedInformation: defaults.GetString("AccountAttributes", "user_defined_information"),

		// TODO ValidationType: ValidationType,

	}
}

func (m *AccountAttributes) WithAcceptanceQualifier(acceptanceQualifier AcceptanceQualifier) *AccountAttributes {

	m.AcceptanceQualifier = acceptanceQualifier

	return m
}

func (m *AccountAttributes) WithAccountClassification(accountClassification string) *AccountAttributes {

	m.AccountClassification = &accountClassification

	return m
}

func (m *AccountAttributes) WithoutAccountClassification() *AccountAttributes {
	m.AccountClassification = nil
	return m
}

func (m *AccountAttributes) WithAccountMatchingOptOut(accountMatchingOptOut bool) *AccountAttributes {

	m.AccountMatchingOptOut = &accountMatchingOptOut

	return m
}

func (m *AccountAttributes) WithoutAccountMatchingOptOut() *AccountAttributes {
	m.AccountMatchingOptOut = nil
	return m
}

func (m *AccountAttributes) WithAccountNumber(accountNumber string) *AccountAttributes {

	m.AccountNumber = accountNumber

	return m
}

func (m *AccountAttributes) WithAlternativeBankAccountNames(alternativeBankAccountNames []string) *AccountAttributes {

	m.AlternativeBankAccountNames = alternativeBankAccountNames

	return m
}

func (m *AccountAttributes) WithAlternativeNames(alternativeNames []string) *AccountAttributes {

	m.AlternativeNames = alternativeNames

	return m
}

func (m *AccountAttributes) WithBankAccountName(bankAccountName string) *AccountAttributes {

	m.BankAccountName = bankAccountName

	return m
}

func (m *AccountAttributes) WithBankID(bankID string) *AccountAttributes {

	m.BankID = bankID

	return m
}

func (m *AccountAttributes) WithBankIDCode(bankIDCode string) *AccountAttributes {

	m.BankIDCode = bankIDCode

	return m
}

func (m *AccountAttributes) WithBaseCurrency(baseCurrency string) *AccountAttributes {

	m.BaseCurrency = baseCurrency

	return m
}

func (m *AccountAttributes) WithBic(bic string) *AccountAttributes {

	m.Bic = bic

	return m
}

func (m *AccountAttributes) WithCountry(country string) *AccountAttributes {

	m.Country = &country

	return m
}

func (m *AccountAttributes) WithoutCountry() *AccountAttributes {
	m.Country = nil
	return m
}

func (m *AccountAttributes) WithCustomerID(customerID string) *AccountAttributes {

	m.CustomerID = customerID

	return m
}

func (m *AccountAttributes) WithFirstName(firstName string) *AccountAttributes {

	m.FirstName = firstName

	return m
}

func (m *AccountAttributes) WithIban(iban string) *AccountAttributes {

	m.Iban = iban

	return m
}

func (m *AccountAttributes) WithJointAccount(jointAccount bool) *AccountAttributes {

	m.JointAccount = &jointAccount

	return m
}

func (m *AccountAttributes) WithoutJointAccount() *AccountAttributes {
	m.JointAccount = nil
	return m
}

func (m *AccountAttributes) WithName(name []string) *AccountAttributes {

	m.Name = name

	return m
}

func (m *AccountAttributes) WithNameMatchingStatus(nameMatchingStatus string) *AccountAttributes {

	m.NameMatchingStatus = &nameMatchingStatus

	return m
}

func (m *AccountAttributes) WithoutNameMatchingStatus() *AccountAttributes {
	m.NameMatchingStatus = nil
	return m
}

func (m *AccountAttributes) WithOrganisationIdentification(organisationIdentification AccountAttributesOrganisationIdentification) *AccountAttributes {

	m.OrganisationIdentification = &organisationIdentification

	return m
}

func (m *AccountAttributes) WithoutOrganisationIdentification() *AccountAttributes {
	m.OrganisationIdentification = nil
	return m
}

func (m *AccountAttributes) WithPrivateIdentification(privateIdentification AccountAttributesPrivateIdentification) *AccountAttributes {

	m.PrivateIdentification = &privateIdentification

	return m
}

func (m *AccountAttributes) WithoutPrivateIdentification() *AccountAttributes {
	m.PrivateIdentification = nil
	return m
}

func (m *AccountAttributes) WithProcessingService(processingService string) *AccountAttributes {

	m.ProcessingService = processingService

	return m
}

func (m *AccountAttributes) WithReferenceMask(referenceMask string) *AccountAttributes {

	m.ReferenceMask = referenceMask

	return m
}

func (m *AccountAttributes) WithSecondaryIdentification(secondaryIdentification string) *AccountAttributes {

	m.SecondaryIdentification = secondaryIdentification

	return m
}

func (m *AccountAttributes) WithStatus(status string) *AccountAttributes {

	m.Status = status

	return m
}

func (m *AccountAttributes) WithStatusReason(statusReason StatusReason) *AccountAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *AccountAttributes) WithSwitched(switched bool) *AccountAttributes {

	m.Switched = &switched

	return m
}

func (m *AccountAttributes) WithoutSwitched() *AccountAttributes {
	m.Switched = nil
	return m
}

func (m *AccountAttributes) WithTitle(title string) *AccountAttributes {

	m.Title = title

	return m
}

func (m *AccountAttributes) WithUserDefinedData(userDefinedData []*UserDefinedData) *AccountAttributes {

	m.UserDefinedData = userDefinedData

	return m
}

func (m *AccountAttributes) WithUserDefinedInformation(userDefinedInformation string) *AccountAttributes {

	m.UserDefinedInformation = userDefinedInformation

	return m
}

func (m *AccountAttributes) WithValidationType(validationType ValidationType) *AccountAttributes {

	m.ValidationType = validationType

	return m
}

// Validate validates this account attributes
func (m *AccountAttributes) Validate(formats strfmt.Registry) error {
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

	if err := m.validateTitle(formats); err != nil {
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

func (m *AccountAttributes) validateAcceptanceQualifier(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptanceQualifier) { // not required
		return nil
	}

	if err := m.AcceptanceQualifier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("acceptance_qualifier")
		}
		return err
	}

	return nil
}

var accountAttributesTypeAccountClassificationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Personal","Business"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountAttributesTypeAccountClassificationPropEnum = append(accountAttributesTypeAccountClassificationPropEnum, v)
	}
}

const (

	// AccountAttributesAccountClassificationPersonal captures enum value "Personal"
	AccountAttributesAccountClassificationPersonal string = "Personal"

	// AccountAttributesAccountClassificationBusiness captures enum value "Business"
	AccountAttributesAccountClassificationBusiness string = "Business"
)

// prop value enum
func (m *AccountAttributes) validateAccountClassificationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountAttributesTypeAccountClassificationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountAttributes) validateAccountClassification(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountClassification) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccountClassificationEnum("account_classification", "body", *m.AccountClassification); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumber) { // not required
		return nil
	}

	if err := validate.Pattern("account_number", "body", string(m.AccountNumber), `^[A-Z0-9]{0,64}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateAlternativeBankAccountNames(formats strfmt.Registry) error {

	if swag.IsZero(m.AlternativeBankAccountNames) { // not required
		return nil
	}

	iAlternativeBankAccountNamesSize := int64(len(m.AlternativeBankAccountNames))

	if err := validate.MaxItems("alternative_bank_account_names", "body", iAlternativeBankAccountNamesSize, 3); err != nil {
		return err
	}

	for i := 0; i < len(m.AlternativeBankAccountNames); i++ {

		if err := validate.MinLength("alternative_bank_account_names"+"."+strconv.Itoa(i), "body", string(m.AlternativeBankAccountNames[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("alternative_bank_account_names"+"."+strconv.Itoa(i), "body", string(m.AlternativeBankAccountNames[i]), 140); err != nil {
			return err
		}

	}

	return nil
}

func (m *AccountAttributes) validateAlternativeNames(formats strfmt.Registry) error {

	if swag.IsZero(m.AlternativeNames) { // not required
		return nil
	}

	iAlternativeNamesSize := int64(len(m.AlternativeNames))

	if err := validate.MaxItems("alternative_names", "body", iAlternativeNamesSize, 3); err != nil {
		return err
	}

	for i := 0; i < len(m.AlternativeNames); i++ {

		if err := validate.MinLength("alternative_names"+"."+strconv.Itoa(i), "body", string(m.AlternativeNames[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("alternative_names"+"."+strconv.Itoa(i), "body", string(m.AlternativeNames[i]), 140); err != nil {
			return err
		}

	}

	return nil
}

func (m *AccountAttributes) validateBankAccountName(formats strfmt.Registry) error {

	if swag.IsZero(m.BankAccountName) { // not required
		return nil
	}

	if err := validate.MinLength("bank_account_name", "body", string(m.BankAccountName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("bank_account_name", "body", string(m.BankAccountName), 140); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateBankID(formats strfmt.Registry) error {

	if swag.IsZero(m.BankID) { // not required
		return nil
	}

	if err := validate.Pattern("bank_id", "body", string(m.BankID), `^[A-Z0-9]{0,16}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateBankIDCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankIDCode) { // not required
		return nil
	}

	if err := validate.Pattern("bank_id_code", "body", string(m.BankIDCode), `^[A-Z]{0,16}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateBaseCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.BaseCurrency) { // not required
		return nil
	}

	if err := validate.Pattern("base_currency", "body", string(m.BaseCurrency), `^[A-Z]{3}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateBic(formats strfmt.Registry) error {

	if swag.IsZero(m.Bic) { // not required
		return nil
	}

	if err := validate.Pattern("bic", "body", string(m.Bic), `^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	if err := validate.Pattern("country", "body", string(*m.Country), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateCustomerID(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := validate.Pattern("customer_id", "body", string(m.CustomerID), `^[a-zA-Z0-9-$@., ]{0,256}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateFirstName(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstName) { // not required
		return nil
	}

	if err := validate.MinLength("first_name", "body", string(m.FirstName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("first_name", "body", string(m.FirstName), 40); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateIban(formats strfmt.Registry) error {

	if swag.IsZero(m.Iban) { // not required
		return nil
	}

	if err := validate.Pattern("iban", "body", string(m.Iban), `^[A-Z]{2}[0-9]{2}[A-Z0-9]{0,64}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	iNameSize := int64(len(m.Name))

	if err := validate.MaxItems("name", "body", iNameSize, 4); err != nil {
		return err
	}

	for i := 0; i < len(m.Name); i++ {

		if err := validate.MinLength("name"+"."+strconv.Itoa(i), "body", string(m.Name[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("name"+"."+strconv.Itoa(i), "body", string(m.Name[i]), 140); err != nil {
			return err
		}

	}

	return nil
}

var accountAttributesTypeNameMatchingStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["supported","switched","opted_out","not_supported"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountAttributesTypeNameMatchingStatusPropEnum = append(accountAttributesTypeNameMatchingStatusPropEnum, v)
	}
}

const (

	// AccountAttributesNameMatchingStatusSupported captures enum value "supported"
	AccountAttributesNameMatchingStatusSupported string = "supported"

	// AccountAttributesNameMatchingStatusSwitched captures enum value "switched"
	AccountAttributesNameMatchingStatusSwitched string = "switched"

	// AccountAttributesNameMatchingStatusOptedOut captures enum value "opted_out"
	AccountAttributesNameMatchingStatusOptedOut string = "opted_out"

	// AccountAttributesNameMatchingStatusNotSupported captures enum value "not_supported"
	AccountAttributesNameMatchingStatusNotSupported string = "not_supported"
)

// prop value enum
func (m *AccountAttributes) validateNameMatchingStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountAttributesTypeNameMatchingStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountAttributes) validateNameMatchingStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.NameMatchingStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameMatchingStatusEnum("name_matching_status", "body", *m.NameMatchingStatus); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateOrganisationIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationIdentification) { // not required
		return nil
	}

	if m.OrganisationIdentification != nil {
		if err := m.OrganisationIdentification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organisation_identification")
			}
			return err
		}
	}

	return nil
}

func (m *AccountAttributes) validatePrivateIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivateIdentification) { // not required
		return nil
	}

	if m.PrivateIdentification != nil {
		if err := m.PrivateIdentification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("private_identification")
			}
			return err
		}
	}

	return nil
}

func (m *AccountAttributes) validateProcessingService(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessingService) { // not required
		return nil
	}

	if err := validate.MaxLength("processing_service", "body", string(m.ProcessingService), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateReferenceMask(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceMask) { // not required
		return nil
	}

	if err := validate.MaxLength("reference_mask", "body", string(m.ReferenceMask), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateSecondaryIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryIdentification) { // not required
		return nil
	}

	if err := validate.MinLength("secondary_identification", "body", string(m.SecondaryIdentification), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("secondary_identification", "body", string(m.SecondaryIdentification), 140); err != nil {
		return err
	}

	return nil
}

var accountAttributesTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","failed","confirmed","closed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountAttributesTypeStatusPropEnum = append(accountAttributesTypeStatusPropEnum, v)
	}
}

const (

	// AccountAttributesStatusPending captures enum value "pending"
	AccountAttributesStatusPending string = "pending"

	// AccountAttributesStatusFailed captures enum value "failed"
	AccountAttributesStatusFailed string = "failed"

	// AccountAttributesStatusConfirmed captures enum value "confirmed"
	AccountAttributesStatusConfirmed string = "confirmed"

	// AccountAttributesStatusClosed captures enum value "closed"
	AccountAttributesStatusClosed string = "closed"
)

// prop value enum
func (m *AccountAttributes) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountAttributesTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateStatusReason(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusReason) { // not required
		return nil
	}

	if err := m.StatusReason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status_reason")
		}
		return err
	}

	return nil
}

func (m *AccountAttributes) validateTitle(formats strfmt.Registry) error {

	if swag.IsZero(m.Title) { // not required
		return nil
	}

	if err := validate.MinLength("title", "body", string(m.Title), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", string(m.Title), 40); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateUserDefinedData(formats strfmt.Registry) error {

	if swag.IsZero(m.UserDefinedData) { // not required
		return nil
	}

	iUserDefinedDataSize := int64(len(m.UserDefinedData))

	if err := validate.MaxItems("user_defined_data", "body", iUserDefinedDataSize, 5); err != nil {
		return err
	}

	for i := 0; i < len(m.UserDefinedData); i++ {
		if swag.IsZero(m.UserDefinedData[i]) { // not required
			continue
		}

		if m.UserDefinedData[i] != nil {
			if err := m.UserDefinedData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_defined_data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountAttributes) validateUserDefinedInformation(formats strfmt.Registry) error {

	if swag.IsZero(m.UserDefinedInformation) { // not required
		return nil
	}

	if err := validate.MaxLength("user_defined_information", "body", string(m.UserDefinedInformation), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributes) validateValidationType(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationType) { // not required
		return nil
	}

	if err := m.ValidationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("validation_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAttributes) UnmarshalBinary(b []byte) error {
	var res AccountAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
