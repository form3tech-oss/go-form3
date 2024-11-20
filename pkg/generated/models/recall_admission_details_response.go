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

// RecallAdmissionDetailsResponse recall admission details response
// swagger:model RecallAdmissionDetailsResponse
type RecallAdmissionDetailsResponse struct {

	// data
	Data *RecallAdmission `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func RecallAdmissionDetailsResponseWithDefaults(defaults client.Defaults) *RecallAdmissionDetailsResponse {
	return &RecallAdmissionDetailsResponse{

		Data: RecallAdmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *RecallAdmissionDetailsResponse) WithData(data RecallAdmission) *RecallAdmissionDetailsResponse {

	m.Data = &data

	return m
}

func (m *RecallAdmissionDetailsResponse) WithoutData() *RecallAdmissionDetailsResponse {
	m.Data = nil
	return m
}

func (m *RecallAdmissionDetailsResponse) WithLinks(links Links) *RecallAdmissionDetailsResponse {

	m.Links = &links

	return m
}

func (m *RecallAdmissionDetailsResponse) WithoutLinks() *RecallAdmissionDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this recall admission details response
func (m *RecallAdmissionDetailsResponse) Validate(formats strfmt.Registry) error {
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

func (m *RecallAdmissionDetailsResponse) validateData(formats strfmt.Registry) error {

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

func (m *RecallAdmissionDetailsResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *RecallAdmissionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallAdmissionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res RecallAdmissionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallAdmissionDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
