package main

import (
	"example.com/user-service/src/application/usecase"
	"example.com/user-service/src/infra/repository"
	"example.com/user-service/src/infra/rest"
	"example.com/user-service/src/infra/rest/controller"
	"example.com/user-service/src/infra/rest/middleware"
	"example.com/user-service/src/infra/service"
	"github.com/joho/godotenv"
)

var envFilePath string

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbService := service.NewDbService()
	dbService.RunDbMigrations()

	hashService := &service.BcryptHashService{}
	authTokenService := &service.JwtAuthTokenService{}

	UserRepository := &repository.UserRepository{Db: dbService.Db}

	SignupUseCase := &usecase.SignupUseCase{UserRepository: UserRepository, HashService: hashService, AuthTokenService: authTokenService}
	LoginUseCase := &usecase.LoginUseCase{UserRepository: UserRepository, HashService: hashService, AuthTokenService: authTokenService}
	GetUserUseCase := &usecase.GetUserUseCase{UserRepository: UserRepository}

	RestMiddleware := &middleware.RestMiddleware{AuthTokenService: authTokenService}
	UserController := &controller.UserController{SignupUseCase: SignupUseCase, LoginUseCase: LoginUseCase, GetUserUseCase: GetUserUseCase}

	httpServer := rest.HttpServer{UserController: UserController, RestMiddleware: RestMiddleware}

	router := httpServer.ServeHttp()
	router.Run()
}
