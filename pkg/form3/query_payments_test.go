package form3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryPayments(t *testing.T) {
	f3 := NewFromEnv()
	_, err := f3.Payments.ListPayments().
		WithFilterAmount("2500.00").
		WithoutFilterOrganisationID().
		Do()
	require.Nil(t, err)
}
