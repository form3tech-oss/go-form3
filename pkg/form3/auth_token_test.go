package form3_test

import (
	"net/url"
	"os"
	"testing"

	"github.com/form3tech-oss/go-form3/v4/pkg/form3"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenAuth(t *testing.T) {
	if os.Getenv("FORM3_CLIENT_ID") == "" || os.Getenv("FORM3_CLIENT_SECRET") == "" {
		t.Skip("WARN: FORM3_CLIENT_ID or FORM3_CLIENT_SECRET not set, skipping test")
	}

	u, err := uuid.Parse(os.Getenv("FORM3_CLIENT_ID"))
	require.NoError(t, err)

	baseURL, err := url.Parse(os.Getenv("FORM3_HOST"))
	require.NoError(t, err)

	f3, err := form3.New(
		form3.WithBaseURL(*baseURL),
		form3.WithTokenTransport(
			form3.WithClientID(u),
			form3.WithClientSecret(os.Getenv("FORM3_CLIENT_SECRET"))))
	require.NoError(t, err)

	req := f3.Users.GetUser()
	req.UserID = strfmt.UUID(os.Getenv("FORM3_CLIENT_ID"))

	resp, err := req.Do()

	require.Nil(t, err)
	assert.Equal(t, resp.Data.ID.String(), os.Getenv("FORM3_CLIENT_ID"))
}
