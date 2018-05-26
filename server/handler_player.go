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
			true,
			types.ExamplePlayer(),
			types.ExampleStatus(true, true),
			app.playerCreate,
		},
		{
			"playerGet",
			"GET",
			"/store/playerGet",
			true,
			"?name=John, ?name=" + bson.NewObjectId(),
			types.ExamplePlayer(),
			app.playerGet,
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

	var player types.Player
	if params.Name != "" {
		player, err = app.store.PlayerGetByName(params.Name)
	} else if params.ID != "" {
		player, err = app.store.PlayerGetByID(bson.ObjectIdHex(params.ID))
	} else {
		status = types.NewStatus(nil, false, "must specify `name` or `id` query field")
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

	return types.NewStatus(nil, true, ""), app.store.PlayerUpdate(player.ID, player)
}
