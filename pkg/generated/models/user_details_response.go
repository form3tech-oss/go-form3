// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserDetailsResponse user details response
// swagger:model UserDetailsResponse
type UserDetailsResponse struct {

	// data
	// Required: true
	Data *User `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func UserDetailsResponseWithDefaults(defaults client.Defaults) *UserDetailsResponse {
	return &UserDetailsResponse{

		Data: UserWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *UserDetailsResponse) WithData(data User) *UserDetailsResponse {

	m.Data = &data

	return m
}

func (m *UserDetailsResponse) WithoutData() *UserDetailsResponse {
	m.Data = nil
	return m
}

func (m *UserDetailsResponse) WithLinks(links Links) *UserDetailsResponse {

	m.Links = &links

	return m
}

func (m *UserDetailsResponse) WithoutLinks() *UserDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this user details response
func (m *UserDetailsResponse) Validate(formats strfmt.Registry) error {
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

func (m *UserDetailsResponse) validateData(formats strfmt.Registry) error {

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

func (m *UserDetailsResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *UserDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserDetailsResponse) UnmarshalBinary(b []byte) error {
	var res UserDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *UserDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
