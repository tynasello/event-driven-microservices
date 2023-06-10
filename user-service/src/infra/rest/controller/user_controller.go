package controller

import (
	"net/http"

	"example.com/user-service/src/application/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	SignupUseCase  *usecase.SignupUseCase
	GetUserUseCase *usecase.GetUserUseCase
	LoginUseCase   *usecase.LoginUseCase
}

func (u UserController) SignupUser(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.BindJSON(&requestBody)

	signupResult := u.SignupUseCase.Execute(requestBody.Username, requestBody.Password)

	if signupResult.IsFailure {
		signupResultErrorMessage, _ := signupResult.GetErrorMessage()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": *signupResultErrorMessage,
		})
		return
	}

	accessToken, _ := signupResult.GetValue()

	c.SetCookie("access-token", *accessToken, 60*60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func (u UserController) LoginUser(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.BindJSON(&requestBody)

	loginResult := u.LoginUseCase.Execute(requestBody.Username, requestBody.Password)

	if loginResult.IsFailure {
		loginResultErrorMessage, _ := loginResult.GetErrorMessage()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": *loginResultErrorMessage,
		})
		return
	}

	accessToken, _ := loginResult.GetValue()

	c.SetCookie("access-token", *accessToken, 60*60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func (u UserController) GetUser(c *gin.Context) {
	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	getUserResult := u.GetUserUseCase.Execute(username)

	if getUserResult.IsFailure {
		getUserResultErrorMessage, _ := getUserResult.GetErrorMessage()
		c.JSON(http.StatusInternalServerError, gin.H{"message": *getUserResultErrorMessage})
		return
	}

	user, _ := getUserResult.GetValue()

	c.JSON(http.StatusOK, gin.H{"user": *user})
}

func (u UserController) AuthenticateUser(c *gin.Context) {
	username := c.GetString("username")

	c.JSON(http.StatusOK, gin.H{"username": username})
}
