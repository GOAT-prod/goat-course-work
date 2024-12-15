package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/GOAT-prod/goathttp/client"
	"github.com/jmoiron/sqlx"
	"net/http"
	"warehouse-service/api"
	"warehouse-service/cluster/clientservice"
	"warehouse-service/database"
	"warehouse-service/database/broker"
	"warehouse-service/repository"
	"warehouse-service/service"
	"warehouse-service/settings"

	"github.com/GOAT-prod/goatlogger"
)

type App struct {
	ctx    context.Context
	logger goatlogger.Logger
	cfg    settings.Config

	server *api.Server

	postgres *sqlx.DB
	producer *broker.Producer

	clientServiceClient *clientservice.Client

	warehouseService service.WareHouse

	warehouseRepository repository.Warehouse
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

func (a *App) initKafka() {
	producer, err := broker.NewProducer(a.cfg.Databases.Kafka.Address, a.cfg.Databases.Kafka.Topic)
	if err != nil {
		a.logger.Panic(fmt.Sprintf("не удалось подключиться к broker: %v", err))
	}

	a.producer = producer
}

func (a *App) initRepositories() {
	a.warehouseRepository = repository.NewWarehouseRepository(a.postgres)
}

func (a *App) initClients() {
	a.clientServiceClient = clientservice.NewClient(client.NewBaseClient(a.cfg.Cluster.ClientService))
}

func (a *App) initServices() {
	a.warehouseService = service.NewWarehouseService(a.warehouseRepository, a.clientServiceClient, a.producer)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.cfg.Port)
	router.SetupRoutes(a.logger, a.warehouseService)

	a.server = api.NewServer(a.ctx, router)
}
