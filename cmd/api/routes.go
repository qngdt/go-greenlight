package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /v1/healthcheck", app.healthCheckHandler)
	router.HandleFunc("GET /v1/movies", app.listMovieHandler)
	router.HandleFunc("POST /v1/movies", app.createMovieHandler)
	router.HandleFunc("GET /v1/movies/{id}", app.showMovieHandler)
	router.HandleFunc("PATCH /v1/movies/{id}", app.updateMovieHandler)
	router.HandleFunc("DELETE /v1/movies/{id}", app.deleteMovieHandler)

	router.HandleFunc("POST /v1/users", app.registerUserHandler)
	router.HandleFunc("PUT /v1/users/activated", app.activateUserHandler)

	return app.recoverPanic(app.rateLimit(router))
}
