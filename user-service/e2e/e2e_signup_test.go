package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/user-service/src/infra/service"
)

func TestSignupEndpoint(t *testing.T) {

	e2eTestEnv := service.SetupE2eTestEnv(t)
	testServer := httptest.NewServer(e2eTestEnv.Router)
	defer testServer.Close()
	httpClient := &http.Client{}

	shouldSignupUser(t, e2eTestEnv, testServer, httpClient)
	shouldNotSignupUserWithExistingUsername(t, e2eTestEnv, testServer, httpClient)

}

func shouldSignupUser(t *testing.T, e2eTestEnv *service.E2eTestEnv, testServer *httptest.Server, client *http.Client) {

	requestBody := map[string]string{
		"username": "testusername",
		"password": "testpassword",
	}
	requestJsonBody, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest(http.MethodPost, testServer.URL+"/signup", bytes.NewBuffer(requestJsonBody))
	request.Header.Set("Content-Type", "application/json")

	response, responseErr := client.Do(request)
	if responseErr != nil {
		t.Errorf("Error making request: %v", responseErr)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, response.StatusCode)
	}

	accessTokenExists := false
	for _, cookie := range response.Cookies() {
		if cookie.Name == "access-token" && cookie.Value != "" {

			accessTokenExists = true
			break
		}
	}

	if !accessTokenExists {
		t.Errorf("Expected cookie '%s' not found in the response", "access-token")
	}

}

func shouldNotSignupUserWithExistingUsername(t *testing.T, e2eTestEnv *service.E2eTestEnv, testServer *httptest.Server, client *http.Client) {
	e2eTestEnv.SignupUseCase.Execute("testusername2", "testpassword")

	requestBody := map[string]string{
		"username": "testusername2",
		"password": "testpassword",
	}
	requestJsonBody, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest(http.MethodPost, testServer.URL+"/signup", bytes.NewBuffer(requestJsonBody))
	request.Header.Set("Content-Type", "application/json")

	response, responseErr := client.Do(request)
	if responseErr != nil {
		t.Errorf("Error making request: %v", responseErr)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, response.StatusCode)
	}

	accessTokenExists := false
	for _, cookie := range response.Cookies() {
		if cookie.Name == "access-token" && cookie.Value != "" {
			accessTokenExists = true
			break
		}
	}

	if accessTokenExists {
		t.Errorf("Unexpected cookie '%s' not found in the response", "access-token")
	}
}
