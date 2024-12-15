package userhandlers

import (
	"api-gateway/cluster/userservice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// UpdateUserHandler godoc
// @Summary Update an existing user
// @Description This endpoint allows you to update the details of an existing user in the system.
// @Tags users
// @Accept json
// @Produce json
// @Param user body userservice.User true "User information to be updated"
// @Success 200 {object} userservice.User "Updated user information"
// @Failure 400 {string}  string "Bad Request"
// @Failure 403 {string}  string "Forbidden"
// @Router /user [put]
// @Security LogisticAuth
func UpdateUserHandler(logger goatlogger.Logger, userClient *userservice.Client) goathttp.Handler {
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

		updatedUser, err := userClient.UpdateUser(ctx, user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, updatedUser); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err.Error())
			return
		}
	}
}
