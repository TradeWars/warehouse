package server

import (
	"fmt"
	"io"
	"net/url"

	"github.com/TradeWars/warehouse/types"
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
	return types.NewStatus(app.handlers, true, fmt.Sprintf("ssc version: %s, good luck out there survivors!", version)), nil
}
