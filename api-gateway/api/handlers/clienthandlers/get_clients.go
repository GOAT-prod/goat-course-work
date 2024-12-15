package clienthandlers

import (
	"api-gateway/cluster/clientservice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// GetClientsHandler godoc
// @Summary Get list of clients
// @Description This endpoint retrieves a list of clients from the client service.
// @Tags clients
// @Accept json
// @Produce json
// @Success 200 {array} clientservice.ClientInfo "OK"
// @Failure 400 {string}  string "Bad Request"
// @Failure 403 {string}  string "Forbidden"
// @Router /client/all [get]
// @Security LogisticAuth
func GetClientsHandler(logger goatlogger.Logger, clientService *clientservice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		clients, err := clientService.GetClients(ctx)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, clients); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
