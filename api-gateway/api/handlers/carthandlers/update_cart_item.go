package carthandlers

import (
	"api-gateway/cluster/cart"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// UpdateCartItemHandler обновляет элемент корзины.
// @Summary      Обновить элемент корзины
// @Description  Обновляет количество, размер или другие параметры элемента в корзине.
// @Tags         Корзина
// @Accept       json
// @Produce      json
// @Param        item  body  cart.Item  true  "Данные элемента корзины для обновления"
// @Success      200   {string}  string  "Элемент корзины успешно обновлен"
// @Failure      400   {string}  string  "Ошибка обновления элемента корзины"
// @Failure      403   {string}  string  "Доступ запрещен"
// @Router       /cart/item [put]
func UpdateCartItemHandler(logger goatlogger.Logger, cartClient *cart.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		var item cart.Item
		if err = json.ReadRequest(r, &item); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = cartClient.UpdateCartItem(ctx, item); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
