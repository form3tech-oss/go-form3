package tests

import (
	"github.com/form3tech-oss/go-form3/pkg/form3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueryPayments(t *testing.T) {
	f3 := form3.New()
	_, err := f3.Payments.ListPayments().
		WithFilterAmount("2500.00").
		WithoutFilterOrganisationID().
		Do()
	require.Nil(t, err)
}
