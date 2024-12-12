package handlers

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"user-service/service"
)

func DeleteUserHandler(logger goatlogger.Logger, userService service.User) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusUnauthorized)
			logger.Error(fmt.Sprintf("не удалось получить контекст запроса: %v", err))
			return
		}

		userId, err := parseUserId(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("пришел не валидный userId: %v", err))
			return
		}

		if err = userService.DeleteUser(ctx, userId); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(fmt.Sprintf("не удалось удалить пользователя: %v", err))
		}
	}
}

func parseUserId(r *http.Request) (userId int, err error) {
	return strconv.Atoi(mux.Vars(r)["id"])
}
