package authhandlers

import (
	"api-gateway/cluster/authservice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// RegistrationHandler handles user registration and issues access tokens.
//
// @Summary      User registration
// @Description  Registers a new user and returns access and refresh tokens.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        registerData  body      authservice.RegisterData  true  "User registration data"
// @Success      200           {object}  authservice.Tokens        "Access and refresh tokens"
// @Failure      403           {string}  string                    "Forbidden - Invalid registration data or request"
// @Router       /auth/register [post]
func RegistrationHandler(logger goatlogger.Logger, authClient *authservice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		var registerData authservice.RegisterData
		if err = json.ReadRequest(r, &registerData); err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		tokens, err := authClient.Register(ctx, registerData)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, tokens); err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}
	}
}
