// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIError Api error
// swagger:model ApiError
type APIError struct {

	// error code
	// Format: uuid
	ErrorCode strfmt.UUID `json:"error_code,omitempty"`

	// error message
	ErrorMessage string `json:"error_message,omitempty"`
}

func APIErrorWithDefaults(defaults client.Defaults) *APIError {
	return &APIError{

		ErrorCode: defaults.GetStrfmtUUID("ApiError", "error_code"),

		ErrorMessage: defaults.GetString("ApiError", "error_message"),
	}
}

func (m *APIError) WithErrorCode(errorCode strfmt.UUID) *APIError {

	m.ErrorCode = errorCode

	return m
}

func (m *APIError) WithErrorMessage(errorMessage string) *APIError {

	m.ErrorMessage = errorMessage

	return m
}

// Validate validates this Api error
func (m *APIError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIError) validateErrorCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorCode) { // not required
		return nil
	}

	if err := validate.FormatOf("error_code", "body", "uuid", m.ErrorCode.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIError) UnmarshalBinary(b []byte) error {
	var res APIError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *APIError) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
