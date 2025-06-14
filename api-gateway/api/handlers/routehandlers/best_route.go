package routehandlers

import (
	"api-gateway/cluster/route"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// GetBestRouteHandler получение самого быстрого пути
// @Summary получение самого быстрого пути
// @Description получение самого быстрого пути
// @Tags Route
// @Produce json
// @Success 200 {object} route.Route "Best route"
// @Failure 400 {string} string "Invalid request or failed to process the response"
// @Failure 403 {string} string "Forbidden - context creation failed"
// @Router /routes/route/best [post]
// @Security LogisticAuth
func GetBestRouteHandler(logger goatlogger.Logger, client *route.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		var locations []route.ServiceLocation
		if err = json.ReadRequest(r, &locations); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		bestRoute, err := client.GetBestRoute(ctx, locations)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusCreated, bestRoute); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
