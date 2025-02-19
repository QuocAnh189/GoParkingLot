package http

import (
	"github.com/gin-gonic/gin"
	"goparking/domains/io_history/service"
	"goparking/pkgs/response"
)

type IOHistoryHandler struct {
	service service.IIOHistoryService
}

func NewIOHistoryHandler(service service.IIOHistoryService) *IOHistoryHandler {
	return &IOHistoryHandler{service: service}
}

func (h *IOHistoryHandler) GetListIOHistories(c *gin.Context) {
	response.JSON(c, 200, "Get List IOHistories")
}

func (h *IOHistoryHandler) Entrance(c *gin.Context) {
	response.JSON(c, 200, "Entrance")
}

func (h *IOHistoryHandler) Exit(c *gin.Context) {
	response.JSON(c, 200, "Exit")
}
