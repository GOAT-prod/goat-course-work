package clienthandlers

import (
	"api-gateway/cluster/clientservice"
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// DeleteClientHandler handles the deletion of a client and requires authorization.
//
// @Summary      Delete client
// @Description  Deletes a client by their ID. Authorization is required to access this endpoint.
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        id  path      int  true  "Client ID"
// @Success      200  {string}  string  "Successfully deleted client"
// @Failure      403  {string}  string  "Forbidden - Invalid token or unauthorized request"
// @Failure      400  {string}  string  "Bad Request - Invalid client ID"
// @Router       /client/{id} [delete]
// @Security LogisticAuth
func DeleteClientHandler(logger goatlogger.Logger, clientService *clientservice.Client) goathttp.Handler {
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

		if err = clientService.DeleteClient(ctx, clientId); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
