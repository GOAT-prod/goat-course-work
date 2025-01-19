package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func parseParams(r *http.Request) (int, time.Time, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return 0, time.Time{}, err
	}

	t, err := time.Parse(time.DateOnly, mux.Vars(r)["date"])
	if err != nil {
		return 0, time.Time{}, err
	}
	
	return id, t, nil
}
