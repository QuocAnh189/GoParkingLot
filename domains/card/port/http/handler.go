package http

import (
	"github.com/gin-gonic/gin"
	"goparking/domains/card/dto"
	"goparking/domains/card/model"
	"goparking/domains/card/service"
	"goparking/internals/libs/logger"
	"goparking/pkgs/response"
	"goparking/pkgs/utils"
	"net/http"
)

type CardHandler struct {
	service service.ICardService
}

func NewCardHandler(service service.ICardService) *CardHandler {
	return &CardHandler{service: service}
}

//		@Summary	 Retrieve a list of cards
//	 @Description Fetches a paginated list of cards based on the provided filter parameters.
//		@Tags		 Cards
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the list of cards"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/cards [get]
func (h *CardHandler) GetListCards(c *gin.Context) {
	var req dto.ListCardRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get query", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	cards, pagination, err := h.service.GetCards(c, &req)
	if err != nil {
		logger.Error("Failed to get cards", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get cards")
		return
	}

	var res dto.ListCardResponse
	utils.MapStruct(&res.Cards, cards)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a card by its ID
//	 @Description Fetches the details of a specific card based on the provided card ID.
//		@Tags		 Cards
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the card"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/cards/{id} [get]
func (h *CardHandler) GetCard(c *gin.Context) {
	var res model.Card

	cardId := c.Param("id")
	card, err := h.service.GetCardById(c, cardId)
	if err != nil {
		logger.Error("Failed to get card detail: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	utils.MapStruct(&res, card)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Create a new card
//	 @Description Creates a new card based on the provided details.
//		@Tags		 Cards
//		@Produce	 json
//		@Param		 _	body	dto.CreateCardRequest	  true	"Body"
//		@Success	 201	{object}	response.Response	"Card created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/cards [post]
func (h *CardHandler) CreateCard(c *gin.Context) {
	var req dto.CreateCardRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	if err := h.service.CreateCard(c, &req); err != nil {
		logger.Error("Failed to create card", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	response.JSON(c, http.StatusCreated, true)
}

//		@Summary	 Update a card
//	 @Description Update a new card based on the provided details.
//		@Tags		 Cards
//		@Produce	 json
//		@Param		 _	body	dto.UpdateCardRequest	  true	"Body"
//		@Success	 201	{object}	response.Response	"Card updated successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/cards/{id} [put]
func (h *CardHandler) UpdateCard(c *gin.Context) {
	var req dto.UpdateCardRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	cardId := c.Param("id")

	if err := h.service.UpdateCard(c, cardId, &req); err != nil {
		logger.Error("Failed to update card", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	response.JSON(c, http.StatusOK, true)
}

//		@Summary	 Delete a card
//	 @Description Delete a new card based on the provided details.
//		@Tags		 Cards
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"Delete updated successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/cards/{id} [delete]
func (h *CardHandler) DeleteCard(c *gin.Context) {
	cardId := c.Param("id")

	err := h.service.DeleteCard(c, cardId)

	if err != nil {
		logger.Error("Failed to delete cards: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, true)
}
