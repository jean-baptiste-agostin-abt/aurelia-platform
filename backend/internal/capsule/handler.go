package capsule

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CapsuleInput struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	FamilyID uint   `json:"family_id" binding:"required"`
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/capsules", func(c *gin.Context) { createCapsule(c, db) })
	r.GET("/capsules/:id", func(c *gin.Context) { getCapsule(c, db) })
}

func createCapsule(c *gin.Context, db *gorm.DB) {
	var in CapsuleInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cap := Capsule{Title: in.Title, Content: in.Content, FamilyID: in.FamilyID}
	if err := db.Create(&cap).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cap)
}

func getCapsule(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var cap Capsule
	if err := db.First(&cap, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, cap)
}
