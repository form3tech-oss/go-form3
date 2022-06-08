// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubscriptionDeliveryErrorResponse subscription delivery error response
// swagger:model SubscriptionDeliveryErrorResponse
type SubscriptionDeliveryErrorResponse struct {

	// data
	Data SubscriptionDeliveryError `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func SubscriptionDeliveryErrorResponseWithDefaults(defaults client.Defaults) *SubscriptionDeliveryErrorResponse {
	return &SubscriptionDeliveryErrorResponse{

		Data: *SubscriptionDeliveryErrorWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *SubscriptionDeliveryErrorResponse) WithData(data *SubscriptionDeliveryError) *SubscriptionDeliveryErrorResponse {

	m.Data = *data

	return m
}

func (m *SubscriptionDeliveryErrorResponse) WithLinks(links Links) *SubscriptionDeliveryErrorResponse {

	m.Links = &links

	return m
}

func (m *SubscriptionDeliveryErrorResponse) WithoutLinks() *SubscriptionDeliveryErrorResponse {
	m.Links = nil
	return m
}

// Validate validates this subscription delivery error response
func (m *SubscriptionDeliveryErrorResponse) Validate(formats strfmt.Registry) error {
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

func (m *SubscriptionDeliveryErrorResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if err := m.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data")
		}
		return err
	}

	return nil
}

func (m *SubscriptionDeliveryErrorResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *SubscriptionDeliveryErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionDeliveryErrorResponse) UnmarshalBinary(b []byte) error {
	var res SubscriptionDeliveryErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SubscriptionDeliveryErrorResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
