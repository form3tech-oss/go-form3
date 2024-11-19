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

// QueryUltimateEntity query ultimate entity
// swagger:model QueryUltimateEntity
type QueryUltimateEntity struct {

	// Additional address line of the debtor/beneficiary address
	// Max Length: 70
	AdditionalAddressLine string `json:"additional_address_line,omitempty"`

	// Ultimate debtor/beneficiary birth city
	// Max Length: 35
	BirthCity string `json:"birth_city,omitempty"`

	// Ultimate debtor/beneficiary birth country. ISO 3166 format country code
	// Pattern: ^[A-Z]{2,2}$
	BirthCountry string `json:"birth_country,omitempty"`

	// Ultimate debtor/beneficiary birth date. Formatted ISO 8601 format YYYY-MM-DD
	// Format: date
	BirthDate *strfmt.Date `json:"birth_date,omitempty"`

	// Ultimate debtor/beneficiary birth province
	// Max Length: 35
	BirthProvince string `json:"birth_province,omitempty"`

	// Building number of the debtor/beneficiary address
	// Max Length: 16
	BuildingNumber string `json:"building_number,omitempty"`

	// City/Town of the debtor/beneficiary address
	// Max Length: 35
	City string `json:"city,omitempty"`

	// Country of ultimate debtor/beneficiary address. ISO 3166 format country code
	// Pattern: ^[A-Z]{2,2}$
	Country string `json:"country,omitempty"`

	// Country of residence of the ultimate debtor/beneficiary, ISO 3166 format country code
	// Pattern: ^[A-Z]{2,2}$
	CountryOfResidence string `json:"country_of_residence,omitempty"`

	// Ultimate debtor/beneficiary name
	// Max Length: 140
	Name string `json:"name,omitempty"`

	// Organisation identification of an ultimate debtor/beneficiary, in the case that the ultimate debtor/beneficiary is an organisation and not a private person.
	// Max Length: 35
	OrganisationIdentification string `json:"organisation_identification,omitempty"`

	// The code that specifies the type of `organisation_identification`
	OrganisationIdentificationCode string `json:"organisation_identification_code,omitempty"`

	// Issuer of the `organisation_identification`
	// Max Length: 35
	OrganisationIdentificationIssuer string `json:"organisation_identification_issuer,omitempty"`

	// The code that specifies the scheme of `organisation_identification`
	// Max Length: 35
	OrganisationIdentificationScheme string `json:"organisation_identification_scheme,omitempty"`

	// Array for additional ID(s) of ultimate organisation
	OrganisationIdentifications []*QueryBeneficiaryDebtorOrganisationIdentification `json:"organisation_identifications,omitempty"`

	// Post code of the debtor/beneficiary address
	// Max Length: 16
	PostCode string `json:"post_code,omitempty"`

	// private identification
	PrivateIdentification *QueryPrivateIdentification `json:"private_identification,omitempty"`

	// Province of the debtor/beneficiary address
	// Max Length: 35
	Province string `json:"province,omitempty"`

	// Street name of the debtor/beneficiary address
	// Max Length: 70
	StreetName string `json:"street_name,omitempty"`
}

func QueryUltimateEntityWithDefaults(defaults client.Defaults) *QueryUltimateEntity {
	return &QueryUltimateEntity{

		AdditionalAddressLine: defaults.GetString("QueryUltimateEntity", "additional_address_line"),

		BirthCity: defaults.GetString("QueryUltimateEntity", "birth_city"),

		BirthCountry: defaults.GetString("QueryUltimateEntity", "birth_country"),

		BirthDate: defaults.GetStrfmtDatePtr("QueryUltimateEntity", "birth_date"),

		BirthProvince: defaults.GetString("QueryUltimateEntity", "birth_province"),

		BuildingNumber: defaults.GetString("QueryUltimateEntity", "building_number"),

		City: defaults.GetString("QueryUltimateEntity", "city"),

		Country: defaults.GetString("QueryUltimateEntity", "country"),

		CountryOfResidence: defaults.GetString("QueryUltimateEntity", "country_of_residence"),

		Name: defaults.GetString("QueryUltimateEntity", "name"),

		OrganisationIdentification: defaults.GetString("QueryUltimateEntity", "organisation_identification"),

		OrganisationIdentificationCode: defaults.GetString("QueryUltimateEntity", "organisation_identification_code"),

		OrganisationIdentificationIssuer: defaults.GetString("QueryUltimateEntity", "organisation_identification_issuer"),

		OrganisationIdentificationScheme: defaults.GetString("QueryUltimateEntity", "organisation_identification_scheme"),

		OrganisationIdentifications: make([]*QueryBeneficiaryDebtorOrganisationIdentification, 0),

		PostCode: defaults.GetString("QueryUltimateEntity", "post_code"),

		PrivateIdentification: QueryPrivateIdentificationWithDefaults(defaults),

		Province: defaults.GetString("QueryUltimateEntity", "province"),

		StreetName: defaults.GetString("QueryUltimateEntity", "street_name"),
	}
}

