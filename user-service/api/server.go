package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"user-service/api/handlers"
	"user-service/service"

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

func (r *Router) SetupRoutes(logger goatlogger.Logger, userService service.User) {
	r.router.HandleFunc("/users", handlers.GetUsersHandler(logger, userService)).Methods(http.MethodGet)
	r.router.HandleFunc("/user/{id}", handlers.GetUserHandler(logger, userService)).Methods(http.MethodGet)
	r.router.HandleFunc("/user/check", handlers.CheckUserExistHandler(logger, userService)).Methods(http.MethodGet)
	r.router.HandleFunc("/user", handlers.AddUserHandler(logger, userService)).Methods(http.MethodPost)
	r.router.HandleFunc("/user", handlers.UpdateUserHandler(logger, userService)).Methods(http.MethodPut)
	r.router.HandleFunc("/user/{id}", handlers.DeleteUserHandler(logger, userService)).Methods(http.MethodDelete)
	r.router.HandleFunc("/user/registration", handlers.RegisterHandler(logger, userService)).Methods(http.MethodPost)
}
