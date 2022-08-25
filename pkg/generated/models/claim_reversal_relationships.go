// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClaimReversalRelationships claim reversal relationships
// swagger:model ClaimReversalRelationships
type ClaimReversalRelationships struct {

	// claim
	Claim *ClaimReversalRelationshipsClaim `json:"claim,omitempty"`

	// claim reversal submission
	ClaimReversalSubmission *ClaimReversalRelationshipsClaimReversalSubmission `json:"claim_reversal_submission,omitempty"`
}

func ClaimReversalRelationshipsWithDefaults(defaults client.Defaults) *ClaimReversalRelationships {
	return &ClaimReversalRelationships{

		Claim: ClaimReversalRelationshipsClaimWithDefaults(defaults),

		ClaimReversalSubmission: ClaimReversalRelationshipsClaimReversalSubmissionWithDefaults(defaults),
	}
}

func (m *ClaimReversalRelationships) WithClaim(claim ClaimReversalRelationshipsClaim) *ClaimReversalRelationships {

	m.Claim = &claim

	return m
}

func (m *ClaimReversalRelationships) WithoutClaim() *ClaimReversalRelationships {
	m.Claim = nil
	return m
}

func (m *ClaimReversalRelationships) WithClaimReversalSubmission(claimReversalSubmission ClaimReversalRelationshipsClaimReversalSubmission) *ClaimReversalRelationships {

	m.ClaimReversalSubmission = &claimReversalSubmission

	return m
}

func (m *ClaimReversalRelationships) WithoutClaimReversalSubmission() *ClaimReversalRelationships {
	m.ClaimReversalSubmission = nil
	return m
}

// Validate validates this claim reversal relationships
func (m *ClaimReversalRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClaim(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClaimReversalSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimReversalRelationships) validateClaim(formats strfmt.Registry) error {

	if swag.IsZero(m.Claim) { // not required
		return nil
	}

	if m.Claim != nil {
		if err := m.Claim.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("claim")
			}
			return err
		}
	}

	return nil
}

func (m *ClaimReversalRelationships) validateClaimReversalSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.ClaimReversalSubmission) { // not required
		return nil
	}

	if m.ClaimReversalSubmission != nil {
		if err := m.ClaimReversalSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("claim_reversal_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimReversalRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimReversalRelationships) UnmarshalBinary(b []byte) error {
	var res ClaimReversalRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimReversalRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ClaimReversalRelationshipsClaim claim reversal relationships claim
// swagger:model ClaimReversalRelationshipsClaim
type ClaimReversalRelationshipsClaim struct {

	// data
	Data []*Claim `json:"data"`
}

func ClaimReversalRelationshipsClaimWithDefaults(defaults client.Defaults) *ClaimReversalRelationshipsClaim {
	return &ClaimReversalRelationshipsClaim{

		Data: make([]*Claim, 0),
	}
}

func (m *ClaimReversalRelationshipsClaim) WithData(data []*Claim) *ClaimReversalRelationshipsClaim {

	m.Data = data

	return m
}

// Validate validates this claim reversal relationships claim
func (m *ClaimReversalRelationshipsClaim) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimReversalRelationshipsClaim) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claim" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimReversalRelationshipsClaim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimReversalRelationshipsClaim) UnmarshalBinary(b []byte) error {
	var res ClaimReversalRelationshipsClaim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimReversalRelationshipsClaim) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ClaimReversalRelationshipsClaimReversalSubmission claim reversal relationships claim reversal submission
// swagger:model ClaimReversalRelationshipsClaimReversalSubmission
type ClaimReversalRelationshipsClaimReversalSubmission struct {

	// data
	Data []*ClaimReversalSubmission `json:"data"`
}

func ClaimReversalRelationshipsClaimReversalSubmissionWithDefaults(defaults client.Defaults) *ClaimReversalRelationshipsClaimReversalSubmission {
	return &ClaimReversalRelationshipsClaimReversalSubmission{

		Data: make([]*ClaimReversalSubmission, 0),
	}
}

func (m *ClaimReversalRelationshipsClaimReversalSubmission) WithData(data []*ClaimReversalSubmission) *ClaimReversalRelationshipsClaimReversalSubmission {

	m.Data = data

	return m
}

// Validate validates this claim reversal relationships claim reversal submission
func (m *ClaimReversalRelationshipsClaimReversalSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimReversalRelationshipsClaimReversalSubmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claim_reversal_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimReversalRelationshipsClaimReversalSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimReversalRelationshipsClaimReversalSubmission) UnmarshalBinary(b []byte) error {
	var res ClaimReversalRelationshipsClaimReversalSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimReversalRelationshipsClaimReversalSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
