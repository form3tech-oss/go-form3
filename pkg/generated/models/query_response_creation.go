// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QueryResponseCreation query response creation
// swagger:model QueryResponseCreation
type QueryResponseCreation struct {

	// data
	// Required: true
	Data *NewQueryResponse `json:"data"`
}

func QueryResponseCreationWithDefaults(defaults client.Defaults) *QueryResponseCreation {
	return &QueryResponseCreation{

		Data: NewQueryResponseWithDefaults(defaults),
	}
}

func (m *QueryResponseCreation) WithData(data NewQueryResponse) *QueryResponseCreation {

	m.Data = &data

	return m
}

func (m *QueryResponseCreation) WithoutData() *QueryResponseCreation {
	m.Data = nil
	return m
}

// Validate validates this query response creation
func (m *QueryResponseCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponseCreation) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResponseCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseCreation) UnmarshalBinary(b []byte) error {
	var res QueryResponseCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseCreation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
