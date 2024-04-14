package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /v1/healthcheck", app.healthCheckHandler)
	router.HandleFunc("GET /v1/movies", app.requirePermission("movies:read", app.listMovieHandler))
	router.HandleFunc("POST /v1/movies", app.requirePermission("movies:right", app.createMovieHandler))
	router.HandleFunc("GET /v1/movies/{id}", app.requirePermission("movies:read", app.showMovieHandler))
	router.HandleFunc("PATCH /v1/movies/{id}", app.requirePermission("movies:write", app.updateMovieHandler))
	router.HandleFunc("DELETE /v1/movies/{id}", app.requirePermission("movies:write", app.deleteMovieHandler))

	router.HandleFunc("POST /v1/users", app.registerUserHandler)
	router.HandleFunc("PUT /v1/users/activated", app.activateUserHandler)

	router.HandleFunc("POST /v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
