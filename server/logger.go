package server

import (
	"os"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	var config zap.Config

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	testing, _ := strconv.ParseBool(os.Getenv("TESTING"))

	if testing {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeTime = nil
		config.DisableCaller = true
	} else {
		config = zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		if debug {
			dyn := zap.NewAtomicLevel()
			dyn.SetLevel(zap.DebugLevel)
			config.Level = dyn
		}
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		panic(err)
	}

	logger.Debug("debug logging active")
}
