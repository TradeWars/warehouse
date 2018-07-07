package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"

	"github.com/TradeWars/warehouse/server"
)

func main() {
	config := &server.Config{}
	err := envconfig.Process("WAREHOUSE", config)
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
