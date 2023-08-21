package preset

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDoRequest(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      "mockAccessToken",
	}

	req, _ := http.NewRequest("GET", mockServer.URL+"/v1/test", nil)
	responseBody, err := client.doRequest(req, nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedBody := []byte(`{"message": "success"}`)
	if !bytes.Equal(responseBody, expectedBody) {
		t.Errorf("Expected response body %s, but got %s", expectedBody, responseBody)
	}
}

func TestNewClient(t *testing.T) {
	tokenName := "mockTokenName"
	secret := "mockSecret"

	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a 200 OK response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"payload":{"access_token":"mockAccessToken"}}`))
	}))
	defer mockServer.Close()

	baseURL := mockServer.URL
	client, err := NewClient(&baseURL, &tokenName, &secret)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if client.Token != "mockAccessToken" {
		t.Errorf("Expected token %s, but got %s", "mockAccessToken", client.Token)
	}

	if client.BaseURL != baseURL {
		t.Errorf("Expected base URL %s, but got %s", baseURL, client.BaseURL)
	}
}