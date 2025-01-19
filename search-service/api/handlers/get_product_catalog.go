package handlers

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"search-service/service"
	"strconv"
)

func GetProductCatalog(logger goatlogger.Logger, searchService service.Search) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusForbidden)
			return
		}

		productId, err := strconv.Atoi(mux.Vars(r)["productId"])
		if err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		product, err := searchService.GetProductCatalog(ctx, productId)
		if err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, product); err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
