package v1

import (
	"net/http"

	"github.com/Alexander272/mattermost_bot/internal/models"
	"github.com/Alexander272/mattermost_bot/internal/models/response"
	"github.com/Alexander272/mattermost_bot/internal/services"
	"github.com/gin-gonic/gin"
)

type MostHandler struct {
	service services.Most
}

func NewMostHandler(service services.Most) *MostHandler {
	return &MostHandler{service: service}
}

func (h *Handler) initMost(api *gin.RouterGroup) {
	handlers := NewMostHandler(h.services.Most)

	masks := api.Group("/mattermost")
	{
		masks.POST("/send", handlers.sendMessage)
	}
}

func (h *MostHandler) sendMessage(c *gin.Context) {
	var dto models.Message
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Данные имеют не верный формат")
		return
	}

	if _, err := h.service.Send(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка. "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, response.IdResponse{Message: "Сообщение отправлено"})
}
