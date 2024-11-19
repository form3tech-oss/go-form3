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

// PrometheusMetricData prometheus metric data
// swagger:model prometheusMetricData
type PrometheusMetricData struct {

	// result
	Result PrometheusResult `json:"result,omitempty"`

	// result type
	ResultType string `json:"resultType,omitempty"`
}

func PrometheusMetricDataWithDefaults(defaults client.Defaults) *PrometheusMetricData {
	return &PrometheusMetricData{

		// TODO Result: PrometheusResult,

		ResultType: defaults.GetString("prometheusMetricData", "resultType"),
	}
}

func (m *PrometheusMetricData) WithResult(result PrometheusResult) *PrometheusMetricData {

	m.Result = result

	return m
}

func (m *PrometheusMetricData) WithResultType(resultType string) *PrometheusMetricData {

	m.ResultType = resultType

	return m
}

// Validate validates this prometheus metric data
func (m *PrometheusMetricData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrometheusMetricData) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if err := m.Result.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("result")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusMetricData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusMetricData) UnmarshalBinary(b []byte) error {
	var res PrometheusMetricData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PrometheusMetricData) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
