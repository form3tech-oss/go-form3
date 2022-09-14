// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CallbackURI callback URI
// swagger:model CallbackURI
type CallbackURI struct {

	// Type of transport: queue, public http or private http for specific cloud
	// Required: true
	CallbackTransport *CallbackTransport `json:"callback_transport"`

	// URI that will be called with the notification
	// Required: true
	// Pattern: ^[A-Za-z0-9 .,@:\&\?=\/\-_]*$
	CallbackURI *string `json:"callback_uri"`
}

func CallbackURIWithDefaults(defaults client.Defaults) *CallbackURI {
	return &CallbackURI{

		// TODO CallbackTransport: CallbackTransport,

		CallbackURI: defaults.GetStringPtr("CallbackURI", "callback_uri"),
	}
}

func (m *CallbackURI) WithCallbackTransport(callbackTransport CallbackTransport) *CallbackURI {

	m.CallbackTransport = &callbackTransport

	return m
}

func (m *CallbackURI) WithoutCallbackTransport() *CallbackURI {
	m.CallbackTransport = nil
	return m
}

func (m *CallbackURI) WithCallbackURI(callbackURI string) *CallbackURI {

	m.CallbackURI = &callbackURI

	return m
}

func (m *CallbackURI) WithoutCallbackURI() *CallbackURI {
	m.CallbackURI = nil
	return m
}

// Validate validates this callback URI
func (m *CallbackURI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallbackTransport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallbackURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallbackURI) validateCallbackTransport(formats strfmt.Registry) error {

	if err := validate.Required("callback_transport", "body", m.CallbackTransport); err != nil {
		return err
	}

	if m.CallbackTransport != nil {
		if err := m.CallbackTransport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callback_transport")
			}
			return err
		}
	}

	return nil
}

func (m *CallbackURI) validateCallbackURI(formats strfmt.Registry) error {

	if err := validate.Required("callback_uri", "body", m.CallbackURI); err != nil {
		return err
	}

	if err := validate.Pattern("callback_uri", "body", string(*m.CallbackURI), `^[A-Za-z0-9 .,@:\&\?=\/\-_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CallbackURI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallbackURI) UnmarshalBinary(b []byte) error {
	var res CallbackURI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *CallbackURI) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
