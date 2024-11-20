// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Credential credential
// swagger:model Credential
type Credential struct {

	// client id
	ClientID string `json:"client_id,omitempty"`
}

func CredentialWithDefaults(defaults client.Defaults) *Credential {
	return &Credential{

		ClientID: defaults.GetString("Credential", "client_id"),
	}
}

func (m *Credential) WithClientID(clientID string) *Credential {

	m.ClientID = clientID

	return m
}

// Validate validates this credential
func (m *Credential) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Credential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Credential) UnmarshalBinary(b []byte) error {
	var res Credential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Credential) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
