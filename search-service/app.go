package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/GOAT-prod/goathttp/client"
	"github.com/GOAT-prod/goatlogger"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"search-service/api"
	"search-service/cluster/warehouse"
	"search-service/database"
	"search-service/repository"
	"search-service/service"
	"search-service/settings"
	"time"
)

type App struct {
	mainCtx context.Context
	logger  goatlogger.Logger
	cfg     settings.Config

	server *api.Server

	mongoClient *mongo.Client
	redisClient *redis.Client

	filterRepository repository.Filter
	cacheRepository  repository.Cache

	warehouseClient *warehouse.Client

	searchService service.Search
}

func NewApp(mainCtx context.Context, logger goatlogger.Logger, cfg settings.Config) *App {
	return &App{
		mainCtx: mainCtx,
		logger:  logger,
		cfg:     cfg,
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
	if err := a.mongoClient.Disconnect(a.mainCtx); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось отключиться от mongoDB: %v", err))
	}

	if err := a.redisClient.Close(); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось отключиться от redis: %v", err))
	}
}

func (a *App) initDatabases() {
	a.initMongo()
	a.initRedis()
}

func (a *App) initMongo() {
	mongoCtx, cancelFunc := context.WithTimeout(a.mainCtx, 15*time.Second)
	defer cancelFunc()

	mongoClient, err := database.MongoConnect(mongoCtx, a.cfg.Databases.Mongo.Connection)
	if err != nil {
		a.logger.Panic(fmt.Sprintf("не удалось подключиться к mongoDb, ошибка: %v", err))
	}

	a.mongoClient = mongoClient
}

func (a *App) initRedis() {
	redisCtx, cancelFunc := context.WithTimeout(a.mainCtx, 15*time.Second)
	defer cancelFunc()

	redisClient, err := database.NewRedisClient(redisCtx, a.cfg.Databases.Redis)
	if err != nil {
		a.logger.Panic(fmt.Sprintf("не удалось подключиться к redis: %v", err))
	}

	a.redisClient = redisClient
}

func (a *App) initRepositories() {
	a.filterRepository = repository.NewFilterRepository(a.mongoClient, a.cfg.Databases.Mongo.Database, a.cfg.Databases.Mongo.Connection)
	a.cacheRepository = repository.NewCacheRepository(a.redisClient)
}

func (a *App) initClients() {
	a.warehouseClient = warehouse.NewClient(client.NewBaseClient(a.cfg.Cluster.WarehouseService))
}

func (a *App) initServices() {
	a.searchService = service.New(a.filterRepository, a.cacheRepository, a.warehouseClient)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.cfg.Port)
	router.SetupRoutes(a.logger, a.searchService)

	a.server = api.NewServer(a.mainCtx, router)
}
