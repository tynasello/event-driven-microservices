package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAuthTokenService struct{}

var accessTokenSecret = []byte("secretKey")

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (j JwtAuthTokenService) GenerateToken(username string, expirationHours time.Duration) (string, error) {
	claims := JwtCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationHours)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenString, err := token.SignedString(accessTokenSecret)
	return signedTokenString, err
}

func (j JwtAuthTokenService) ValidateToken(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(accessTokenSecret), nil
		})
	if err != nil {
		return "", err
	}

	// check if token is signed and not expireed, and claims contains a username
	if claims, ok := token.Claims.(*JwtCustomClaims); !ok || !token.Valid || claims.Username == "" || claims.ExpiresAt.Time.Before(time.Now()) {

		return "", errors.New("Invalid token")
	} else {
		return claims.Username, nil
	}
}
