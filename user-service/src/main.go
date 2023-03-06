package main

import (
	"os"

	"example.com/user-service/src/infra"
	"example.com/user-service/src/infra/rest"
	"example.com/user-service/src/infra/rest/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db := infra.Connect(os.Getenv("DB_URI"))
	infra.Migrate(db)

	restMiddleware := middleware.RestMiddleware{}
	hashService := infra.BcryptHashService{}
	authTokenService := infra.JwtAuthTokenService{}

	rest.ServeHTTP(db, restMiddleware, hashService, authTokenService)
}
