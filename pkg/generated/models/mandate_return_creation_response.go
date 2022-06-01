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
)

// MandateReturnCreationResponse mandate return creation response
// swagger:model MandateReturnCreationResponse
type MandateReturnCreationResponse struct {

	// data
	Data *MandateReturn `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func MandateReturnCreationResponseWithDefaults(defaults client.Defaults) *MandateReturnCreationResponse {
	return &MandateReturnCreationResponse{

		Data: MandateReturnWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *MandateReturnCreationResponse) WithData(data MandateReturn) *MandateReturnCreationResponse {

	m.Data = &data

	return m
}

func (m *MandateReturnCreationResponse) WithoutData() *MandateReturnCreationResponse {
	m.Data = nil
	return m
}

func (m *MandateReturnCreationResponse) WithLinks(links Links) *MandateReturnCreationResponse {

	m.Links = &links

	return m
}

func (m *MandateReturnCreationResponse) WithoutLinks() *MandateReturnCreationResponse {
	m.Links = nil
	return m
}

// Validate validates this mandate return creation response
func (m *MandateReturnCreationResponse) Validate(formats strfmt.Registry) error {
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

func (m *MandateReturnCreationResponse) validateData(formats strfmt.Registry) error {

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

func (m *MandateReturnCreationResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *MandateReturnCreationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateReturnCreationResponse) UnmarshalBinary(b []byte) error {
	var res MandateReturnCreationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateReturnCreationResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
