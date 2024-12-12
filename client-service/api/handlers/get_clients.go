package handlers

import (
	"client-service/service"
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

func GetClientsHandler(logger goatlogger.Logger, clientService service.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			logger.Error(fmt.Sprintf("не удалось получить контекст запроса: %v", err))
			return
		}

		clients, err := clientService.GetClients(ctx)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось получить клиентов: %v", err))
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, clients); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось записать ответ: %v", err))
			return
		}
	}
}
