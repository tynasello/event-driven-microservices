package infra

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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationHours * time.Hour)),
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

	// check if token is not expired and claims contains a username
	if claims, ok := token.Claims.(*JwtCustomClaims); !ok || !token.Valid || claims.Username == "" {
		return "", errors.New("Invalid token")
	} else {
		return claims.Username, nil
	}
}
