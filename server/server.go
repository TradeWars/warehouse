// Package server provides a HTTP listener for handling requests from either a
// game server or any form of interface to the data.
package server

import (
	"context"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"github.com/Southclaws/ScavengeSurviveCore/cache"
	"github.com/Southclaws/ScavengeSurviveCore/storage"
	"github.com/Southclaws/ScavengeSurviveCore/timeline"
	"github.com/Southclaws/ScavengeSurviveCore/types"
)

var version string

// Config stores static configuration
type Config struct {
	Temporary bool   `default:"true"`
	Bind      string `default:"0.0.0.0:7788"`
	Auth      string `default:"cunning_fox"`
	MongoHost string `split_words:"true" required:"false"`
	MongoPort string `split_words:"true" required:"false"`
	MongoName string `split_words:"true" required:"false"`
	MongoUser string `split_words:"true" required:"false"`
	MongoPass string `split_words:"true" required:"false"`
}

// App stores and controls program state
type App struct {
	config     *Config
	handlers   map[string][]Route
	validator  *validator.Validate
	store      types.Storer
	timeline   types.Timeline
	httpServer *http.Server
	ctx        context.Context
	cancel     context.CancelFunc
}

// Initialise performs all the necessary actions to bootstrap the application
// into a runnable state ready for starting with app.Start
func Initialise(config *Config) (app *App, err error) {
	logger.Debug("initialising ssc-server with debug logging",
		zap.String("version", version),
		zap.Any("config", config))

	app = &App{
		config:    config,
		validator: validator.New(),
	}
	app.handlers = map[string][]Route{
		"index":  app.indexRoutes(),
		"player": app.playerRoutes(),
		"admin":  app.adminRoutes(),
		"report": app.reportRoutes(),
		"ban":    app.banRoutes(),
	}

	if config.Temporary {
		logger.Debug("temporary mode is active, no data will persist!")
		app.store = cache.New()
	} else {
		if config.Bind == "" {
			err = errors.New("app not in temporary mode but SSC_BIND is not set")
			return
		}
		if config.Auth == "" {
			err = errors.New("app not in temporary mode but SSC_AUTH is not set")
			return
		}
		if config.MongoHost == "" {
			err = errors.New("app not in temporary mode but SSC_MONGO_HOST is not set")
			return
		}
		if config.MongoPort == "" {
			err = errors.New("app not in temporary mode but SSC_MONGO_PORT is not set")
			return
		}

		app.store, err = storage.New(storage.Config{
			Host: config.MongoHost,
			Port: config.MongoPort,
			Name: config.MongoName,
			User: config.MongoUser,
			Pass: config.MongoPass,
		})
		if err != nil {
			err = errors.Wrap(err, "failed to connect to storage")
			return
		}

		app.timeline, err = timeline.New(timeline.Config{
			Host: "localhost",
			Port: "5432",
			Name: "ssc",
			User: "default",
			Pass: "default",
		})
		if err != nil {
			err = errors.Wrap(err, "failed to connect to event store")
			return
		}
	}
	app.ctx, app.cancel = context.WithCancel(context.Background())

	router := mux.NewRouter().StrictSlash(true)
	for name, routes := range app.handlers {
		logger.Debug("loaded handler",
			zap.String("name", name),
			zap.Int("routes", len(routes)))

		for _, route := range routes {
			router.Methods(route.Method).
				Path(route.Path).
				Name(route.Name).
				Handler(app.Authenticator(route.handler))

			logger.Debug("registered handler route",
				zap.String("name", route.Name),
				zap.String("method", route.Method),
				zap.String("path", route.Path))
		}
	}

	app.httpServer = &http.Server{
		Addr: app.config.Bind,
		Handler: handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"HEAD", "GET", "POST", "PUT", "OPTIONS"}),
		)(router),
	}

	logger.Debug("initialisation complete")

	return
}

// Start fires up the HTTP server and blocks until failure
func (app *App) Start() (err error) {
	return app.httpServer.ListenAndServe()
}

// Stop gracefully shuts down the application
func (app *App) Stop() (err error) {
	app.cancel()
	return app.httpServer.Close()
}
