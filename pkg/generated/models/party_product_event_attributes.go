// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PartyProductEventAttributes party product event attributes
// swagger:model PartyProductEventAttributes
type PartyProductEventAttributes struct {

	// provider status
	ProviderStatus string `json:"provider_status,omitempty"`

	// provider status items
	ProviderStatusItems []*ProviderStatusItem `json:"provider_status_items"`

	// status
	Status string `json:"status,omitempty"`
}

func PartyProductEventAttributesWithDefaults(defaults client.Defaults) *PartyProductEventAttributes {
	return &PartyProductEventAttributes{

		ProviderStatus: defaults.GetString("PartyProductEventAttributes", "provider_status"),

		ProviderStatusItems: make([]*ProviderStatusItem, 0),

		Status: defaults.GetString("PartyProductEventAttributes", "status"),
	}
}

func (m *PartyProductEventAttributes) WithProviderStatus(providerStatus string) *PartyProductEventAttributes {

	m.ProviderStatus = providerStatus

	return m
}

func (m *PartyProductEventAttributes) WithProviderStatusItems(providerStatusItems []*ProviderStatusItem) *PartyProductEventAttributes {

	m.ProviderStatusItems = providerStatusItems

	return m
}

func (m *PartyProductEventAttributes) WithStatus(status string) *PartyProductEventAttributes {

	m.Status = status

	return m
}

// Validate validates this party product event attributes
func (m *PartyProductEventAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProviderStatusItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyProductEventAttributes) validateProviderStatusItems(formats strfmt.Registry) error {

	if swag.IsZero(m.ProviderStatusItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ProviderStatusItems); i++ {
		if swag.IsZero(m.ProviderStatusItems[i]) { // not required
			continue
		}

		if m.ProviderStatusItems[i] != nil {
			if err := m.ProviderStatusItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("provider_status_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyProductEventAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyProductEventAttributes) UnmarshalBinary(b []byte) error {
	var res PartyProductEventAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyProductEventAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
