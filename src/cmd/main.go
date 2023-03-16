package main

import (
	"github.com/cassa10/arq2-tp1/src/infrastructure/api"
	"github.com/cassa10/arq2-tp1/src/infrastructure/config"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
)

func main() {
	conf := config.LoadConfig()
	baseLogger := logger.New(conf.Environment, conf.LogLevel)

	app := api.NewApplication(baseLogger, conf)

	baseLogger.Fatal(app.Run())
}
