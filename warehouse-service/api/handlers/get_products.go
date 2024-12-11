package handlers

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	"github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
	"warehouse-service/service"
)

func GetProductsHandler(logger goatlogger.Logger, warehouseService service.WareHouse) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		products, err := warehouseService.GetProducts(ctx)
		if err != nil {
			logger.Error(fmt.Sprintf("не удалось получить список продуктов: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, products); err != nil {
			logger.Error(fmt.Sprintf("не удалось записать ответ: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
