package servicemocks

import "time"

type JwtAuthTokenServiceMock struct{}

func (jm *JwtAuthTokenServiceMock) GenerateToken(username string, expirationHours time.Duration) (string, error) {
	return username, nil
}

func (jm *JwtAuthTokenServiceMock) ValidateToken(username string) (string, error) {
	return username, nil
}
