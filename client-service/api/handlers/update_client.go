package handlers

import (
	"client-service/domain"
	"client-service/service"
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

func UpdateClientHandler(logger goatlogger.Logger, clientService service.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			logger.Error(fmt.Sprintf("не удалось получить контекст запроса: %v", err))
			return
		}

		var client domain.Client
		if err = json.ReadRequest(r, &client); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось прочитать тело запроса: %v", err))
			return
		}

		if err = clientService.UpdateClient(ctx, client); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось обновить клиента: %v", err))
			return
		}
	}
}
