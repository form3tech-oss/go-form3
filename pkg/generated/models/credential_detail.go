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

// CredentialDetail credential detail
// swagger:model CredentialDetail
type CredentialDetail struct {

	// client id
	ClientID string `json:"client_id,omitempty"`

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`

	// rate limit
	RateLimit *CredentialDetailRateLimit `json:"rate_limit,omitempty"`

	// role
	Role *string `json:"role,omitempty"`
}

func CredentialDetailWithDefaults(defaults client.Defaults) *CredentialDetail {
	return &CredentialDetail{

		ClientID: defaults.GetString("CredentialDetail", "client_id"),

		ClientSecret: defaults.GetString("CredentialDetail", "client_secret"),

		RateLimit: CredentialDetailRateLimitWithDefaults(defaults),

		Role: defaults.GetStringPtr("CredentialDetail", "role"),
	}
}

func (m *CredentialDetail) WithClientID(clientID string) *CredentialDetail {

	m.ClientID = clientID

	return m
}

func (m *CredentialDetail) WithClientSecret(clientSecret string) *CredentialDetail {

	m.ClientSecret = clientSecret

	return m
}

func (m *CredentialDetail) WithRateLimit(rateLimit CredentialDetailRateLimit) *CredentialDetail {

	m.RateLimit = &rateLimit

	return m
}

func (m *CredentialDetail) WithoutRateLimit() *CredentialDetail {
	m.RateLimit = nil
	return m
}

func (m *CredentialDetail) WithRole(role string) *CredentialDetail {

	m.Role = &role

	return m
}

func (m *CredentialDetail) WithoutRole() *CredentialDetail {
	m.Role = nil
	return m
}

// Validate validates this credential detail
func (m *CredentialDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRateLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialDetail) validateRateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.RateLimit) { // not required
		return nil
	}

	if m.RateLimit != nil {
		if err := m.RateLimit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rate_limit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialDetail) UnmarshalBinary(b []byte) error {
	var res CredentialDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *CredentialDetail) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// CredentialDetailRateLimit credential detail rate limit
// swagger:model CredentialDetailRateLimit
type CredentialDetailRateLimit struct {

	// block interval seconds
	// Minimum: 1
	BlockIntervalSeconds int64 `json:"block_interval_seconds,omitempty"`

	// interval seconds
	// Minimum: 1
	IntervalSeconds int64 `json:"interval_seconds,omitempty"`

	// rate
	// Minimum: 1
	Rate int64 `json:"rate,omitempty"`
}

func CredentialDetailRateLimitWithDefaults(defaults client.Defaults) *CredentialDetailRateLimit {
	return &CredentialDetailRateLimit{

		BlockIntervalSeconds: defaults.GetInt64("CredentialDetailRateLimit", "block_interval_seconds"),

		IntervalSeconds: defaults.GetInt64("CredentialDetailRateLimit", "interval_seconds"),

		Rate: defaults.GetInt64("CredentialDetailRateLimit", "rate"),
	}
}

func (m *CredentialDetailRateLimit) WithBlockIntervalSeconds(blockIntervalSeconds int64) *CredentialDetailRateLimit {

	m.BlockIntervalSeconds = blockIntervalSeconds

	return m
}

func (m *CredentialDetailRateLimit) WithIntervalSeconds(intervalSeconds int64) *CredentialDetailRateLimit {

	m.IntervalSeconds = intervalSeconds

	return m
}

func (m *CredentialDetailRateLimit) WithRate(rate int64) *CredentialDetailRateLimit {

	m.Rate = rate

	return m
}

// Validate validates this credential detail rate limit
func (m *CredentialDetailRateLimit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockIntervalSeconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntervalSeconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialDetailRateLimit) validateBlockIntervalSeconds(formats strfmt.Registry) error {

	if swag.IsZero(m.BlockIntervalSeconds) { // not required
		return nil
	}

	if err := validate.MinimumInt("rate_limit"+"."+"block_interval_seconds", "body", int64(m.BlockIntervalSeconds), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *CredentialDetailRateLimit) validateIntervalSeconds(formats strfmt.Registry) error {

	if swag.IsZero(m.IntervalSeconds) { // not required
		return nil
	}

	if err := validate.MinimumInt("rate_limit"+"."+"interval_seconds", "body", int64(m.IntervalSeconds), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *CredentialDetailRateLimit) validateRate(formats strfmt.Registry) error {

	if swag.IsZero(m.Rate) { // not required
		return nil
	}

	if err := validate.MinimumInt("rate_limit"+"."+"rate", "body", int64(m.Rate), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialDetailRateLimit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialDetailRateLimit) UnmarshalBinary(b []byte) error {
	var res CredentialDetailRateLimit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *CredentialDetailRateLimit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
