package main

import (
	"goparking/configs"
	"goparking/database"
	"goparking/internals/libs/logger"
	"goparking/internals/libs/validation"
	"sync"

	httpServer "goparking/internals/server/http"
)

//	@title			GoParking Swagger API
//	@version		1.0
//	@description	Swagger API for GoShop.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Tran Phuoc Anh Quoc
//	@contact.email	anhquoc18092003@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	cfg := configs.LoadConfig(".")
	logger.Initialize(cfg.Environment)

	db, err := database.NewDatabase(cfg.DatabaseURI)
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}
	validator := validation.New()

	// Initialize HTTP server
	httpSvr := httpServer.NewServer(validator, db)

	var wg sync.WaitGroup
	wg.Add(1)

	// Run HTTP server
	go func() {
		defer wg.Done()
		if err := httpSvr.Run(); err != nil {
			logger.Fatal("Running HTTP server error:", err)
		}
	}()

	wg.Wait()
}
