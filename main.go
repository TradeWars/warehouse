package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"

	"github.com/Southclaws/ScavengeSurviveCore/server"
)

func main() {
	config := &server.Config{}
	err := envconfig.Process("SSC", config)
	if err != nil {
		panic(err)
	}

	app, err := server.Initialise(config)
	if err != nil {
		panic(err)
	}

	err = app.Start()

	panic(err)
}
