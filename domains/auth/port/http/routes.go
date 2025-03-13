package http

import (
	"goparking/database"
	"goparking/domains/auth/repository"
	"goparking/domains/auth/service"
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
	userRepository := repository.NewUserRepository(sqlDB)
	userService := service.NewUserService(validator, userRepository, minioClient, cache, token)
	authHandler := NewAuthHandler(userService)

	authMiddleware := middleware.NewAuthMiddleware(token).TokenAuth(cache)

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/signin", authHandler.SignIn)
		authRouter.POST("/signup", authHandler.SignUp)
		authRouter.POST("/signout", authMiddleware, authHandler.SignOut)
	}
}
