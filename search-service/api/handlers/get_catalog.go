package handlers

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/json"
	goathttp "github.com/GOAT-prod/goathttp/server"
	"github.com/GOAT-prod/goatlogger"
	"github.com/google/uuid"
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
			logger.Error(err.Error())
			w.WriteHeader(http.StatusForbidden)
			return
		}

		filters, err := parseFilters(r)
		if err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		catalog, err := searchService.GetCatalog(ctx, uuid.NewString(), filters)
		if err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = json.WriteResponse(w, http.StatusOK, catalog); err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func parseFilters(r *http.Request) (domain.AppliedFilters, error) {
	queryParams := r.URL.Query()

	sizes := queryParams["size"]
	intSizes := make([]int, 0, len(sizes))

	for _, s := range sizes {
		i, err := strconv.Atoi(s)
		if err != nil {
			return domain.AppliedFilters{}, err
		}

		intSizes = append(intSizes, i)
	}

	minPrice, _ := decimal.NewFromString(queryParams.Get("minPrice"))
	maxPrice, _ := decimal.NewFromString(queryParams.Get("maxPrice"))
	page, _ := strconv.Atoi(queryParams.Get("page"))
	limit, _ := strconv.Atoi(queryParams.Get("limit"))

	if maxPrice.Equal(decimal.Zero) {
		maxPrice = decimal.NewFromFloat32(1000000)
	}

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	return domain.AppliedFilters{
		Page:     page,
		Limit:    limit,
		Size:     intSizes,
		Color:    queryParams["color"],
		Brand:    queryParams["brand"],
		Material: queryParams["material"],
		MinPrice: minPrice,
		MaxPrice: maxPrice,
	}, nil
}
