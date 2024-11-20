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

// PrivateIdentification private identification
// swagger:model PrivateIdentification
type PrivateIdentification struct {

	// Private Identification of an debtor/beneficiary or ultimate debtor/beneficiary
	Identification string `json:"identification,omitempty"`

	// Issuer of the `identification`
	IdentificationIssuer string `json:"identification_issuer,omitempty"`

	// Scheme of the `identification`
	IdentificationScheme string `json:"identification_scheme,omitempty"`

	// The code that specifies the type of `identification`
	IdentificationSchemeCode string `json:"identification_scheme_code,omitempty"`
}

func PrivateIdentificationWithDefaults(defaults client.Defaults) *PrivateIdentification {
	return &PrivateIdentification{

		Identification: defaults.GetString("PrivateIdentification", "identification"),

		IdentificationIssuer: defaults.GetString("PrivateIdentification", "identification_issuer"),

		IdentificationScheme: defaults.GetString("PrivateIdentification", "identification_scheme"),

		IdentificationSchemeCode: defaults.GetString("PrivateIdentification", "identification_scheme_code"),
	}
}

func (m *PrivateIdentification) WithIdentification(identification string) *PrivateIdentification {

	m.Identification = identification

	return m
}

func (m *PrivateIdentification) WithIdentificationIssuer(identificationIssuer string) *PrivateIdentification {

	m.IdentificationIssuer = identificationIssuer

	return m
}

func (m *PrivateIdentification) WithIdentificationScheme(identificationScheme string) *PrivateIdentification {

	m.IdentificationScheme = identificationScheme

	return m
}

func (m *PrivateIdentification) WithIdentificationSchemeCode(identificationSchemeCode string) *PrivateIdentification {

	m.IdentificationSchemeCode = identificationSchemeCode

	return m
}

// Validate validates this private identification
func (m *PrivateIdentification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrivateIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateIdentification) UnmarshalBinary(b []byte) error {
	var res PrivateIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PrivateIdentification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
