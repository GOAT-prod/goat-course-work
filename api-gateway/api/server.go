package api

import (
	"api-gateway/api/handlers"
	"api-gateway/api/handlers/authhandlers"
	"api-gateway/api/handlers/carthandlers"
	"api-gateway/api/handlers/clienthandlers"
	"api-gateway/api/handlers/orderhandlers"
	"api-gateway/api/handlers/reporthandlers"
	"api-gateway/api/handlers/requesthandlers"
	"api-gateway/api/handlers/searchhandlers"
	"api-gateway/api/handlers/userhandlers"
	"api-gateway/api/handlers/warehousehandlers"
	"api-gateway/cluster/authservice"
	"api-gateway/cluster/cart"
	"api-gateway/cluster/clientservice"
	"api-gateway/cluster/order"
	"api-gateway/cluster/report"
	"api-gateway/cluster/request"
	"api-gateway/cluster/search"
	"api-gateway/cluster/userservice"
	"api-gateway/cluster/warehousesevice"
	"context"
	"fmt"
	"net"
	"net/http"

	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
)

type Router struct {
	port   int
	router *mux.Router
}

type Server struct {
	server *http.Server
}

func NewServer(ctx context.Context, router *Router) *Server {
	return &Server{
		server: &http.Server{
			Addr: fmt.Sprintf(":%d", router.port),
			BaseContext: func(_ net.Listener) context.Context {
				return ctx
			},
			Handler: router.router,
		},
	}
}

