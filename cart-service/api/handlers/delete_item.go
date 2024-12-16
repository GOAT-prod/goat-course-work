package handlers

import (
	"cart-service/service"
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DeleteItemHandler(logger goatlogger.Logger, cartService service.Cart) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(fmt.Sprintf("не удалось получить контекст запроса: %v", err))
			return
		}

		itemId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = cartService.RemoveCartItem(ctx, itemId); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
