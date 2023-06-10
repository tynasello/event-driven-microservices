package tests

import (
	"testing"

	"example.com/user-service/src/infra/service"
)

func TestHash(t *testing.T) {
	hashService := service.BcryptHashService{}
	hash, err := hashService.Hash("test")
	if err != nil {
		t.Errorf("Error while hashing: %v", err)
	}
	if hash == "" {
		t.Errorf("Hashed value is empty")
	}
}

func TestValidateHash(t *testing.T) {
	hashService := service.BcryptHashService{}

	testCases := []struct {
		hashedValue string
		rawValue    string
		shouldPass  bool
	}{
		{
			hashedValue: func() string {
				value, _ := hashService.Hash("test")
				return value
			}(),
			rawValue:   "test",
			shouldPass: true},
		{
			hashedValue: func() string {
				value, _ := hashService.Hash("test1234")
				return value
			}(),
			rawValue:   "test1234",
			shouldPass: true,
		},
		{
			hashedValue: func() string {
				value, _ := hashService.Hash("test")
				return value
			}(),
			rawValue:   "test1234",
			shouldPass: false,
		},
	}

	for _, tc := range testCases {
		err := hashService.ValidateHash(tc.hashedValue, tc.rawValue)
		if (err == nil) != tc.shouldPass {
			if tc.shouldPass {
				t.Errorf("Expected pass but got error for hashed value:%s, and raw value: %s", tc.hashedValue, tc.rawValue)
			} else {
				t.Errorf("Expected error but got pass for hashed value: %s, and raw value: %s", tc.hashedValue, tc.rawValue)
			}
		}
	}
}
