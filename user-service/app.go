package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/GOAT-prod/goathttp/client"
	"github.com/jmoiron/sqlx"
	"net/http"
	"user-service/api"
	"user-service/cluster/notifier"
	"user-service/database"
	"user-service/repository"
	"user-service/service"
	"user-service/settings"

	"github.com/GOAT-prod/goatlogger"
)

type App struct {
	ctx    context.Context
	logger goatlogger.Logger
	cfg    settings.Config

	server *api.Server

	postgres *sqlx.DB

	userRepository repository.User
	roleRepository repository.Role
	userService    service.User
	notifierClient *notifier.Client
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
	a.userRepository = repository.NewUserRepository(a.postgres)
	a.roleRepository = repository.NewRoleRepository(a.postgres)
}

func (a *App) initCluster() {
	a.notifierClient = notifier.NewClient(client.NewBaseClient(a.cfg.Cluster.NotifierService))
}

func (a *App) initServices() {
	a.userService = service.NewUserService(a.userRepository, a.roleRepository, a.notifierClient)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.cfg.Port)
	router.SetupRoutes(a.logger, a.userService)

	a.server = api.NewServer(a.ctx, router)
}
