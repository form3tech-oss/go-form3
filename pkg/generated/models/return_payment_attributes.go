// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReturnPaymentAttributes return payment attributes
// swagger:model ReturnPaymentAttributes
type ReturnPaymentAttributes struct {

	// Block to represent a Financial Institution/agent in the payment chain
	Agents []*Agent `json:"agents,omitempty"`

	// The amount to return which should correspond to the amount of the original payment
	// Pattern: ^[0-9]{0,20}(?:\.[0-9]{1,10})?$
	Amount string `json:"amount,omitempty"`

	// Unique identifier for organisations collecting payments
	ClearingID string `json:"clearing_id,omitempty"`

	// ISO currency code for transaction amount
	Currency string `json:"currency,omitempty"`

	// Time a payment was released from being held due to a limit breach
	// Read Only: true
	// Format: date-time
	LimitBreachEndDatetime *strfmt.DateTime `json:"limit_breach_end_datetime,omitempty"`

	// Start time a payment was held due to a limit breach
	// Read Only: true
	// Format: date-time
	LimitBreachStartDatetime *strfmt.DateTime `json:"limit_breach_start_datetime,omitempty"`

	// reason
	Reason *string `json:"reason,omitempty"`

	// The return [reason code](http://draft-api-docs.form3.tech/api.html#enumerations-payment-return-codes)
	ReturnCode string `json:"return_code,omitempty"`

	// A unique reference to the return payment instruction. If not supplied, the value is generated by Form3.
	SchemeTransactionID string `json:"scheme_transaction_id,omitempty"`

	// settlement
	Settlement *Settlement `json:"settlement,omitempty"`

	// The scheme-specific unique transaction ID. Populated by the scheme.
	UniqueSchemeID string `json:"unique_scheme_id,omitempty"`

	// All purpose list of key-value pairs specific data stored on the return.
	// Max Items: 10
	UserDefinedData []*UserDefinedData `json:"user_defined_data,omitempty"`
}

func ReturnPaymentAttributesWithDefaults(defaults client.Defaults) *ReturnPaymentAttributes {
	return &ReturnPaymentAttributes{

		Agents: make([]*Agent, 0),

		Amount: defaults.GetString("ReturnPaymentAttributes", "amount"),

		ClearingID: defaults.GetString("ReturnPaymentAttributes", "clearing_id"),

		Currency: defaults.GetString("ReturnPaymentAttributes", "currency"),

		LimitBreachEndDatetime: defaults.GetStrfmtDateTimePtr("ReturnPaymentAttributes", "limit_breach_end_datetime"),

		LimitBreachStartDatetime: defaults.GetStrfmtDateTimePtr("ReturnPaymentAttributes", "limit_breach_start_datetime"),

		Reason: defaults.GetStringPtr("ReturnPaymentAttributes", "reason"),

		ReturnCode: defaults.GetString("ReturnPaymentAttributes", "return_code"),

		SchemeTransactionID: defaults.GetString("ReturnPaymentAttributes", "scheme_transaction_id"),

		Settlement: SettlementWithDefaults(defaults),

		UniqueSchemeID: defaults.GetString("ReturnPaymentAttributes", "unique_scheme_id"),

		UserDefinedData: make([]*UserDefinedData, 0),
	}
}

func (m *ReturnPaymentAttributes) WithAgents(agents []*Agent) *ReturnPaymentAttributes {

	m.Agents = agents

	return m
}

func (m *ReturnPaymentAttributes) WithAmount(amount string) *ReturnPaymentAttributes {

	m.Amount = amount

	return m
}

func (m *ReturnPaymentAttributes) WithClearingID(clearingID string) *ReturnPaymentAttributes {

	m.ClearingID = clearingID

	return m
}

func (m *ReturnPaymentAttributes) WithCurrency(currency string) *ReturnPaymentAttributes {

	m.Currency = currency

	return m
}

