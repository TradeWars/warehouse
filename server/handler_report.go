package server

import (
	"io"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func (app App) reportRoutes() []Route {
	return []Route{
		{
			"ReportCreate",
			"POST",
			"/report",
			true,
			types.ExampleReport(),
			types.ExampleStatus(true, true),
			app.reportCreate,
		},
		{
			"ReportRemove",
			"DELETE",
			"/report",
			true,
			nil,
			types.ExampleReport(),
			app.reportRemove,
		},
		{
			"ReportGetList",
			"GET",
			"/reports",
			true,
			nil,
			nil,
			app.reportGetList,
		},
		{
			"ReportGetInfo",
			"GET",
			"/report",
			true,
			nil,
			nil,
			app.reportGetInfo,
		},
	}
}

func (app App) reportCreate(r io.Reader) (status types.Status, err error) {
	return
}

func (app App) reportRemove(r io.Reader) (status types.Status, err error) {
	return
}

func (app App) reportGetList(r io.Reader) (status types.Status, err error) {
	return
}

func (app App) reportGetInfo(r io.Reader) (status types.Status, err error) {
	return
}
