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

func GetClientHandler(logger goatlogger.Logger, clientService service.Client) goathttp.Handler {
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

		client, err := clientService.GetClient(ctx, clientId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось получить клиента: %v", err))
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, client); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось записать ответ: %v", err))
			return
		}
	}
}
