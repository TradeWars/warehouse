package server

import (
	"io"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func banRoutes(app App) []Route {
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

func (app App) banCreate(r io.ReadCloser) (status types.Status, err error) {
	return
}

func (app App) banRemove(r io.ReadCloser) (status types.Status, err error) {
	return
}

func (app App) banUpdate(r io.ReadCloser) (status types.Status, err error) {
	return
}

func (app App) banGetList(r io.ReadCloser) (status types.Status, err error) {
	return
}

func (app App) banGetInfo(r io.ReadCloser) (status types.Status, err error) {
	return
}
