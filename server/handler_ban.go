package server

import (
	"io"
	"net/url"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func (app App) banRoutes() []Route {
	return []Route{
		{
			"BanIO_Create",
			"POST",
			"/ban",
			true,
			nil,
			nil,
			app.banCreate,
		},
		{
			"BanIO_Remove",
			"DELETE",
			"/ban",
			true,
			nil,
			nil,
			app.banRemove,
		},
		{
			"BanIO_Update",
			"PATCH",
			"/ban",
			true,
			nil,
			nil,
			app.banUpdate,
		},
		{
			"BanIO_GetList",
			"GET",
			"/bans",
			true,
			nil,
			nil,
			app.banGetList,
		},
		{
			"BanIO_GetInfo",
			"GET",
			"/ban",
			true,
			nil,
			nil,
			app.banGetInfo,
		},
	}
}

func (app App) banCreate(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}

func (app App) banRemove(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}

func (app App) banUpdate(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}

func (app App) banGetList(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}

func (app App) banGetInfo(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}
