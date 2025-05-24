package legacyguard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/legacyguard/trigger", func(c *gin.Context) { trigger(c, db) })
}

func trigger(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{"status": "legacy guard triggered"})
}
