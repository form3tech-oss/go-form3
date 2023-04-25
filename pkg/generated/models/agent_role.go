// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// AgentRole Role of the agent in the payment chain. Enum of pre-defined values, new values can be added when needed
// swagger:model AgentRole
type AgentRole string

const (

	// AgentRoleIntermediaryAgent1 captures enum value "IntermediaryAgent1"
	AgentRoleIntermediaryAgent1 AgentRole = "IntermediaryAgent1"

	// AgentRoleIntermediaryAgent2 captures enum value "IntermediaryAgent2"
	AgentRoleIntermediaryAgent2 AgentRole = "IntermediaryAgent2"

	// AgentRoleIntermediaryAgent3 captures enum value "IntermediaryAgent3"
	AgentRoleIntermediaryAgent3 AgentRole = "IntermediaryAgent3"

	// AgentRoleInstructingAgent captures enum value "InstructingAgent"
	AgentRoleInstructingAgent AgentRole = "InstructingAgent"

	// AgentRoleInstructedAgent captures enum value "InstructedAgent"
	AgentRoleInstructedAgent AgentRole = "InstructedAgent"

	// AgentRoleAdditionalBeneficiaryPartyAgentInformation captures enum value "AdditionalBeneficiaryPartyAgentInformation"
	AgentRoleAdditionalBeneficiaryPartyAgentInformation AgentRole = "AdditionalBeneficiaryPartyAgentInformation"

	// AgentRoleCreditorAgent captures enum value "CreditorAgent"
	AgentRoleCreditorAgent AgentRole = "CreditorAgent"

	// AgentRoleDebtorAgent captures enum value "DebtorAgent"
	AgentRoleDebtorAgent AgentRole = "DebtorAgent"
)

// for schema
var agentRoleEnum []interface{}

func init() {
	var res []AgentRole
	if err := json.Unmarshal([]byte(`["IntermediaryAgent1","IntermediaryAgent2","IntermediaryAgent3","InstructingAgent","InstructedAgent","AdditionalBeneficiaryPartyAgentInformation","CreditorAgent","DebtorAgent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		agentRoleEnum = append(agentRoleEnum, v)
	}
}

func (m AgentRole) validateAgentRoleEnum(path, location string, value AgentRole) error {
	if err := validate.Enum(path, location, value, agentRoleEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this agent role
func (m AgentRole) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAgentRoleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentRole) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
