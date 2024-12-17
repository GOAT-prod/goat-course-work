package main

import (
	"context"
	"errors"
	"fmt"
	goatclient "github.com/GOAT-prod/goathttp/client"
	"github.com/jmoiron/sqlx"
	"net/http"
	"order-service/api"
	"order-service/cluster/warehouse"
	"order-service/settings"

	"github.com/GOAT-prod/goatlogger"
)

type App struct {
	ctx    context.Context
	logger goatlogger.Logger
	cfg    settings.Config

	server *api.Server

	postgres *sqlx.DB

	warehouseClient *warehouse.Client
}

func NewApp(ctx context.Context, logger goatlogger.Logger, cfg settings.Config) *App {
	return &App{
		ctx:    ctx,
		logger: logger,
		cfg:    cfg,
	}
}

func (a *App) Start() {
	go func() {
		if err := a.server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error(fmt.Sprintf("приложение неожиданно остановлено, ошибка: %v", err))
		}
	}()
}

func (a *App) Stop(_ context.Context) {
	if err := a.postgres.Close(); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось закрыть подключение к базе: %v", err))
	}
}

func (a *App) initDatabases() {
	a.initPostgres()
}

func (a *App) initPostgres() {
	a.postgres = database.ConnectPostgres(a.cfg.Databases.Postgres)

	if err := database.RunMigrations(a.postgres); err != nil {
		a.logger.Panic(fmt.Sprintf("не удалось прогнать миграции: %v", err))
	}
}

func (a *App) initRepositories() {
}

func (a *App) initClients() {
	a.warehouseClient = warehouse.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.WarehouseClient))
}

func (a *App) initServices() {
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.cfg.Port)
	router.SetupRoutes(a.logger)

	a.server = api.NewServer(a.ctx, router)
}
