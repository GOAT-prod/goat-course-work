package carthandlers

import (
	"api-gateway/cluster/cart"
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// ClearCartHandler очищает корзину пользователя.
// @Summary      Очистить корзину
// @Description  Удаляет все товары из корзины текущего пользователя.
// @Tags         Корзина
// @Produce      json
// @Success      200  {string}  string  "Корзина успешно очищена"
// @Failure      400  {string}  string  "Ошибка очистки корзины"
// @Failure      403  {string}  string  "Доступ запрещен"
// @Router       /cart/clear [delete]
func ClearCartHandler(logger goatlogger.Logger, cartClient *cart.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		if err = cartClient.ClearCart(ctx); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
