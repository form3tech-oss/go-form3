package client

import (
	"github.com/go-openapi/strfmt"
)

func NewEmptyDefaults() *EmptyDefaults {
	return &EmptyDefaults{}
}

type EmptyDefaults struct{}

func (*EmptyDefaults) GetStrfmtUUID(objectName, attributeName string) strfmt.UUID {
	var a strfmt.UUID

	return a
}

func (d *EmptyDefaults) GetStrfmtUUIDPtr(objectName, attributeName string) *strfmt.UUID {
	return nil
}

func (*EmptyDefaults) GetStrfmtDateTime(objectName, attributeName string) strfmt.DateTime {
	var a strfmt.DateTime

	return a
}

func (*EmptyDefaults) GetStrfmtDateTimePtr(objectName, attributeName string) *strfmt.DateTime {
	return nil
}

func (*EmptyDefaults) GetStrfmtDate(objectName, attributeName string) strfmt.Date {
	var a strfmt.Date

	return a
}

func (*EmptyDefaults) GetStrfmtDatePtr(objectName, attributeName string) *strfmt.Date {
	return nil
}

func (*EmptyDefaults) GetStrfmtURI(objectName, attributeName string) strfmt.URI {
	var a strfmt.URI

	return a
}

func (*EmptyDefaults) GetStrfmtURIPtr(objectName, attributeName string) *strfmt.URI {
	return nil
}

func (*EmptyDefaults) GetString(objectName, attributeName string) string {
	return ""
}

func (*EmptyDefaults) GetStringArray(objectName, attributeName string) []string {
	var a []string

	return a
}

func (*EmptyDefaults) GetStringPtr(objectName, attributeName string) *string {
	return nil
}

func (*EmptyDefaults) GetBool(objectName, attributeName string) bool {
	return true
}

func (*EmptyDefaults) GetBoolPtr(objectName, attributeName string) *bool {
	return nil
}

func (*EmptyDefaults) GetInt64(objectName, attributeName string) int64 {
	return int64(44)
}

func (*EmptyDefaults) GetInt64Ptr(objectName, attributeName string) *int64 {
	return nil
}

func (*EmptyDefaults) GetInt32(objectName, attributeName string) int32 {
	return int32(44)
}

func (*EmptyDefaults) GetFloat64(objectName, attributeName string) float64 {
	return float64(44)
}

func (*EmptyDefaults) GetFloat64Ptr(objectName, attributeName string) *float64 {
	return nil
}

func (*EmptyDefaults) GetMapStringInterface(objectName, attributeName string) map[string]interface{} {
	return nil
}
