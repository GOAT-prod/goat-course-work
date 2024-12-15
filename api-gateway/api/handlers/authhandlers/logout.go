package authhandlers

import (
	"api-gateway/cluster/authservice"
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"net/http"
)

// LogoutHandler handles user logout by invalidating the user's refresh token.
//
// @Summary      User logout
// @Description  Invalidates the user's session by revoking the refresh token.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200        {string}  string                 "Successfully logged out"
// @Failure      403        {string}  string                 "Forbidden - Invalid token or request"
// @Router       /auth/logout [post]
func LogoutHandler(logger goatlogger.Logger, authClient *authservice.Client) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		refreshToken, err := r.Cookie("refresh_token")
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}

		if err = authClient.Logout(ctx, refreshToken.Value); err != nil {
			w.WriteHeader(http.StatusForbidden)
			logger.Error(err.Error())
			return
		}
	}
}
