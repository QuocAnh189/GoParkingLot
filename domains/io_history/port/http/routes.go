package http

import (
	"github.com/gin-gonic/gin"
	"goparking/database"
	cardRepository "goparking/domains/card/repository"
	"goparking/domains/io_history/repository"
	"goparking/domains/io_history/service"
	"goparking/internals/libs/validation"
	"goparking/pkgs/minio"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation, minioClient *minio.MinioClient) {
	ioHistoryRepository := repository.NewIOHistoryRepository(sqlDB)
	cardRepository := cardRepository.NewCardRepository(sqlDB)
	ioHistoryService := service.NewIOHistoryService(validator, ioHistoryRepository, minioClient, cardRepository)
	ioHistoryHandler := NewIOHistoryHandler(ioHistoryService)

	ioHistoryRoute := r.Group("/io-histories")
	{
		ioHistoryRoute.GET("", ioHistoryHandler.GetListIOHistories)
		ioHistoryRoute.POST("/entrance", ioHistoryHandler.Entrance)
		ioHistoryRoute.POST("/exit", ioHistoryHandler.Exit)
	}

}
