package warehousehandlers

import (
	"api-gateway/cluster/warehousesevice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// GetMaterialsHandler godoc
// @Summary Get list of materials
// @Description This endpoint retrieves a list of all materials from the warehouse.
// @Tags materials
// @Accept json
// @Produce json
// @Success 200 {array} warehousesevice.ProductMaterial "List of materials"
// @Failure 400 {string}  string "Bad Request"
// @Failure 403 {string}  string "Forbidden"
// @Router /products/materials [get]
// @Security LogisticAuth
func GetMaterialsHandler(logger goatlogger.Logger, warehouseClient *warehousesevice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		materials, err := warehouseClient.GetAllMaterials(ctx)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, materials); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
