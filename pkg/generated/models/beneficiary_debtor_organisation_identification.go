// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BeneficiaryDebtorOrganisationIdentification beneficiary debtor organisation identification
// swagger:model BeneficiaryDebtorOrganisationIdentification
type BeneficiaryDebtorOrganisationIdentification struct {

	// Another ID of the organisation
	Identification string `json:"identification,omitempty"`

	// Code for the type of other ID of organisation
	IdentificationCode string `json:"identification_code,omitempty"`

	// Issuer of the other ID of organisation
	IdentificationIssuer string `json:"identification_issuer,omitempty"`

	// Custom internal code for the type of other ID of organisation
	IdentificationScheme string `json:"identification_scheme,omitempty"`
}

func BeneficiaryDebtorOrganisationIdentificationWithDefaults(defaults client.Defaults) *BeneficiaryDebtorOrganisationIdentification {
	return &BeneficiaryDebtorOrganisationIdentification{

		Identification: defaults.GetString("BeneficiaryDebtorOrganisationIdentification", "identification"),

		IdentificationCode: defaults.GetString("BeneficiaryDebtorOrganisationIdentification", "identification_code"),

		IdentificationIssuer: defaults.GetString("BeneficiaryDebtorOrganisationIdentification", "identification_issuer"),

		IdentificationScheme: defaults.GetString("BeneficiaryDebtorOrganisationIdentification", "identification_scheme"),
	}
}

func (m *BeneficiaryDebtorOrganisationIdentification) WithIdentification(identification string) *BeneficiaryDebtorOrganisationIdentification {

	m.Identification = identification

	return m
}

func (m *BeneficiaryDebtorOrganisationIdentification) WithIdentificationCode(identificationCode string) *BeneficiaryDebtorOrganisationIdentification {

	m.IdentificationCode = identificationCode

	return m
}

func (m *BeneficiaryDebtorOrganisationIdentification) WithIdentificationIssuer(identificationIssuer string) *BeneficiaryDebtorOrganisationIdentification {

	m.IdentificationIssuer = identificationIssuer

	return m
}

func (m *BeneficiaryDebtorOrganisationIdentification) WithIdentificationScheme(identificationScheme string) *BeneficiaryDebtorOrganisationIdentification {

	m.IdentificationScheme = identificationScheme

	return m
}

// Validate validates this beneficiary debtor organisation identification
func (m *BeneficiaryDebtorOrganisationIdentification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BeneficiaryDebtorOrganisationIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BeneficiaryDebtorOrganisationIdentification) UnmarshalBinary(b []byte) error {
	var res BeneficiaryDebtorOrganisationIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *BeneficiaryDebtorOrganisationIdentification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
