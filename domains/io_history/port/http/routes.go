package http

import (
	"github.com/gin-gonic/gin"
	"goparking/database"
	"goparking/domains/io_history/repository"
	"goparking/domains/io_history/service"
	"goparking/internals/libs/validation"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	ioHistoryRepository := repository.NewIOHistoryRepository(sqlDB)
	ioHistoryService := service.NewIOHistoryService(validator, ioHistoryRepository)
	ioHistoryHandler := NewIOHistoryHandler(ioHistoryService)

	ioHistoryRoute := r.Group("/io-histories")
	{
		ioHistoryRoute.GET("", ioHistoryHandler.GetListIOHistories)
		ioHistoryRoute.POST("/entrance/:id", ioHistoryHandler.Entrance)
		ioHistoryRoute.POST("/exit/:id", ioHistoryHandler.Exit)
	}

}
