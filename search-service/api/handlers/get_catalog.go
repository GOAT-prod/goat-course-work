package handlers

import (
	"github.com/GOAT-prod/goatcontext"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/shopspring/decimal"
	"net/http"
	"search-service/domain"
	"search-service/service"
	"strconv"
)

func GetCatalogHandler(logger goatlogger.Logger, searchService service.Search) goathttp.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := goatcontext.New(r)
		if err != nil || !ctx.IsAuthorized() {
			w.WriteHeader(http.StatusForbidden)
			return
		}
	}
}

func parseQueryParams(r *http.Request) (domain.AppliedFilters, error) {
	queryParams := r.URL.Query()

	sizes := queryParams["size"]
	intSizes := make([]int, 0, len(sizes))

	for _, s := range sizes {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		intSizes = append(intSizes, i)
	}

	return domain.AppliedFilters{
		Size:     intSizes,
		Color:    queryParams["color"],
		Brand:    queryParams["brand"],
		Material: queryParams["material"],
		MinPrice: decimal.Decimal{},
		MaxPrice: decimal.Decimal{},
	}, nil
}
