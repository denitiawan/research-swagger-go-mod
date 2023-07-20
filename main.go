package main

import (
	"denitiawan/research-swagger-gomod-gin/common/helper"
	"denitiawan/research-swagger-gomod-gin/config"
	_ "denitiawan/research-swagger-gomod-gin/docs"
	"denitiawan/research-swagger-gomod-gin/router"
	"github.com/rs/zerolog/log"
	"net/http"
)

// --------------------------------------------------
// annotations		: Annotation used for Swagger-UI
//					and will be mapping to folder and files (./root/docs/**)
// docs import		: import ( _ "denitiawan/research-swagger-gomod-gin/docs" )
//					will be used for update all values on all files inside that folder
//					when you run syntax (swag init)
// url swagger-ui 	: http://localhost:8899/nexsoft/doc/api/swagger/index.html
// --------------------------------------------------
// @version		1.1.0
// @title 		Nexsoft Demo Swagger in GO (gin)
// @description How to implement swagger-ui in Go and gin http client project
// @host 		localhost:8899
// @BasePath 	/api
func main() {

	log.Info().Msg("Started Server!")

	// # Database Connection
	db := config.DatabaseConnection()

	// # Database Migration
	//config.DatabaseMigration(db)

	// # Route Initialization
	appRoute := router.NewRouterInit()

	// # Base Router
	basePath := appRoute.Group("/api")

	// # API Swagger Routing
	router.SwaggerRouting(appRoute)

	// # API Welcome
	router.APIWelcome(appRoute)

	// # ALL APIs Routing
	router.APIRouting(db, basePath)

	// # Server
	server := &http.Server{
		Addr:    ":8899",
		Handler: appRoute,
	}

	// # Serve
	err := server.ListenAndServe()
	helper.ErrorPanic(err)

	defer db.Close()
}