func (m *ReturnPaymentAttributes) WithLimitBreachEndDatetime(limitBreachEndDatetime strfmt.DateTime) *ReturnPaymentAttributes {

	m.LimitBreachEndDatetime = &limitBreachEndDatetime

	return m
}

func (m *ReturnPaymentAttributes) WithoutLimitBreachEndDatetime() *ReturnPaymentAttributes {
	m.LimitBreachEndDatetime = nil
	return m
}

func (m *ReturnPaymentAttributes) WithLimitBreachStartDatetime(limitBreachStartDatetime strfmt.DateTime) *ReturnPaymentAttributes {

	m.LimitBreachStartDatetime = &limitBreachStartDatetime

	return m
}

func (m *ReturnPaymentAttributes) WithoutLimitBreachStartDatetime() *ReturnPaymentAttributes {
	m.LimitBreachStartDatetime = nil
	return m
}

func (m *ReturnPaymentAttributes) WithReason(reason string) *ReturnPaymentAttributes {

	m.Reason = &reason

	return m
}

func (m *ReturnPaymentAttributes) WithoutReason() *ReturnPaymentAttributes {
	m.Reason = nil
	return m
}

func (m *ReturnPaymentAttributes) WithReturnCode(returnCode string) *ReturnPaymentAttributes {

	m.ReturnCode = returnCode

	return m
}

func (m *ReturnPaymentAttributes) WithSchemeTransactionID(schemeTransactionID string) *ReturnPaymentAttributes {

	m.SchemeTransactionID = schemeTransactionID

	return m
}

func (m *ReturnPaymentAttributes) WithSettlement(settlement Settlement) *ReturnPaymentAttributes {

	m.Settlement = &settlement

	return m
}

func (m *ReturnPaymentAttributes) WithoutSettlement() *ReturnPaymentAttributes {
	m.Settlement = nil
	return m
}

func (m *ReturnPaymentAttributes) WithUniqueSchemeID(uniqueSchemeID string) *ReturnPaymentAttributes {

	m.UniqueSchemeID = uniqueSchemeID

	return m
}

func (m *ReturnPaymentAttributes) WithUserDefinedData(userDefinedData []*UserDefinedData) *ReturnPaymentAttributes {

	m.UserDefinedData = userDefinedData

	return m
}

// Validate validates this return payment attributes
func (m *ReturnPaymentAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitBreachEndDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitBreachStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDefinedData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnPaymentAttributes) validateAgents(formats strfmt.Registry) error {

	if swag.IsZero(m.Agents) { // not required
		return nil
	}

	for i := 0; i < len(m.Agents); i++ {
		if swag.IsZero(m.Agents[i]) { // not required
			continue
		}

		if m.Agents[i] != nil {
			if err := m.Agents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReturnPaymentAttributes) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := validate.Pattern("amount", "body", string(m.Amount), `^[0-9]{0,20}(?:\.[0-9]{1,10})?$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnPaymentAttributes) validateLimitBreachEndDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitBreachEndDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("limit_breach_end_datetime", "body", "date-time", m.LimitBreachEndDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnPaymentAttributes) validateLimitBreachStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitBreachStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("limit_breach_start_datetime", "body", "date-time", m.LimitBreachStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnPaymentAttributes) validateSettlement(formats strfmt.Registry) error {

	if swag.IsZero(m.Settlement) { // not required
		return nil
	}

	if m.Settlement != nil {
		if err := m.Settlement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settlement")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnPaymentAttributes) validateUserDefinedData(formats strfmt.Registry) error {

	if swag.IsZero(m.UserDefinedData) { // not required
		return nil
	}

	iUserDefinedDataSize := int64(len(m.UserDefinedData))

	if err := validate.MaxItems("user_defined_data", "body", iUserDefinedDataSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m.UserDefinedData); i++ {
		if swag.IsZero(m.UserDefinedData[i]) { // not required
			continue
		}

		if m.UserDefinedData[i] != nil {
			if err := m.UserDefinedData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_defined_data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnPaymentAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnPaymentAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnPaymentAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnPaymentAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
