package server

import (
	"fmt"
	"goparking/pkgs/minio"
	"goparking/pkgs/token"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "goparking/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"goparking/database"
	"goparking/internals/libs/logger"
	"goparking/internals/libs/validation"
	"net/http"

	"goparking/configs"
	"goparking/pkgs/redis"

	authHttp "goparking/domains/auth/port/http"
	cardHttp "goparking/domains/card/port/http"
	ioHistoryHttp "goparking/domains/io_history/port/http"
)

type Server struct {
	engine      *gin.Engine
	cfg         *configs.Config
	validator   validation.Validation
	db          database.IDatabase
	minioClient *minio.MinioClient
	cache       redis.IRedis
	tokenMarker token.IMarker
}

func NewServer(
	validator validation.Validation,
	db database.IDatabase,
	minioClient *minio.MinioClient,
	cache redis.IRedis,
	tokenMarker token.IMarker,
) *Server {
	return &Server{
		engine:      gin.Default(),
		cfg:         configs.GetConfig(),
		validator:   validator,
		db:          db,
		minioClient: minioClient,
		cache:       cache,
		tokenMarker: tokenMarker,
	}
}

func (s Server) Run() error {
	_ = s.engine.SetTrustedProxies(nil)
	if s.cfg.Environment == configs.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	//s.engine.Use(cors.Default())
	s.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "access-control-allow-origin", "access-control-allow-headers"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	if err := s.MapRoutes(); err != nil {
		logger.Fatalf("MapRoutes Error: %v", err)
	}

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Goparking API"})
	})

	//Start http server
	logger.Info("HTTP server is listening on PORT: ", s.cfg.HttpPort)
	if err := s.engine.Run(fmt.Sprintf(":%d", s.cfg.HttpPort)); err != nil {
		logger.Fatalf("Running HTTP server: %v", err)
	}

	return nil
}

func (s Server) GetEngine() *gin.Engine {
	return s.engine
}

func (s Server) MapRoutes() error {
	routesV1 := s.engine.Group("/api/v1")

	authHttp.Routes(routesV1, s.db, s.validator, s.minioClient, s.cache, s.tokenMarker)
	cardHttp.Routes(routesV1, s.db, s.validator, s.cache, s.tokenMarker)
	ioHistoryHttp.Routes(routesV1, s.db, s.validator, s.minioClient, s.cache, s.tokenMarker)

	return nil
}
