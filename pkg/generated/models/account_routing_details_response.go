// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccountRoutingDetailsResponse account routing details response
// swagger:model AccountRoutingDetailsResponse
type AccountRoutingDetailsResponse struct {

	// data
	Data *AccountRouting `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

// line 140

func AccountRoutingDetailsResponseWithDefaults(defaults client.Defaults) *AccountRoutingDetailsResponse {
	return &AccountRoutingDetailsResponse{

		Data: AccountRoutingWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *AccountRoutingDetailsResponse) WithData(data AccountRouting) *AccountRoutingDetailsResponse {

	m.Data = &data

	return m
}

func (m *AccountRoutingDetailsResponse) WithoutData() *AccountRoutingDetailsResponse {
	m.Data = nil
	return m
}

func (m *AccountRoutingDetailsResponse) WithLinks(links Links) *AccountRoutingDetailsResponse {

	m.Links = &links

	return m
}

func (m *AccountRoutingDetailsResponse) WithoutLinks() *AccountRoutingDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this account routing details response
func (m *AccountRoutingDetailsResponse) Validate(formats strfmt.Registry) error {
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

func (m *AccountRoutingDetailsResponse) validateData(formats strfmt.Registry) error {

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

func (m *AccountRoutingDetailsResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *AccountRoutingDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRoutingDetailsResponse) UnmarshalBinary(b []byte) error {
	var res AccountRoutingDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountRoutingDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
