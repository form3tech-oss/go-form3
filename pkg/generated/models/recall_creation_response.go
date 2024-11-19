// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecallCreationResponse recall creation response
// swagger:model RecallCreationResponse
type RecallCreationResponse struct {

	// data
	Data *Recall `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func RecallCreationResponseWithDefaults(defaults client.Defaults) *RecallCreationResponse {
	return &RecallCreationResponse{

		Data: RecallWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *RecallCreationResponse) WithData(data Recall) *RecallCreationResponse {

	m.Data = &data

	return m
}

func (m *RecallCreationResponse) WithoutData() *RecallCreationResponse {
	m.Data = nil
	return m
}

func (m *RecallCreationResponse) WithLinks(links Links) *RecallCreationResponse {

	m.Links = &links

	return m
}

func (m *RecallCreationResponse) WithoutLinks() *RecallCreationResponse {
	m.Links = nil
	return m
}

// Validate validates this recall creation response
func (m *RecallCreationResponse) Validate(formats strfmt.Registry) error {
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

func (m *RecallCreationResponse) validateData(formats strfmt.Registry) error {

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

func (m *RecallCreationResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *RecallCreationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallCreationResponse) UnmarshalBinary(b []byte) error {
	var res RecallCreationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallCreationResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
