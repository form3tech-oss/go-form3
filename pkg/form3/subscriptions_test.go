package form3_test

import (
	"os"
	"testing"

	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateSubscriptions(t *testing.T) {
	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	organisationId := strfmt.UUID(uuid.MustParse(os.Getenv("FORM3_ORGANISATION_ID")).String())
	id := strfmt.UUID(uuid.New().String())
	transport := "http"
	callbackUri := "https://google.com/callback"
	eventType := "created"
	recordType := "PaymentSubmission"

	req := f3.Subscriptions.CreateSubscription()
	req.WithData(models.Subscription{
		ID:             &id,
		OrganisationID: &organisationId,
		Attributes: &models.SubscriptionAttributes{
			CallbackTransport: &transport,
			CallbackURI:       &callbackUri,
			UserID:            strfmt.UUID(uuid.New().String()),
			EventType:         &eventType,
			RecordType:        &recordType,
		},
	})

	resp, err := req.Do()
	require.NoError(t, err)
	assert.Equal(t, id.String(), resp.Data.ID.String())

	_, err = f3.Subscriptions.DeleteSubscription().WithID(id).WithVersion(0).Do()
	require.NoError(t, err)
}

func TestListSubscriptions(t *testing.T) {
	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	units, err := f3.Subscriptions.ListSubscriptions().Do()
	require.Nil(t, err)
	assert.NotNil(t, units.Data)
}
