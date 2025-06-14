package service

import (
	"github.com/GOAT-prod/goatcontext"
	"route-search-service/domain"
)

type Route interface {
	GetShortestRoute(ctx goatcontext.Context, serviceLocations []domain.ServiceLocation) (domain.Route, error)
}
