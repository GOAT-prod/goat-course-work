package userhandlers

import (
	"api-gateway/cluster/userservice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// GetUsersHandler godoc
// @Summary Get list of users
// @Description This endpoint retrieves a list of all users from the system.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} userservice.User "List of users"
// @Failure 400 {string}  string "Bad Request"
// @Failure 403 {string}  string "Forbidden"
// @Router /user/all [get]
// @Security LogisticAuth
func GetUsersHandler(logger goatlogger.Logger, userClient *userservice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		users, err := userClient.GetUsers(ctx)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, users); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
