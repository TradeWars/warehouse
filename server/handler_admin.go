package server

import (
	"io"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func (app App) adminRoutes() []Route {
	return []Route{
		{
			"AdminSet",
			"POST",
			"/store/admin/{id}",
			true,
			map[string]int{"level": 3},
			types.ExampleStatus(true, true),
			app.adminSet,
		},
	}
}

func (app App) adminSet(r io.Reader) (status types.Status, err error) {
	return
}
