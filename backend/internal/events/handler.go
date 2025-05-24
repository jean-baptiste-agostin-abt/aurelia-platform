package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo Repository
}

func NewHandler(r Repository) *Handler {
	return &Handler{repo: r}
}

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	r.GET("/events", h.listEvents)
	r.GET("/feed", h.listEvents)
}

func (h *Handler) listEvents(c *gin.Context) {
	evs, err := h.repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, evs)
}
