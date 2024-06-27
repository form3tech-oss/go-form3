package form3_test

import (
	"os"
	"testing"
)

func skipWhenCredentialsMissing(t *testing.T) {
	if os.Getenv("FORM3_CLIENT_ID") == "" || os.Getenv("FORM3_CLIENT_SECRET") == "" {
		t.Skip("WARN: FORM3_CLIENT_ID or FORM3_CLIENT_SECRET not set, skipping test")
	}
}
