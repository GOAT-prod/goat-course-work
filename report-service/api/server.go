package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"report-service/api/handlers"
	"report-service/service"

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

func (r *Router) SetupRoutes(logger goatlogger.Logger, reportService service.Report) {
	r.router.HandleFunc("/report/user/{id}/{date}", handlers.GetUserReportHandler(logger, reportService)).Methods(http.MethodGet)
	r.router.HandleFunc("/report/factory/{id}/{date}", handlers.GetFactoryReportHandler(logger, reportService)).Methods(http.MethodGet)
}
