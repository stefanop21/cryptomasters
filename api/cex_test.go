package api_test

import (
	"testing"

	"github.com/stefanop21/cryptomasters/api"
)

func TestApi(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Not error received")
	}
}