func NewRouter(logger goatlogger.Logger, port int) *Router {
	router := mux.NewRouter()
	router.Use(goathttp.CommonJsonMiddleware, goathttp.CORSMiddleware, goathttp.PanicRecoveryMiddleware(logger))

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodGet, http.MethodOptions)

	return &Router{
		port:   port,
		router: router,
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (r *Router) SetupRoutes(logger goatlogger.Logger, authClient *authservice.Client, userClient *userservice.Client, clientService *clientservice.Client,
	warehouseClient *warehousesevice.Client, cartClient *cart.Client, orderClient *order.Client, searchClient *search.Client, requestClient *request.Client, reportClient *report.Client) {
	r.router.PathPrefix("/swagger/").Handler(handlers.SwaggerHandler())

	//	auth-service
	{
		authSubRouter := r.router.PathPrefix("/auth").Subrouter()
		authSubRouter.HandleFunc("/login", authhandlers.LoginHandler(logger, authClient)).Methods(http.MethodPost, http.MethodOptions)
		authSubRouter.HandleFunc("/logout", authhandlers.LogoutHandler(logger, authClient)).Methods(http.MethodPost, http.MethodOptions)
		authSubRouter.HandleFunc("/register", authhandlers.RegistrationHandler(logger, authClient)).Methods(http.MethodPost, http.MethodOptions)

	}

	//	user-service
	{
		userSubRouter := r.router.PathPrefix("/user").Subrouter()
		userSubRouter.Use(goathttp.AuthMiddleware)
		userSubRouter.HandleFunc("/all", userhandlers.GetUsersHandler(logger, userClient)).Methods(http.MethodGet, http.MethodOptions)
		userSubRouter.HandleFunc("/{id}", userhandlers.GetUserHandler(logger, userClient)).Methods(http.MethodGet, http.MethodOptions)
		userSubRouter.HandleFunc("", userhandlers.AddUserHandler(logger, userClient)).Methods(http.MethodPost, http.MethodOptions)
		userSubRouter.HandleFunc("", userhandlers.UpdateUserHandler(logger, userClient)).Methods(http.MethodPut, http.MethodOptions)
		userSubRouter.HandleFunc("/{id}", userhandlers.DeleteUserHandler(logger, userClient)).Methods(http.MethodDelete, http.MethodOptions)
	}

	//	client-service
	{
		clientSubRouter := r.router.PathPrefix("/client").Subrouter()
		clientSubRouter.Use(goathttp.AuthMiddleware)
		clientSubRouter.HandleFunc("/all", clienthandlers.GetClientsHandler(logger, clientService)).Methods(http.MethodGet, http.MethodOptions)
		clientSubRouter.HandleFunc("/{id}", clienthandlers.GetClientHandler(logger, clientService)).Methods(http.MethodGet, http.MethodOptions)
		clientSubRouter.HandleFunc("", clienthandlers.UpdateClientHandler(logger, clientService)).Methods(http.MethodPut, http.MethodOptions)
		clientSubRouter.HandleFunc("/{id}", clienthandlers.DeleteClientHandler(logger, clientService)).Methods(http.MethodDelete, http.MethodOptions)

	}

	//	warehouse-service
	{
		warehouseSubRouter := r.router.PathPrefix("/products").Subrouter()
		warehouseSubRouter.Use(goathttp.AuthMiddleware)
		warehouseSubRouter.HandleFunc("/materials", warehousehandlers.GetMaterialsHandler(logger, warehouseClient)).Methods(http.MethodGet, http.MethodOptions)
		warehouseSubRouter.HandleFunc("", warehousehandlers.GetProductsHandler(logger, warehouseClient)).Methods(http.MethodGet, http.MethodOptions)
		warehouseSubRouter.HandleFunc("/{id}", warehousehandlers.GetProductHandler(logger, warehouseClient)).Methods(http.MethodGet, http.MethodOptions)
		warehouseSubRouter.HandleFunc("", warehousehandlers.AddProductsHandler(logger, warehouseClient)).Methods(http.MethodPost, http.MethodOptions)
		warehouseSubRouter.HandleFunc("", warehousehandlers.UpdateProductsHandler(logger, warehouseClient)).Methods(http.MethodPut, http.MethodOptions)
		warehouseSubRouter.HandleFunc("", warehousehandlers.DeleteProductsHandler(logger, warehouseClient)).Methods(http.MethodDelete, http.MethodOptions)
	}

	//	cart-service
	{
		cartSubRouter := r.router.PathPrefix("/cart").Subrouter()
		cartSubRouter.Use(goathttp.AuthMiddleware)
		cartSubRouter.HandleFunc("", carthandlers.GetCartHandler(logger, cartClient)).Methods(http.MethodGet, http.MethodOptions)
		cartSubRouter.HandleFunc("/item", carthandlers.AddCartItemHandler(logger, cartClient)).Methods(http.MethodPost, http.MethodOptions)
		cartSubRouter.HandleFunc("/item", carthandlers.UpdateCartItemHandler(logger, cartClient)).Methods(http.MethodPut, http.MethodOptions)
		cartSubRouter.HandleFunc("/item/{id}", carthandlers.DeleteCartItemHandler(logger, cartClient)).Methods(http.MethodDelete, http.MethodOptions)
		cartSubRouter.HandleFunc("/clear", carthandlers.ClearCartHandler(logger, cartClient)).Methods(http.MethodDelete, http.MethodOptions)
	}

	//	order-service
	{
		orderSubRouter := r.router.PathPrefix("/order").Subrouter()
		orderSubRouter.Use(goathttp.AuthMiddleware)
		orderSubRouter.HandleFunc("", orderhandlers.CreateOrderHandler(logger, orderClient)).Methods(http.MethodPost, http.MethodOptions)
		orderSubRouter.HandleFunc("/all", orderhandlers.GetUserOrdersHandler(logger, orderClient)).Methods(http.MethodGet, http.MethodOptions)
	}

	//	search-service
	{
		searchSubRouter := r.router.PathPrefix("/search").Subrouter()
		searchSubRouter.Use(goathttp.AuthMiddleware)
		searchSubRouter.HandleFunc("/filters", searchhandlers.GetActiveFiltersHandler(logger, searchClient)).Methods(http.MethodGet, http.MethodOptions)
		searchSubRouter.HandleFunc("/catalog", searchhandlers.GetCatalogHandler(logger, searchClient)).Methods(http.MethodGet, http.MethodOptions)
		searchSubRouter.HandleFunc("/catalog/product/{id}", searchhandlers.GetProductCatalogHandler(logger, searchClient)).Methods(http.MethodGet, http.MethodOptions)
	}

	//	request-service
	{
		requestSubRouter := r.router.PathPrefix("/request").Subrouter()
		requestSubRouter.Use(goathttp.AuthMiddleware)
		requestSubRouter.HandleFunc("/all", requesthandlers.GetRequestsHandler(logger, requestClient)).Methods(http.MethodGet, http.MethodOptions)
		requestSubRouter.HandleFunc("/{requestId}/{status}", requesthandlers.UpdateRequestStatusHandler(logger, requestClient)).Methods(http.MethodPut, http.MethodOptions)
	}

	//	report-service
	{
		reportSubRouter := r.router.PathPrefix("/report").Subrouter()
		reportSubRouter.Use(goathttp.AuthMiddleware)
		reportSubRouter.HandleFunc("/sell/{userId}/{date}", reporthandlers.GetSellReportHandlers(logger, reportClient)).Methods(http.MethodGet, http.MethodOptions)
		reportSubRouter.HandleFunc("/order/{userId}/{date}", reporthandlers.GetOrderReportHandler(logger, reportClient)).Methods(http.MethodGet, http.MethodOptions)
	}
}
