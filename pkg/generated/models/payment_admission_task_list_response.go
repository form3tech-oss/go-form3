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

// PaymentAdmissionTaskListResponse payment admission task list response
// swagger:model PaymentAdmissionTaskListResponse
type PaymentAdmissionTaskListResponse struct {

	// data
	Data []*PaymentAdmissionTask `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func PaymentAdmissionTaskListResponseWithDefaults(defaults client.Defaults) *PaymentAdmissionTaskListResponse {
	return &PaymentAdmissionTaskListResponse{

		Data: make([]*PaymentAdmissionTask, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *PaymentAdmissionTaskListResponse) WithData(data []*PaymentAdmissionTask) *PaymentAdmissionTaskListResponse {

	m.Data = data

	return m
}

func (m *PaymentAdmissionTaskListResponse) WithLinks(links Links) *PaymentAdmissionTaskListResponse {

	m.Links = &links

	return m
}

func (m *PaymentAdmissionTaskListResponse) WithoutLinks() *PaymentAdmissionTaskListResponse {
	m.Links = nil
	return m
}

// Validate validates this payment admission task list response
func (m *PaymentAdmissionTaskListResponse) Validate(formats strfmt.Registry) error {
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

func (m *PaymentAdmissionTaskListResponse) validateData(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionTaskListResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *PaymentAdmissionTaskListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionTaskListResponse) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionTaskListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionTaskListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
