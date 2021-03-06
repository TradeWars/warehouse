package server

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-resty/resty"
)

var client *resty.Client
var app *App
var appCache *App

func TestMain(m *testing.M) {
	var err error

	config := &Config{
		Temporary: false,
		Bind:      "0.0.0.0:7788",
		Auth:      "secret_key",
		MongoHost: "localhost",
		MongoPort: "27017",
		MongoName: "ss",
	}
	configCache := &Config{
		Temporary: true,
	}

	client = resty.New().
		SetRESTMode().
		SetHeader("Authorization", "secret_key").
		SetHostURL("http://" + config.Bind)

	app, err = Initialise(config)
	if err != nil {
		panic(err)
	}
	appCache, err = Initialise(configCache)
	if err != nil {
		panic(err)
	}
	logger.Info("clearing database before running tests")
	err = app.store.DeleteEverythingPermanently()
	if err != nil {
		panic(err)
	}

	go app.Start()

	os.Exit(m.Run())

	err = app.Stop()
	if err != nil {
		fmt.Println(err)
	}
}
