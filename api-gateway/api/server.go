package api

import (
	"api-gateway/api/handlers"
	"api-gateway/api/handlers/authhandlers"
	"api-gateway/api/handlers/carthandlers"
	"api-gateway/api/handlers/clienthandlers"
	"api-gateway/api/handlers/userhandlers"
	"api-gateway/api/handlers/warehousehandlers"
	"api-gateway/cluster/authservice"
	"api-gateway/cluster/cart"
	"api-gateway/cluster/clientservice"
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

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodGet)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodGet)

	return &Router{
		port:   port,
		router: router,
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (r *Router) SetupRoutes(logger goatlogger.Logger, authClient *authservice.Client, userClient *userservice.Client, clientService *clientservice.Client, warehouseClient *warehousesevice.Client, cartClient *cart.Client) {
	r.router.PathPrefix("/swagger/").Handler(handlers.SwaggerHandler())

	//	auth-service
	authSubRouter := r.router.PathPrefix("/auth").Subrouter()
	authSubRouter.HandleFunc("/login", authhandlers.LoginHandler(logger, authClient)).Methods(http.MethodPost)
	authSubRouter.HandleFunc("/logout", authhandlers.LogoutHandler(logger, authClient)).Methods(http.MethodPost)
	authSubRouter.HandleFunc("/register", authhandlers.RegistrationHandler(logger, authClient)).Methods(http.MethodPost)

	//	user-service
	userSubRouter := r.router.PathPrefix("/user").Subrouter()
	userSubRouter.Use(goathttp.AuthMiddleware)
	userSubRouter.HandleFunc("/all", userhandlers.GetUsersHandler(logger, userClient)).Methods(http.MethodGet)
	userSubRouter.HandleFunc("/{id}", userhandlers.GetUserHandler(logger, userClient)).Methods(http.MethodGet)
	userSubRouter.HandleFunc("/", userhandlers.AddUserHandler(logger, userClient)).Methods(http.MethodPost)
	userSubRouter.HandleFunc("/", userhandlers.UpdateUserHandler(logger, userClient)).Methods(http.MethodPut)
	userSubRouter.HandleFunc("/{id}", userhandlers.DeleteUserHandler(logger, userClient)).Methods(http.MethodDelete)

	//	client-service
	clientSubRouter := r.router.PathPrefix("/client").Subrouter()
	clientSubRouter.Use(goathttp.AuthMiddleware)
	clientSubRouter.HandleFunc("/all", clienthandlers.GetClientsHandler(logger, clientService)).Methods(http.MethodGet)
	clientSubRouter.HandleFunc("/{id}", clienthandlers.GetClientHandler(logger, clientService)).Methods(http.MethodGet)
	clientSubRouter.HandleFunc("/", clienthandlers.UpdateClientHandler(logger, clientService)).Methods(http.MethodPut)
	clientSubRouter.HandleFunc("/{id}", clienthandlers.DeleteClientHandler(logger, clientService)).Methods(http.MethodDelete)

	//	warehouse-service
	warehouseSubRouter := r.router.PathPrefix("/products").Subrouter()
	warehouseSubRouter.Use(goathttp.AuthMiddleware)
	warehouseSubRouter.HandleFunc("/", warehousehandlers.GetProductsHandler(logger, warehouseClient)).Methods(http.MethodGet)
	warehouseSubRouter.HandleFunc("/{id}", warehousehandlers.GetProductHandler(logger, warehouseClient)).Methods(http.MethodGet)
	warehouseSubRouter.HandleFunc("/", warehousehandlers.AddProductsHandler(logger, warehouseClient)).Methods(http.MethodPost)
	warehouseSubRouter.HandleFunc("/", warehousehandlers.UpdateProductsHandler(logger, warehouseClient)).Methods(http.MethodPut)
	warehouseSubRouter.HandleFunc("/", warehousehandlers.DeleteProductsHandler(logger, warehouseClient)).Methods(http.MethodDelete)
	warehouseSubRouter.HandleFunc("/materials", warehousehandlers.GetMaterialsHandler(logger, warehouseClient)).Methods(http.MethodGet)

	//	cart-service
	cartSubRouter := r.router.PathPrefix("/cart").Subrouter()
	cartSubRouter.Use(goathttp.AuthMiddleware)
	cartSubRouter.HandleFunc("/", carthandlers.GetCartHandler(logger, cartClient)).Methods(http.MethodGet)
	cartSubRouter.HandleFunc("/item", carthandlers.AddCartItemHandler(logger, cartClient)).Methods(http.MethodPost)
	cartSubRouter.HandleFunc("/item", carthandlers.UpdateCartItemHandler(logger, cartClient)).Methods(http.MethodPut)
	cartSubRouter.HandleFunc("/item/{id}", carthandlers.DeleteCartItemHandler(logger, cartClient)).Methods(http.MethodDelete)
	cartSubRouter.HandleFunc("/clear", carthandlers.ClearCartHandler(logger, cartClient)).Methods(http.MethodDelete)
}
