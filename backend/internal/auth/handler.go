package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/yourorg/aurelia-backend/internal/user"
	"github.com/yourorg/aurelia-backend/pkg/crypto"
)

type SignupInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput SignupInput

var jwtKey = []byte("secret")

func RegisterRoutes(r *gin.RouterGroup, repo user.Repository) {
	r.POST("/auth/signup", func(c *gin.Context) { signup(c, repo) })
	r.POST("/auth/login", func(c *gin.Context) { login(c, repo) })
}

func signup(c *gin.Context, repo user.Repository) {
	var in SignupInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashed, _ := crypto.HashPassword(in.Password)
	u := user.User{Email: in.Email, Password: hashed}
	if err := repo.Create(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "signup successful"})
}

func login(c *gin.Context, repo user.Repository) {
	var in LoginInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := repo.FindByEmail(in.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	if !crypto.CheckPasswordHash(in.Password, u.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, _ := token.SignedString(jwtKey)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
