// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SchemeMessageEntryItem scheme message entry item
// swagger:model SchemeMessageEntryItem
type SchemeMessageEntryItem struct {

	// key
	// Required: true
	Key string `json:"key"`

	// value
	Value string `json:"value,omitempty"`
}

// line 140

func SchemeMessageEntryItemWithDefaults(defaults client.Defaults) *SchemeMessageEntryItem {
	return &SchemeMessageEntryItem{

		Key: defaults.GetString("SchemeMessageEntryItem", "key"),

		Value: defaults.GetString("SchemeMessageEntryItem", "value"),
	}
}

func (m *SchemeMessageEntryItem) WithKey(key string) *SchemeMessageEntryItem {

	m.Key = key

	return m
}

func (m *SchemeMessageEntryItem) WithValue(value string) *SchemeMessageEntryItem {

	m.Value = value

	return m
}

// Validate validates this scheme message entry item
func (m *SchemeMessageEntryItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeMessageEntryItem) validateKey(formats strfmt.Registry) error {

	if err := validate.RequiredString("key", "body", string(m.Key)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeMessageEntryItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeMessageEntryItem) UnmarshalBinary(b []byte) error {
	var res SchemeMessageEntryItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeMessageEntryItem) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
