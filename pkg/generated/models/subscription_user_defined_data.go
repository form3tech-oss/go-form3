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

// SubscriptionUserDefinedData subscription user defined data
// swagger:model SubscriptionUserDefinedData
type SubscriptionUserDefinedData struct {

	// A user-defined key, up to 45 characters. An entry always contains both `key` and `value` fields.
	// Required: true
	// Max Length: 45
	// Min Length: 1
	Key *string `json:"key"`

	// A user-defined value, up to 45 characters. An entry always contains both `key` and `value` fields.
	// Required: true
	// Max Length: 45
	// Min Length: 1
	Value *string `json:"value"`
}

func SubscriptionUserDefinedDataWithDefaults(defaults client.Defaults) *SubscriptionUserDefinedData {
	return &SubscriptionUserDefinedData{

		Key: defaults.GetStringPtr("SubscriptionUserDefinedData", "key"),

		Value: defaults.GetStringPtr("SubscriptionUserDefinedData", "value"),
	}
}

func (m *SubscriptionUserDefinedData) WithKey(key string) *SubscriptionUserDefinedData {

	m.Key = &key

	return m
}

func (m *SubscriptionUserDefinedData) WithoutKey() *SubscriptionUserDefinedData {
	m.Key = nil
	return m
}

func (m *SubscriptionUserDefinedData) WithValue(value string) *SubscriptionUserDefinedData {

	m.Value = &value

	return m
}

func (m *SubscriptionUserDefinedData) WithoutValue() *SubscriptionUserDefinedData {
	m.Value = nil
	return m
}

// Validate validates this subscription user defined data
func (m *SubscriptionUserDefinedData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionUserDefinedData) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	if err := validate.MinLength("key", "body", string(*m.Key), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("key", "body", string(*m.Key), 45); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionUserDefinedData) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.MinLength("value", "body", string(*m.Value), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("value", "body", string(*m.Value), 45); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionUserDefinedData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionUserDefinedData) UnmarshalBinary(b []byte) error {
	var res SubscriptionUserDefinedData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SubscriptionUserDefinedData) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
