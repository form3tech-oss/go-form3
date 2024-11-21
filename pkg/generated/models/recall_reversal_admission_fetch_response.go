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

// RecallReversalAdmissionFetchResponse recall reversal admission fetch response
// swagger:model RecallReversalAdmissionFetchResponse
type RecallReversalAdmissionFetchResponse struct {

	// data
	Data *RecallReversalAdmissionFetch `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func RecallReversalAdmissionFetchResponseWithDefaults(defaults client.Defaults) *RecallReversalAdmissionFetchResponse {
	return &RecallReversalAdmissionFetchResponse{

		Data: RecallReversalAdmissionFetchWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *RecallReversalAdmissionFetchResponse) WithData(data RecallReversalAdmissionFetch) *RecallReversalAdmissionFetchResponse {

	m.Data = &data

	return m
}

func (m *RecallReversalAdmissionFetchResponse) WithoutData() *RecallReversalAdmissionFetchResponse {
	m.Data = nil
	return m
}

func (m *RecallReversalAdmissionFetchResponse) WithLinks(links Links) *RecallReversalAdmissionFetchResponse {

	m.Links = &links

	return m
}

func (m *RecallReversalAdmissionFetchResponse) WithoutLinks() *RecallReversalAdmissionFetchResponse {
	m.Links = nil
	return m
}

// Validate validates this recall reversal admission fetch response
func (m *RecallReversalAdmissionFetchResponse) Validate(formats strfmt.Registry) error {
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

func (m *RecallReversalAdmissionFetchResponse) validateData(formats strfmt.Registry) error {

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

func (m *RecallReversalAdmissionFetchResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *RecallReversalAdmissionFetchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallReversalAdmissionFetchResponse) UnmarshalBinary(b []byte) error {
	var res RecallReversalAdmissionFetchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallReversalAdmissionFetchResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}