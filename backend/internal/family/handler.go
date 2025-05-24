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

func RegisterRoutes(r *gin.RouterGroup, repo Repository) {
	r.POST("/families", func(c *gin.Context) { createFamily(c, repo) })
}

func createFamily(c *gin.Context, repo Repository) {
	var in FamilyInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fam := Family{Name: in.Name}
	if err := repo.Create(&fam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, FamilyResponse{ID: fam.ID, Name: fam.Name})
}
