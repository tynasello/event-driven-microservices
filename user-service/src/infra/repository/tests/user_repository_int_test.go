package tests

import (
	"testing"

	"example.com/user-service/src/domain/entity"
	"example.com/user-service/src/infra/repository"
	"example.com/user-service/src/infra/service"
)

func TestCreateUser(t *testing.T) {
	service.LoadTestEnv()
	testDbService, cleanup := service.SetupTestDb(t)
	defer cleanup()
	userRepository := &repository.UserRepository{Db: testDbService.Db}

	testCases := []struct {
		user          entity.User
		expectSuccess bool
	}{
		{entity.User{Username: "testuser1", Password: "testpassword1"}, true},
		{entity.User{Username: "testuser2", Password: "testpassword1"}, true},
		{entity.User{Username: "testuser2", Password: "testpassword1"}, false},
	}

	for _, tc := range testCases {
		userCreatedResult := userRepository.Create(tc.user)
		if userCreatedResult.IsSuccess != tc.expectSuccess {
			t.Errorf("Expected success to be %v but got %v", tc.expectSuccess, userCreatedResult.IsSuccess)
		}
		if tc.expectSuccess && userCreatedResult.IsSuccess {
			createdUser, _ := userCreatedResult.GetValue()
			if createdUser.Id == 0 {
				t.Errorf("Expected created user to have an id but got empty id")
			}
			if createdUser.Username != tc.user.Username {
				t.Errorf("Expected created user to have username %s but got %s", tc.user.Username, createdUser.Username)
			}
			if createdUser.Password != tc.user.Password {
				t.Errorf("Expected created user to have password %s but got %s", tc.user.Password, createdUser.Password)
			}
			foundUserResult := userRepository.GetByUsername(tc.user.Username)
			if !foundUserResult.IsSuccess {
				t.Errorf("Expected to find user but got error")
			}
			if foundUserResult.IsSuccess {
				foundUser, _ := foundUserResult.GetValue()
				if foundUser.Id == 0 {
					t.Errorf("Expected found user to have an id but got empty id")
				}
				if foundUser.Username != tc.user.Username {
					t.Errorf("Expected found user to have username %s but got %s", tc.user.Username, foundUser.Username)
				}
				if foundUser.Password != tc.user.Password {
					t.Errorf("Expected found user to have password %s but got %s", tc.user.Password, foundUser.Password)
				}
			}
		}

	}
}

func TestGetUserById(t *testing.T) {
	service.LoadTestEnv()
	testDbService, cleanup := service.SetupTestDb(t)
	defer cleanup()
	userRepository := repository.UserRepository{Db: testDbService.Db}

	testCases := []struct {
		username      string
		userToCreate  entity.User
		expectSuccess bool
	}{
		{
			"testuser1",
			entity.User{Username: "testuser1", Password: "testpassword1"},
			true,
		},
		{
			"testuser2",
			entity.User{},
			false,
		},
	}

	for _, tc := range testCases {
		if tc.expectSuccess {
			userRepository.Create(tc.userToCreate)
		}
		existingUserResult := userRepository.GetByUsername(tc.username)
		if existingUserResult.IsSuccess != tc.expectSuccess {
			t.Errorf("Expected success to be %v but got %v", tc.expectSuccess, existingUserResult.IsSuccess)
		}
		if tc.expectSuccess && existingUserResult.IsSuccess {
			existingUser, _ := existingUserResult.GetValue()
			if existingUser.Id == 0 {
				t.Errorf("Expected created user to have an id but got empty id")
			}
			if existingUser.Username != tc.userToCreate.Username {
				t.Errorf("Expected created user to have username %s but got %s", tc.userToCreate.Username, existingUser.Username)
			}
			if existingUser.Password != tc.userToCreate.Password {
				t.Errorf("Expected created user to have password %s but got %s", tc.userToCreate.Password, existingUser.Password)
			}
		}

	}
}
