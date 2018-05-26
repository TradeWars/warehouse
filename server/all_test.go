package server

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-resty/resty"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var client *resty.Client
var app *App

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	config := &Config{}
	err = envconfig.Process("SSC", config)
	if err != nil {
		panic(err)
	}

	client = resty.New().
		SetRESTMode().
		SetHeader("Authorization", "secret_key").
		SetHostURL("http://" + config.Bind)

	app, err = Initialise(config)
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
