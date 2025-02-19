package http

import (
	"github.com/gin-gonic/gin"
	"goparking/database"
	"goparking/domains/auth/repository"
	"goparking/domains/auth/service"
	"goparking/internals/libs/validation"
	"goparking/pkgs/minio"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation, minioClient *minio.MinioClient) {
	userRepository := repository.NewUserRepository(sqlDB)
	userService := service.NewUserService(validator, userRepository, minioClient)
	authHandler := NewAuthHandler(userService)

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/signin", authHandler.SignIn)
		authRouter.POST("/signup", authHandler.SignUp)
	}
}
