// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UltimateEntity ultimate entity
// swagger:model UltimateEntity
type UltimateEntity struct {

	// Ultimate debtor/beneficiary address
	Address []string `json:"address,omitempty"`

	// Ultimate debtor/beneficiary birth city
	BirthCity string `json:"birth_city,omitempty"`

	// Ultimate debtor/beneficiary birth country. ISO 3166 format country code
	BirthCountry string `json:"birth_country,omitempty"`

	// Ultimate debtor/beneficiary birth date. Formatted ISO 8601 format YYYY-MM-DD
	// Format: date
	BirthDate *strfmt.Date `json:"birth_date,omitempty"`

	// Ultimate debtor/beneficiary birth province
	BirthProvince string `json:"birth_province,omitempty"`

	// Country of ultimate debtor/beneficiary address. ISO 3166 format country code
	Country string `json:"country,omitempty"`

	// Ultimate beneficiary name
	Name string `json:"name,omitempty"`

	// Organisation identification of an ultimate debtor/beneficiary, in the case that the ultimate debtor/beneficiary is an organisation and not a private person.
	OrganisationIdentification string `json:"organisation_identification,omitempty"`

	// The code that specifies the type of `organisation_identification`
	OrganisationIdentificationCode string `json:"organisation_identification_code,omitempty"`

	// Issuer of the `organisation_identification`
	OrganisationIdentificationIssuer string `json:"organisation_identification_issuer,omitempty"`

	// The code that specifies the scheme of `organisation_identification`
	OrganisationIdentificationScheme string `json:"organisation_identification_scheme,omitempty"`

	// private identification
	PrivateIdentification *PrivateIdentification `json:"private_identification,omitempty"`
}

// line 140

func UltimateEntityWithDefaults(defaults client.Defaults) *UltimateEntity {
	return &UltimateEntity{

		Address: make([]string, 0),

		BirthCity: defaults.GetString("UltimateEntity", "birth_city"),

		BirthCountry: defaults.GetString("UltimateEntity", "birth_country"),

		BirthDate: defaults.GetStrfmtDatePtr("UltimateEntity", "birth_date"),

		BirthProvince: defaults.GetString("UltimateEntity", "birth_province"),

		Country: defaults.GetString("UltimateEntity", "country"),

		Name: defaults.GetString("UltimateEntity", "name"),

		OrganisationIdentification: defaults.GetString("UltimateEntity", "organisation_identification"),

		OrganisationIdentificationCode: defaults.GetString("UltimateEntity", "organisation_identification_code"),

		OrganisationIdentificationIssuer: defaults.GetString("UltimateEntity", "organisation_identification_issuer"),

		OrganisationIdentificationScheme: defaults.GetString("UltimateEntity", "organisation_identification_scheme"),

		PrivateIdentification: PrivateIdentificationWithDefaults(defaults),
	}
}

func (m *UltimateEntity) WithAddress(address []string) *UltimateEntity {

	m.Address = address

	return m
}

func (m *UltimateEntity) WithBirthCity(birthCity string) *UltimateEntity {

	m.BirthCity = birthCity

	return m
}

func (m *UltimateEntity) WithBirthCountry(birthCountry string) *UltimateEntity {

	m.BirthCountry = birthCountry

	return m
}

func (m *UltimateEntity) WithBirthDate(birthDate strfmt.Date) *UltimateEntity {

	m.BirthDate = &birthDate

	return m
}

func (m *UltimateEntity) WithoutBirthDate() *UltimateEntity {
	m.BirthDate = nil
	return m
}

func (m *UltimateEntity) WithBirthProvince(birthProvince string) *UltimateEntity {

	m.BirthProvince = birthProvince

	return m
}

func (m *UltimateEntity) WithCountry(country string) *UltimateEntity {

	m.Country = country

	return m
}

func (m *UltimateEntity) WithName(name string) *UltimateEntity {

	m.Name = name

	return m
}

func (m *UltimateEntity) WithOrganisationIdentification(organisationIdentification string) *UltimateEntity {

	m.OrganisationIdentification = organisationIdentification

	return m
}

func (m *UltimateEntity) WithOrganisationIdentificationCode(organisationIdentificationCode string) *UltimateEntity {

	m.OrganisationIdentificationCode = organisationIdentificationCode

	return m
}

func (m *UltimateEntity) WithOrganisationIdentificationIssuer(organisationIdentificationIssuer string) *UltimateEntity {

	m.OrganisationIdentificationIssuer = organisationIdentificationIssuer

	return m
}

func (m *UltimateEntity) WithOrganisationIdentificationScheme(organisationIdentificationScheme string) *UltimateEntity {

	m.OrganisationIdentificationScheme = organisationIdentificationScheme

	return m
}

func (m *UltimateEntity) WithPrivateIdentification(privateIdentification PrivateIdentification) *UltimateEntity {

	m.PrivateIdentification = &privateIdentification

	return m
}

func (m *UltimateEntity) WithoutPrivateIdentification() *UltimateEntity {
	m.PrivateIdentification = nil
	return m
}

// Validate validates this ultimate entity
func (m *UltimateEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UltimateEntity) validateBirthDate(formats strfmt.Registry) error {

	if swag.IsZero(m.BirthDate) { // not required
		return nil
	}

	if err := validate.FormatOf("birth_date", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UltimateEntity) validatePrivateIdentification(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *UltimateEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UltimateEntity) UnmarshalBinary(b []byte) error {
	var res UltimateEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *UltimateEntity) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
