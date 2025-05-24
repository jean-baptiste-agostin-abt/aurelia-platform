package family

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FamilyInput struct {
	Name string `json:"name" binding:"required"`
}

type FamilyResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Handler struct {
	repo Repository
}

func NewHandler(r Repository) *Handler {
	return &Handler{repo: r}
}

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	r.POST("/families", h.createFamily)
}

func (h *Handler) createFamily(c *gin.Context) {
	var in FamilyInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fam := NewFamily(in.Name)
	if err := h.repo.Create(fam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, FamilyResponse{ID: fam.ID, Name: fam.Name})
}
