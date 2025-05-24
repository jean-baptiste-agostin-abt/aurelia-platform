package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/events", func(c *gin.Context) { listEvents(c, db) })
	r.GET("/feed", func(c *gin.Context) { listEvents(c, db) })
}

func listEvents(c *gin.Context, db *gorm.DB) {
	var evs []Event
	if err := db.Find(&evs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, evs)
}
