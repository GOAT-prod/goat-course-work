package handlers

import (
	"client-service/service"
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

func DeleteClientHandler(logger goatlogger.Logger, clientService service.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			logger.Error(fmt.Sprintf("не удалось получить контекст запроса: %v", err))
			return
		}

		clientId, err := parseClientId(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("пришел не валидный clientId: %v", err))
			return
		}

		if err = clientService.DeleteClient(ctx, clientId); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось удалить клиента: %v", err))
		}
	}
}
