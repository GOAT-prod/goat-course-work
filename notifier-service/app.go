package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
	"notifier-service/api"
	"notifier-service/client"
	"notifier-service/service"
	"notifier-service/settings"
)

type App struct {
	mainCtx context.Context
	logger  goatlogger.Logger
	cfg     settings.Settings

	server *api.Server

	smtpClient *client.Smtp
	sender     service.Sender
}

func NewApp(mainCtx context.Context, cfg settings.Settings, logger goatlogger.Logger) *App {
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

func (a *App) Stop(_ context.Context) {}

func (a *App) initClient() {
	a.smtpClient = client.NewSmtp(a.cfg.SmtpCredentials.Host, a.cfg.SmtpCredentials.From, a.cfg.SmtpCredentials.Password, a.cfg.SmtpCredentials.Port)
}

func (a *App) initService() {
	a.sender = service.NewSender(a.cfg.SmtpCredentials.To, a.cfg.SmtpCredentials.From, a.smtpClient)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic(fmt.Sprintf("сервер уже запущен"))
	}

	router := api.NewRouter(a.logger, a.cfg.Port)
	router.SetupRoutes(a.logger, a.sender)

	a.server = api.NewServer(a.mainCtx, router)
}
