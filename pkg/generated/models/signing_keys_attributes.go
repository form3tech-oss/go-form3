// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SigningKeysAttributes signing keys attributes
// swagger:model SigningKeysAttributes
type SigningKeysAttributes struct {

	// public key
	// Required: true
	PublicKey *string `json:"public_key"`
}

func SigningKeysAttributesWithDefaults(defaults client.Defaults) *SigningKeysAttributes {
	return &SigningKeysAttributes{

		PublicKey: defaults.GetStringPtr("SigningKeysAttributes", "public_key"),
	}
}

func (m *SigningKeysAttributes) WithPublicKey(publicKey string) *SigningKeysAttributes {

	m.PublicKey = &publicKey

	return m
}

func (m *SigningKeysAttributes) WithoutPublicKey() *SigningKeysAttributes {
	m.PublicKey = nil
	return m
}

// Validate validates this signing keys attributes
func (m *SigningKeysAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SigningKeysAttributes) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("public_key", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SigningKeysAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SigningKeysAttributes) UnmarshalBinary(b []byte) error {
	var res SigningKeysAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SigningKeysAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
