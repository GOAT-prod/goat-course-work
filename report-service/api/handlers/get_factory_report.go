package handlers

import (
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
	"report-service/service"
)

func GetFactoryReportHandler(logger goatlogger.Logger, reportService service.Report) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		id, date, err := parseParams(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		report, err := reportService.GetFactoryReport(ctx, id, date)

		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

		if err = report.Write(w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error(err.Error())
			return
		}
	}
}
