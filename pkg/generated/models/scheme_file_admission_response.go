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
)

// SchemeFileAdmissionResponse scheme file admission response
// swagger:model SchemeFileAdmissionResponse
type SchemeFileAdmissionResponse struct {

	// data
	Data *SchemeFileAdmission `json:"data,omitempty"`

	// links
	Links *SchemeFileLinks `json:"links,omitempty"`
}

func SchemeFileAdmissionResponseWithDefaults(defaults client.Defaults) *SchemeFileAdmissionResponse {
	return &SchemeFileAdmissionResponse{

		Data: SchemeFileAdmissionWithDefaults(defaults),

		// TODO Links: SchemeFileLinks,

	}
}

func (m *SchemeFileAdmissionResponse) WithData(data SchemeFileAdmission) *SchemeFileAdmissionResponse {

	m.Data = &data

	return m
}

func (m *SchemeFileAdmissionResponse) WithoutData() *SchemeFileAdmissionResponse {
	m.Data = nil
	return m
}

func (m *SchemeFileAdmissionResponse) WithLinks(links SchemeFileLinks) *SchemeFileAdmissionResponse {

	m.Links = &links

	return m
}

func (m *SchemeFileAdmissionResponse) WithoutLinks() *SchemeFileAdmissionResponse {
	m.Links = nil
	return m
}

// Validate validates this scheme file admission response
func (m *SchemeFileAdmissionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeFileAdmissionResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
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

func (m *SchemeFileAdmissionResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeFileAdmissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeFileAdmissionResponse) UnmarshalBinary(b []byte) error {
	var res SchemeFileAdmissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeFileAdmissionResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
