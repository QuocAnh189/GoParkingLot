package http

import (
	"github.com/gin-gonic/gin"
	"goparking/domains/io_history/dto"
	"goparking/domains/io_history/service"
	"goparking/internals/libs/logger"
	"goparking/pkgs/response"
	"goparking/pkgs/utils"
	"net/http"
)

type IOHistoryHandler struct {
	service service.IIOHistoryService
}

func NewIOHistoryHandler(service service.IIOHistoryService) *IOHistoryHandler {
	return &IOHistoryHandler{service: service}
}

//		@Summary	 Retrieve a list of io_histories
//	 @Description Fetches a paginated list of io_histories based on the provided filter parameters.
//		@Tags		 IOHistory
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the list of io_histories"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/io-histories [get]
func (h *IOHistoryHandler) GetListIOHistories(c *gin.Context) {
	var req dto.ListIOHistoryRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get query", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	ioHistories, pagination, err := h.service.GetIOHistories(c, &req)
	if err != nil {
		logger.Error("Failed to get ioHistories", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get ioHistories")
		return
	}

	var res dto.ListIOHistoryResponse
	utils.MapStruct(&res.IOHistories, ioHistories)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Implement entrance
//	 @Description Fetches a paginated list of cards based on the provided filter parameters.
//		@Tags		 IOHistory
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"IOHistory created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/io-histories/entrance [post]
func (h *IOHistoryHandler) Entrance(c *gin.Context) {
	var req dto.CreateIoHistoryRequest

	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	ioHistory, err := h.service.Entrance(c, &req)
	if err != nil {
		logger.Error("Failed to create io_history", err)
		switch err.Error() {
		case "same type IN":
			response.Error(c, http.StatusConflict, err, "same type IN")
			return
		case "invalid card":
			response.Error(c, http.StatusConflict, err, "invalid card")
			return
		}
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	response.JSON(c, http.StatusCreated, ioHistory)
}

//		@Summary	 Implement entrance
//	 @Description Fetches a paginated list of cards based on the provided filter parameters.
//		@Tags		 IOHistory
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"IOHistory created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/io-histories/exit [post]
func (h *IOHistoryHandler) Exit(c *gin.Context) {
	var req dto.CreateIoHistoryRequest

	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	outHistory, inHistory, dataCard, err := h.service.Exit(c, &req)
	if err != nil {
		logger.Error("Failed to create io_history", err)
		switch err.Error() {
		case "same type OUT":
			response.Error(c, http.StatusConflict, err, "same type OUT")
			return
		case "invalid card":
			response.Error(c, http.StatusConflict, err, "invalid card")
			return
		case "the car is not in the yard":
			response.Error(c, http.StatusConflict, err, "the car is not in the yard")
			return
		}
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.ExitResponse
	res.DataIn = inHistory
	res.DataOut = outHistory
	res.DataCard = dataCard
	response.JSON(c, http.StatusCreated, res)
}
