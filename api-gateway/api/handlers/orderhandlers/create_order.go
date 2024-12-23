package orderhandlers

import (
	"api-gateway/cluster/order"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// CreateOrderHandler создает заказ на основе идентификаторов элементов корзины.
// @Summary      Создать заказ
// @Description  Создает новый заказ, используя переданные идентификаторы элементов корзины.
// @Tags         Заказы
// @Accept       json
// @Produce      json
// @Param        cartItemIds  body  []int  true  "Список идентификаторов элементов корзины"
// @Success      200   {string}  string  "Заказ успешно создан"
// @Failure      400   {string}  string  "Ошибка при создании заказа"
// @Failure      403   {string}  string  "Доступ запрещен"
// @Router       /order [post]
// @Security LogisticAuth
func CreateOrderHandler(logger goatlogger.Logger, orderClient *order.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		var cartItemIds []int
		if err = json.ReadRequest(r, &cartItemIds); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = orderClient.CreateOrder(ctx, cartItemIds); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
