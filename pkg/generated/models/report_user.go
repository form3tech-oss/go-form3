// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReportUser report user
// swagger:model ReportUser
type ReportUser struct {

	// Identifies a user or organization
	UserID string `json:"user_id,omitempty"`

	// Type of identifier for the user or organization
	UserIDCode string `json:"user_id_code,omitempty"`
}

func ReportUserWithDefaults(defaults client.Defaults) *ReportUser {
	return &ReportUser{

		UserID: defaults.GetString("ReportUser", "user_id"),

		UserIDCode: defaults.GetString("ReportUser", "user_id_code"),
	}
}

func (m *ReportUser) WithUserID(userID string) *ReportUser {

	m.UserID = userID

	return m
}

func (m *ReportUser) WithUserIDCode(userIDCode string) *ReportUser {

	m.UserIDCode = userIDCode

	return m
}

// Validate validates this report user
func (m *ReportUser) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportUser) UnmarshalBinary(b []byte) error {
	var res ReportUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportUser) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
