package handlers

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	"github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
	"warehouse-service/domain"
	"warehouse-service/service"
)

func UpdateProductsHandler(logger goatlogger.Logger, warehouseService service.WareHouse) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var products []domain.Product
		if err = json.ReadRequest(r, &products); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось прочитать обновляемые продукты: %v", err))
			return
		}

		if err = warehouseService.UpdateProducts(ctx, products); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось обновить продукты: %v", err))
			return
		}
	}
}
