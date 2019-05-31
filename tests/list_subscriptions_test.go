package tests

import (
	"github.com/form3tech-oss/go-form3/pkg/form3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListSubscriptions(t *testing.T) {
	f3 := form3.New()
	units, err := f3.Subscriptions.ListSubscriptions().Do()
	require.Nil(t, err)
	require.NotNil(t, units.Data)
}
