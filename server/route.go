package server

import (
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

// Route represents an API route and its associated handler function
type Route struct {
	Name          string      `json:"name"`
	Path          string      `json:"path"`
	Method        string      `json:"method"`
	Authenticated bool        `json:"authenticated"`
	Accepts       interface{} `json:"accepts"`
	Returns       interface{} `json:"returns"`
	handler       EndpointHandler
}

// EndpointHandler wraps a HTTP handler function with app-specific args/returns
type EndpointHandler func(io.ReadCloser) (types.Status, error)

// ServeHTTP implements the necessary chaining functionality for HTTP middleware
func (f EndpointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := f(r.Body)
	if err != nil {
		logger.Error("request handler failed",
			zap.Error(err))
	}

	err = json.NewEncoder(w).Encode(status)
	if err != nil {
		logger.Error("failed to encode response",
			zap.Error(err))
		http.Error(w, "failed to encode response", 500)
	}
}
