// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/go-openapi/strfmt"
)

// RequestForInformationQuerySubTypes request for information query sub types
// swagger:model RequestForInformationQuerySubTypes
type RequestForInformationQuerySubTypes []string

// Validate validates this request for information query sub types
func (m RequestForInformationQuerySubTypes) Validate(formats strfmt.Registry) error {
	return nil
}
func (m *RequestForInformationQuerySubTypes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}