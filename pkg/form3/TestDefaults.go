package form3

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type TestDefaults struct {
	OrganisationId *strfmt.UUID
}

func (*TestDefaults) GetStrfmtUUID(objectName, attributeName string) strfmt.UUID {
	var a strfmt.UUID
	return a
}

func (d *TestDefaults) GetStrfmtUUIDPtr(objectName, attributeName string) *strfmt.UUID {
	if "organisation_id" == attributeName {
		return d.OrganisationId
	}
	if "id" == attributeName {
		u := strfmt.UUID(uuid.Must(uuid.NewRandom()).String())
		return &u
	}
	return nil
}

func (*TestDefaults) GetStrfmtDateTime(objectName, attributeName string) strfmt.DateTime {
	var a strfmt.DateTime
	return a
}

func (*TestDefaults) GetStrfmtDateTimePtr(objectName, attributeName string) *strfmt.DateTime {
	return nil
}

func (*TestDefaults) GetStrfmtDate(objectName, attributeName string) strfmt.Date {
	var a strfmt.Date
	return a
}

func (*TestDefaults) GetStrfmtDatePtr(objectName, attributeName string) *strfmt.Date {
	return nil
}

func (*TestDefaults) GetStrfmtURI(objectName, attributeName string) strfmt.URI {
	var a strfmt.URI
	return a
}

func (*TestDefaults) GetStrfmtURIPtr(objectName, attributeName string) *strfmt.URI {
	return nil
}

func (*TestDefaults) GetString(objectName, attributeName string) string {
	return ""
}

func (*TestDefaults) GetStringArray(objectName, attributeName string) []string {
	var a []string
	return a
}

func (*TestDefaults) GetStringPtr(objectName, attributeName string) *string {
	return nil
}

func (*TestDefaults) GetBool(objectName, attributeName string) bool {
	return true
}

func (*TestDefaults) GetBoolPtr(objectName, attributeName string) *bool {
	return nil
}

func (*TestDefaults) GetInt64(objectName, attributeName string) int64 {
	return int64(44)
}

func (*TestDefaults) GetInt64Ptr(objectName, attributeName string) *int64 {
	return nil
}

func (*TestDefaults) GetInt32(objectName, attributeName string) int32 {
	return int32(44)
}

func NewTestDefaults() *TestDefaults {
	return &TestDefaults{}
}
