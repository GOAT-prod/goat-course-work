package main

import (
	"api-gateway/api"
	"api-gateway/cluster/authservice"
	"api-gateway/cluster/cart"
	"api-gateway/cluster/clientservice"
	"api-gateway/cluster/order"
	"api-gateway/cluster/report"
	"api-gateway/cluster/request"
	"api-gateway/cluster/route"
	"api-gateway/cluster/search"
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
	cartServiceClient      *cart.Client
	orderServiceCLinet     *order.Client
	searchServiceClient    *search.Client
	requestServiceClient   *request.Client
	reportServiceClient    *report.Client
	routeServiceClient     *route.Client
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
	a.cartServiceClient = cart.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.CartService))
	a.orderServiceCLinet = order.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.OrderService))
	a.searchServiceClient = search.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.SearchService))
	a.requestServiceClient = request.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.RequestService))
	a.reportServiceClient = report.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.ReportService))
	a.routeServiceClient = route.NewClient(goatclient.NewBaseClient(a.cfg.Cluster.RouteService))
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.cfg.Port)
	router.SetupRoutes(
		a.logger,
		a.authServiceClient,
		a.userServiceClient,
		a.clientServiceClient,
		a.warehouseServiceClient,
		a.cartServiceClient,
		a.orderServiceCLinet,
		a.searchServiceClient,
		a.requestServiceClient,
		a.reportServiceClient,
		a.routeServiceClient)

	a.server = api.NewServer(a.ctx, router)
}
