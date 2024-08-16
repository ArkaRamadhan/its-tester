package controllers

import (
	"project-gin/initializers"
	"project-gin/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	if username == "" || password == "" {
		c.JSON(400, gin.H{"error": "Username and password are required"})
		return
	}

	var dbUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := initializers.DB.Where("username = ?", username).First(&dbUser).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}

	if password != dbUser.Password {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := generateJWT(dbUser.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("your_secret_key"))
}

func AdminLogin(c *gin.Context) {
	var auth Auth
	if err := c.BindJSON(&auth); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var admin models.Admin
	if err := initializers.DB.Where("username = ?", auth.Username).First(&admin).Error; err != nil {
		c.JSON(401, gin.H{"error": "invalid username or password"})
		return
	}

	if admin.Password != auth.Password {
		c.JSON(401, gin.H{"error": "invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": admin.Username,
		"role":     admin.Role,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Token{Token: tokenString})
}
