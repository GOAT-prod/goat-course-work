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

func GetDetailedProductsHandler(logger goatlogger.Logger, warehouseService service.WareHouse) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var ids []int
		if err = json.ReadRequest(r, &ids); err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		detailedProducts, err := warehouseService.GetDetailedProductsInfos(ctx, ids)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось получить продукт: %v", err))
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, detailedProducts); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось записать ответ: %v", err))
			return
		}
	}
}
