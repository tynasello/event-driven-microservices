package service

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"example.com/user-service/src/application/usecase"
	"example.com/user-service/src/infra/model"
	"example.com/user-service/src/infra/repository"
	"example.com/user-service/src/infra/rest"
	"example.com/user-service/src/infra/rest/controller"
	"example.com/user-service/src/infra/rest/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// ProjectRoot folder of this project
	ProjectRoot = filepath.Join(filepath.Dir(b), "../../..")
)

func LoadTestEnv() {

	err := godotenv.Load(ProjectRoot + "/.env")
	if err != nil {
		panic("Error loading .env file")
	}
}

func SetupTestDb(t *testing.T) (*TestDbService, func()) {
	testDbService := NewTestDbService()
	testDbService.RunDbMigrations()
	testDbService.ClearDb()

	return testDbService, testDbService.ClearDb
}

type TestDbService struct {
	Db *gorm.DB
}

func NewTestDbService() *TestDbService {
	dbs := &TestDbService{}

	testDbUri := os.Getenv("TEST_DB_URI")

	db, err := gorm.Open(mysql.Open(testDbUri), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbs.Db = db

	return dbs
}

func (dbs TestDbService) RunDbMigrations() {
	dbs.Db.AutoMigrate(&model.UserModel{})
}

func (dbs TestDbService) ClearDb() {
	println("Clearing test db...")
	tables := []string{"user"}
	dbs.Db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	for _, table := range tables {
		err := dbs.Db.Exec("DELETE FROM " + table).Error
		if err != nil {
			fmt.Println("failed to drop database: %w", err)
		}
	}
	dbs.Db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

type E2eTestEnv struct {
	Router         *gin.Engine
	SignupUseCase  *usecase.SignupUseCase
	LoginUseCase   *usecase.LoginUseCase
	GetUserUseCase *usecase.GetUserUseCase
	UserRepository *repository.UserRepository
}

func SetupE2eTestEnv(t *testing.T) *E2eTestEnv {

	LoadTestEnv()
	testDbService, cleanup := SetupTestDb(t)
	defer cleanup()
	userRepository := &repository.UserRepository{Db: testDbService.Db}

	hashService := &BcryptHashService{}
	authTokenService := &JwtAuthTokenService{}

	signupUseCase := &usecase.SignupUseCase{UserRepository: userRepository, HashService: hashService, AuthTokenService: authTokenService}
	loginUseCase := &usecase.LoginUseCase{UserRepository: userRepository, HashService: hashService, AuthTokenService: authTokenService}
	getUserUseCase := &usecase.GetUserUseCase{UserRepository: userRepository}

	restMiddleware := &middleware.RestMiddleware{AuthTokenService: authTokenService}
	userController := &controller.UserController{SignupUseCase: signupUseCase, LoginUseCase: loginUseCase, GetUserUseCase: getUserUseCase}

	httpServer := &rest.HttpServer{UserController: userController, RestMiddleware: restMiddleware}
	router := httpServer.ServeHttp()

	return &E2eTestEnv{Router: router, SignupUseCase: signupUseCase, LoginUseCase: loginUseCase, GetUserUseCase: getUserUseCase, UserRepository: userRepository}
}
