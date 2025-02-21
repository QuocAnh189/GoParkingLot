package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goparking/pkgs/minio"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "goparking/docs"

	"goparking/database"
	"goparking/internals/libs/logger"
	"goparking/internals/libs/validation"
	"net/http"

	"goparking/configs"

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
}

func NewServer(validator validation.Validation, db database.IDatabase, minioClient *minio.MinioClient) *Server {
	return &Server{
		engine:      gin.Default(),
		cfg:         configs.GetConfig(),
		validator:   validator,
		db:          db,
		minioClient: minioClient,
	}
}

func (s Server) Run() error {
	_ = s.engine.SetTrustedProxies(nil)
	if s.cfg.Environment == configs.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	s.engine.Use(cors.Default())

	if err := s.MapRoutes(); err != nil {
		logger.Fatalf("MapRoutes Error: %v", err)
	}

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to GoHub API"})
	})

	// Start http server
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

	routesV1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"Test": "Call api"})
	})

	authHttp.Routes(routesV1, s.db, s.validator, s.minioClient)
	cardHttp.Routes(routesV1, s.db, s.validator)
	ioHistoryHttp.Routes(routesV1, s.db, s.validator, s.minioClient)

	return nil
}
