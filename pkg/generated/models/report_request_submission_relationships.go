// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportRequestSubmissionRelationships report request submission relationships
// swagger:model ReportRequestSubmissionRelationships
type ReportRequestSubmissionRelationships struct {

	// report request
	ReportRequest *ReportRequestSubmissionRelationshipsReportRequest `json:"report_request,omitempty"`
}

func ReportRequestSubmissionRelationshipsWithDefaults(defaults client.Defaults) *ReportRequestSubmissionRelationships {
	return &ReportRequestSubmissionRelationships{

		ReportRequest: ReportRequestSubmissionRelationshipsReportRequestWithDefaults(defaults),
	}
}

func (m *ReportRequestSubmissionRelationships) WithReportRequest(reportRequest ReportRequestSubmissionRelationshipsReportRequest) *ReportRequestSubmissionRelationships {

	m.ReportRequest = &reportRequest

	return m
}

func (m *ReportRequestSubmissionRelationships) WithoutReportRequest() *ReportRequestSubmissionRelationships {
	m.ReportRequest = nil
	return m
}

// Validate validates this report request submission relationships
func (m *ReportRequestSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReportRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportRequestSubmissionRelationships) validateReportRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportRequest) { // not required
		return nil
	}

	if m.ReportRequest != nil {
		if err := m.ReportRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report_request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportRequestSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportRequestSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res ReportRequestSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportRequestSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReportRequestSubmissionRelationshipsReportRequest report request submission relationships report request
// swagger:model ReportRequestSubmissionRelationshipsReportRequest
type ReportRequestSubmissionRelationshipsReportRequest struct {

	// data
	// Required: true
	Data []*ReportRequest `json:"data"`
}

func ReportRequestSubmissionRelationshipsReportRequestWithDefaults(defaults client.Defaults) *ReportRequestSubmissionRelationshipsReportRequest {
	return &ReportRequestSubmissionRelationshipsReportRequest{

		Data: make([]*ReportRequest, 0),
	}
}

func (m *ReportRequestSubmissionRelationshipsReportRequest) WithData(data []*ReportRequest) *ReportRequestSubmissionRelationshipsReportRequest {

	m.Data = data

	return m
}

// Validate validates this report request submission relationships report request
func (m *ReportRequestSubmissionRelationshipsReportRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportRequestSubmissionRelationshipsReportRequest) validateData(formats strfmt.Registry) error {

	if err := validate.Required("report_request"+"."+"data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("report_request" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportRequestSubmissionRelationshipsReportRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportRequestSubmissionRelationshipsReportRequest) UnmarshalBinary(b []byte) error {
	var res ReportRequestSubmissionRelationshipsReportRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportRequestSubmissionRelationshipsReportRequest) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
