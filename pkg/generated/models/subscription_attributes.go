// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SubscriptionAttributes subscription attributes
// swagger:model SubscriptionAttributes
type SubscriptionAttributes struct {

	// Deprecated. Please use callback_uris instead
	CallbackTransport CallbackTransport `json:"callback_transport,omitempty"`

	// Deprecated. Please use callback_uris instead
	// Pattern: ^[A-Za-z0-9 .,@:\&\?=\/\-_]*$
	CallbackURI string `json:"callback_uri,omitempty"`

	// callback uris
	CallbackUris []*CallbackURI `json:"callback_uris"`

	// deactivated
	Deactivated bool `json:"deactivated,omitempty"`

	// event type
	// Required: true
	// Pattern: ^[A-Za-z_-]*$
	EventType *string `json:"event_type"`

	// filter
	Filter string `json:"filter,omitempty"`

	// record type
	// Required: true
	// Pattern: ^[A-Za-z_-]*$
	RecordType *string `json:"record_type"`

	// All purpose list of key-value pairs to store specific data for the associated subscription.
	// Max Items: 5
	UserDefinedData []*SubscriptionUserDefinedData `json:"user_defined_data,omitempty"`

	// user id
	// Read Only: true
	// Format: uuid
	UserID strfmt.UUID `json:"user_id,omitempty"`
}

func SubscriptionAttributesWithDefaults(defaults client.Defaults) *SubscriptionAttributes {
	return &SubscriptionAttributes{

		// TODO CallbackTransport: CallbackTransport,

		CallbackURI: defaults.GetString("SubscriptionAttributes", "callback_uri"),

		CallbackUris: make([]*CallbackURI, 0),

		Deactivated: defaults.GetBool("SubscriptionAttributes", "deactivated"),

		EventType: defaults.GetStringPtr("SubscriptionAttributes", "event_type"),

		Filter: defaults.GetString("SubscriptionAttributes", "filter"),

		RecordType: defaults.GetStringPtr("SubscriptionAttributes", "record_type"),

		UserDefinedData: make([]*SubscriptionUserDefinedData, 0),

		UserID: defaults.GetStrfmtUUID("SubscriptionAttributes", "user_id"),
	}
}

func (m *SubscriptionAttributes) WithCallbackTransport(callbackTransport CallbackTransport) *SubscriptionAttributes {

	m.CallbackTransport = callbackTransport

	return m
}

func (m *SubscriptionAttributes) WithCallbackURI(callbackURI string) *SubscriptionAttributes {

	m.CallbackURI = callbackURI

	return m
}

func (m *SubscriptionAttributes) WithCallbackUris(callbackUris []*CallbackURI) *SubscriptionAttributes {

	m.CallbackUris = callbackUris

	return m
}

func (m *SubscriptionAttributes) WithDeactivated(deactivated bool) *SubscriptionAttributes {

	m.Deactivated = deactivated

	return m
}

func (m *SubscriptionAttributes) WithEventType(eventType string) *SubscriptionAttributes {

	m.EventType = &eventType

	return m
}

func (m *SubscriptionAttributes) WithoutEventType() *SubscriptionAttributes {
	m.EventType = nil
	return m
}

func (m *SubscriptionAttributes) WithFilter(filter string) *SubscriptionAttributes {

	m.Filter = filter

	return m
}

func (m *SubscriptionAttributes) WithRecordType(recordType string) *SubscriptionAttributes {

	m.RecordType = &recordType

	return m
}

func (m *SubscriptionAttributes) WithoutRecordType() *SubscriptionAttributes {
	m.RecordType = nil
	return m
}

func (m *SubscriptionAttributes) WithUserDefinedData(userDefinedData []*SubscriptionUserDefinedData) *SubscriptionAttributes {

	m.UserDefinedData = userDefinedData

	return m
}

func (m *SubscriptionAttributes) WithUserID(userID strfmt.UUID) *SubscriptionAttributes {

	m.UserID = userID

	return m
}

// Validate validates this subscription attributes
func (m *SubscriptionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallbackTransport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallbackURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallbackUris(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDefinedData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionAttributes) validateCallbackTransport(formats strfmt.Registry) error {

	if swag.IsZero(m.CallbackTransport) { // not required
		return nil
	}

	if err := m.CallbackTransport.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("callback_transport")
		}
		return err
	}

	return nil
}

func (m *SubscriptionAttributes) validateCallbackURI(formats strfmt.Registry) error {

	if swag.IsZero(m.CallbackURI) { // not required
		return nil
	}

	if err := validate.Pattern("callback_uri", "body", string(m.CallbackURI), `^[A-Za-z0-9 .,@:\&\?=\/\-_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionAttributes) validateCallbackUris(formats strfmt.Registry) error {

	if swag.IsZero(m.CallbackUris) { // not required
		return nil
	}

	for i := 0; i < len(m.CallbackUris); i++ {
		if swag.IsZero(m.CallbackUris[i]) { // not required
			continue
		}

		if m.CallbackUris[i] != nil {
			if err := m.CallbackUris[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("callback_uris" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubscriptionAttributes) validateEventType(formats strfmt.Registry) error {

	if err := validate.Required("event_type", "body", m.EventType); err != nil {
		return err
	}

	if err := validate.Pattern("event_type", "body", string(*m.EventType), `^[A-Za-z_-]*$`); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionAttributes) validateRecordType(formats strfmt.Registry) error {

	if err := validate.Required("record_type", "body", m.RecordType); err != nil {
		return err
	}

	if err := validate.Pattern("record_type", "body", string(*m.RecordType), `^[A-Za-z_-]*$`); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionAttributes) validateUserDefinedData(formats strfmt.Registry) error {

	if swag.IsZero(m.UserDefinedData) { // not required
		return nil
	}

	iUserDefinedDataSize := int64(len(m.UserDefinedData))

	if err := validate.MaxItems("user_defined_data", "body", iUserDefinedDataSize, 5); err != nil {
		return err
	}

	for i := 0; i < len(m.UserDefinedData); i++ {
		if swag.IsZero(m.UserDefinedData[i]) { // not required
			continue
		}

		if m.UserDefinedData[i] != nil {
			if err := m.UserDefinedData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_defined_data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubscriptionAttributes) validateUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.UserID) { // not required
		return nil
	}

	if err := validate.FormatOf("user_id", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionAttributes) UnmarshalBinary(b []byte) error {
	var res SubscriptionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SubscriptionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
