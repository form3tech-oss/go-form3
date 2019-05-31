// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserRoleListResponse user role list response
// swagger:model UserRoleListResponse
type UserRoleListResponse struct {

	// data
	Data []*Role `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func UserRoleListResponseWithDefaults(defaults client.Defaults) *UserRoleListResponse {
	return &UserRoleListResponse{

		Data: make([]*Role, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *UserRoleListResponse) WithData(data []*Role) *UserRoleListResponse {

	m.Data = data

	return m
}

func (m *UserRoleListResponse) WithLinks(links Links) *UserRoleListResponse {

	m.Links = &links

	return m
}

func (m *UserRoleListResponse) WithoutLinks() *UserRoleListResponse {
	m.Links = nil
	return m
}

// Validate validates this user role list response
func (m *UserRoleListResponse) Validate(formats strfmt.Registry) error {
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

func (m *UserRoleListResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
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

func (m *UserRoleListResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *UserRoleListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRoleListResponse) UnmarshalBinary(b []byte) error {
	var res UserRoleListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *UserRoleListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
