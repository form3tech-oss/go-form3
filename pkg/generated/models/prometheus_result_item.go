// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PrometheusResultItem prometheus result item
// swagger:model prometheusResultItem
type PrometheusResultItem struct {

	// metric
	Metric PrometheusMetric `json:"metric,omitempty"`

	// value
	Value []interface{} `json:"value"`
}

func PrometheusResultItemWithDefaults(defaults client.Defaults) *PrometheusResultItem {
	return &PrometheusResultItem{

		// TODO Metric: PrometheusMetric,

		Value: make([]interface{}, 0),
	}
}

func (m *PrometheusResultItem) WithMetric(metric PrometheusMetric) *PrometheusResultItem {

	m.Metric = metric

	return m
}

func (m *PrometheusResultItem) WithValue(value []interface{}) *PrometheusResultItem {

	m.Value = value

	return m
}

// Validate validates this prometheus result item
func (m *PrometheusResultItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrometheusResultItem) validateMetric(formats strfmt.Registry) error {

	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	if err := m.Metric.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metric")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusResultItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusResultItem) UnmarshalBinary(b []byte) error {
	var res PrometheusResultItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PrometheusResultItem) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
