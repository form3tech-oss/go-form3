package form3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListSubscriptions(t *testing.T) {
	f3 := NewFromEnv()
	units, err := f3.Subscriptions.ListSubscriptions().Do()
	require.Nil(t, err)
	require.NotNil(t, units.Data)
}
