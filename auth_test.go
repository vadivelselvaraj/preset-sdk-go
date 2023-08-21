package preset

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAccessToken(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful authentication response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"payload": {"access_token": "mockAccessToken"}}`))
	}))
	defer mockServer.Close()

	client := &PresetClient{
		BaseURL:    mockServer.URL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Auth: AuthStruct{
			TokenName: "mockTokenName",
			Secret:    "mockSecret",
		},
	}

	authResponse, err := client.GetAccessToken()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedToken := "mockAccessToken"
	if authResponse.Payload.AccessToken != expectedToken {
		t.Errorf("Expected access token %s, but got %s", expectedToken, authResponse.Payload.AccessToken)
	}
}

