package main

import (
	"api-gateway/api"
	"api-gateway/cluster/authservice"
	"api-gateway/cluster/clientservice"
	"api-gateway/cluster/userservice"
	"api-gateway/cluster/warehousesevice"
	"api-gateway/settings"
	"context"
	"errors"
	"fmt"
	goatclient "github.com/GOAT-prod/goathttp/client"
	"net/http"

	"github.com/GOAT-prod/goatlogger"
)

type App struct {
	ctx    context.Context
	logger goatlogger.Logger
	cfg    settings.Config

	server *api.Server

	authServiceClient      *authservice.Client
	userServiceClient      *userservice.Client
	clientServiceClient    *clientservice.Client
	warehouseServiceClient *warehousesevice.Client
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

}

func (a *App) initClients() {
	a.authServiceClient = authservice.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.AuthService))
	a.userServiceClient = userservice.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.UserService))
	a.clientServiceClient = clientservice.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.ClientService))
	a.warehouseServiceClient = warehousesevice.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.WareHouseService))
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.cfg.Port)
	router.SetupRoutes(a.logger, a.authServiceClient, a.userServiceClient, a.clientServiceClient, a.warehouseServiceClient)

	a.server = api.NewServer(a.ctx, router)
}
