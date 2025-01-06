package handlers

import (
	"errors"
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"request-service/domain"
	"request-service/service"
	"strconv"
)

func UpdateRequestStatusHandler(logger goatlogger.Logger, requestService service.Request) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		id, status, err := getIdAndStatusFromRequest(r)
		if err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = requestService.UpdateStatus(ctx, id, status); err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func getIdAndStatusFromRequest(r *http.Request) (int, domain.Status, error) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, "", err
	}

	status, ok := vars["status"]
	if !ok {
		return 0, "", errors.New("status not found")
	}

	domainStatus, ok := domain.StatusToDomain[status]
	if !ok {
		return 0, "", errors.New("invalid status")
	}

	return id, domainStatus, nil
}
