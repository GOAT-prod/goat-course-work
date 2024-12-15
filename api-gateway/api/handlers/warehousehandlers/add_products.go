package warehousehandlers

import (
	"api-gateway/cluster/warehousesevice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// AddProductsHandler godoc
// @Summary Add products to the warehouse
// @Description This endpoint allows you to add a list of products to the warehouse.
// @Tags products
// @Accept json
// @Produce json
// @Param products body []warehousesevice.Product true "List of products to be added"
// @Success 200 "Products successfully added"
// @Failure 400 {string}  string "Bad Request"
// @Failure 403 {string}  string "Forbidden"
// @Router /products [post]
// @Security LogisticAuth
func AddProductsHandler(logger goatlogger.Logger, warehouseClient *warehousesevice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		var products []warehousesevice.Product
		if err = json.ReadRequest(r, &products); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = warehouseClient.AddProducts(ctx, products); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
