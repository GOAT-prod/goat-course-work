package reporthandlers

import (
	"api-gateway/cluster/report"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetOrderReportHandler godoc
// @Summary Get order report
// @Description Retrieves the order report for a specific user by their ID.
// @Tags Reports
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} object "Order report data"
// @Failure 400 {string} string "Invalid request or failed to process the response"
// @Failure 403 {string} string "Forbidden - context creation failed"
// @Router /reports/order/{userId} [get]
// @Security LogisticAuth
func GetOrderReportHandler(logger goatlogger.Logger, reportClient *report.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		userId, err := strconv.Atoi(mux.Vars(r)["userId"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		orderReport, err := reportClient.GetOrderReport(ctx, userId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, orderReport); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
