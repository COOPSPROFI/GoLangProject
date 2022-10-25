package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/COOPSPROFI/GoLangProject/pkg/configs"
	"github.com/COOPSPROFI/GoLangProject/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// Get name,username,password off req body
	var request struct {
		Name     string
		Username string
		Password string
	}
	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create the user
	user := models.User{Name: request.Name, Username: request.Username, Password: string(hash)}
	result := configs.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{})
}

func Login(c *gin.Context) {
	// Get the email and pass off request body
	var request struct {
		Name     string
		Username string
		Password string
	}
	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		return
	}

	// Look up requested user
	var user models.User
	configs.DB.First(&user, "username = ?", request.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid to create token",
		})
		return
	}

	// Respond it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.Set("user", &models.User{})
	c.JSON(http.StatusOK, gin.H{
		"message": "success unauthorized",
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
