package server

import (
	"io"
	"net/url"
	"time"

	"github.com/globalsign/mgo/bson"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func (app App) banRoutes() []Route {
	return []Route{
		{
			"banCreate",
			"POST",
			"/store/banCreate",
			types.ExampleBan(),
			types.ExampleStatus(bson.NewObjectId(), true),
			app.banCreate,
		},
		{
			"banArchive",
			"PATCH",
			"/store/banArchive",
			"id, archive",
			types.ExampleStatus(nil, true),
			app.banArchive,
		},
		{
			"banUpdate",
			"PATCH",
			"/store/banUpdate",
			"id, archive",
			types.ExampleStatus(nil, true),
			app.banUpdate,
		},
		{
			"banGetList",
			"GET",
			"/store/banGetList",
			"pagesize, page, archived, by, of, from, to",
			[]types.Ban{types.ExampleBan(), types.ExampleBan()},
			app.banGetList,
		},
		{
			"banGetInfo",
			"GET",
			"/store/banGetInfo",
			"id",
			types.ExampleBan(),
			app.banGetInfo,
		},
	}
}

func (app App) banCreate(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}

type banArchiveParams struct {
	ID      string
	Archive bool
}

func (app App) banArchive(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}

func (app App) banUpdate(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}

type banGetListParams struct {
	PageSize, Page int
	Archived       bool
	By, Of         string
	From, To       time.Time
}

func (app App) banGetList(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}

func (app App) banGetInfo(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}
