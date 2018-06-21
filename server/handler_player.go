package server

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"

	"github.com/dyninc/qstring"
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
			types.ExamplePlayer(),
			types.ExampleStatus(bson.NewObjectId(), true),
			app.playerCreate,
		},
		{
			"playerGet",
			"GET",
			"/store/playerGet",
			"?name=John, ?id=" + bson.NewObjectId(),
			types.ExamplePlayer(),
			app.playerGet,
		},
		{
			"playerUpdate",
			"PATCH",
			"/store/playerUpdate",
			types.ExamplePlayer(),
			types.ExampleStatus(nil, true),
			app.playerUpdate,
		},
	}
}

func (app *App) playerCreate(r io.Reader, query url.Values) (status types.Status, err error) {
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

type playerGetParams struct {
	Name string
	ID   string
}

func (app *App) playerGet(r io.Reader, query url.Values) (status types.Status, err error) {
	params := playerGetParams{}
	err = qstring.Unmarshal(query, &params)
	if err != nil {
		return
	}

	logger.Debug("received request playerGet",
		zap.Any("params", params))

	var player types.Player
	if params.Name != "" {
		player, err = app.store.PlayerGetByName(params.Name)
	} else if params.ID != "" {
		if !bson.IsObjectIdHex(params.ID) {
			status = types.NewStatus(nil, false, "invalid id format")
			return
		} else {
			player, err = app.store.PlayerGetByID(bson.ObjectIdHex(params.ID))
		}
	} else {
		status = types.NewStatus(nil, false, "id or name not specified")
		return
	}

	if err == nil {
		status = types.NewStatus(player, true, "")
	} else if err.Error() == "not found" {
		err = nil
		status = types.NewStatus(nil, false, "not found")
	}

	return
}

func (app *App) playerUpdate(r io.Reader, query url.Values) (status types.Status, err error) {
	var player types.Player
	err = json.NewDecoder(r).Decode(&player)
	if err != nil {
		return
	}
	err = app.validator.Struct(player)
	if err != nil {
		return types.NewStatusValidationError(err.(validator.ValidationErrors)), nil
	}

	logger.Debug("received request playerUpdate",
		zap.Any("player", player))

	return types.NewStatus(nil, true, ""), app.store.PlayerUpdate(player.ID, player)
}
