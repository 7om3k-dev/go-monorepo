package store_test

import (
	"testing"

	"github.com/7om3k/link-share/auth-server/store"
)

const (
	existingClientId    = "oauth-client-1"
	notExistingClientId = "oauth-client-999999"
)

func TestGetClientDataSuccessful(t *testing.T) {
	got := store.GetClientData(existingClientId)

	if got.IsZeroValue() {
		t.Errorf("store.GetClientData(%s) = %s; want non zero value ClientData", existingClientId, "zero value ClientData")
	}
}

func TestGetClientDataInvalidClientId(t *testing.T) {
	got := store.GetClientData(notExistingClientId)

	if !got.IsZeroValue() {
		t.Errorf("store.GetClientData(%s) = %s; want zero value ClientData", notExistingClientId, "non zero value ClientData")
	}
}
