package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goparking/database"
	"goparking/internals/libs/logger"
	"goparking/internals/libs/validation"

	"goparking/configs"
)

type Server struct {
	engine    *gin.Engine
	cfg       *configs.Config
	validator validation.Validation
	db        database.IDatabase
}

func NewServer(validator validation.Validation, db database.IDatabase) *Server {
	return &Server{
		engine:    gin.Default(),
		cfg:       configs.GetConfig(),
		validator: validator,
		db:        db,
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

	return nil
}
