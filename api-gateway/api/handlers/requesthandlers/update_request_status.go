package requesthandlers

import (
	"api-gateway/cluster/request"
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// UpdateRequestStatusHandler godoc
// @Summary Update request status
// @Description Updates the status of a specific request by its ID.
// @Tags Requests
// @Param requestId path int true "Request ID"
// @Param status path string true "New status of the request"
// @Success 204 "Status updated successfully"
// @Failure 400 {string} string "Invalid request or failed to process the response"
// @Failure 403 {string} string "Forbidden - context creation failed"
// @Router /requests/{requestId}/status/{status} [put]
// @Security LogisticAuth
func UpdateRequestStatusHandler(logger goatlogger.Logger, requestClient *request.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		requestId, err := strconv.Atoi(mux.Vars(r)["requestId"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		status := mux.Vars(r)["status"]

		if err = requestClient.UpdateRequestStatus(ctx, requestId, status); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
