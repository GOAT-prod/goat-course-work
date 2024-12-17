package orderhandlers

import (
	"api-gateway/cluster/order"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// GetUserOrdersHandler получает список заказов пользователя.
// @Summary      Получить заказы пользователя
// @Description  Возвращает список всех заказов, связанных с текущим авторизованным пользователем.
// @Tags         Заказы
// @Accept       json
// @Produce      json
// @Success      200   {array}  order.Order  "Список заказов пользователя"
// @Failure      400   {string}  string  "Ошибка при получении заказов пользователя"
// @Failure      403   {string}  string  "Доступ запрещен"
// @Router       /order/all [get]
func GetUserOrdersHandler(logger goatlogger.Logger, orderClient *order.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		orders, err := orderClient.GetUserOrders(ctx)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, orders); err != nil {
			logger.Error(err.Error())
			return
		}
	}
}
