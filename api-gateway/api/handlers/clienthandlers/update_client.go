package clienthandlers

import (
	"api-gateway/cluster/clientservice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// UpdateClientHandler godoc
// @Summary Update an existing client
// @Description This endpoint updates an existing client's information in the system.
// @Tags clients
// @Accept json
// @Produce json
// @Param client body clientservice.ClientInfo true "Client information to be updated"
// @Success 200 {object} clientservice.ClientInfo "Updated client information"
// @Failure 400 {string}  string "Bad Request"
// @Failure 403 {string}  string "Forbidden"
// @Router /client [put]
// @Security LogisticAuth
func UpdateClientHandler(logger goatlogger.Logger, clientService *clientservice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		var client clientservice.ClientInfo
		if err = json.ReadRequest(r, &client); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = clientService.UpdateClient(ctx, client); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
