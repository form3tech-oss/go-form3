// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/go-openapi/strfmt"
)

// ReturnSubmissionTaskGroup Identifies the task group the task belongs to
// swagger:model ReturnSubmissionTaskGroup
type ReturnSubmissionTaskGroup string

// Validate validates this return submission task group
func (m ReturnSubmissionTaskGroup) Validate(formats strfmt.Registry) error {
	return nil
}
func (m *ReturnSubmissionTaskGroup) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}