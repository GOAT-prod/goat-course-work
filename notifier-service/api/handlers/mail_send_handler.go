package handlers

import (
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
	"notifier-service/domain"
	"notifier-service/service"
)

func MainSendHandler(logger goatlogger.Logger, sender service.Sender) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		var mailMessage domain.Mail
		if err := json.ReadRequest(r, &mailMessage); err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := sender.Send(mailMessage); err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
