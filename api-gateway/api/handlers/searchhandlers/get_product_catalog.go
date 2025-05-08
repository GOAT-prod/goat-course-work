package searchhandlers

import (
	"api-gateway/cluster/search"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetProductCatalogHandler godoc
// @Summary Get product catalog
// @Description Retrieves the catalog for a specific product by its ID.
// @Tags Catalog
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} search.Product "Product catalog data"
// @Failure 400 {string} string "Invalid request or failed to process the response"
// @Failure 403 {string} string "Forbidden - context creation failed"
// @Router /search/catalog/product/{id} [get]
// @Security LogisticAuth
func GetProductCatalogHandler(logger goatlogger.Logger, searchClient *search.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		id, _ := mux.Vars(r)["id"]

		productId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		productCatalog, err := searchClient.GetProductCatalog(ctx, productId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, productCatalog); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
