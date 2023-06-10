package e2e

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/user-service/src/infra/service"
)

func TestRouter(t *testing.T) {
	e2eTestEnv := service.SetupE2eTestEnv(t)
	testServer := httptest.NewServer(e2eTestEnv.Router)
	defer testServer.Close()

	shouldPongPing(t, testServer)
}

func shouldPongPing(t *testing.T, testServer *httptest.Server) {
	response, _ := http.Get(testServer.URL + "/")
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, response.StatusCode)
	}

	var responseBody map[string]interface{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	if message, ok := responseBody["message"].(string); !ok || message != "pong" {
		t.Errorf("Expected message %q but got %q", "pong", message)
	}
}
