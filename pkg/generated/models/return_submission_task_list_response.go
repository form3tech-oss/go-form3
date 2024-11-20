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

// ReturnSubmissionTaskListResponse return submission task list response
// swagger:model ReturnSubmissionTaskListResponse
type ReturnSubmissionTaskListResponse struct {

	// data
	Data []*ReturnSubmissionTask `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func ReturnSubmissionTaskListResponseWithDefaults(defaults client.Defaults) *ReturnSubmissionTaskListResponse {
	return &ReturnSubmissionTaskListResponse{

		Data: make([]*ReturnSubmissionTask, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *ReturnSubmissionTaskListResponse) WithData(data []*ReturnSubmissionTask) *ReturnSubmissionTaskListResponse {

	m.Data = data

	return m
}

func (m *ReturnSubmissionTaskListResponse) WithLinks(links Links) *ReturnSubmissionTaskListResponse {

	m.Links = &links

	return m
}

func (m *ReturnSubmissionTaskListResponse) WithoutLinks() *ReturnSubmissionTaskListResponse {
	m.Links = nil
	return m
}

// Validate validates this return submission task list response
func (m *ReturnSubmissionTaskListResponse) Validate(formats strfmt.Registry) error {
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

func (m *ReturnSubmissionTaskListResponse) validateData(formats strfmt.Registry) error {

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

func (m *ReturnSubmissionTaskListResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *ReturnSubmissionTaskListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionTaskListResponse) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionTaskListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnSubmissionTaskListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
