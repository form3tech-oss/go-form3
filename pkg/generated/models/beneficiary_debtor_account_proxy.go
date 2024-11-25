// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BeneficiaryDebtorAccountProxy beneficiary debtor account proxy
// swagger:model BeneficiaryDebtorAccountProxy
type BeneficiaryDebtorAccountProxy struct {

	// Name of the identification scheme, in a free text form. When used proxy_id_code must not be present
	// Max Length: 35
	Proxy *string `json:"proxy,omitempty"`

	// Identification used to indicate the account identification under another specified name
	// Max Length: 2048
	ProxyID string `json:"proxy_id,omitempty"`

	// Name of the identification scheme, in a coded form as published in an external list. When used proxy field must not be present
	// Max Length: 4
	ProxyIDCode *string `json:"proxy_id_code,omitempty"`
}

func BeneficiaryDebtorAccountProxyWithDefaults(defaults client.Defaults) *BeneficiaryDebtorAccountProxy {
	return &BeneficiaryDebtorAccountProxy{

		Proxy: defaults.GetStringPtr("BeneficiaryDebtorAccountProxy", "proxy"),

		ProxyID: defaults.GetString("BeneficiaryDebtorAccountProxy", "proxy_id"),

		ProxyIDCode: defaults.GetStringPtr("BeneficiaryDebtorAccountProxy", "proxy_id_code"),
	}
}

func (m *BeneficiaryDebtorAccountProxy) WithProxy(proxy string) *BeneficiaryDebtorAccountProxy {

	m.Proxy = &proxy

	return m
}

func (m *BeneficiaryDebtorAccountProxy) WithoutProxy() *BeneficiaryDebtorAccountProxy {
	m.Proxy = nil
	return m
}

func (m *BeneficiaryDebtorAccountProxy) WithProxyID(proxyID string) *BeneficiaryDebtorAccountProxy {

	m.ProxyID = proxyID

	return m
}

func (m *BeneficiaryDebtorAccountProxy) WithProxyIDCode(proxyIDCode string) *BeneficiaryDebtorAccountProxy {

	m.ProxyIDCode = &proxyIDCode

	return m
}

func (m *BeneficiaryDebtorAccountProxy) WithoutProxyIDCode() *BeneficiaryDebtorAccountProxy {
	m.ProxyIDCode = nil
	return m
}

// Validate validates this beneficiary debtor account proxy
func (m *BeneficiaryDebtorAccountProxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyIDCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BeneficiaryDebtorAccountProxy) validateProxy(formats strfmt.Registry) error {

	if swag.IsZero(m.Proxy) { // not required
		return nil
	}

	if err := validate.MaxLength("proxy", "body", string(*m.Proxy), 35); err != nil {
		return err
	}

	return nil
}

func (m *BeneficiaryDebtorAccountProxy) validateProxyID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProxyID) { // not required
		return nil
	}

	if err := validate.MaxLength("proxy_id", "body", string(m.ProxyID), 2048); err != nil {
		return err
	}

	return nil
}

func (m *BeneficiaryDebtorAccountProxy) validateProxyIDCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ProxyIDCode) { // not required
		return nil
	}

	if err := validate.MaxLength("proxy_id_code", "body", string(*m.ProxyIDCode), 4); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BeneficiaryDebtorAccountProxy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BeneficiaryDebtorAccountProxy) UnmarshalBinary(b []byte) error {
	var res BeneficiaryDebtorAccountProxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *BeneficiaryDebtorAccountProxy) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
