package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, repo Repository) {
	r.GET("/events", func(c *gin.Context) { listEvents(c, repo) })
	r.GET("/feed", func(c *gin.Context) { listEvents(c, repo) })
}

func listEvents(c *gin.Context, repo Repository) {
	evs, err := repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, evs)
}
