package server

import (
	"io"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func playerRoutes(app App) []Route {
	return []Route{
		{
			"PlayerCreate",
			"POST",
			"/store/player/{id}",
			true,
			types.ExamplePlayer(),
			types.ExampleStatus(true, true),
			app.playerCreate,
		},
		{
			"PlayerGet",
			"GET",
			"/store/player/{id}",
			true,
			nil,
			types.ExamplePlayer(),
			app.playerGet,
		},
		{
			"PlayerUpdate",
			"PATCH",
			"/store/player/{id}",
			true,
			types.ExamplePlayer(),
			types.ExampleStatus(true, true),
			app.playerUpdate,
		},
	}
}

func (app App) playerCreate(r io.ReadCloser) (status types.Status, err error) {
	return
}

func (app App) playerGet(r io.ReadCloser) (status types.Status, err error) {
	return
}

func (app App) playerUpdate(r io.ReadCloser) (status types.Status, err error) {
	return
}
