package publicapi_test

import (
	"testing"

	"github.com/jh1104/publicapi"
)

func TestNewClient(t *testing.T) {
	client := publicapi.NewClient("key")
	if client.HTTPClient == nil {
		t.Error("want HTTPClient to be set, got nil")
	}
}
