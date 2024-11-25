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

// MandateReturnSubmissionDetailsResponse mandate return submission details response
// swagger:model MandateReturnSubmissionDetailsResponse
type MandateReturnSubmissionDetailsResponse struct {

	// data
	Data *MandateReturnSubmission `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func MandateReturnSubmissionDetailsResponseWithDefaults(defaults client.Defaults) *MandateReturnSubmissionDetailsResponse {
	return &MandateReturnSubmissionDetailsResponse{

		Data: MandateReturnSubmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *MandateReturnSubmissionDetailsResponse) WithData(data MandateReturnSubmission) *MandateReturnSubmissionDetailsResponse {

	m.Data = &data

	return m
}

func (m *MandateReturnSubmissionDetailsResponse) WithoutData() *MandateReturnSubmissionDetailsResponse {
	m.Data = nil
	return m
}

func (m *MandateReturnSubmissionDetailsResponse) WithLinks(links Links) *MandateReturnSubmissionDetailsResponse {

	m.Links = &links

	return m
}

func (m *MandateReturnSubmissionDetailsResponse) WithoutLinks() *MandateReturnSubmissionDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this mandate return submission details response
func (m *MandateReturnSubmissionDetailsResponse) Validate(formats strfmt.Registry) error {
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

func (m *MandateReturnSubmissionDetailsResponse) validateData(formats strfmt.Registry) error {

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

func (m *MandateReturnSubmissionDetailsResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *MandateReturnSubmissionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateReturnSubmissionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res MandateReturnSubmissionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateReturnSubmissionDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
