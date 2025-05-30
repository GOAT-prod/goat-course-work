package main

import (
	"api-gateway/settings"
	"context"
	"fmt"
	"github.com/GOAT-prod/goatsettings"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GOAT-prod/goatlogger"
)

// @title api-gateway
// @version 1.0
// @description Прослойка для взаимодействия с логистическим сервисом
// @securityDefinitions.apikey LogisticAuth
// @in header
// @name Authorization

func main() {
	config, err := settings.Parse()
	if err != nil {
		panic(err)
	}

	// BLYAT
	logger := goatlogger.New(config.AppName)

	mainCtx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	app := NewApp(mainCtx, logger, config)
	app.initClients()
	app.initServer()
	app.Start()

	logger.Info(fmt.Sprintf("приложение запушено, порт: %d, конфиг: %s.json", config.Port, goatsettings.GetEnv()))

	waitTerminate(mainCtx, app.Stop)
	logger.Info("приложение остановлено")
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
