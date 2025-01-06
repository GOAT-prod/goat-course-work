package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"warehouse-service/api/handlers"
	"warehouse-service/service"

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

func (r *Router) SetupRoutes(logger goatlogger.Logger, warehouse service.WareHouse) {
	r.router.HandleFunc("/products", handlers.GetProductsHandler(logger, warehouse)).Methods(http.MethodGet)
	r.router.HandleFunc("/product/{productId}", handlers.GetDetailedProductHandler(logger, warehouse)).Methods(http.MethodGet)
	r.router.HandleFunc("/materials/list", handlers.GetMaterialsList(logger, warehouse)).Methods(http.MethodGet)
	r.router.HandleFunc("/products", handlers.AddProductsHandler(logger, warehouse)).Methods(http.MethodPost)
	r.router.HandleFunc("/products", handlers.UpdateProductsHandler(logger, warehouse)).Methods(http.MethodPut)
	r.router.HandleFunc("/products", handlers.DeleteProductsHandler(logger, warehouse)).Methods(http.MethodDelete)
	r.router.HandleFunc("/products/items", handlers.GetProductItemsInfoHandler(logger, warehouse)).Methods(http.MethodPost)
	r.router.HandleFunc("/products/detailed", handlers.GetDetailedProductsHandler(logger, warehouse)).Methods(http.MethodPost)
}
