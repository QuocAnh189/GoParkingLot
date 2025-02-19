package http

import (
	"github.com/gin-gonic/gin"
	"goparking/database"
	"goparking/domains/card/repository"
	"goparking/domains/card/service"
	"goparking/internals/libs/validation"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	cardRepository := repository.NewCardRepository(sqlDB)
	cardService := service.NewCardService(validator, cardRepository)
	cardHandler := NewCardHandler(cardService)

	cardRoute := r.Group("/cards")
	{
		cardRoute.GET("", cardHandler.GetListCards)
		cardRoute.GET("/:id", cardHandler.GetCard)
		cardRoute.POST("", cardHandler.CreateCard)
		cardRoute.PUT("/:id", cardHandler.UpdateCard)
		cardRoute.DELETE("/:id", cardHandler.DeleteCard)
	}

}
