// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportRequestFilter report request filter
// swagger:model ReportRequestFilter
type ReportRequestFilter struct {

	// date from
	// Format: date
	DateFrom *strfmt.Date `json:"date_from,omitempty"`

	// date to
	// Format: date
	DateTo *strfmt.Date `json:"date_to,omitempty"`

	// report users
	// Required: true
	ReportUsers []*ReportUser `json:"report_users"`
}

func ReportRequestFilterWithDefaults(defaults client.Defaults) *ReportRequestFilter {
	return &ReportRequestFilter{

		DateFrom: defaults.GetStrfmtDatePtr("ReportRequestFilter", "date_from"),

		DateTo: defaults.GetStrfmtDatePtr("ReportRequestFilter", "date_to"),

		ReportUsers: make([]*ReportUser, 0),
	}
}

func (m *ReportRequestFilter) WithDateFrom(dateFrom strfmt.Date) *ReportRequestFilter {

	m.DateFrom = &dateFrom

	return m
}

func (m *ReportRequestFilter) WithoutDateFrom() *ReportRequestFilter {
	m.DateFrom = nil
	return m
}

func (m *ReportRequestFilter) WithDateTo(dateTo strfmt.Date) *ReportRequestFilter {

	m.DateTo = &dateTo

	return m
}

func (m *ReportRequestFilter) WithoutDateTo() *ReportRequestFilter {
	m.DateTo = nil
	return m
}

func (m *ReportRequestFilter) WithReportUsers(reportUsers []*ReportUser) *ReportRequestFilter {

	m.ReportUsers = reportUsers

	return m
}

// Validate validates this report request filter
func (m *ReportRequestFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportRequestFilter) validateDateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.DateFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("date_from", "body", "date", m.DateFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportRequestFilter) validateDateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTo) { // not required
		return nil
	}

	if err := validate.FormatOf("date_to", "body", "date", m.DateTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportRequestFilter) validateReportUsers(formats strfmt.Registry) error {

	if err := validate.Required("report_users", "body", m.ReportUsers); err != nil {
		return err
	}

	for i := 0; i < len(m.ReportUsers); i++ {
		if swag.IsZero(m.ReportUsers[i]) { // not required
			continue
		}

		if m.ReportUsers[i] != nil {
			if err := m.ReportUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("report_users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportRequestFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportRequestFilter) UnmarshalBinary(b []byte) error {
	var res ReportRequestFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportRequestFilter) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
