package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezashahnazar/GolangAuth/initializers"
	"github.com/rezashahnazar/GolangAuth/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}
	// hash pass
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Hash password",
		})
		return
	}
	// create user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Create user.",
		})
		return
	}

	// Respond

	c.JSON(http.StatusOK, gin.H{})
}
