package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/client"
	"github.com/GOAT-prod/goatlogger"
	"github.com/jmoiron/sqlx"
	"net/http"
	"request-service/api"
	"request-service/cluster/notifier"
	"request-service/cluster/warehouse"
	"request-service/database"
	"request-service/kafka"
	"request-service/repository"
	"request-service/service"
	"request-service/settings"
)

type App struct {
	ctx    context.Context
	logger goatlogger.Logger
	cfg    settings.Config

	server *api.Server

	postgres               *sqlx.DB
	approveProductConsumer *kafka.Consumer
	supplyProductConsumer  *kafka.Consumer

	warehouseClient *warehouse.Client
	notifierClient  *notifier.Client

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

	go a.approveProductConsumer.Consume(goatcontext.Context{}, a.logger)
	go a.supplyProductConsumer.Consume(goatcontext.Context{}, a.logger)
}

func (a *App) Stop(_ context.Context) {
	if err := a.postgres.Close(); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось закрыть подключение к базе: %v", err))
	}

	if err := a.approveProductConsumer.Stop(); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось закрыть консюмер подтверждения продутов: %v", err))
	}

	if err := a.supplyProductConsumer.Stop(); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось закрыть консюмер необходимых поставок: %v", err))
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

func (a *App) initKafka() {
	approveProductMessageHandler := kafka.NewMessageHandler(a.requestRepository)
	supplyProductMessageHandler := kafka.NewMessageHandler(a.requestRepository)

	approveProductConsumer, err := kafka.NewConsumer(approveProductMessageHandler, a.cfg.Kafka.Address, a.cfg.Kafka.ProductApproveTopic)
	if err != nil {
		a.logger.Panic(fmt.Sprintf("не инициализировать космюмер подтверждения продуктов: %v", err))
	}

	supplyProductConsumer, err := kafka.NewConsumer(supplyProductMessageHandler, a.cfg.Kafka.Address, a.cfg.Kafka.SupplyProductsTopic)
	if err != nil {
		a.logger.Panic(fmt.Sprintf("не инициализировать космюмер поставки продуктов: %v", err))
	}

	a.approveProductConsumer = approveProductConsumer
	a.supplyProductConsumer = supplyProductConsumer
}

func (a *App) initRepositories() {
	a.requestRepository = repository.NewRequestRepository(a.postgres)
}

func (a *App) initClients() {
	a.warehouseClient = warehouse.NewClient(client.NewBaseClient(a.cfg.Cluster.WarehouseService))
	a.notifierClient = notifier.NewClient(client.NewBaseClient(a.cfg.Cluster.NotifierService))
}

func (a *App) initServices() {
	a.requestService = service.NewRequestService(a.requestRepository, a.warehouseClient, a.notifierClient)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.cfg.Port)
	router.SetupRoutes(a.logger, a.requestService)

	a.server = api.NewServer(a.ctx, router)
}
