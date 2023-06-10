package tests

import (
	"testing"
	"time"

	"example.com/user-service/src/infra/service"
)

func TestGenerateToken(t *testing.T) {
	accessTokenService := service.JwtAuthTokenService{}

	testCases := []struct {
		username        string
		expirationHours time.Duration
		returnsError    bool
	}{
		{
			username:        "test",
			expirationHours: 1 * time.Hour,
			returnsError:    false,
		},
		{
			username:        "",
			expirationHours: 1 * time.Hour,
			returnsError:    false,
		},
	}

	for _, tc := range testCases {
		signedToken, err := accessTokenService.GenerateToken(tc.username, tc.expirationHours)
		if tc.returnsError && err == nil {
			t.Errorf("Expected error but got nil")
		} else if !tc.returnsError && err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		if !tc.returnsError {
			username, err := accessTokenService.ValidateToken(signedToken)
			t.Log(err)
			if username != tc.username {
				t.Errorf("Expected token to be signed with username %v but got %v", tc.username, username)
			}
		}
	}
}

func TestValidateToken(t *testing.T) {
	accessTokenService := service.JwtAuthTokenService{}

	testCases := []struct {
		signedToken      string
		returnsError     bool
		expectedUsername string
	}{
		{
			signedToken: func() string {
				token, _ := accessTokenService.GenerateToken("test", 1*time.Hour)
				return token
			}(),
			returnsError:     false,
			expectedUsername: "test",
		},
		{
			signedToken: func() string {
				token, _ := accessTokenService.GenerateToken("", 1*time.Hour)
				return token
			}(),
			returnsError:     true,
			expectedUsername: "",
		},
		{
			signedToken: func() string {
				token, _ := accessTokenService.GenerateToken("test", 0*time.Hour)
				return token
			}(),
			returnsError:     true,
			expectedUsername: "",
		},
	}

	for _, tc := range testCases {
		username, err := accessTokenService.ValidateToken(tc.signedToken)
		if tc.returnsError && err == nil {
			t.Errorf("Expected error but got nil")
		} else if !tc.returnsError && err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		if !tc.returnsError {
			if username != tc.expectedUsername {
				t.Errorf("Expected username to be %v but got %v", tc.expectedUsername, username)
			}
		}
	}
}
