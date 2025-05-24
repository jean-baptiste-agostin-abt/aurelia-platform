package legacyguard

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, repo Repository) {
	r.POST("/legacyguard/trigger", func(c *gin.Context) { trigger(c, repo) })
}

func trigger(c *gin.Context, repo Repository) {
	lg := LegacyGuard{Triggered: true}
	_ = repo.Create(&lg)
	c.JSON(http.StatusOK, gin.H{"status": "legacy guard triggered"})
}
