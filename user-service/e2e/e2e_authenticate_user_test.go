package e2e

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/user-service/src/infra/service"
)

func TestAuthenticateUserEndpoint(t *testing.T) {

	e2eTestEnv := service.SetupE2eTestEnv(t)
	testServer := httptest.NewServer(e2eTestEnv.Router)
	defer testServer.Close()
	httpClient := &http.Client{}

	doesNotAuthenticateUserWithNoAccessTokenCookieInRequest(t, e2eTestEnv, testServer, httpClient)
	doesNotAuthenticateUserWithInvalidAccessTokenCookieInRequest(t, e2eTestEnv, testServer, httpClient)
	// doesNotAuthenticateUserWithExpiredAccessTokenCookieInRequest(t, e2eTestEnv, testServer, httpClient)
	authenticatesUserWithValidAccessToken(t, e2eTestEnv, testServer, httpClient)

}

func doesNotAuthenticateUserWithNoAccessTokenCookieInRequest(t *testing.T, e2eTestEnv *service.E2eTestEnv, testServer *httptest.Server, httpClient *http.Client) {
	request, _ := http.NewRequest(http.MethodGet, testServer.URL+"/authenticate-user", nil)
	response, responseErr := httpClient.Do(request)
	if responseErr != nil {
		t.Errorf("Error making request: %v", responseErr)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, response.StatusCode)
	}

	var responseBody map[string]interface{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	expectedMessage := "No access token cookie found"
	if message, ok := responseBody["message"].(string); !ok || message != expectedMessage {
		t.Errorf("Expected message %q but got %q", expectedMessage, message)
	}

}

func doesNotAuthenticateUserWithInvalidAccessTokenCookieInRequest(t *testing.T, e2eTestEnv *service.E2eTestEnv, testServer *httptest.Server, httpClient *http.Client) {
	request, _ := http.NewRequest(http.MethodGet, testServer.URL+"/authenticate-user", nil)
	request.AddCookie(&http.Cookie{Name: "access-token", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"})
	response, responseErr := httpClient.Do(request)
	if responseErr != nil {
		t.Errorf("Error making request: %v", responseErr)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, response.StatusCode)
	}
}

func doesNotAuthenticateUserWithExpiredAccessTokenCookieInRequest(t *testing.T, e2eTestEnv *service.E2eTestEnv, testServer *httptest.Server, httpClient *http.Client) {
	request, _ := http.NewRequest(http.MethodGet, testServer.URL+"/authenticate-user", nil)
	request.AddCookie(&http.Cookie{Name: "access-token", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFpb2ZiZXdvZmJvaWIiLCJleHAiOjE2ODY0MzYxMjR9.rvU-WYkyKzHkZei_CAuy51C_6t5aDyiwjccwurDgpUY"})
	response, _ := httpClient.Do(request)

	if response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, response.StatusCode)
	}
}

func authenticatesUserWithValidAccessToken(t *testing.T, e2eTestEnv *service.E2eTestEnv, testServer *httptest.Server, httpClient *http.Client) {
	createdUserResult := e2eTestEnv.SignupUseCase.Execute("testusername", "testpassword")
	createdUserResultingAccessToken, _ := createdUserResult.GetValue()

	request, _ := http.NewRequest(http.MethodGet, testServer.URL+"/authenticate-user", nil)
	request.AddCookie(&http.Cookie{Name: "access-token", Value: *createdUserResultingAccessToken})
	response, responseErr := httpClient.Do(request)

	if responseErr != nil {
		t.Errorf("Error making request: %v", responseErr)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, response.StatusCode)
	}

	var authenticateUserWithValidAccessTokenCookieResponseBody map[string]interface{}
	json.NewDecoder(response.Body).Decode(&authenticateUserWithValidAccessTokenCookieResponseBody)

	expectedUsername := "testusername"
	if username, ok := authenticateUserWithValidAccessTokenCookieResponseBody["username"].(string); !ok || username != expectedUsername {
		t.Errorf("Expected username %q but got %q", expectedUsername, username)
	}

}
