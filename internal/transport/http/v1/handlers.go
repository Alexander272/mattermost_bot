package v1

import (
	"github.com/Alexander272/mattermost_bot/internal/config"
	"github.com/Alexander272/mattermost_bot/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Services
	// middleware *middleware.Middleware
}

func NewHandler(services *services.Services) *Handler {
	// middleware.CookieName = CookieName

	return &Handler{
		services: services,
		// middleware: middleware,
	}
}

func (h *Handler) Init(conf *config.Config, api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		v1.GET("/", h.notImplemented)
	}

	h.initMost(v1)
}

func (h *Handler) notImplemented(c *gin.Context) {}
