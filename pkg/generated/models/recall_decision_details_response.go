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

// RecallDecisionDetailsResponse recall decision details response
// swagger:model RecallDecisionDetailsResponse
type RecallDecisionDetailsResponse struct {

	// data
	Data *RecallDecision `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func RecallDecisionDetailsResponseWithDefaults(defaults client.Defaults) *RecallDecisionDetailsResponse {
	return &RecallDecisionDetailsResponse{

		Data: RecallDecisionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *RecallDecisionDetailsResponse) WithData(data RecallDecision) *RecallDecisionDetailsResponse {

	m.Data = &data

	return m
}

func (m *RecallDecisionDetailsResponse) WithoutData() *RecallDecisionDetailsResponse {
	m.Data = nil
	return m
}

func (m *RecallDecisionDetailsResponse) WithLinks(links Links) *RecallDecisionDetailsResponse {

	m.Links = &links

	return m
}

func (m *RecallDecisionDetailsResponse) WithoutLinks() *RecallDecisionDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this recall decision details response
func (m *RecallDecisionDetailsResponse) Validate(formats strfmt.Registry) error {
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

func (m *RecallDecisionDetailsResponse) validateData(formats strfmt.Registry) error {

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

func (m *RecallDecisionDetailsResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *RecallDecisionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res RecallDecisionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
