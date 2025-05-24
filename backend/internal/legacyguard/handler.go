package legacyguard

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
	r.POST("/legacyguard/trigger", h.trigger)
}

func (h *Handler) trigger(c *gin.Context) {
	lg := NewLegacyGuard(0)
	if err := h.repo.Trigger(lg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "legacy guard triggered"})
}
