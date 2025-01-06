package handlers

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
	"user-service/domain"
	"user-service/service"
)

func UpdateUserPasswordHandler(logger goatlogger.Logger, userService service.User) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(fmt.Sprintf("не удалось получить контекст запроса: %v", err))
			return
		}

		var updatePasswordRequest domain.UpdatePasswordRequest
		if err = json.ReadRequest(r, &updatePasswordRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
		
		if err = userService.UpdateUserPassword(ctx, updatePasswordRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
