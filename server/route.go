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
	Method        string      `json:"method"`
	Path          string      `json:"path"`
	Authenticated bool        `json:"authenticated"`
	Accepts       interface{} `json:"accepts"`
	Returns       interface{} `json:"returns"`
	handler       EndpointHandler
}

// EndpointHandler wraps a HTTP handler function with app-specific args/returns
type EndpointHandler func(io.Reader) (types.Status, error)

// ServeHTTP implements the necessary chaining functionality for HTTP middleware
func (f EndpointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := f(r.Body)
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
