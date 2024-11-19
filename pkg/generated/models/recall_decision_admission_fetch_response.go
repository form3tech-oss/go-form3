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
)

// RecallDecisionAdmissionFetchResponse recall decision admission fetch response
// swagger:model RecallDecisionAdmissionFetchResponse
type RecallDecisionAdmissionFetchResponse struct {

	// data
	Data *RecallDecisionAdmissionFetch `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func RecallDecisionAdmissionFetchResponseWithDefaults(defaults client.Defaults) *RecallDecisionAdmissionFetchResponse {
	return &RecallDecisionAdmissionFetchResponse{

		Data: RecallDecisionAdmissionFetchWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *RecallDecisionAdmissionFetchResponse) WithData(data RecallDecisionAdmissionFetch) *RecallDecisionAdmissionFetchResponse {

	m.Data = &data

	return m
}

func (m *RecallDecisionAdmissionFetchResponse) WithoutData() *RecallDecisionAdmissionFetchResponse {
	m.Data = nil
	return m
}

func (m *RecallDecisionAdmissionFetchResponse) WithLinks(links Links) *RecallDecisionAdmissionFetchResponse {

	m.Links = &links

	return m
}

func (m *RecallDecisionAdmissionFetchResponse) WithoutLinks() *RecallDecisionAdmissionFetchResponse {
	m.Links = nil
	return m
}

// Validate validates this recall decision admission fetch response
func (m *RecallDecisionAdmissionFetchResponse) Validate(formats strfmt.Registry) error {
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

func (m *RecallDecisionAdmissionFetchResponse) validateData(formats strfmt.Registry) error {

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

func (m *RecallDecisionAdmissionFetchResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *RecallDecisionAdmissionFetchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionAdmissionFetchResponse) UnmarshalBinary(b []byte) error {
	var res RecallDecisionAdmissionFetchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionAdmissionFetchResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
