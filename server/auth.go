package server

import (
	"net/http"
)

// Authenticator is a middleware layer for requests that require authentication
func (app *App) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("received authenticated request")

		next.ServeHTTP(w, r)
	})
}
