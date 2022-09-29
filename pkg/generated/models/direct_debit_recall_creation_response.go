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
)

// DirectDebitRecallCreationResponse direct debit recall creation response
// swagger:model DirectDebitRecallCreationResponse
type DirectDebitRecallCreationResponse struct {

	// data
	Data *DirectDebitRecall `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func DirectDebitRecallCreationResponseWithDefaults(defaults client.Defaults) *DirectDebitRecallCreationResponse {
	return &DirectDebitRecallCreationResponse{

		Data: DirectDebitRecallWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *DirectDebitRecallCreationResponse) WithData(data DirectDebitRecall) *DirectDebitRecallCreationResponse {

	m.Data = &data

	return m
}

func (m *DirectDebitRecallCreationResponse) WithoutData() *DirectDebitRecallCreationResponse {
	m.Data = nil
	return m
}

func (m *DirectDebitRecallCreationResponse) WithLinks(links Links) *DirectDebitRecallCreationResponse {

	m.Links = &links

	return m
}

func (m *DirectDebitRecallCreationResponse) WithoutLinks() *DirectDebitRecallCreationResponse {
	m.Links = nil
	return m
}

// Validate validates this direct debit recall creation response
func (m *DirectDebitRecallCreationResponse) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitRecallCreationResponse) validateData(formats strfmt.Registry) error {

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

func (m *DirectDebitRecallCreationResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *DirectDebitRecallCreationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecallCreationResponse) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecallCreationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecallCreationResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
