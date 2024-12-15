package authhandlers

import (
	"api-gateway/cluster/authservice"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// LoginHandler handles user authentication and issues access tokens.
//
// @Summary      User login
// @Description  Authenticates a user and returns access and refresh tokens.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginData  body      authservice.LoginData  true  "User login credentials"
// @Success      200        {object}  authservice.Tokens    "Access and refresh tokens"
// @Failure      403        {string}  string                 "Forbidden - Invalid credentials or request"
// @Router       /auth/login [post]
func LoginHandler(logger goatlogger.Logger, authClient *authservice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		var loginData authservice.LoginData
		if err = json.ReadRequest(r, &loginData); err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		tokens, err := authClient.Login(ctx, loginData)
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
