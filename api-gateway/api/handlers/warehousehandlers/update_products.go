package warehousehandlers

import (
	"api-gateway/cluster/warehousesevice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// UpdateProductsHandler godoc
// @Summary Update a list of products
// @Description This endpoint allows you to update a list of products in the warehouse.
// @Tags products
// @Accept json
// @Produce json
// @Param products body []warehousesevice.Product true "List of products to update"
// @Success 200 "Products successfully updated"
// @Failure 400 {string}  string "Bad Request"
// @Failure 403 {string}  string "Forbidden"
// @Router /products [put]
// @Security LogisticAuth
func UpdateProductsHandler(logger goatlogger.Logger, warehouseClient *warehousesevice.Client) goathttp.Handler {
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

		if err = warehouseClient.UpdateProducts(ctx, products); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
