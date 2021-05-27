package form3_test

import (
	"os"
	"testing"

	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenBasedAuth(t *testing.T) {
	if os.Getenv("FORM3_CLIENT_ID") == "" || os.Getenv("FORM3_CLIENT_SECRET") == "" {
		t.Skip("WARN: FORM3_CLIENT_ID or FORM3_CLIENT_SECRET not set, skipping test")
	}

	f3, err := form3.NewWithTokenBasedAuthFromEnv()
	require.NoError(t, err)

	req := f3.Users.GetUser()
	req.UserID = strfmt.UUID(os.Getenv("FORM3_CLIENT_ID"))

	resp, err := req.Do()
	require.Nil(t, err)
	assert.Equal(t, resp.Data.ID.String(), os.Getenv("FORM3_CLIENT_ID"))
}
