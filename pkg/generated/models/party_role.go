// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PartyRole Role of the party. Enum of pre-defined values, new values can be added when needed
// swagger:model PartyRole
type PartyRole string

const (

	// PartyRoleInitiatingParty captures enum value "InitiatingParty"
	PartyRoleInitiatingParty PartyRole = "InitiatingParty"

	// PartyRoleCaseCreator captures enum value "CaseCreator"
	PartyRoleCaseCreator PartyRole = "CaseCreator"

	// PartyRoleOriginator captures enum value "Originator"
	PartyRoleOriginator PartyRole = "Originator"
)

// for schema
var partyRoleEnum []interface{}

func init() {
	var res []PartyRole
	if err := json.Unmarshal([]byte(`["InitiatingParty","CaseCreator","Originator"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyRoleEnum = append(partyRoleEnum, v)
	}
}

func (m PartyRole) validatePartyRoleEnum(path, location string, value PartyRole) error {
	if err := validate.Enum(path, location, value, partyRoleEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this party role
func (m PartyRole) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePartyRoleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyRole) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}