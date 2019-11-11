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

// AccountRoutingDetailsListResponse account routing details list response
// swagger:model AccountRoutingDetailsListResponse
type AccountRoutingDetailsListResponse struct {

	// data
	Data []*AccountRouting `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

// line 140

func AccountRoutingDetailsListResponseWithDefaults(defaults client.Defaults) *AccountRoutingDetailsListResponse {
	return &AccountRoutingDetailsListResponse{

		Data: make([]*AccountRouting, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *AccountRoutingDetailsListResponse) WithData(data []*AccountRouting) *AccountRoutingDetailsListResponse {

	m.Data = data

	return m
}

func (m *AccountRoutingDetailsListResponse) WithLinks(links Links) *AccountRoutingDetailsListResponse {

	m.Links = &links

	return m
}

func (m *AccountRoutingDetailsListResponse) WithoutLinks() *AccountRoutingDetailsListResponse {
	m.Links = nil
	return m
}

// Validate validates this account routing details list response
func (m *AccountRoutingDetailsListResponse) Validate(formats strfmt.Registry) error {
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

func (m *AccountRoutingDetailsListResponse) validateData(formats strfmt.Registry) error {

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

func (m *AccountRoutingDetailsListResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *AccountRoutingDetailsListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRoutingDetailsListResponse) UnmarshalBinary(b []byte) error {
	var res AccountRoutingDetailsListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountRoutingDetailsListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
