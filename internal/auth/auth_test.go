package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    req, err := http.NewRequest("GET", "/test", nil)
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Authorization", "ApiKey abcd1234")

    apiKey, err := GetAPIKey(req.Header)
    if err != nil {
        t.Fatal(err)
    }

    if apiKey != "abcd1234" {
        t.Errorf("Expected API key to be 'abcd1234', got %v", apiKey)
    }
}
