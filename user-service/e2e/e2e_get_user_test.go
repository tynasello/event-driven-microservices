package e2e

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/user-service/src/infra/service"
)

func TestGetUserEndpoint(t *testing.T) {

	e2eTestEnv := service.SetupE2eTestEnv(t)
	testServer := httptest.NewServer(e2eTestEnv.Router)
	defer testServer.Close()
	client := &http.Client{}

	getsExistingUser(t, e2eTestEnv, testServer, client)

}

func getsExistingUser(t *testing.T, e2eTestEnv *service.E2eTestEnv, testServer *httptest.Server, client *http.Client) {

	createsUserResult := e2eTestEnv.SignupUseCase.Execute("testusername", "testpassword")
	createdUserAccessToken, _ := createsUserResult.GetValue()

	request, _ := http.NewRequest(http.MethodGet, testServer.URL+"/get-user", nil)
	accessTokenCookie := &http.Cookie{
		Name:  "access-token",
		Value: *createdUserAccessToken,
	}
	request.AddCookie(accessTokenCookie)

	response, responseErr := client.Do(request)
	if responseErr != nil {
		t.Errorf("Error making request: %v", responseErr)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, response.StatusCode)
	}

	var getUserResponseBody map[string]interface{}
	json.NewDecoder(response.Body).Decode(&getUserResponseBody)

	// Check the presence of the "user" field
	userField, ok := getUserResponseBody["user"]
	if !ok {
		t.Error("Expected field 'user' not found in the response")
	}

	userMap, _ := userField.(map[string]interface{})
	usernameField, ok := userMap["Username"]
	if !ok {
		t.Error("Expected field 'username' not found in the 'user' field")
	}
	username, _ := usernameField.(string)
	if username != "testusername" {
		t.Errorf("Expected username to be %s, but got %s", "testusername", username)
	}

	idField, ok := userMap["Id"]
	if !ok {
		t.Error("Expected field 'id' not found in the 'user' field")
	}
	t.Log(idField)
	if idField == "" {
		t.Error("Expected id to be not empty")
	}

	passwordField, ok := userMap["Password"]
	if !ok {
		t.Error("Expected field 'password' not found in the 'user' field")
	}
	password, _ := passwordField.(string)
	if password == "" {
		t.Error("Expected password to not be empty")
	}

}
