package server

import (
	"context"
	"net/http"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Config stores static configuration
type Config struct {
	Bind string `split_words:"true" required:"true"` // bind interface
}

// App stores and controls program state
type App struct {
	config   Config
	handlers map[string][]Route
	db       *mgo.Database
	ctx      context.Context
	cancel   context.CancelFunc
}

// Start fires up a HTTP server and routes API calls to the database manager
func Start(config Config) {
	logger.Debug("initialising ssc-server with debug logging", zap.Any("config", config))

	app := App{
		config: config,
	}

	app.ctx, app.cancel = context.WithCancel(context.Background())

	logger.Debug("configured routes, starting listener")

	router := mux.NewRouter().StrictSlash(true)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"HEAD", "GET", "POST", "PUT", "OPTIONS"})

	app.handlers = map[string][]Route{
		"player": playerRoutes(app),
		"report": reportRoutes(app),
		"ban":    banRoutes(app),
	}

	for name, routes := range app.handlers {
		logger.Debug("loaded handler",
			zap.String("name", name),
			zap.Int("routes", len(routes)))

		for _, route := range routes {
			if route.Authenticated {
				router.Methods(route.Method).
					Path(route.Path).
					Name(route.Name).
					Handler(app.Authenticator(route.handler))
			} else {
				router.Methods(route.Method).
					Path(route.Path).
					Name(route.Name).
					Handler(route.handler)
			}

			logger.Debug("registered handler route",
				zap.String("name", route.Name),
				zap.String("method", route.Method),
				zap.String("path", route.Path))
		}
	}

	err := http.ListenAndServe(app.config.Bind, handlers.CORS(headersOk, originsOk, methodsOk)(router))

	logger.Fatal("unexpected termination",
		zap.Error(err))
}
