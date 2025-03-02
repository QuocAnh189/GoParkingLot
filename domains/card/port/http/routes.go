package http

import (
	"github.com/gin-gonic/gin"
	"goparking/database"
	"goparking/domains/card/repository"
	"goparking/domains/card/service"
	"goparking/internals/libs/validation"
	"goparking/pkgs/middleware"
	"goparking/pkgs/redis"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation, cache redis.IRedis) {
	cardRepository := repository.NewCardRepository(sqlDB)
	cardService := service.NewCardService(validator, cardRepository)
	cardHandler := NewCardHandler(cardService)

	authMiddleware := middleware.JWTAuth(cache)
	cardRoute := r.Group("/cards").Use(authMiddleware)
	{
		cardRoute.GET("/", cardHandler.GetListCards)
		cardRoute.GET("/:id", cardHandler.GetCard)
		cardRoute.POST("/", cardHandler.CreateCard)
		cardRoute.PUT("/:id", cardHandler.UpdateCard)
		cardRoute.DELETE("/:id", cardHandler.DeleteCard)
	}

}
