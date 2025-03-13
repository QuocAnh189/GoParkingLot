package http

import (
	"goparking/database"
	cardRepository "goparking/domains/card/repository"
	"goparking/domains/io_history/repository"
	"goparking/domains/io_history/service"
	"goparking/internals/libs/validation"
	"goparking/pkgs/middleware"
	"goparking/pkgs/minio"
	"goparking/pkgs/redis"
	"goparking/pkgs/token"

	"github.com/gin-gonic/gin"
)

func Routes(
	r *gin.RouterGroup,
	sqlDB database.IDatabase,
	validator validation.Validation,
	minioClient *minio.MinioClient,
	cache redis.IRedis,
	token token.IMarker,
) {
	ioHistoryRepository := repository.NewIOHistoryRepository(sqlDB)
	cardRepository := cardRepository.NewCardRepository(sqlDB)
	ioHistoryService := service.NewIOHistoryService(validator, ioHistoryRepository, minioClient, cardRepository)
	ioHistoryHandler := NewIOHistoryHandler(ioHistoryService)

	authMiddleware := middleware.NewAuthMiddleware(token).TokenAuth(cache)

	ioHistoryRoute := r.Group("/io-histories").Use(authMiddleware)
	{
		ioHistoryRoute.GET("/", ioHistoryHandler.GetListIOHistories)
		ioHistoryRoute.POST("/entrance", ioHistoryHandler.Entrance)
		ioHistoryRoute.POST("/exit", ioHistoryHandler.Exit)
	}
}
