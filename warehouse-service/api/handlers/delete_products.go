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

func DeleteProductsHandler(logger goatlogger.Logger, warehouseService service.WareHouse) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var productIds []int
		if err = json.ReadRequest(r, &productIds); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось прочитать ids удалемых продуктов: %v", err))
			return
		}

		if err = warehouseService.DeleteProducts(ctx, productIds); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось удалить продукты: %v", err))
			return
		}
	}
}
