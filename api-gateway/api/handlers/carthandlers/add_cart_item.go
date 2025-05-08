package carthandlers

import (
	"api-gateway/cluster/cart"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// AddCartItemHandler добавляет новый элемент в корзину.
// @Summary      Добавить товар в корзину
// @Description  Позволяет добавить новый товар в корзину пользователя.
// @Tags         Корзина
// @Accept       json
// @Produce      json
// @Param        item  body  cart.Item  true  "Данные товара для добавления в корзину"
// @Success      200  {object}  cart.Item  "Добавленный товар"
// @Failure      400  {string}  string  "Ошибка ввода или бизнес-логики"
// @Failure      403  {string}  string  "Доступ запрещен"
// @Router       /cart/item [post]
// @Security LogisticAuth
func AddCartItemHandler(logger goatlogger.Logger, cartClient *cart.Client) goathttp.Handler {
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

		addedItem, err := cartClient.AddCartItem(ctx, item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, addedItem); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