func (m *QueryUltimateEntity) WithAdditionalAddressLine(additionalAddressLine string) *QueryUltimateEntity {

	m.AdditionalAddressLine = additionalAddressLine

	return m
}

func (m *QueryUltimateEntity) WithBirthCity(birthCity string) *QueryUltimateEntity {

	m.BirthCity = birthCity

	return m
}

func (m *QueryUltimateEntity) WithBirthCountry(birthCountry string) *QueryUltimateEntity {

	m.BirthCountry = birthCountry

	return m
}

func (m *QueryUltimateEntity) WithBirthDate(birthDate strfmt.Date) *QueryUltimateEntity {

	m.BirthDate = &birthDate

	return m
}

func (m *QueryUltimateEntity) WithoutBirthDate() *QueryUltimateEntity {
	m.BirthDate = nil
	return m
}

func (m *QueryUltimateEntity) WithBirthProvince(birthProvince string) *QueryUltimateEntity {

	m.BirthProvince = birthProvince

	return m
}

func (m *QueryUltimateEntity) WithBuildingNumber(buildingNumber string) *QueryUltimateEntity {

	m.BuildingNumber = buildingNumber

	return m
}

func (m *QueryUltimateEntity) WithCity(city string) *QueryUltimateEntity {

	m.City = city

	return m
}

func (m *QueryUltimateEntity) WithCountry(country string) *QueryUltimateEntity {

	m.Country = country

	return m
}

func (m *QueryUltimateEntity) WithCountryOfResidence(countryOfResidence string) *QueryUltimateEntity {

	m.CountryOfResidence = countryOfResidence

	return m
}

func (m *QueryUltimateEntity) WithName(name string) *QueryUltimateEntity {

	m.Name = name

	return m
}

func (m *QueryUltimateEntity) WithOrganisationIdentification(organisationIdentification string) *QueryUltimateEntity {

	m.OrganisationIdentification = organisationIdentification

	return m
}

func (m *QueryUltimateEntity) WithOrganisationIdentificationCode(organisationIdentificationCode string) *QueryUltimateEntity {

	m.OrganisationIdentificationCode = organisationIdentificationCode

	return m
}

func (m *QueryUltimateEntity) WithOrganisationIdentificationIssuer(organisationIdentificationIssuer string) *QueryUltimateEntity {

	m.OrganisationIdentificationIssuer = organisationIdentificationIssuer

	return m
}

func (m *QueryUltimateEntity) WithOrganisationIdentificationScheme(organisationIdentificationScheme string) *QueryUltimateEntity {

	m.OrganisationIdentificationScheme = organisationIdentificationScheme

	return m
}

func (m *QueryUltimateEntity) WithOrganisationIdentifications(organisationIdentifications []*QueryBeneficiaryDebtorOrganisationIdentification) *QueryUltimateEntity {

	m.OrganisationIdentifications = organisationIdentifications

	return m
}

func (m *QueryUltimateEntity) WithPostCode(postCode string) *QueryUltimateEntity {

	m.PostCode = postCode

	return m
}

func (m *QueryUltimateEntity) WithPrivateIdentification(privateIdentification QueryPrivateIdentification) *QueryUltimateEntity {

	m.PrivateIdentification = &privateIdentification

	return m
}

