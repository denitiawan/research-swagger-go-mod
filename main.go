package main

import (
	"denitiawan/research-swagger-gomod-gin/common/error"
	"denitiawan/research-swagger-gomod-gin/config"
	_ "denitiawan/research-swagger-gomod-gin/docs"
	"denitiawan/research-swagger-gomod-gin/router"
	"fmt"
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
// @title 		Demo Swagger-UI (GO+GIN) for Nexsoft Project
// @description Implement swagger-ui on Go project with gin (web framework) + JWT Authorization
// @host 		localhost:8899
// @BasePath 	/api

// ------showing authorize button (but validation jwt is not working)---------
// @Security Authorization
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @schemes http
// ------showing authorize button (but validation jwt is not working)---------
func main() {

	log.Info().Msg("Try to start Application!")

	// App Config
	appConfig, err := config.LoadConfig(".", "dev")
	error.ErrorPanic(err)

	// # Database Connection
	db := config.DatabaseConnection(appConfig)

	// # Database Migration
	config.DatabaseMigration(appConfig, db)

	// # Route Initialization
	appRoute := router.NewRouterInit()

	// # Base Router
	basePath := appRoute.Group("/api")

	// # API Swagger Routing
	config.SwaggerRouting(appRoute)

	// # API Welcome
	router.APIWelcome(appRoute)

	// # ALL APIs Routing
	router.APIRouting(appConfig, db, basePath)
	fmt.Println("Server Config")
	fmt.Println(appConfig)

	// # Server
	server := &http.Server{
		Addr:    ":" + appConfig.ServerPort,
		Handler: appRoute,
	}
	log.Info().Msg("Yea Boy!.. Application is running!")

	// # Serve
	err = server.ListenAndServe()
	error.ErrorPanic(err)

	defer db.Close()

}
