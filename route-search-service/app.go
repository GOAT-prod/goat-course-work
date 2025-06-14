package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/GOAT-prod/goathttp/client"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
	"route-search-service/api"
	"route-search-service/maps"
	"route-search-service/service"
	"route-search-service/settings"
)

type App struct {
	mainCtx context.Context
	logger  goatlogger.Logger
	config  settings.Config

	server *api.Server

	mapsClient *maps.Client

	routeService service.Route
}

func NewApp(ctx context.Context, config settings.Config, logger goatlogger.Logger) *App {
	return &App{
		mainCtx: ctx,
		logger:  logger,
		config:  config,
	}
}

func (a *App) Start() {
	go func() {
		if err := a.server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error(fmt.Sprintf("приложение неожиданно остановлено, ошибка: %v", err))
		}
	}()
}

func (a *App) initClients() {
	a.mapsClient = maps.NewClient(a.config.MapsApiKey, client.NewBaseClient(a.config.MapsApiUrl))
}

func (a *App) initServices() {
	a.routeService = service.NewRouteService(a.mapsClient)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.config.Port)
	router.SetupRoutes(a.logger, a.routeService)

	a.server = api.NewServer(a.mainCtx, router)
}

func (a *App) Stop(shutdownCtx context.Context) {
	if err := a.server.Shutdown(shutdownCtx); err != nil {
		a.logger.Error(err.Error())
	}
}
