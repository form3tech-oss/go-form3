package tests

import (
	"testing"

	"github.com/form3tech-oss/go-form3/pkg/form3"
	"github.com/stretchr/testify/require"
)

func TestListSubscriptions(t *testing.T) {
	f3 := form3.NewFromEnv()
	units, err := f3.Subscriptions.ListSubscriptions().Do()
	require.Nil(t, err)
	require.NotNil(t, units.Data)
}
