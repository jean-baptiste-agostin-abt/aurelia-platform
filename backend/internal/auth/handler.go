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

type Handler struct {
	users user.Repository
}

func NewHandler(repo user.Repository) *Handler {
	return &Handler{users: repo}
}

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	r.POST("/auth/signup", h.signup)
	r.POST("/auth/login", h.login)
}

func (h *Handler) signup(c *gin.Context) {
	var in SignupInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashed, _ := crypto.HashPassword(in.Password)
	u := user.NewUser(in.Email, hashed, 0)
	if err := h.users.Create(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "signup successful"})
}

func (h *Handler) login(c *gin.Context) {
	var in LoginInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := h.users.FindByEmail(in.Email)
	if err != nil || !crypto.CheckPasswordHash(in.Password, u.Password) {
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
