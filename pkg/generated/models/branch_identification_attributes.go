// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BranchIdentificationAttributes branch identification attributes
// swagger:model BranchIdentificationAttributes
type BranchIdentificationAttributes struct {

	// Used to validate where expecting an exact match against payment's reference information.
	// Required: true
	// Max Length: 35
	// Min Length: 1
	SecondaryIdentification *string `json:"secondary_identification"`
}

func BranchIdentificationAttributesWithDefaults(defaults client.Defaults) *BranchIdentificationAttributes {
	return &BranchIdentificationAttributes{

		SecondaryIdentification: defaults.GetStringPtr("BranchIdentificationAttributes", "secondary_identification"),
	}
}

func (m *BranchIdentificationAttributes) WithSecondaryIdentification(secondaryIdentification string) *BranchIdentificationAttributes {

	m.SecondaryIdentification = &secondaryIdentification

	return m
}

func (m *BranchIdentificationAttributes) WithoutSecondaryIdentification() *BranchIdentificationAttributes {
	m.SecondaryIdentification = nil
	return m
}

// Validate validates this branch identification attributes
func (m *BranchIdentificationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecondaryIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchIdentificationAttributes) validateSecondaryIdentification(formats strfmt.Registry) error {

	if err := validate.Required("secondary_identification", "body", m.SecondaryIdentification); err != nil {
		return err
	}

	if err := validate.MinLength("secondary_identification", "body", string(*m.SecondaryIdentification), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("secondary_identification", "body", string(*m.SecondaryIdentification), 35); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchIdentificationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchIdentificationAttributes) UnmarshalBinary(b []byte) error {
	var res BranchIdentificationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *BranchIdentificationAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
