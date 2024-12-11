package handlers

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	goatjson "github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
	"warehouse-service/service"
)

func GetMaterialsList(logger goatlogger.Logger, warehouseService service.WareHouse) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		materials, err := warehouseService.GetMaterialsList(ctx)
		if err != nil {
			logger.Error(fmt.Sprintf("не удалось получить список материалов: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = goatjson.WriteResponse(w, http.StatusOK, materials); err != nil {
			logger.Error(fmt.Sprintf("не удалось записать ответ: %v", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
