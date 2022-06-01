// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountAttributesOrganisationIdentification account attributes organisation identification
// swagger:model AccountAttributesOrganisationIdentification
type AccountAttributesOrganisationIdentification struct {

	// actors
	Actors []*AccountAttributesOrganisationIdentificationActor `json:"actors"`

	// address
	Address []string `json:"address"`

	// city
	// Max Length: 35
	// Min Length: 1
	City string `json:"city,omitempty"`

	// country
	// Pattern: ^[A-Z]{2}$
	Country string `json:"country,omitempty"`

	// identification
	// Max Length: 140
	// Min Length: 1
	Identification string `json:"identification,omitempty"`

	// identification issuer
	IdentificationIssuer string `json:"identification_issuer,omitempty"`

	// identification scheme
	// Max Length: 35
	// Min Length: 1
	IdentificationScheme string `json:"identification_scheme,omitempty"`

	// identification scheme code
	// Max Length: 35
	// Min Length: 1
	IdentificationSchemeCode string `json:"identification_scheme_code,omitempty"`

	// registration number
	RegistrationNumber string `json:"registration_number,omitempty"`

	// ISO 3166-1 code used to identify the domicile of the account
	// Pattern: ^[A-Z]{2}$
	TaxResidency string `json:"tax_residency,omitempty"`
}

func AccountAttributesOrganisationIdentificationWithDefaults(defaults client.Defaults) *AccountAttributesOrganisationIdentification {
	return &AccountAttributesOrganisationIdentification{

		Actors: make([]*AccountAttributesOrganisationIdentificationActor, 0),

		Address: make([]string, 0),

		City: defaults.GetString("AccountAttributesOrganisationIdentification", "city"),

		Country: defaults.GetString("AccountAttributesOrganisationIdentification", "country"),

		Identification: defaults.GetString("AccountAttributesOrganisationIdentification", "identification"),

		IdentificationIssuer: defaults.GetString("AccountAttributesOrganisationIdentification", "identification_issuer"),

		IdentificationScheme: defaults.GetString("AccountAttributesOrganisationIdentification", "identification_scheme"),

		IdentificationSchemeCode: defaults.GetString("AccountAttributesOrganisationIdentification", "identification_scheme_code"),

		RegistrationNumber: defaults.GetString("AccountAttributesOrganisationIdentification", "registration_number"),

		TaxResidency: defaults.GetString("AccountAttributesOrganisationIdentification", "tax_residency"),
	}
}

func (m *AccountAttributesOrganisationIdentification) WithActors(actors []*AccountAttributesOrganisationIdentificationActor) *AccountAttributesOrganisationIdentification {

	m.Actors = actors

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithAddress(address []string) *AccountAttributesOrganisationIdentification {

	m.Address = address

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithCity(city string) *AccountAttributesOrganisationIdentification {

	m.City = city

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithCountry(country string) *AccountAttributesOrganisationIdentification {

	m.Country = country

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithIdentification(identification string) *AccountAttributesOrganisationIdentification {

	m.Identification = identification

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithIdentificationIssuer(identificationIssuer string) *AccountAttributesOrganisationIdentification {

	m.IdentificationIssuer = identificationIssuer

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithIdentificationScheme(identificationScheme string) *AccountAttributesOrganisationIdentification {

	m.IdentificationScheme = identificationScheme

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithIdentificationSchemeCode(identificationSchemeCode string) *AccountAttributesOrganisationIdentification {

	m.IdentificationSchemeCode = identificationSchemeCode

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithRegistrationNumber(registrationNumber string) *AccountAttributesOrganisationIdentification {

	m.RegistrationNumber = registrationNumber

	return m
}

func (m *AccountAttributesOrganisationIdentification) WithTaxResidency(taxResidency string) *AccountAttributesOrganisationIdentification {

	m.TaxResidency = taxResidency

	return m
}

// Validate validates this account attributes organisation identification
func (m *AccountAttributesOrganisationIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentificationScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentificationSchemeCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxResidency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateActors(formats strfmt.Registry) error {

	if swag.IsZero(m.Actors) { // not required
		return nil
	}

	for i := 0; i < len(m.Actors); i++ {
		if swag.IsZero(m.Actors[i]) { // not required
			continue
		}

		if m.Actors[i] != nil {
			if err := m.Actors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	for i := 0; i < len(m.Address); i++ {

		if err := validate.MinLength("address"+"."+strconv.Itoa(i), "body", string(m.Address[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("address"+"."+strconv.Itoa(i), "body", string(m.Address[i]), 140); err != nil {
			return err
		}

	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateCity(formats strfmt.Registry) error {

	if swag.IsZero(m.City) { // not required
		return nil
	}

	if err := validate.MinLength("city", "body", string(m.City), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("city", "body", string(m.City), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if err := validate.Pattern("country", "body", string(m.Country), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.Identification) { // not required
		return nil
	}

	if err := validate.MinLength("identification", "body", string(m.Identification), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("identification", "body", string(m.Identification), 140); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateIdentificationScheme(formats strfmt.Registry) error {

	if swag.IsZero(m.IdentificationScheme) { // not required
		return nil
	}

	if err := validate.MinLength("identification_scheme", "body", string(m.IdentificationScheme), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("identification_scheme", "body", string(m.IdentificationScheme), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateIdentificationSchemeCode(formats strfmt.Registry) error {

	if swag.IsZero(m.IdentificationSchemeCode) { // not required
		return nil
	}

	if err := validate.MinLength("identification_scheme_code", "body", string(m.IdentificationSchemeCode), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("identification_scheme_code", "body", string(m.IdentificationSchemeCode), 35); err != nil {
		return err
	}

	return nil
}

func (m *AccountAttributesOrganisationIdentification) validateTaxResidency(formats strfmt.Registry) error {

	if swag.IsZero(m.TaxResidency) { // not required
		return nil
	}

	if err := validate.Pattern("tax_residency", "body", string(m.TaxResidency), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAttributesOrganisationIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAttributesOrganisationIdentification) UnmarshalBinary(b []byte) error {
	var res AccountAttributesOrganisationIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAttributesOrganisationIdentification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
