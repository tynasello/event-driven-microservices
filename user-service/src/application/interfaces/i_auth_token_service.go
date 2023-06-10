package interfaces

import "time"

type IAuthTokenService interface {
	GenerateToken(username string, expirationHours time.Duration) (generatedToken string, err error)
	ValidateToken(signedToken string) (username string, err error)
}
