package client

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type TestDefaults struct {
	EmptyDefaults
	OrganisationId *strfmt.UUID
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

func NewTestDefaults() *TestDefaults {
	return &TestDefaults{
		EmptyDefaults: *NewEmptyDefaults(),
	}
}
