package carthandlers

import (
	"api-gateway/cluster/cart"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// GetCartHandler получает содержимое корзины пользователя.
// @Summary      Получить корзину пользователя
// @Description  Возвращает содержимое корзины текущего пользователя.
// @Tags         Корзина
// @Produce      json
// @Success      200  {object}  cart.Cart  "Содержимое корзины пользователя"
// @Failure      400  {string}  string  "Ошибка получения корзины"
// @Failure      403  {string}  string  "Доступ запрещен"
// @Router       /cart [get]
// @Security LogisticAuth
func GetCartHandler(logger goatlogger.Logger, cartClient *cart.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		userCart, err := cartClient.GetCart(ctx)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, userCart); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
