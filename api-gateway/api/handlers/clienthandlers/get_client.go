package clienthandlers

import (
	"api-gateway/cluster/clientservice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetClientHandler retrieves client details by their ID and requires authorization.
//
// @Summary      Get client
// @Description  Retrieves client details by their ID. Authorization is required to access this endpoint.
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        id  path      int  true  "Client ID"
// @Success      200  {object}  clientservice.ClientInfo  "Client details"
// @Failure      403  {string}  string                 "Forbidden - Invalid token or unauthorized request"
// @Failure      400  {string}  string                 "Bad Request - Invalid client ID"
// @Router       /client/{id} [get]
// @Security LogisticAuth
func GetClientHandler(logger goatlogger.Logger, clientService *clientservice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		clientId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		client, err := clientService.GetClientById(ctx, clientId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, client); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
