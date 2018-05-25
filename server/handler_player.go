package server

import (
	"encoding/json"
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

func (app App) playerCreate(r io.Reader) (status types.Status, err error) {
	var player types.Player
	err = json.NewDecoder(r).Decode(&player)
	if err != nil {
		return
	}
	err = app.validator.Struct(player)
	if err != nil {
		return
	}

	return types.NewStatus(nil, true, ""), app.store.PlayerCreate(player)
}

func (app App) playerGet(r io.Reader) (status types.Status, err error) {
	return
}

func (app App) playerUpdate(r io.Reader) (status types.Status, err error) {
	return
}
