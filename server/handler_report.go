package server

import (
	"encoding/json"
	"io"
	"net/url"
	"time"

	"github.com/dyninc/qstring"
	"github.com/globalsign/mgo/bson"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"github.com/Southclaws/ScavengeSurviveCore/types"
)

func (app *App) reportRoutes() []Route {
	return []Route{
		{
			"reportCreate",
			"POST",
			"/store/reportCreate",
			types.ExampleReport(),
			types.ExampleStatus(bson.NewObjectId(), true),
			app.reportCreate,
		},
		{
			"reportArchive",
			"PATCH",
			"/store/reportArchive",
			"?id=" + bson.NewObjectId(),
			types.ExampleReport(),
			app.reportArchive,
		},
		{
			"reportGetList",
			"GET",
			"/store/reportGetList",
			nil,
			[]types.Report{types.ExampleReport(), types.ExampleReport()},
			app.reportGetList,
		},
		{
			"reportGet",
			"GET",
			"/store/reportGet",
			"?id=" + bson.NewObjectId(),
			types.ExampleReport(),
			app.reportGet,
		},
	}
}

func (app *App) reportCreate(r io.Reader, query url.Values) (status types.Status, err error) {
	var report types.Report
	err = json.NewDecoder(r).Decode(&report)
	if err != nil {
		return
	}
	err = app.validator.Struct(report)
	if err != nil {
		return types.NewStatusValidationError(err.(validator.ValidationErrors)), nil
	}

	logger.Debug("received request reportCreate",
		zap.Any("report", report))

	id, err := app.store.ReportCreate(report)
	return types.NewStatus(id, true, ""), err
}

type reportArchiveParams struct {
	ID      string
	Archive bool
}

func (app *App) reportArchive(r io.Reader, query url.Values) (status types.Status, err error) {
	params := reportArchiveParams{}
	err = qstring.Unmarshal(query, &params)
	if err != nil {
		return
	}

	logger.Debug("received request reportArchive",
		zap.Any("params", params))

	return types.NewStatus(nil, true, ""), app.store.ReportArchive(bson.ObjectIdHex(params.ID), params.Archive)
}

type reportGetListParams struct {
	PageSize, Page   int
	Archived, NoRead bool
	By, Of           string
	From, To         time.Time
}

func (app *App) reportGetList(r io.Reader, query url.Values) (status types.Status, err error) {
	params := reportGetListParams{}
	err = qstring.Unmarshal(query, &params)
	if err != nil {
		return
	}

	logger.Debug("received request reportGetList",
		zap.Any("params", params))

	var byID, ofID bson.ObjectId
	if bson.IsObjectIdHex(params.By) {
		byID = bson.ObjectIdHex(params.By)
	}
	if bson.IsObjectIdHex(params.Of) {
		ofID = bson.ObjectIdHex(params.Of)
	}

	var from, to *time.Time
	if !params.From.Equal(time.Time{}) {
		from = &params.From
	}
	if !params.To.Equal(time.Time{}) {
		to = &params.To
	}

	list, err := app.store.ReportGetList(
		params.PageSize, params.Page,
		params.Archived, params.NoRead,
		byID, ofID,
		from, to,
	)
	return types.NewStatus(list, true, ""), err
}

type reportGetParams struct {
	ID string
}

func (app *App) reportGet(r io.Reader, query url.Values) (status types.Status, err error) {
	params := reportGetParams{}
	err = qstring.Unmarshal(query, &params)
	if err != nil {
		return
	}

	if !bson.IsObjectIdHex(params.ID) {
		status = types.NewStatus(nil, false, "id not specified")
		return
	}

	logger.Debug("received request reportGet",
		zap.Any("params", params))

	report, err := app.store.ReportGet(bson.ObjectIdHex(params.ID))

	if err == nil {
		status = types.NewStatus(report, true, "")
	} else if err.Error() == "not found" {
		err = nil
		status = types.NewStatus(nil, false, "not found")
	}

	return
}
