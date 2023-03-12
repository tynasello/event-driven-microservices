package rest

import (
	"net/http"
	"time"

	"example.com/user-service/src/application/interfaces"
	"example.com/user-service/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignupUser(c *gin.Context, db *gorm.DB, hashService interfaces.IHashService, authTokenService interfaces.IAuthTokenService) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.BindJSON(&requestBody)

	hashedPassword, err := hashService.Hash(requestBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error hashing user password",
		})
		return
	}

	user := models.User{
		Username: requestBody.Username,
		Password: hashedPassword,
	}

	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating a user",
		})
		return
	}

	accessToken, err := authTokenService.GenerateToken(user.Username, 1*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error generating access token",
		})
	}

	c.SetCookie("access-token", accessToken, 60*60, "/", "localhost", false, true)

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func LoginUser(c *gin.Context, db *gorm.DB, hashService interfaces.IHashService, authTokenService interfaces.IAuthTokenService) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.BindJSON(&requestBody)

	var user models.User

	result := db.Where("username = ?", requestBody.Username).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	err := hashService.ValidateHash(user.Password, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	accessToken, err := authTokenService.GenerateToken(user.Username, 1*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error generating access token",
		})
	}

	c.SetCookie("access-token", accessToken, 60*60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func GetUser(c *gin.Context, db *gorm.DB, authTokenService interfaces.IAuthTokenService) {
	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var user models.User

	result := db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func AuthenticateUser(c *gin.Context, db *gorm.DB, authTokenService interfaces.IAuthTokenService) {
	username := c.GetString("username")

	c.JSON(http.StatusOK, gin.H{"username": username})
}
