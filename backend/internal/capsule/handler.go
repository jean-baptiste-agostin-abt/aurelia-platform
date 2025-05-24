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

func RegisterRoutes(r *gin.RouterGroup, repo Repository) {
	r.POST("/capsules", func(c *gin.Context) { createCapsule(c, repo) })
	r.GET("/capsules/:id", func(c *gin.Context) { getCapsule(c, repo) })
}

func createCapsule(c *gin.Context, repo Repository) {
	var in CapsuleInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cap := Capsule{Title: in.Title, Content: in.Content, FamilyID: in.FamilyID}
	if err := repo.Create(&cap); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cap)
}

func getCapsule(c *gin.Context, repo Repository) {
	id := c.Param("id")
	cap, err := repo.Find(parseID(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, cap)
}

func parseID(idStr string) uint {
	var id uint
	fmt.Sscan(idStr, &id)
	return id
}
