// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OrganisationAttributes organisation attributes
// swagger:model OrganisationAttributes
type OrganisationAttributes struct {

	// Name of the organisation
	Name string `json:"name,omitempty"`
}

func OrganisationAttributesWithDefaults(defaults client.Defaults) *OrganisationAttributes {
	return &OrganisationAttributes{

		Name: defaults.GetString("OrganisationAttributes", "name"),
	}
}

func (m *OrganisationAttributes) WithName(name string) *OrganisationAttributes {

	m.Name = name

	return m
}

// Validate validates this organisation attributes
func (m *OrganisationAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganisationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganisationAttributes) UnmarshalBinary(b []byte) error {
	var res OrganisationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *OrganisationAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
