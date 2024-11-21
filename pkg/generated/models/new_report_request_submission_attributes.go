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

// NewReportRequestSubmissionAttributes new report request submission attributes
// swagger:model NewReportRequestSubmissionAttributes
type NewReportRequestSubmissionAttributes struct {

	// status
	// Enum: ["pending"]
	Status string `json:"status,omitempty"`
}

func NewReportRequestSubmissionAttributesWithDefaults(defaults client.Defaults) *NewReportRequestSubmissionAttributes {
	return &NewReportRequestSubmissionAttributes{

		Status: defaults.GetString("NewReportRequestSubmissionAttributes", "status"),
	}
}

func (m *NewReportRequestSubmissionAttributes) WithStatus(status string) *NewReportRequestSubmissionAttributes {

	m.Status = status

	return m
}

// Validate validates this new report request submission attributes
func (m *NewReportRequestSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var newReportRequestSubmissionAttributesTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newReportRequestSubmissionAttributesTypeStatusPropEnum = append(newReportRequestSubmissionAttributesTypeStatusPropEnum, v)
	}
}

const (

	// NewReportRequestSubmissionAttributesStatusPending captures enum value "pending"
	NewReportRequestSubmissionAttributesStatusPending string = "pending"
)

// prop value enum
func (m *NewReportRequestSubmissionAttributes) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, newReportRequestSubmissionAttributesTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewReportRequestSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewReportRequestSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewReportRequestSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res NewReportRequestSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewReportRequestSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}