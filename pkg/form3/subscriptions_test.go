package form3_test

import (
	"os"
	"testing"

	"github.com/form3tech-oss/go-form3/v6/pkg/form3"
	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateSubscriptionsDeprecatedCallbackUriParams(t *testing.T) {
	transport := models.CallbackTransportHTTP
	callbackURI := "https://webhook.site/db28eac0-d0a5-4078-b9be-0e4617471b01"
	eventType := "created"
	recordType := "PaymentSubmission"

	testCreateAndUpdateSubscriptions(t, &models.SubscriptionAttributes{
		CallbackTransport: transport,
		CallbackURI:       callbackURI,
		UserID:            strfmt.UUID(uuid.New().String()),
		EventType:         &eventType,
		RecordType:        &recordType,
	})
}

func TestCreateSubscriptionsNewCallbackUriParams(t *testing.T) {
	transport := models.CallbackTransportHTTP
	callbackURI := "https://webhook.site/db28eac0-d0a5-4078-b9be-0e4617471b01"
	eventType := "created"
	recordType := "PaymentSubmission"

	testCreateAndUpdateSubscriptions(t, &models.SubscriptionAttributes{
		CallbackUris: []*models.CallbackURI{{
			CallbackTransport: &transport,
			CallbackURI:       &callbackURI,
		}},
		UserID:     strfmt.UUID(uuid.New().String()),
		EventType:  &eventType,
		RecordType: &recordType,
	})
}

func TestCreateSubscriptionsNewCallbackUriParamsHttpAwsPrivate(t *testing.T) {
	transport := models.CallbackTransportHTTPAwsPrivate
	callbackURI := "https://webhook.site/db28eac0-d0a5-4078-b9be-0e4617471b01"
	eventType := "created"
	recordType := "PaymentSubmission"

	testCreateAndUpdateSubscriptions(t, &models.SubscriptionAttributes{
		CallbackUris: []*models.CallbackURI{{
			CallbackTransport: &transport,
			CallbackURI:       &callbackURI,
		}},
		UserID:     strfmt.UUID(uuid.New().String()),
		EventType:  &eventType,
		RecordType: &recordType,
	})
}

func testCreateAndUpdateSubscriptions(t *testing.T, attributes *models.SubscriptionAttributes) {
	skipWhenCredentialsMissing(t)

	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	organisationID := strfmt.UUID(uuid.MustParse(os.Getenv("FORM3_ORGANISATION_ID")).String())
	id := strfmt.UUID(uuid.New().String())

	// create
	req := f3.Subscriptions.CreateSubscription()
	req.WithData(models.Subscription{
		ID:             &id,
		OrganisationID: &organisationID,
		Attributes:     attributes,
	})
	resp, err := req.Do()
	require.NoError(t, err)
	assert.Equal(t, id.String(), resp.Data.ID.String())

	// update
	updateReq := f3.Subscriptions.ModifySubscription()
	version := int64(0)
	updateReq.
		WithID(id).
		WithData(models.SubscriptionUpdate{
			ID:             &id,
			OrganisationID: &organisationID,
			Attributes: &models.SubscriptionUpdateAttributes{
				Deactivated: true,
			},
			Version: &version,
		})
	updateResp, err := updateReq.Do()
	require.NoError(t, err)
	assert.Equal(t, id.String(), updateResp.Data.ID.String())
	assert.Equal(t, true, updateResp.Data.Attributes.Deactivated)

	//delete
	_, err = f3.Subscriptions.DeleteSubscription().WithID(id).WithVersion(version + 1).Do()
	require.NoError(t, err)
}

func TestListSubscriptions(t *testing.T) {
	skipWhenCredentialsMissing(t)

	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	units, err := f3.Subscriptions.ListSubscriptions().Do()
	require.Nil(t, err)
	assert.NotNil(t, units.Data)
}
