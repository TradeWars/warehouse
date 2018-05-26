package server

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/globalsign/mgo/bson"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func (app *App) playerRoutes() []Route {
	return []Route{
		{
			"playerCreate",
			"POST",
			"/store/playerCreate",
			true,
			types.ExamplePlayer(),
			types.ExampleStatus(true, true),
			app.playerCreate,
		},
		{
			"playerGetByName",
			"GET",
			"/store/playerGetByName",
			true,
			"John",
			types.ExamplePlayer(),
			app.playerGetByName,
		},
		{
			"playerGetByID",
			"GET",
			"/store/playerGetByID",
			true,
			bson.NewObjectId(),
			types.ExamplePlayer(),
			app.playerGetByID,
		},
		{
			"playerUpdate",
			"PATCH",
			"/store/playerUpdate",
			true,
			types.ExamplePlayer(),
			types.ExampleStatus(true, true),
			app.playerUpdate,
		},
	}
}

func (app *App) playerCreate(r io.Reader) (status types.Status, err error) {
	var player types.Player
	err = json.NewDecoder(r).Decode(&player)
	if err != nil {
		return
	}
	err = app.validator.Struct(player)
	if err != nil {
		return types.NewStatusValidationError(err.(validator.ValidationErrors)), nil
	}

	logger.Debug("received request playerCreate",
		zap.Any("player", player))

	id, err := app.store.PlayerCreate(player)
	if err != nil && strings.HasPrefix(err.Error(), "E11000") {
		return types.NewStatus(nil, false, "player name already registered"), nil
	}
	return types.NewStatus(id, true, ""), err
}

func (app *App) playerGetByName(r io.Reader) (status types.Status, err error) {
	var name string
	err = json.NewDecoder(r).Decode(&name)
	if err != nil {
		return
	}

	player, err := app.store.PlayerGetByName(name)
	status = types.NewStatus(player, true, "")

	return
}

func (app *App) playerGetByID(r io.Reader) (status types.Status, err error) {
	var id bson.ObjectId
	err = json.NewDecoder(r).Decode(&id)
	if err != nil {
		return
	}

	player, err := app.store.PlayerGetByID(id)
	status = types.NewStatus(player, true, "")

	return
}

func (app *App) playerUpdate(r io.Reader) (status types.Status, err error) {
	var player types.Player
	err = json.NewDecoder(r).Decode(&player)
	if err != nil {
		return
	}
	err = app.validator.Struct(player)
	if err != nil {
		return types.NewStatusValidationError(err.(validator.ValidationErrors)), nil
	}

	return types.NewStatus(nil, true, ""), app.store.PlayerUpdate(player.ID, player)
}
