package api

import (
	"client-service/api/handlers"
	"client-service/service"
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

func (r *Router) SetupRoutes(logger goatlogger.Logger, clientService service.Client) {
	r.router.HandleFunc("/clients", handlers.GetClientsHandler(logger, clientService)).Methods(http.MethodGet)
	r.router.HandleFunc("/client/{id}", handlers.GetClientHandler(logger, clientService)).Methods(http.MethodGet)
	r.router.HandleFunc("/client", handlers.AddClientHandler(logger, clientService)).Methods(http.MethodPost)
	r.router.HandleFunc("/client", handlers.UpdateClientHandler(logger, clientService)).Methods(http.MethodPut)
	r.router.HandleFunc("/client/{id}", handlers.DeleteClientHandler(logger, clientService)).Methods(http.MethodDelete)
	r.router.HandleFunc("/info/short", handlers.GetInfoShort(logger, clientService)).Methods(http.MethodPost)
}
