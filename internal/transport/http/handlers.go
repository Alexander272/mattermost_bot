package http

import (
	"net/http"

	"github.com/Alexander272/mattermost_bot/internal/config"
	"github.com/Alexander272/mattermost_bot/internal/services"
	httpV1 "github.com/Alexander272/mattermost_bot/internal/transport/http/v1"
	"github.com/Alexander272/mattermost_bot/pkg/limiter"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(conf *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		limiter.Limit(conf.Limiter.RPS, conf.Limiter.Burst, conf.Limiter.TTL),
	)

	// Init router
	router.GET("/api/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(conf, router)

	return router
}

func (h *Handler) initAPI(conf *config.Config, router *gin.Engine) {
	handlerV1 := httpV1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(conf, api)
	}
}
