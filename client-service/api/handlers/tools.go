package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func parseClientId(r *http.Request) (int, error) {
	return strconv.Atoi(mux.Vars(r)["id"])
}
