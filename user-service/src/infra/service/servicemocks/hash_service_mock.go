package servicemocks

import "errors"

type HashServiceMock struct{}

func (hs *HashServiceMock) Hash(password string) (string, error) {
	return "SOMEHASHTHATWILLNEVEREQUALAPASSWORD", nil
}

func (hs *HashServiceMock) ValidateHash(hashedValue string, rawValue string) error {
	if hashedValue != rawValue {
		return errors.New("Invalid hash")
	}
	return nil
}
