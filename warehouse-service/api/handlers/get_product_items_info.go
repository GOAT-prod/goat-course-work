package handlers

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
	"warehouse-service/service"
)

func GetProductItemsInfoHandler(logger goatlogger.Logger, warehouseService service.WareHouse) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var ids []int
		if err = json.ReadRequest(r, &ids); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось получить список id: %v", err))
			return
		}

		infos, err := warehouseService.GetProductItemsInfo(ctx, ids)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось получить список информацию о вариантах продукции: %v", err))
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, infos); err != nil {
			logger.Error(fmt.Sprintf("не удалось записать ответ: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
