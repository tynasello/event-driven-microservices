package infra

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptHashService struct{}

func (hs BcryptHashService) Hash(value string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	if err != nil {
		return "", err
	}
	hashedValue := string(hashedBytes)
	return hashedValue, nil
}

func (hs BcryptHashService) ValidateHash(value string, providedValue string) error {
	err := bcrypt.CompareHashAndPassword([]byte(value), []byte(providedValue))
	if err != nil {
		return err
	}
	return nil
}
