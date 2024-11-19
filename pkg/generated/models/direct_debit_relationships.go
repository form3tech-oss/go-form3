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
)

// DirectDebitRelationships direct debit relationships
// swagger:model DirectDebitRelationships
type DirectDebitRelationships struct {

	// direct debit admission
	DirectDebitAdmission *DirectDebitRelationshipsDirectDebitAdmission `json:"direct_debit_admission,omitempty"`

	// direct debit decision
	DirectDebitDecision *DirectDebitRelationshipsDirectDebitDecision `json:"direct_debit_decision,omitempty"`

	// direct debit recall
	DirectDebitRecall *DirectDebitRelationshipsDirectDebitRecall `json:"direct_debit_recall,omitempty"`

	// direct debit return
	DirectDebitReturn *DirectDebitRelationshipsDirectDebitReturn `json:"direct_debit_return,omitempty"`

	// direct debit reversal
	DirectDebitReversal *DirectDebitRelationshipsDirectDebitReversal `json:"direct_debit_reversal,omitempty"`

	// direct debit submission
	DirectDebitSubmission *DirectDebitRelationshipsDirectDebitSubmission `json:"direct_debit_submission,omitempty"`

	// mandate
	Mandate *DirectDebitRelationshipsMandate `json:"mandate,omitempty"`
}

func DirectDebitRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitRelationships {
	return &DirectDebitRelationships{

		DirectDebitAdmission: DirectDebitRelationshipsDirectDebitAdmissionWithDefaults(defaults),

		DirectDebitDecision: DirectDebitRelationshipsDirectDebitDecisionWithDefaults(defaults),

		DirectDebitRecall: DirectDebitRelationshipsDirectDebitRecallWithDefaults(defaults),

		DirectDebitReturn: DirectDebitRelationshipsDirectDebitReturnWithDefaults(defaults),

		DirectDebitReversal: DirectDebitRelationshipsDirectDebitReversalWithDefaults(defaults),

		DirectDebitSubmission: DirectDebitRelationshipsDirectDebitSubmissionWithDefaults(defaults),

		Mandate: DirectDebitRelationshipsMandateWithDefaults(defaults),
	}
}

func (m *DirectDebitRelationships) WithDirectDebitAdmission(directDebitAdmission DirectDebitRelationshipsDirectDebitAdmission) *DirectDebitRelationships {

	m.DirectDebitAdmission = &directDebitAdmission

	return m
}

func (m *DirectDebitRelationships) WithoutDirectDebitAdmission() *DirectDebitRelationships {
	m.DirectDebitAdmission = nil
	return m
}

func (m *DirectDebitRelationships) WithDirectDebitDecision(directDebitDecision DirectDebitRelationshipsDirectDebitDecision) *DirectDebitRelationships {

	m.DirectDebitDecision = &directDebitDecision

	return m
}

func (m *DirectDebitRelationships) WithoutDirectDebitDecision() *DirectDebitRelationships {
	m.DirectDebitDecision = nil
	return m
}

func (m *DirectDebitRelationships) WithDirectDebitRecall(directDebitRecall DirectDebitRelationshipsDirectDebitRecall) *DirectDebitRelationships {

	m.DirectDebitRecall = &directDebitRecall

	return m
}

func (m *DirectDebitRelationships) WithoutDirectDebitRecall() *DirectDebitRelationships {
	m.DirectDebitRecall = nil
	return m
}

func (m *DirectDebitRelationships) WithDirectDebitReturn(directDebitReturn DirectDebitRelationshipsDirectDebitReturn) *DirectDebitRelationships {

	m.DirectDebitReturn = &directDebitReturn

	return m
}

func (m *DirectDebitRelationships) WithoutDirectDebitReturn() *DirectDebitRelationships {
	m.DirectDebitReturn = nil
	return m
}

func (m *DirectDebitRelationships) WithDirectDebitReversal(directDebitReversal DirectDebitRelationshipsDirectDebitReversal) *DirectDebitRelationships {

	m.DirectDebitReversal = &directDebitReversal

	return m
}

func (m *DirectDebitRelationships) WithoutDirectDebitReversal() *DirectDebitRelationships {
	m.DirectDebitReversal = nil
	return m
}

func (m *DirectDebitRelationships) WithDirectDebitSubmission(directDebitSubmission DirectDebitRelationshipsDirectDebitSubmission) *DirectDebitRelationships {

	m.DirectDebitSubmission = &directDebitSubmission

	return m
}

func (m *DirectDebitRelationships) WithoutDirectDebitSubmission() *DirectDebitRelationships {
	m.DirectDebitSubmission = nil
	return m
}

