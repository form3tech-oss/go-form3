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
)

// SubscriptionDeliveryErrorListResponse subscription delivery error list response
// swagger:model SubscriptionDeliveryErrorListResponse
type SubscriptionDeliveryErrorListResponse struct {

	// data
	Data []SubscriptionDeliveryError `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func SubscriptionDeliveryErrorListResponseWithDefaults(defaults client.Defaults) *SubscriptionDeliveryErrorListResponse {
	return &SubscriptionDeliveryErrorListResponse{

		Data: make([]SubscriptionDeliveryError, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *SubscriptionDeliveryErrorListResponse) WithData(data []SubscriptionDeliveryError) *SubscriptionDeliveryErrorListResponse {

	m.Data = data

	return m
}

func (m *SubscriptionDeliveryErrorListResponse) WithLinks(links Links) *SubscriptionDeliveryErrorListResponse {

	m.Links = &links

	return m
}

func (m *SubscriptionDeliveryErrorListResponse) WithoutLinks() *SubscriptionDeliveryErrorListResponse {
	m.Links = nil
	return m
}

// Validate validates this subscription delivery error list response
func (m *SubscriptionDeliveryErrorListResponse) Validate(formats strfmt.Registry) error {
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

func (m *SubscriptionDeliveryErrorListResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {

		if err := m.Data[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SubscriptionDeliveryErrorListResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *SubscriptionDeliveryErrorListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionDeliveryErrorListResponse) UnmarshalBinary(b []byte) error {
	var res SubscriptionDeliveryErrorListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SubscriptionDeliveryErrorListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
