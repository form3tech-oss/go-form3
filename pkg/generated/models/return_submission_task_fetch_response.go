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

// ReturnSubmissionTaskFetchResponse return submission task fetch response
// swagger:model ReturnSubmissionTaskFetchResponse
type ReturnSubmissionTaskFetchResponse struct {

	// data
	Data *ReturnSubmissionTaskFetch `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func ReturnSubmissionTaskFetchResponseWithDefaults(defaults client.Defaults) *ReturnSubmissionTaskFetchResponse {
	return &ReturnSubmissionTaskFetchResponse{

		Data: ReturnSubmissionTaskFetchWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *ReturnSubmissionTaskFetchResponse) WithData(data ReturnSubmissionTaskFetch) *ReturnSubmissionTaskFetchResponse {

	m.Data = &data

	return m
}

func (m *ReturnSubmissionTaskFetchResponse) WithoutData() *ReturnSubmissionTaskFetchResponse {
	m.Data = nil
	return m
}

func (m *ReturnSubmissionTaskFetchResponse) WithLinks(links Links) *ReturnSubmissionTaskFetchResponse {

	m.Links = &links

	return m
}

func (m *ReturnSubmissionTaskFetchResponse) WithoutLinks() *ReturnSubmissionTaskFetchResponse {
	m.Links = nil
	return m
}

// Validate validates this return submission task fetch response
func (m *ReturnSubmissionTaskFetchResponse) Validate(formats strfmt.Registry) error {
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

func (m *ReturnSubmissionTaskFetchResponse) validateData(formats strfmt.Registry) error {

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

func (m *ReturnSubmissionTaskFetchResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *ReturnSubmissionTaskFetchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionTaskFetchResponse) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionTaskFetchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnSubmissionTaskFetchResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