func (m *DirectDebitRelationships) WithMandate(mandate DirectDebitRelationshipsMandate) *DirectDebitRelationships {

	m.Mandate = &mandate

	return m
}

func (m *DirectDebitRelationships) WithoutMandate() *DirectDebitRelationships {
	m.Mandate = nil
	return m
}

// Validate validates this direct debit relationships
func (m *DirectDebitRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebitAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitDecision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitRecall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReturn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitSubmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMandate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRelationships) validateDirectDebitAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitAdmission) { // not required
		return nil
	}

	if m.DirectDebitAdmission != nil {
		if err := m.DirectDebitAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_debit_admission")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitRelationships) validateDirectDebitDecision(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitDecision) { // not required
		return nil
	}

	if m.DirectDebitDecision != nil {
		if err := m.DirectDebitDecision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_debit_decision")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitRelationships) validateDirectDebitRecall(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitRecall) { // not required
		return nil
	}

	if m.DirectDebitRecall != nil {
		if err := m.DirectDebitRecall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_debit_recall")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitRelationships) validateDirectDebitReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReturn) { // not required
		return nil
	}

	if m.DirectDebitReturn != nil {
		if err := m.DirectDebitReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_debit_return")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitRelationships) validateDirectDebitReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReversal) { // not required
		return nil
	}

	if m.DirectDebitReversal != nil {
		if err := m.DirectDebitReversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_debit_reversal")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitRelationships) validateDirectDebitSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitSubmission) { // not required
		return nil
	}

	if m.DirectDebitSubmission != nil {
		if err := m.DirectDebitSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_debit_submission")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitRelationships) validateMandate(formats strfmt.Registry) error {

	if swag.IsZero(m.Mandate) { // not required
		return nil
	}

	if m.Mandate != nil {
		if err := m.Mandate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mandate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRelationshipsDirectDebitAdmission direct debit relationships direct debit admission
// swagger:model DirectDebitRelationshipsDirectDebitAdmission
type DirectDebitRelationshipsDirectDebitAdmission struct {

	// data
	Data []*DirectDebitAdmission `json:"data"`
}

func DirectDebitRelationshipsDirectDebitAdmissionWithDefaults(defaults client.Defaults) *DirectDebitRelationshipsDirectDebitAdmission {
	return &DirectDebitRelationshipsDirectDebitAdmission{

		Data: make([]*DirectDebitAdmission, 0),
	}
}

func (m *DirectDebitRelationshipsDirectDebitAdmission) WithData(data []*DirectDebitAdmission) *DirectDebitRelationshipsDirectDebitAdmission {

	m.Data = data

	return m
}

// Validate validates this direct debit relationships direct debit admission
func (m *DirectDebitRelationshipsDirectDebitAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRelationshipsDirectDebitAdmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("direct_debit_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitAdmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitRelationshipsDirectDebitAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRelationshipsDirectDebitAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRelationshipsDirectDebitDecision direct debit relationships direct debit decision
// swagger:model DirectDebitRelationshipsDirectDebitDecision
type DirectDebitRelationshipsDirectDebitDecision struct {

	// data
	Data []*DirectDebitDecision `json:"data"`
}

func DirectDebitRelationshipsDirectDebitDecisionWithDefaults(defaults client.Defaults) *DirectDebitRelationshipsDirectDebitDecision {
	return &DirectDebitRelationshipsDirectDebitDecision{

		Data: make([]*DirectDebitDecision, 0),
	}
}

func (m *DirectDebitRelationshipsDirectDebitDecision) WithData(data []*DirectDebitDecision) *DirectDebitRelationshipsDirectDebitDecision {

	m.Data = data

	return m
}

// Validate validates this direct debit relationships direct debit decision
func (m *DirectDebitRelationshipsDirectDebitDecision) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRelationshipsDirectDebitDecision) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("direct_debit_decision" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitDecision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitDecision) UnmarshalBinary(b []byte) error {
	var res DirectDebitRelationshipsDirectDebitDecision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRelationshipsDirectDebitDecision) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRelationshipsDirectDebitRecall direct debit relationships direct debit recall
// swagger:model DirectDebitRelationshipsDirectDebitRecall
type DirectDebitRelationshipsDirectDebitRecall struct {

	// data
	Data []*DirectDebitRecall `json:"data"`
}

func DirectDebitRelationshipsDirectDebitRecallWithDefaults(defaults client.Defaults) *DirectDebitRelationshipsDirectDebitRecall {
	return &DirectDebitRelationshipsDirectDebitRecall{

		Data: make([]*DirectDebitRecall, 0),
	}
}

func (m *DirectDebitRelationshipsDirectDebitRecall) WithData(data []*DirectDebitRecall) *DirectDebitRelationshipsDirectDebitRecall {

	m.Data = data

	return m
}

// Validate validates this direct debit relationships direct debit recall
func (m *DirectDebitRelationshipsDirectDebitRecall) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRelationshipsDirectDebitRecall) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("direct_debit_recall" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitRecall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitRecall) UnmarshalBinary(b []byte) error {
	var res DirectDebitRelationshipsDirectDebitRecall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRelationshipsDirectDebitRecall) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRelationshipsDirectDebitReturn direct debit relationships direct debit return
// swagger:model DirectDebitRelationshipsDirectDebitReturn
type DirectDebitRelationshipsDirectDebitReturn struct {

	// data
	Data []*DirectDebitReturn `json:"data"`
}

func DirectDebitRelationshipsDirectDebitReturnWithDefaults(defaults client.Defaults) *DirectDebitRelationshipsDirectDebitReturn {
	return &DirectDebitRelationshipsDirectDebitReturn{

		Data: make([]*DirectDebitReturn, 0),
	}
}

func (m *DirectDebitRelationshipsDirectDebitReturn) WithData(data []*DirectDebitReturn) *DirectDebitRelationshipsDirectDebitReturn {

	m.Data = data

	return m
}

// Validate validates this direct debit relationships direct debit return
func (m *DirectDebitRelationshipsDirectDebitReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRelationshipsDirectDebitReturn) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("direct_debit_return" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitReturn) UnmarshalBinary(b []byte) error {
	var res DirectDebitRelationshipsDirectDebitReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRelationshipsDirectDebitReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRelationshipsDirectDebitReversal direct debit relationships direct debit reversal
// swagger:model DirectDebitRelationshipsDirectDebitReversal
type DirectDebitRelationshipsDirectDebitReversal struct {

	// data
	Data []*DirectDebitReversal `json:"data"`
}

func DirectDebitRelationshipsDirectDebitReversalWithDefaults(defaults client.Defaults) *DirectDebitRelationshipsDirectDebitReversal {
	return &DirectDebitRelationshipsDirectDebitReversal{

		Data: make([]*DirectDebitReversal, 0),
	}
}

func (m *DirectDebitRelationshipsDirectDebitReversal) WithData(data []*DirectDebitReversal) *DirectDebitRelationshipsDirectDebitReversal {

	m.Data = data

	return m
}

// Validate validates this direct debit relationships direct debit reversal
func (m *DirectDebitRelationshipsDirectDebitReversal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRelationshipsDirectDebitReversal) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("direct_debit_reversal" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitReversal) UnmarshalBinary(b []byte) error {
	var res DirectDebitRelationshipsDirectDebitReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRelationshipsDirectDebitReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRelationshipsDirectDebitSubmission The submission resource related to the direct debit
// swagger:model DirectDebitRelationshipsDirectDebitSubmission
type DirectDebitRelationshipsDirectDebitSubmission struct {

	// data
	Data []*DirectDebitSubmission `json:"data"`
}

func DirectDebitRelationshipsDirectDebitSubmissionWithDefaults(defaults client.Defaults) *DirectDebitRelationshipsDirectDebitSubmission {
	return &DirectDebitRelationshipsDirectDebitSubmission{

		Data: make([]*DirectDebitSubmission, 0),
	}
}

func (m *DirectDebitRelationshipsDirectDebitSubmission) WithData(data []*DirectDebitSubmission) *DirectDebitRelationshipsDirectDebitSubmission {

	m.Data = data

	return m
}

// Validate validates this direct debit relationships direct debit submission
func (m *DirectDebitRelationshipsDirectDebitSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRelationshipsDirectDebitSubmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("direct_debit_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRelationshipsDirectDebitSubmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitRelationshipsDirectDebitSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRelationshipsDirectDebitSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRelationshipsMandate direct debit relationships mandate
// swagger:model DirectDebitRelationshipsMandate
type DirectDebitRelationshipsMandate struct {

	// data
	Data []*DirectDebitMandate `json:"data"`
}

func DirectDebitRelationshipsMandateWithDefaults(defaults client.Defaults) *DirectDebitRelationshipsMandate {
	return &DirectDebitRelationshipsMandate{

		Data: make([]*DirectDebitMandate, 0),
	}
}

func (m *DirectDebitRelationshipsMandate) WithData(data []*DirectDebitMandate) *DirectDebitRelationshipsMandate {

	m.Data = data

	return m
}

// Validate validates this direct debit relationships mandate
func (m *DirectDebitRelationshipsMandate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRelationshipsMandate) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("mandate" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRelationshipsMandate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRelationshipsMandate) UnmarshalBinary(b []byte) error {
	var res DirectDebitRelationshipsMandate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRelationshipsMandate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
