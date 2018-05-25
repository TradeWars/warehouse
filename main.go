package main

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/Southclaws/ScavengeSurviveCore/server"
)

func main() {
	config := server.Config{}
	err := envconfig.Process("SSC", &config)
	if err != nil {
		panic(err)
	}

	server.Start(config)
}
