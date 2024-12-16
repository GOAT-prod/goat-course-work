package main

import (
	"cart-service/api"
	"cart-service/cluster/warehouse"
	"cart-service/database"
	"cart-service/repository"
	"cart-service/service"
	"cart-service/settings"
	"context"
	"errors"
	"fmt"
	goatclient "github.com/GOAT-prod/goathttp/client"
	"github.com/jmoiron/sqlx"
	"net/http"

	"github.com/GOAT-prod/goatlogger"
)

type App struct {
	ctx    context.Context
	logger goatlogger.Logger
	cfg    settings.Config

	server *api.Server

	postgres *sqlx.DB

	warehouseClient *warehouse.Client
	cartService     service.Cart
	cartRepository  repository.Cart
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
	a.cartRepository = repository.NewCartRepository(a.postgres)
}

func (a *App) initClients() {
	a.warehouseClient = warehouse.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.WarehouseClient))
}

func (a *App) initServices() {
	a.cartService = service.NewCartServiceImpl(a.cartRepository, a.warehouseClient)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.cfg.Port)
	router.SetupRoutes(a.logger, a.cartService)

	a.server = api.NewServer(a.ctx, router)
}
