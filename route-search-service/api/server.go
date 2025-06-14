package api

import (
	"context"
	"fmt"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"route-search-service/api/handlers"
	"route-search-service/service"
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

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (r *Router) SetupRoutes(logger goatlogger.Logger, routeService service.Route) {
	r.router.HandleFunc("/route/best", handlers.GetShortestRouteHandler(logger, routeService)).Methods(http.MethodPost)
}
