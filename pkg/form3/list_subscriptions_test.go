package form3_test

import (
	"testing"

	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListSubscriptions(t *testing.T) {
	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	units, err := f3.Subscriptions.ListSubscriptions().Do()
	require.Nil(t, err)
	assert.NotNil(t, units.Data)
}
