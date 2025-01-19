package reporthandlers

import (
	"api-gateway/cluster/report"
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"time"
)

// GetSellReportHandlers godoc
// @Summary Get sell report
// @Description Retrieves the sell report for a specific user by their ID.
// @Tags Reports
// @Produce json
// @Param userId path int true "User ID"
// @Param date path string true "Date"
// @Success 200 {object} object "Sell report data"
// @Failure 400 {string} string "Invalid request or failed to process the response"
// @Failure 403 {string} string "Forbidden - context creation failed"
// @Router /reports/sell/{userId}/{date} [get]
// @Security LogisticAuth
func GetSellReportHandlers(logger goatlogger.Logger, reportClient *report.Client) goathttp.Handler {
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

		date, err := time.Parse("2006-01-02", mux.Vars(r)["date"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		sellReport, err := reportClient.GetSellReport(ctx, userId, date)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

		if _, err = io.Copy(w, sellReport); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error(err.Error())
			return
		}
	}
}
