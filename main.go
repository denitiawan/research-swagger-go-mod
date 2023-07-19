package main

import (
	"denitiawan/research-swagger-gomod-gin/common/helper"
	"denitiawan/research-swagger-gomod-gin/config"
	"denitiawan/research-swagger-gomod-gin/router"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {

	log.Info().Msg("Started Server!")

	// # Database Connection
	db := config.DatabaseConnection()

	// # Database Migration
	//config.DatabaseMigration(db)

	// # Routes Config
	routes := router.NewRouterInit()

	// # Swagger
	router.SwaggerRouting("http://localhost:8899", routes)

	// # API Registration
	router.APIRouting(db, routes)

	// # Server
	server := &http.Server{
		Addr:    ":8899",
		Handler: routes,
	}

	// # Serve
	err := server.ListenAndServe()
	helper.ErrorPanic(err)

	defer db.Close()
}
