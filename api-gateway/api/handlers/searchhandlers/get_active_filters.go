package searchhandlers

import (
	"api-gateway/cluster/search"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// GetActiveFiltersHandler godoc
// @Summary Get active filters
// @Description Retrieves a list of active filters from the search service.
// @Tags Filters
// @Produce json
// @Success 200 {object} []search.Filter "List of active filters"
// @Failure 400 {string} string "Invalid request or failed to process the response"
// @Failure 403 {string} string "Forbidden - context creation failed"
// @Router /search/filters [get]
// @Security LogisticAuth
func GetActiveFiltersHandler(logger goatlogger.Logger, searchClient *search.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		filters, err := searchClient.GetActiveFilters(ctx)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, filters); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
