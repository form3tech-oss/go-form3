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
	"github.com/go-openapi/validate"
)

// QueryResponseAdmissionAttributes query response admission attributes
// swagger:model QueryResponseAdmissionAttributes
type QueryResponseAdmissionAttributes struct {

	// status
	// Required: true
	Status *QueryResponseAdmissionStatus `json:"status"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`
}

func QueryResponseAdmissionAttributesWithDefaults(defaults client.Defaults) *QueryResponseAdmissionAttributes {
	return &QueryResponseAdmissionAttributes{

		// TODO Status: QueryResponseAdmissionStatus,

		StatusReason: defaults.GetString("QueryResponseAdmissionAttributes", "status_reason"),
	}
}

func (m *QueryResponseAdmissionAttributes) WithStatus(status QueryResponseAdmissionStatus) *QueryResponseAdmissionAttributes {

	m.Status = &status

	return m
}

func (m *QueryResponseAdmissionAttributes) WithoutStatus() *QueryResponseAdmissionAttributes {
	m.Status = nil
	return m
}

func (m *QueryResponseAdmissionAttributes) WithStatusReason(statusReason string) *QueryResponseAdmissionAttributes {

	m.StatusReason = statusReason

	return m
}

// Validate validates this query response admission attributes
func (m *QueryResponseAdmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponseAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResponseAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res QueryResponseAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
