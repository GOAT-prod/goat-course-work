package searchhandlers

import (
	"api-gateway/cluster/search"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// GetCatalogHandler godoc
// @Summary Get catalog
// @Description Retrieves the catalog based on the provided query parameters.
// @Tags Catalog
// @Produce json
// @Param filter query string false "Filter query parameter"
// @Param sort query string false "Sort query parameter"
// @Success 200 {object} object "Catalog data"
// @Failure 400 {string} string "Invalid request or failed to process the response"
// @Failure 403 {string} string "Forbidden - context creation failed"
// @Router /catalog [get]
// @Security LogisticAuth
func GetCatalogHandler(logger goatlogger.Logger, searchClient *search.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		catalog, err := searchClient.GetCatalog(ctx, r.URL.Query())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, catalog); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
