// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RoleDetailsListResponse role details list response
// swagger:model RoleDetailsListResponse
type RoleDetailsListResponse struct {

	// data
	// Required: true
	Data []*Role `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func RoleDetailsListResponseWithDefaults(defaults client.Defaults) *RoleDetailsListResponse {
	return &RoleDetailsListResponse{

		Data: make([]*Role, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *RoleDetailsListResponse) WithData(data []*Role) *RoleDetailsListResponse {

	m.Data = data

	return m
}

func (m *RoleDetailsListResponse) WithLinks(links Links) *RoleDetailsListResponse {

	m.Links = &links

	return m
}

func (m *RoleDetailsListResponse) WithoutLinks() *RoleDetailsListResponse {
	m.Links = nil
	return m
}

// Validate validates this role details list response
func (m *RoleDetailsListResponse) Validate(formats strfmt.Registry) error {
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

func (m *RoleDetailsListResponse) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RoleDetailsListResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *RoleDetailsListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleDetailsListResponse) UnmarshalBinary(b []byte) error {
	var res RoleDetailsListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RoleDetailsListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
