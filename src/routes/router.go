package routes

import "net/http"

func AppRouter() *http.ServeMux {
	router := http.NewServeMux()

	UserRoutes(router)

	return router
}
