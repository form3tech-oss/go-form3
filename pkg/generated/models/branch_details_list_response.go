// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BranchDetailsListResponse branch details list response
// swagger:model BranchDetailsListResponse
type BranchDetailsListResponse struct {

	// data
	Data []*Branch `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func BranchDetailsListResponseWithDefaults(defaults client.Defaults) *BranchDetailsListResponse {
	return &BranchDetailsListResponse{

		Data: make([]*Branch, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *BranchDetailsListResponse) WithData(data []*Branch) *BranchDetailsListResponse {

	m.Data = data

	return m
}

func (m *BranchDetailsListResponse) WithLinks(links Links) *BranchDetailsListResponse {

	m.Links = &links

	return m
}

func (m *BranchDetailsListResponse) WithoutLinks() *BranchDetailsListResponse {
	m.Links = nil
	return m
}

// Validate validates this branch details list response
func (m *BranchDetailsListResponse) Validate(formats strfmt.Registry) error {
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

func (m *BranchDetailsListResponse) validateData(formats strfmt.Registry) error {

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

func (m *BranchDetailsListResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *BranchDetailsListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchDetailsListResponse) UnmarshalBinary(b []byte) error {
	var res BranchDetailsListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *BranchDetailsListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
