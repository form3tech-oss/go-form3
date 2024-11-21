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

// PaymentSubmissionFetchResponse payment submission fetch response
// swagger:model PaymentSubmissionFetchResponse
type PaymentSubmissionFetchResponse struct {

	// data
	Data *PaymentSubmissionFetch `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func PaymentSubmissionFetchResponseWithDefaults(defaults client.Defaults) *PaymentSubmissionFetchResponse {
	return &PaymentSubmissionFetchResponse{

		Data: PaymentSubmissionFetchWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *PaymentSubmissionFetchResponse) WithData(data PaymentSubmissionFetch) *PaymentSubmissionFetchResponse {

	m.Data = &data

	return m
}

func (m *PaymentSubmissionFetchResponse) WithoutData() *PaymentSubmissionFetchResponse {
	m.Data = nil
	return m
}

func (m *PaymentSubmissionFetchResponse) WithLinks(links Links) *PaymentSubmissionFetchResponse {

	m.Links = &links

	return m
}

func (m *PaymentSubmissionFetchResponse) WithoutLinks() *PaymentSubmissionFetchResponse {
	m.Links = nil
	return m
}

// Validate validates this payment submission fetch response
func (m *PaymentSubmissionFetchResponse) Validate(formats strfmt.Registry) error {
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

func (m *PaymentSubmissionFetchResponse) validateData(formats strfmt.Registry) error {

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

func (m *PaymentSubmissionFetchResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *PaymentSubmissionFetchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentSubmissionFetchResponse) UnmarshalBinary(b []byte) error {
	var res PaymentSubmissionFetchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentSubmissionFetchResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}