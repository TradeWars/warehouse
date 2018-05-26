package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

// Route represents an API route and its associated handler function
type Route struct {
	Name    string      `json:"name"`
	Method  string      `json:"method"`
	Path    string      `json:"path"`
	Accepts interface{} `json:"accepts"`
	Returns interface{} `json:"returns"`
	handler EndpointHandler
}

// Authenticator is a middleware layer for requests that require authentication
func (app *App) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		access := r.Header.Get("Authorization")
		if access != app.config.Auth {
			var status types.Status
			if access == "" {
				status = types.NewStatus(nil, false, "header 'Authorization' is absent from request")
			} else {
				status = types.NewStatus(nil, false, "header 'Authorization' incorrect")
			}

			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(status)
			if err != nil {
				logger.Error("failed to encode response",
					zap.Error(err))
				http.Error(w, "failed to encode response", 500)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

// EndpointHandler wraps a HTTP handler function with app-specific args/returns
type EndpointHandler func(r io.Reader, query url.Values) (types.Status, error)

// ServeHTTP implements the necessary chaining functionality for HTTP middleware
func (f EndpointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := f(r.Body, r.URL.Query())
	if err != nil {
		logger.Error("request handler failed",
			zap.Error(err))
		status = types.NewStatus(nil, false, err.Error())
	}

	logger.Debug("responding with status",
		zap.Any("status", status))

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(status)
	if err != nil {
		logger.Error("failed to encode response",
			zap.Error(err))
		http.Error(w, "failed to encode response", 500)
	}
}
