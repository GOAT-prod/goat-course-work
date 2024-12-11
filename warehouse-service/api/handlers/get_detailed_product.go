package handlers

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	"github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"warehouse-service/service"
)

func GetDetailedProductHandler(logger goatlogger.Logger, warehouseService service.WareHouse) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		productId, err := parseProductId(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось прочитать id продукта: %v", err))
			return
		}

		detailedProduct, err := warehouseService.GetDetailedProductsInfo(ctx, productId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось получить продукт: %v", err))
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, detailedProduct); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось записать ответ: %v", err))
			return
		}
	}
}

func parseProductId(r *http.Request) (int, error) {
	productId := mux.Vars(r)["productId"]
	return strconv.Atoi(productId)
}
