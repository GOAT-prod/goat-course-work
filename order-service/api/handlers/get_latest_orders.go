package handlers

import (
	"context"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
	"order-service/service"
)

func GetLatestOrdersHandler(logger goatlogger.Logger, orderService service.Order) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		latestOrders, err := orderService.GetLatestOrders(goatcontext.Context{Context: context.Background()})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, latestOrders); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
