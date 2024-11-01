package routes

import (
	"net/http"

	"github.com/Sahil2k07/go-SOLID/src/controllers"
	"github.com/Sahil2k07/go-SOLID/src/database"
	"github.com/Sahil2k07/go-SOLID/src/middlewares"
	"github.com/Sahil2k07/go-SOLID/src/queries"
	"github.com/Sahil2k07/go-SOLID/src/services"
)

func UserRoutes(router *http.ServeMux) {
	queries := queries.UserQueriesInit(database.DbConnection)

	service := services.UserServiceInit(queries)

	controller := controllers.UserControllerInit(service)

	// Handle routes
	router.HandleFunc("/signup", controller.Signup)
	router.HandleFunc("/login", controller.Login)
	router.Handle("/update-profile", middlewares.Auth(http.HandlerFunc(controller.UpdateProfile)))
}
