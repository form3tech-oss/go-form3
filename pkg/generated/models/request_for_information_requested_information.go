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

// RequestForInformationRequestedInformation request for information requested information
// swagger:model RequestForInformationRequestedInformation
type RequestForInformationRequestedInformation []*RequestForInformationRequestedInformationItems0

// Validate validates this request for information requested information
func (m RequestForInformationRequestedInformation) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestForInformationRequestedInformation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// RequestForInformationRequestedInformationItems0 request for information requested information items0
// swagger:model RequestForInformationRequestedInformationItems0
type RequestForInformationRequestedInformationItems0 struct {

	// additional information
	AdditionalInformation string `json:"additional_information,omitempty"`

	// code
	// Required: true
	Code *string `json:"code"`

	// type
	// Required: true
	// Enum: [missing incorrect]
	Type *string `json:"type"`
}

func RequestForInformationRequestedInformationItems0WithDefaults(defaults client.Defaults) *RequestForInformationRequestedInformationItems0 {
	return &RequestForInformationRequestedInformationItems0{

		AdditionalInformation: defaults.GetString("RequestForInformationRequestedInformationItems0", "additional_information"),

		Code: defaults.GetStringPtr("RequestForInformationRequestedInformationItems0", "code"),

		Type: defaults.GetStringPtr("RequestForInformationRequestedInformationItems0", "type"),
	}
}

func (m *RequestForInformationRequestedInformationItems0) WithAdditionalInformation(additionalInformation string) *RequestForInformationRequestedInformationItems0 {

	m.AdditionalInformation = additionalInformation

	return m
}

func (m *RequestForInformationRequestedInformationItems0) WithCode(code string) *RequestForInformationRequestedInformationItems0 {

	m.Code = &code

	return m
}

func (m *RequestForInformationRequestedInformationItems0) WithoutCode() *RequestForInformationRequestedInformationItems0 {
	m.Code = nil
	return m
}

func (m *RequestForInformationRequestedInformationItems0) WithType(typeVar string) *RequestForInformationRequestedInformationItems0 {

	m.Type = &typeVar

	return m
}

func (m *RequestForInformationRequestedInformationItems0) WithoutType() *RequestForInformationRequestedInformationItems0 {
	m.Type = nil
	return m
}

// Validate validates this request for information requested information items0
func (m *RequestForInformationRequestedInformationItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestForInformationRequestedInformationItems0) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

var requestForInformationRequestedInformationItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["missing","incorrect"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		requestForInformationRequestedInformationItems0TypeTypePropEnum = append(requestForInformationRequestedInformationItems0TypeTypePropEnum, v)
	}
}

const (

	// RequestForInformationRequestedInformationItems0TypeMissing captures enum value "missing"
	RequestForInformationRequestedInformationItems0TypeMissing string = "missing"

	// RequestForInformationRequestedInformationItems0TypeIncorrect captures enum value "incorrect"
	RequestForInformationRequestedInformationItems0TypeIncorrect string = "incorrect"
)

// prop value enum
func (m *RequestForInformationRequestedInformationItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, requestForInformationRequestedInformationItems0TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RequestForInformationRequestedInformationItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestForInformationRequestedInformationItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestForInformationRequestedInformationItems0) UnmarshalBinary(b []byte) error {
	var res RequestForInformationRequestedInformationItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RequestForInformationRequestedInformationItems0) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
