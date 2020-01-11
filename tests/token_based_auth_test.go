package tests

import (
	"testing"

	"github.com/form3tech-oss/go-form3/pkg/form3"
	"github.com/stretchr/testify/require"
)

func TestTokenBasedAuth(t *testing.T) {
	f3 := form3.NewWithTokenBasedAuthFromEnv()

	req := f3.Users.GetUsersHealth()

	resp, err := req.Do()
	require.Nil(t, err)
	require.Equal(t, resp.Status, "up")
}
