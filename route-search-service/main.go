package main

import (
	"context"
	"fmt"
	"github.com/GOAT-prod/goatlogger"
	"github.com/GOAT-prod/goatsettings"
	"os"
	"os/signal"
	"route-search-service/settings"
	"syscall"
	"time"
)

func main() {
	mainCtx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	config, err := settings.Get()
	if err != nil {
		panic(err)
	}

	logger := goatlogger.New(config.AppName)

	app := NewApp(mainCtx, config, logger)

	app.initClients()
	app.initServices()
	app.initServer()

	app.Start()

	logger.Info(fmt.Sprintf("приложение запущено, порт: %d, конфиг: %s.json", config.Port, goatsettings.GetEnv()))

	waitTerminate(mainCtx, app.Stop)

	logger.Info("приложение остановлено")
}

func waitTerminate(mainCtx context.Context, quitFn func(ctx context.Context)) {
	const shutdownTimeout = 30 * time.Second
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	if quitFn == nil {
		return
	}

	quitCtx, cancelQuitCtx := context.WithTimeout(mainCtx, shutdownTimeout)
	defer cancelQuitCtx()

	quitFn(quitCtx)
}
