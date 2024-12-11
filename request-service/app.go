package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/GOAT-prod/goatlogger"
	"github.com/jmoiron/sqlx"
	"net/http"
	"request-service/api"
	"request-service/database"
	"request-service/repository"
	"request-service/service"
	"request-service/settings"
)

type App struct {
	ctx    context.Context
	logger goatlogger.Logger
	cfg    settings.Config

	server *api.Server

	postgres *sqlx.DB

	requestRepository repository.Request
	requestService    service.Request
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
	a.initKafka()
}

func (a *App) initPostgres() {
	a.postgres = database.ConnectPostgres(a.cfg.Databases.Postgres)

	if err := database.RunMigrations(a.postgres); err != nil {
		a.logger.Panic(fmt.Sprintf("не удалось прогнать миграции: %v", err))
	}
}

func (a *App) initKafka() {}

func (a *App) initRepositories() {
	a.requestRepository = repository.NewRequestRepository(a.postgres)
}

func (a *App) initServices() {
	a.requestService = service.NewRequestService(a.requestRepository)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.cfg.Port)
	router.SetupRoutes(a.logger, a.requestService)

	a.server = api.NewServer(a.ctx, router)
}
