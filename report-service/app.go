package main

import (
	"context"
	"errors"
	"fmt"
	goatclient "github.com/GOAT-prod/goathttp/client"
	"github.com/jmoiron/sqlx"
	"net/http"
	"report-service/api"
	"report-service/cluster/order"
	"report-service/database"
	"report-service/repository"
	"report-service/service"
	"report-service/settings"

	"github.com/GOAT-prod/goatlogger"
)

type App struct {
	ctx    context.Context
	logger goatlogger.Logger
	cfg    settings.Config

	server *api.Server

	postgres *sqlx.DB

	orderClient *order.Client

	reportService service.Report
	cronService   *service.CronService

	reportRepository repository.Report
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

	go a.cronService.Run()
}

func (a *App) Stop(_ context.Context) {
	if err := a.postgres.Close(); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось закрыть подключение к базе: %v", err))
	}

	a.cronService.Stop()
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
	a.reportRepository = repository.NewReportRepository(a.postgres)
}

func (a *App) initClients() {
	a.orderClient = order.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.OrderService))
}

func (a *App) initServices() {
	a.reportService = service.NewReportService(a.reportRepository)
	a.cronService = service.NewCronService(a.orderClient, a.reportRepository, a.cfg)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.cfg.Port)
	router.SetupRoutes(a.logger, a.reportService)

	a.server = api.NewServer(a.ctx, router)
}
