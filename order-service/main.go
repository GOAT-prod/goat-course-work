package main

import (
	"context"
	"fmt"
	"github.com/GOAT-prod/goatlogger"
	"github.com/GOAT-prod/goatsettings"
	"github.com/shopspring/decimal"
	"order-service/settings"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	decimalSettings()

	config, err := settings.Parse()
	if err != nil {
		panic(err)
	}

	logger := goatlogger.New(config.AppName)

	mainCtx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	app := NewApp(mainCtx, logger, config)
	app.initDatabases()
	app.initRepositories()
	app.initServices()
	app.initServer()
	app.Start()

	logger.Info(fmt.Sprintf("приложение запушено, порт: %d, конфиг: %s.json", config.Port, goatsettings.GetEnv()))

	waitTerminate(mainCtx, app.Stop)

	logger.Info("приложение остановлено")
}

func decimalSettings() {
	decimal.MarshalJSONWithoutQuotes = true
	decimal.DivisionPrecision = 2
}

func waitTerminate(mainCtx context.Context, quitFn func(ctx context.Context)) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	if quitFn == nil {
		return
	}

	quitCtx, cancelQuitCtx := context.WithTimeout(mainCtx, time.Second*15)
	defer cancelQuitCtx()

	quitFn(quitCtx)
}
