package store_test

import (
	"testing"

	"github.com/7om3k-dev/OAuth/internal/store"
)

const (
	existingClientId    = "oauth-client-1"
	notExistingClientId = "oauth-client-1000000"
)

func TestGetClientDataSuccessful(t *testing.T) {
	_, err := store.GetClientData(existingClientId)

	if err != nil {
		t.Errorf("store.GetClientData(%s) = %s; want ClientData", existingClientId, "client not found error")
	}
}

func TestGetClientDataInvalidClientId(t *testing.T) {
	_, err := store.GetClientData(notExistingClientId)

	if err == nil {
		t.Errorf("store.GetClientData(%s) = %s; want client not found error", notExistingClientId, "nil error")
	}
}
