// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RequestForInformationCreatorParty request for information creator party
// swagger:model RequestForInformationCreatorParty
type RequestForInformationCreatorParty struct {

	// birth city
	BirthCity string `json:"birth_city,omitempty"`

	// birth country
	BirthCountry string `json:"birth_country,omitempty"`

	// birth date
	BirthDate string `json:"birth_date,omitempty"`

	// birth province
	BirthProvince string `json:"birth_province,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organisation identification
	OrganisationIdentification string `json:"organisation_identification,omitempty"`

	// organisation identification code
	OrganisationIdentificationCode string `json:"organisation_identification_code,omitempty"`

	// organisation identification issuer
	OrganisationIdentificationIssuer string `json:"organisation_identification_issuer,omitempty"`

	// organisation identifications
	OrganisationIdentifications []*RequestForInformationCreatorPartyOrganisationIdentificationsItems0 `json:"organisation_identifications,omitempty"`

	// private identification
	PrivateIdentification *RequestForInformationCreatorPartyPrivateIdentification `json:"private_identification,omitempty"`

	// telephone number
	TelephoneNumber string `json:"telephone_number,omitempty"`
}

func RequestForInformationCreatorPartyWithDefaults(defaults client.Defaults) *RequestForInformationCreatorParty {
	return &RequestForInformationCreatorParty{

		BirthCity: defaults.GetString("RequestForInformationCreatorParty", "birth_city"),

		BirthCountry: defaults.GetString("RequestForInformationCreatorParty", "birth_country"),

		BirthDate: defaults.GetString("RequestForInformationCreatorParty", "birth_date"),

		BirthProvince: defaults.GetString("RequestForInformationCreatorParty", "birth_province"),

		Name: defaults.GetString("RequestForInformationCreatorParty", "name"),

		OrganisationIdentification: defaults.GetString("RequestForInformationCreatorParty", "organisation_identification"),

		OrganisationIdentificationCode: defaults.GetString("RequestForInformationCreatorParty", "organisation_identification_code"),

		OrganisationIdentificationIssuer: defaults.GetString("RequestForInformationCreatorParty", "organisation_identification_issuer"),

		OrganisationIdentifications: make([]*RequestForInformationCreatorPartyOrganisationIdentificationsItems0, 0),

		PrivateIdentification: RequestForInformationCreatorPartyPrivateIdentificationWithDefaults(defaults),

		TelephoneNumber: defaults.GetString("RequestForInformationCreatorParty", "telephone_number"),
	}
}

func (m *RequestForInformationCreatorParty) WithBirthCity(birthCity string) *RequestForInformationCreatorParty {

	m.BirthCity = birthCity

	return m
}

func (m *RequestForInformationCreatorParty) WithBirthCountry(birthCountry string) *RequestForInformationCreatorParty {

	m.BirthCountry = birthCountry

	return m
}

func (m *RequestForInformationCreatorParty) WithBirthDate(birthDate string) *RequestForInformationCreatorParty {

	m.BirthDate = birthDate

	return m
}

func (m *RequestForInformationCreatorParty) WithBirthProvince(birthProvince string) *RequestForInformationCreatorParty {

	m.BirthProvince = birthProvince

	return m
}

func (m *RequestForInformationCreatorParty) WithName(name string) *RequestForInformationCreatorParty {

	m.Name = name

	return m
}

func (m *RequestForInformationCreatorParty) WithOrganisationIdentification(organisationIdentification string) *RequestForInformationCreatorParty {

	m.OrganisationIdentification = organisationIdentification

	return m
}

func (m *RequestForInformationCreatorParty) WithOrganisationIdentificationCode(organisationIdentificationCode string) *RequestForInformationCreatorParty {

	m.OrganisationIdentificationCode = organisationIdentificationCode

	return m
}

func (m *RequestForInformationCreatorParty) WithOrganisationIdentificationIssuer(organisationIdentificationIssuer string) *RequestForInformationCreatorParty {

	m.OrganisationIdentificationIssuer = organisationIdentificationIssuer

	return m
}

func (m *RequestForInformationCreatorParty) WithOrganisationIdentifications(organisationIdentifications []*RequestForInformationCreatorPartyOrganisationIdentificationsItems0) *RequestForInformationCreatorParty {

	m.OrganisationIdentifications = organisationIdentifications

	return m
}

func (m *RequestForInformationCreatorParty) WithPrivateIdentification(privateIdentification RequestForInformationCreatorPartyPrivateIdentification) *RequestForInformationCreatorParty {

	m.PrivateIdentification = &privateIdentification

	return m
}

func (m *RequestForInformationCreatorParty) WithoutPrivateIdentification() *RequestForInformationCreatorParty {
	m.PrivateIdentification = nil
	return m
}

func (m *RequestForInformationCreatorParty) WithTelephoneNumber(telephoneNumber string) *RequestForInformationCreatorParty {

	m.TelephoneNumber = telephoneNumber

	return m
}

// Validate validates this request for information creator party
func (m *RequestForInformationCreatorParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganisationIdentifications(formats); err != nil {
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

func (m *RequestForInformationCreatorParty) validateOrganisationIdentifications(formats strfmt.Registry) error {

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

func (m *RequestForInformationCreatorParty) validatePrivateIdentification(formats strfmt.Registry) error {

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
func (m *RequestForInformationCreatorParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestForInformationCreatorParty) UnmarshalBinary(b []byte) error {
	var res RequestForInformationCreatorParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RequestForInformationCreatorParty) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// RequestForInformationCreatorPartyOrganisationIdentificationsItems0 request for information creator party organisation identifications items0
// swagger:model RequestForInformationCreatorPartyOrganisationIdentificationsItems0
type RequestForInformationCreatorPartyOrganisationIdentificationsItems0 struct {

	// identification
	Identification string `json:"identification,omitempty"`

	// identification code
	IdentificationCode string `json:"identification_code,omitempty"`

	// identification issuer
	IdentificationIssuer string `json:"identification_issuer,omitempty"`
}

func RequestForInformationCreatorPartyOrganisationIdentificationsItems0WithDefaults(defaults client.Defaults) *RequestForInformationCreatorPartyOrganisationIdentificationsItems0 {
	return &RequestForInformationCreatorPartyOrganisationIdentificationsItems0{

		Identification: defaults.GetString("RequestForInformationCreatorPartyOrganisationIdentificationsItems0", "identification"),

		IdentificationCode: defaults.GetString("RequestForInformationCreatorPartyOrganisationIdentificationsItems0", "identification_code"),

		IdentificationIssuer: defaults.GetString("RequestForInformationCreatorPartyOrganisationIdentificationsItems0", "identification_issuer"),
	}
}

func (m *RequestForInformationCreatorPartyOrganisationIdentificationsItems0) WithIdentification(identification string) *RequestForInformationCreatorPartyOrganisationIdentificationsItems0 {

	m.Identification = identification

	return m
}

func (m *RequestForInformationCreatorPartyOrganisationIdentificationsItems0) WithIdentificationCode(identificationCode string) *RequestForInformationCreatorPartyOrganisationIdentificationsItems0 {

	m.IdentificationCode = identificationCode

	return m
}

func (m *RequestForInformationCreatorPartyOrganisationIdentificationsItems0) WithIdentificationIssuer(identificationIssuer string) *RequestForInformationCreatorPartyOrganisationIdentificationsItems0 {

	m.IdentificationIssuer = identificationIssuer

	return m
}

// Validate validates this request for information creator party organisation identifications items0
func (m *RequestForInformationCreatorPartyOrganisationIdentificationsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RequestForInformationCreatorPartyOrganisationIdentificationsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestForInformationCreatorPartyOrganisationIdentificationsItems0) UnmarshalBinary(b []byte) error {
	var res RequestForInformationCreatorPartyOrganisationIdentificationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RequestForInformationCreatorPartyOrganisationIdentificationsItems0) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// RequestForInformationCreatorPartyPrivateIdentification request for information creator party private identification
// swagger:model RequestForInformationCreatorPartyPrivateIdentification
type RequestForInformationCreatorPartyPrivateIdentification struct {

	// private identification code
	PrivateIdentificationCode string `json:"private_identification.code,omitempty"`

	// private identification identification
	PrivateIdentificationIdentification string `json:"private_identification.identification,omitempty"`

	// private identification issuer
	PrivateIdentificationIssuer string `json:"private_identification.issuer,omitempty"`
}

func RequestForInformationCreatorPartyPrivateIdentificationWithDefaults(defaults client.Defaults) *RequestForInformationCreatorPartyPrivateIdentification {
	return &RequestForInformationCreatorPartyPrivateIdentification{

		PrivateIdentificationCode: defaults.GetString("RequestForInformationCreatorPartyPrivateIdentification", "private_identification.code"),

		PrivateIdentificationIdentification: defaults.GetString("RequestForInformationCreatorPartyPrivateIdentification", "private_identification.identification"),

		PrivateIdentificationIssuer: defaults.GetString("RequestForInformationCreatorPartyPrivateIdentification", "private_identification.issuer"),
	}
}

func (m *RequestForInformationCreatorPartyPrivateIdentification) WithPrivateIdentificationCode(privateIdentificationCode string) *RequestForInformationCreatorPartyPrivateIdentification {

	m.PrivateIdentificationCode = privateIdentificationCode

	return m
}

func (m *RequestForInformationCreatorPartyPrivateIdentification) WithPrivateIdentificationIdentification(privateIdentificationIdentification string) *RequestForInformationCreatorPartyPrivateIdentification {

	m.PrivateIdentificationIdentification = privateIdentificationIdentification

	return m
}

func (m *RequestForInformationCreatorPartyPrivateIdentification) WithPrivateIdentificationIssuer(privateIdentificationIssuer string) *RequestForInformationCreatorPartyPrivateIdentification {

	m.PrivateIdentificationIssuer = privateIdentificationIssuer

	return m
}

// Validate validates this request for information creator party private identification
func (m *RequestForInformationCreatorPartyPrivateIdentification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RequestForInformationCreatorPartyPrivateIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestForInformationCreatorPartyPrivateIdentification) UnmarshalBinary(b []byte) error {
	var res RequestForInformationCreatorPartyPrivateIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RequestForInformationCreatorPartyPrivateIdentification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}