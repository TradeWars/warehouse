package server

import (
	"io"
	"net/url"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func (app *App) indexRoutes() []Route {
	return []Route{
		{
			"index",
			"GET",
			"/",
			nil,
			"index data",
			app.index,
		},
	}
}

func (app *App) index(r io.Reader, query url.Values) (status types.Status, err error) {
	return types.NewStatus(app.handlers, true, "good luck out there survivors!"), nil
}
