package handlers

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
	"route-search-service/domain"
	"route-search-service/service"
)

func GetShortestRouteHandler(logger goatlogger.Logger, routeService service.Route) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusForbidden)
			return
		}

		var locations []domain.ServiceLocation
		if err = json.ReadRequest(r, &locations); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		bestRoute, err := routeService.GetShortestRoute(ctx, locations)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, bestRoute); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
