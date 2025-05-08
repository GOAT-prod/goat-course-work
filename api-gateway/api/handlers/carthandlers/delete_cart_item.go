package carthandlers

import (
	"api-gateway/cluster/cart"
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// DeleteCartItemHandler удаляет элемент из корзины.
// @Summary      Удалить элемент из корзины
// @Description  Удаляет элемент корзины по его идентификатору.
// @Tags         Корзина
// @Param        id   path      int  true  "ID элемента корзины"
// @Produce      json
// @Success      200  {string}  string  "Элемент успешно удален"
// @Failure      400  {string}  string  "Ошибка удаления элемента корзины"
// @Failure      403  {string}  string  "Доступ запрещен"
// @Router       /cart/item/{id} [delete]
// @Security LogisticAuth
func DeleteCartItemHandler(logger goatlogger.Logger, cartClient *cart.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		cartItemId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = cartClient.DeleteCartItem(ctx, cartItemId); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
