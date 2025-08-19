package auth

import (
	"errors"
	"net/http"
	"testing"
)

// func GetAPIKey(headers http.Header) (string, error)
func TestGetAPIKey(t *testing.T) {

	goodHeader := http.Header{}
	badApiHeader := http.Header{}
	badValHeader := http.Header{}
	badJSONNameHeader := http.Header{}

	goodHeader.Set("Authorization", "ApiKey 1234567890")
	badApiHeader.Set("Authorization", "Authorization 1234567890")
	badValHeader.Set("Authorization", "")
	badJSONNameHeader.Set("Auth", "ApiKey 1234567890")

	// Should pass
	key, err := GetAPIKey(goodHeader)
	if err != nil {
		t.Errorf("key: %v  |  error: %v", key, err)
	}

	// Should fail due to incorrect "ApiKey" text
	key1, err := GetAPIKey(badApiHeader)
	if err == nil {
		t.Errorf("False pass bad ApiKey string!  | key: %v", key1)
	}

	// Should fail due to incorrect "ApiKey" text in header
	key2, err := GetAPIKey(badValHeader)
	if err == nil {
		t.Errorf("False pass no API key!  | key: %v", key2)
	}

	key3, err := GetAPIKey(badJSONNameHeader)
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("expected ErrNoAuthHeaderIncluded, got: %v  |  key: %v", err, key3)
	}
	/*
		if err == nil {
			t.Errorf("False pass JSON written incorrectly  | key: %v", key3)
		}
	*/
	t.Fatalf("This always fails!")
}
