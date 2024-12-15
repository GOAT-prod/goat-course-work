package userhandlers

import (
	"api-gateway/cluster/userservice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// AddUserHandler godoc
// @Summary Add a new user
// @Description This endpoint allows you to add a new user to the system.
// @Tags users
// @Accept json
// @Produce json
// @Param user body userservice.User true "User information to be added"
// @Success 200 {object} userservice.User "User successfully added"
// @Failure 400 {string}  string "Bad Request"
// @Failure 403 {string}  string "Forbidden"
// @Router /user [post]
// @Security LogisticAuth
func AddUserHandler(logger goatlogger.Logger, userClient *userservice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		var user userservice.User
		if err = json.ReadRequest(r, &user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		addedUser, err := userClient.AddUser(ctx, user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, addedUser); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
