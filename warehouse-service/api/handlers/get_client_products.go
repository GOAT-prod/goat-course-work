package handlers

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"warehouse-service/service"
)

func GetClientProductsHandler(logger goatlogger.Logger, warehouseService service.WareHouse) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		clientId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось прочитать id клиента: %v", err))
			return
		}

		products, err := warehouseService.GetClientProduct(ctx, clientId)
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
