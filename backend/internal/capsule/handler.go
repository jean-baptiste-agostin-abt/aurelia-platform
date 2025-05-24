package capsule

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CapsuleInput struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	FamilyID uint   `json:"family_id" binding:"required"`
}

type Handler struct {
	repo Repository
}

func NewHandler(r Repository) *Handler {
	return &Handler{repo: r}
}

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	r.POST("/capsules", h.createCapsule)
	r.GET("/capsules/:id", h.getCapsule)
}

func (h *Handler) createCapsule(c *gin.Context) {
	var in CapsuleInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cap := NewCapsule(in.Title, in.Content, in.FamilyID)
	if err := h.repo.Create(cap); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cap)
}

func (h *Handler) getCapsule(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscan(idParam, &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	cap, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, cap)
}
