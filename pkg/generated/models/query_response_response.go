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
	"github.com/go-openapi/validate"
)

// QueryResponseResponse query response response
// swagger:model QueryResponseResponse
type QueryResponseResponse struct {

	// data
	// Required: true
	Data *QueryResponse `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func QueryResponseResponseWithDefaults(defaults client.Defaults) *QueryResponseResponse {
	return &QueryResponseResponse{

		Data: QueryResponseWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *QueryResponseResponse) WithData(data QueryResponse) *QueryResponseResponse {

	m.Data = &data

	return m
}

func (m *QueryResponseResponse) WithoutData() *QueryResponseResponse {
	m.Data = nil
	return m
}

func (m *QueryResponseResponse) WithLinks(links Links) *QueryResponseResponse {

	m.Links = &links

	return m
}

func (m *QueryResponseResponse) WithoutLinks() *QueryResponseResponse {
	m.Links = nil
	return m
}

// Validate validates this query response response
func (m *QueryResponseResponse) Validate(formats strfmt.Registry) error {
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

func (m *QueryResponseResponse) validateData(formats strfmt.Registry) error {

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

func (m *QueryResponseResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *QueryResponseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseResponse) UnmarshalBinary(b []byte) error {
	var res QueryResponseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