func (m *QueryUltimateEntity) WithoutPrivateIdentification() *QueryUltimateEntity {
	m.PrivateIdentification = nil
	return m
}

func (m *QueryUltimateEntity) WithProvince(province string) *QueryUltimateEntity {

	m.Province = province

	return m
}

func (m *QueryUltimateEntity) WithStreetName(streetName string) *QueryUltimateEntity {

	m.StreetName = streetName

	return m
}

// Validate validates this query ultimate entity
func (m *QueryUltimateEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalAddressLine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthProvince(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildingNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountryOfResidence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationIdentificationIssuer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationIdentificationScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationIdentifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvince(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreetName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryUltimateEntity) validateAdditionalAddressLine(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalAddressLine) { // not required
		return nil
	}

	if err := validate.MaxLength("additional_address_line", "body", string(m.AdditionalAddressLine), 70); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateBirthCity(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthCity) { // not required
		return nil
	}

	if err := validate.MaxLength("birth_city", "body", string(m.BirthCity), 35); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateBirthCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthCountry) { // not required
		return nil
	}

	if err := validate.Pattern("birth_country", "body", string(m.BirthCountry), `^[A-Z]{2,2}$`); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateBirthDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthDate) { // not required
		return nil
	}

	if err := validate.FormatOf("birth_date", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateBirthProvince(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthProvince) { // not required
		return nil
	}

	if err := validate.MaxLength("birth_province", "body", string(m.BirthProvince), 35); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateBuildingNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildingNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("building_number", "body", string(m.BuildingNumber), 16); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateCity(formats strfmt.Registry) error {

	if swag.IsZero(m.City) { // not required
		return nil
	}

	if err := validate.MaxLength("city", "body", string(m.City), 35); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if err := validate.Pattern("country", "body", string(m.Country), `^[A-Z]{2,2}$`); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateCountryOfResidence(formats strfmt.Registry) error {

	if swag.IsZero(m.CountryOfResidence) { // not required
		return nil
	}

	if err := validate.Pattern("country_of_residence", "body", string(m.CountryOfResidence), `^[A-Z]{2,2}$`); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 140); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateOrganisationIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationIdentification) { // not required
		return nil
	}

	if err := validate.MaxLength("organisation_identification", "body", string(m.OrganisationIdentification), 35); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateOrganisationIdentificationIssuer(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationIdentificationIssuer) { // not required
		return nil
	}

	if err := validate.MaxLength("organisation_identification_issuer", "body", string(m.OrganisationIdentificationIssuer), 35); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateOrganisationIdentificationScheme(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationIdentificationScheme) { // not required
		return nil
	}

	if err := validate.MaxLength("organisation_identification_scheme", "body", string(m.OrganisationIdentificationScheme), 35); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateOrganisationIdentifications(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationIdentifications) { // not required
		return nil
	}

	for i := 0; i < len(m.OrganisationIdentifications); i++ {
		if swag.IsZero(m.OrganisationIdentifications[i]) { // not required
			continue
		}

		if m.OrganisationIdentifications[i] != nil {
			if err := m.OrganisationIdentifications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organisation_identifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueryUltimateEntity) validatePostCode(formats strfmt.Registry) error {

	if swag.IsZero(m.PostCode) { // not required
		return nil
	}

	if err := validate.MaxLength("post_code", "body", string(m.PostCode), 16); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validatePrivateIdentification(formats strfmt.Registry) error {

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

func (m *QueryUltimateEntity) validateProvince(formats strfmt.Registry) error {

	if swag.IsZero(m.Province) { // not required
		return nil
	}

	if err := validate.MaxLength("province", "body", string(m.Province), 35); err != nil {
		return err
	}

	return nil
}

func (m *QueryUltimateEntity) validateStreetName(formats strfmt.Registry) error {

	if swag.IsZero(m.StreetName) { // not required
		return nil
	}

	if err := validate.MaxLength("street_name", "body", string(m.StreetName), 70); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryUltimateEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryUltimateEntity) UnmarshalBinary(b []byte) error {
	var res QueryUltimateEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryUltimateEntity) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
