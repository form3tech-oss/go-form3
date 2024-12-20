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

// ReportRequestLink report request link
// swagger:model ReportRequestLink
type ReportRequestLink struct {

	// href
	// Required: true
	// Format: uri
	Href *strfmt.URI `json:"href"`

	// meta
	Meta *ReportRequestLinkMeta `json:"meta,omitempty"`
}

func ReportRequestLinkWithDefaults(defaults client.Defaults) *ReportRequestLink {
	return &ReportRequestLink{

		Href: defaults.GetStrfmtURIPtr("ReportRequestLink", "href"),

		Meta: ReportRequestLinkMetaWithDefaults(defaults),
	}
}

func (m *ReportRequestLink) WithHref(href strfmt.URI) *ReportRequestLink {

	m.Href = &href

	return m
}

func (m *ReportRequestLink) WithoutHref() *ReportRequestLink {
	m.Href = nil
	return m
}

func (m *ReportRequestLink) WithMeta(meta ReportRequestLinkMeta) *ReportRequestLink {

	m.Meta = &meta

	return m
}

func (m *ReportRequestLink) WithoutMeta() *ReportRequestLink {
	m.Meta = nil
	return m
}

// Validate validates this report request link
func (m *ReportRequestLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportRequestLink) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	if err := validate.FormatOf("href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportRequestLink) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportRequestLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportRequestLink) UnmarshalBinary(b []byte) error {
	var res ReportRequestLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportRequestLink) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReportRequestLinkMeta report request link meta
// swagger:model ReportRequestLinkMeta
type ReportRequestLinkMeta struct {

	// content type
	ContentType string `json:"content-type,omitempty"`
}

func ReportRequestLinkMetaWithDefaults(defaults client.Defaults) *ReportRequestLinkMeta {
	return &ReportRequestLinkMeta{

		ContentType: defaults.GetString("ReportRequestLinkMeta", "content-type"),
	}
}

func (m *ReportRequestLinkMeta) WithContentType(contentType string) *ReportRequestLinkMeta {

	m.ContentType = contentType

	return m
}

// Validate validates this report request link meta
func (m *ReportRequestLinkMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportRequestLinkMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportRequestLinkMeta) UnmarshalBinary(b []byte) error {
	var res ReportRequestLinkMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportRequestLinkMeta) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
