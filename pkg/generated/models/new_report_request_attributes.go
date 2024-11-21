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

// NewReportRequestAttributes new report request attributes
// swagger:model NewReportRequestAttributes
type NewReportRequestAttributes struct {

	// filter
	// Required: true
	Filter *NewReportRequestFilter `json:"filter"`

	// payment scheme
	// Required: true
	PaymentScheme string `json:"payment_scheme"`

	// report type
	// Required: true
	ReportType string `json:"report_type"`
}

func NewReportRequestAttributesWithDefaults(defaults client.Defaults) *NewReportRequestAttributes {
	return &NewReportRequestAttributes{

		Filter: NewReportRequestFilterWithDefaults(defaults),

		PaymentScheme: defaults.GetString("NewReportRequestAttributes", "payment_scheme"),

		ReportType: defaults.GetString("NewReportRequestAttributes", "report_type"),
	}
}

func (m *NewReportRequestAttributes) WithFilter(filter NewReportRequestFilter) *NewReportRequestAttributes {

	m.Filter = &filter

	return m
}

func (m *NewReportRequestAttributes) WithoutFilter() *NewReportRequestAttributes {
	m.Filter = nil
	return m
}

func (m *NewReportRequestAttributes) WithPaymentScheme(paymentScheme string) *NewReportRequestAttributes {

	m.PaymentScheme = paymentScheme

	return m
}

func (m *NewReportRequestAttributes) WithReportType(reportType string) *NewReportRequestAttributes {

	m.ReportType = reportType

	return m
}

// Validate validates this new report request attributes
func (m *NewReportRequestAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewReportRequestAttributes) validateFilter(formats strfmt.Registry) error {

	if err := validate.Required("filter", "body", m.Filter); err != nil {
		return err
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *NewReportRequestAttributes) validatePaymentScheme(formats strfmt.Registry) error {

	if err := validate.RequiredString("payment_scheme", "body", string(m.PaymentScheme)); err != nil {
		return err
	}

	return nil
}

func (m *NewReportRequestAttributes) validateReportType(formats strfmt.Registry) error {

	if err := validate.RequiredString("report_type", "body", string(m.ReportType)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewReportRequestAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewReportRequestAttributes) UnmarshalBinary(b []byte) error {
	var res NewReportRequestAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewReportRequestAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}