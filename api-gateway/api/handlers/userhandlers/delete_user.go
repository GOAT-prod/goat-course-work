package userhandlers

import (
	"api-gateway/cluster/userservice"
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// DeleteUserHandler godoc
// @Summary Delete a user
// @Description This endpoint allows you to delete an existing user by their ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID to be deleted"
// @Success 204 "No Content"
// @Failure 400 {string}  string "Bad Request"
// @Failure 403 {string}  string "Forbidden"
// @Router /user/{id} [delete]
// @Security LogisticAuth
func DeleteUserHandler(logger goatlogger.Logger, userClient *userservice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		userId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = userClient.DeleteUser(ctx, userId); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
