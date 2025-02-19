package http

import (
	"github.com/gin-gonic/gin"
	"goparking/domains/card/service"
	"goparking/pkgs/response"
)

type CardHandler struct {
	service service.ICardService
}

func NewCardHandler(service service.ICardService) *CardHandler {
	return &CardHandler{service: service}
}

func (h *CardHandler) GetListCards(c *gin.Context) {
	response.JSON(c, 200, "Get List Cards")
}

func (h *CardHandler) GetCard(c *gin.Context) {
	response.JSON(c, 200, "Get Card")
}

func (h *CardHandler) CreateCard(c *gin.Context) {
	response.JSON(c, 200, "Create Card")
}

func (h *CardHandler) UpdateCard(c *gin.Context) {
	response.JSON(c, 200, "Update Card")
}

func (h *CardHandler) DeleteCard(c *gin.Context) {
	response.JSON(c, 200, "Delete Card")
}
