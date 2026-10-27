package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes sets up the routing configuration for HTTP requests.
// It returns a configured instance of *httprouter.Router.
//
// Available routes:
//   - GET	/v1/healthcheck	: Check the health status of the service
//   - POST /v1/movies		: Create a new movie
//   - GET	/v1/movies/:id	: Get information about a specific movie by its ID
func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return router
}
