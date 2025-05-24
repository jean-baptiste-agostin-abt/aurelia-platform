package family

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type FamilyInput struct {
	Name string `json:"name" binding:"required"`
}

type FamilyResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/families", func(c *gin.Context) { createFamily(c, db) })
}

func createFamily(c *gin.Context, db *gorm.DB) {
	var in FamilyInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fam := Family{Name: in.Name}
	if err := db.Create(&fam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, FamilyResponse{ID: fam.ID, Name: fam.Name})
}
