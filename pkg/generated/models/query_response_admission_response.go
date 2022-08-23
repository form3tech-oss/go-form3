// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QueryResponseAdmissionResponse query response admission response
// swagger:model QueryResponseAdmissionResponse
type QueryResponseAdmissionResponse struct {

	// data
	// Required: true
	Data *QueryResponseAdmission `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func QueryResponseAdmissionResponseWithDefaults(defaults client.Defaults) *QueryResponseAdmissionResponse {
	return &QueryResponseAdmissionResponse{

		Data: QueryResponseAdmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *QueryResponseAdmissionResponse) WithData(data QueryResponseAdmission) *QueryResponseAdmissionResponse {

	m.Data = &data

	return m
}

func (m *QueryResponseAdmissionResponse) WithoutData() *QueryResponseAdmissionResponse {
	m.Data = nil
	return m
}

func (m *QueryResponseAdmissionResponse) WithLinks(links Links) *QueryResponseAdmissionResponse {

	m.Links = &links

	return m
}

func (m *QueryResponseAdmissionResponse) WithoutLinks() *QueryResponseAdmissionResponse {
	m.Links = nil
	return m
}

// Validate validates this query response admission response
func (m *QueryResponseAdmissionResponse) Validate(formats strfmt.Registry) error {
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

func (m *QueryResponseAdmissionResponse) validateData(formats strfmt.Registry) error {

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

func (m *QueryResponseAdmissionResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *QueryResponseAdmissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseAdmissionResponse) UnmarshalBinary(b []byte) error {
	var res QueryResponseAdmissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseAdmissionResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
