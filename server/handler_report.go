package server

import (
	"io"
	"net/url"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func (app App) reportRoutes() []Route {
	return []Route{
		{
			"ReportCreate",
			"POST",
			"/report",
			types.ExampleReport(),
			types.ExampleStatus(true, true),
			app.reportCreate,
		},
		{
			"ReportRemove",
			"DELETE",
			"/report",
			nil,
			types.ExampleReport(),
			app.reportRemove,
		},
		{
			"ReportGetList",
			"GET",
			"/reports",
			nil,
			nil,
			app.reportGetList,
		},
		{
			"ReportGetInfo",
			"GET",
			"/report",
			nil,
			nil,
			app.reportGetInfo,
		},
	}
}

func (app App) reportCreate(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}

func (app App) reportRemove(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}

func (app App) reportGetList(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}

func (app App) reportGetInfo(r io.Reader, query url.Values) (status types.Status, err error) {
	return
}
