package form3_test

import (
	"context"
	"github.com/form3tech-oss/go-form3/pkg/form3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetFXRates(t *testing.T) {
	f3 := form3.NewFromEnv()

	req := f3.FxAPI.GetFXRates()
	req.
		WithFilterPartyID(uuid.New().String()).
		WithFilterSourceCurrency("GBP").
		WithFilterSourceAmount("11.11").
		WithFilterTargetCurrency("EUR").
		WithFilterType("spot").
		WithContext(context.TODO())

	resp, err := req.Do()
	require.Nil(t, err)
	require.NotEmpty(t, resp.Data, "the response rates should not be empty")
	require.Equal(t, "11.11", resp.Data[0].Attributes.Source.Amount)
}
