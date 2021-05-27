package form3_test

import (
	"testing"

	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryPayments(t *testing.T) {
	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	_, err = f3.Payments.ListPayments().
		WithFilterAmount("2500.00").
		WithoutFilterOrganisationID().
		Do()
	assert.Nil(t, err)
}
