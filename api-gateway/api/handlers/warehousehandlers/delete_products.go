package warehousehandlers

import (
	"api-gateway/cluster/warehousesevice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// DeleteProductsHandler godoc
// @Summary Delete products from the warehouse
// @Description This endpoint allows you to delete a list of products from the warehouse by their IDs.
// @Tags products
// @Accept json
// @Produce json
// @Param productIds body []int true "List of product IDs to be deleted"
// @Success 200 "Products successfully deleted"
// @Failure 400 {string}  string "Bad Request"
// @Failure 403 {string}  string "Forbidden"
// @Router /products [delete]
// @Security LogisticAuth
func DeleteProductsHandler(logger goatlogger.Logger, warehouseClient *warehousesevice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		var productIds []int
		if err = json.ReadRequest(r, &productIds); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = warehouseClient.DeleteProducts(ctx, productIds); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
