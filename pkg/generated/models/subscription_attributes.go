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

// SubscriptionAttributes subscription attributes
// swagger:model SubscriptionAttributes
type SubscriptionAttributes struct {

	// callback transport
	// Required: true
	// Enum: [http email queue]
	CallbackTransport *string `json:"callback_transport"`

	// callback uri
	// Required: true
	// Pattern: ^[A-Za-z0-9 .,@:\&\?=\/\-_]*$
	CallbackURI *string `json:"callback_uri"`

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

	// user id
	// Read Only: true
	// Format: uuid
	UserID strfmt.UUID `json:"user_id,omitempty"`
}

// line 140

func SubscriptionAttributesWithDefaults(defaults client.Defaults) *SubscriptionAttributes {
	return &SubscriptionAttributes{

		CallbackTransport: defaults.GetStringPtr("SubscriptionAttributes", "callback_transport"),

		CallbackURI: defaults.GetStringPtr("SubscriptionAttributes", "callback_uri"),

		Deactivated: defaults.GetBool("SubscriptionAttributes", "deactivated"),

		EventType: defaults.GetStringPtr("SubscriptionAttributes", "event_type"),

		Filter: defaults.GetString("SubscriptionAttributes", "filter"),

		RecordType: defaults.GetStringPtr("SubscriptionAttributes", "record_type"),

		UserID: defaults.GetStrfmtUUID("SubscriptionAttributes", "user_id"),
	}
}

func (m *SubscriptionAttributes) WithCallbackTransport(callbackTransport string) *SubscriptionAttributes {

	m.CallbackTransport = &callbackTransport

	return m
}

func (m *SubscriptionAttributes) WithoutCallbackTransport() *SubscriptionAttributes {
	m.CallbackTransport = nil
	return m
}

func (m *SubscriptionAttributes) WithCallbackURI(callbackURI string) *SubscriptionAttributes {

	m.CallbackURI = &callbackURI

	return m
}

func (m *SubscriptionAttributes) WithoutCallbackURI() *SubscriptionAttributes {
	m.CallbackURI = nil
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

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordType(formats); err != nil {
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

var subscriptionAttributesTypeCallbackTransportPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","email","queue"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionAttributesTypeCallbackTransportPropEnum = append(subscriptionAttributesTypeCallbackTransportPropEnum, v)
	}
}

const (

	// SubscriptionAttributesCallbackTransportHTTP captures enum value "http"
	SubscriptionAttributesCallbackTransportHTTP string = "http"

	// SubscriptionAttributesCallbackTransportEmail captures enum value "email"
	SubscriptionAttributesCallbackTransportEmail string = "email"

	// SubscriptionAttributesCallbackTransportQueue captures enum value "queue"
	SubscriptionAttributesCallbackTransportQueue string = "queue"
)

// prop value enum
func (m *SubscriptionAttributes) validateCallbackTransportEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionAttributesTypeCallbackTransportPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionAttributes) validateCallbackTransport(formats strfmt.Registry) error {

	if err := validate.Required("callback_transport", "body", m.CallbackTransport); err != nil {
		return err
	}

	// value enum
	if err := m.validateCallbackTransportEnum("callback_transport", "body", *m.CallbackTransport); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionAttributes) validateCallbackURI(formats strfmt.Registry) error {

	if err := validate.Required("callback_uri", "body", m.CallbackURI); err != nil {
		return err
	}

	if err := validate.Pattern("callback_uri", "body", string(*m.CallbackURI), `^[A-Za-z0-9 .,@:\&\?=\/\-_]*$`); err != nil {
		return err
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
