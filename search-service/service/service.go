package service

import (
	"context"
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/samber/lo"
	"log"
	"search-service/cluster/warehouse"
	"search-service/database"
	"search-service/domain"
	"search-service/repository"
	"sort"
)

type Search interface {
	GetFilters(ctx goatcontext.Context) ([]domain.Filter, error)
	GetCatalog(ctx goatcontext.Context, searchId string, appliedFilters domain.AppliedFilters) (domain.Catalog, error)
	GetProductCatalog(ctx goatcontext.Context, productId int) (domain.Product, error)
}

type SearchService struct {
	filterRepository repository.Filter
	cacheRepository  repository.Cache
	warehouseClient  *warehouse.Client
}

func New(filterRepository repository.Filter, cacheRepository repository.Cache, warehouseClient *warehouse.Client) Search {
	return &SearchService{
		filterRepository: filterRepository,
		cacheRepository:  cacheRepository,
		warehouseClient:  warehouseClient,
	}
}

func (s *SearchService) GetFilters(ctx goatcontext.Context) ([]domain.Filter, error) {
	dbFilters, err := s.filterRepository.GetFilters(ctx)
	if err != nil {
		return nil, err
	}

	return lo.Map(dbFilters, func(item database.Filter, _ int) domain.Filter {
		return domain.Filter{
			Name:          item.Name,
			AllowedValues: item.AllowedValues,
		}
	}), nil
}

func (s *SearchService) GetCatalog(ctx goatcontext.Context, searchId string, appliedFilters domain.AppliedFilters) (domain.Catalog, error) {
	catalog := domain.Catalog{
		Filters:  appliedFilters,
		SearchId: searchId,
	}

	if s.cacheRepository.Check(ctx, searchId) {
		products, err := s.cacheRepository.Get(ctx, searchId)
		if err != nil {
			return domain.Catalog{}, err
		}

		catalog.Products = products

		return catalog, nil
	}

	allProducts, err := s.warehouseClient.GetProducts(ctx)
	if err != nil {
		return domain.Catalog{}, err
	}

	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].Id < allProducts[j].Id
	})

	catalog.Products = applyFilters(allProducts, appliedFilters)

	go func(sid string, p []domain.Product) {
		redisCtx := ctx
		redisCtx.Context = context.Background()
		if redisErr := s.cacheRepository.Set(redisCtx, sid, p); redisErr != nil {
			log.Printf(fmt.Sprintf("[sid::%s] не удалось сохранить выдачу redis : %v", sid, redisErr))
		}
	}(searchId, catalog.Products)

	return catalog, nil
}

func applyFilters(products []domain.Product, appliedFilters domain.AppliedFilters) []domain.Product {
	filteredProducts := make([]domain.Product, 0, len(products))

	for _, product := range products {
		if checkProduct(appliedFilters, product) {
			filteredProductItems := make([]domain.ProductItem, 0, len(product.Items))

			for _, item := range product.Items {
				if checkProductItem(appliedFilters, item) {
					filteredProductItems = append(filteredProductItems, item)
				}
			}

			if len(filteredProductItems) == 0 {
				continue
			}

			filteredProduct := product
			filteredProduct.Items = filteredProductItems
			filteredProducts = append(filteredProducts, filteredProduct)
		}
	}

	return filteredProducts
}

func checkProduct(appliedFilters domain.AppliedFilters, product domain.Product) bool {
	switch true {
	case product.Status != domain.Approved:
		return false
	case len(appliedFilters.Brand) != 0 && !lo.Contains(appliedFilters.Brand, product.BrandName):
		return false
	case appliedFilters.MinPrice.GreaterThan(product.Price) || appliedFilters.MaxPrice.LessThan(product.Price):
		return false
	case len(appliedFilters.Material) != 0 && !lo.Some(product.GetMaterialNames(), appliedFilters.Material):
		return false
	}

	return true
}

func checkProductItem(appliedFilters domain.AppliedFilters, productItem domain.ProductItem) bool {
	switch true {
	case productItem.StockCount == 0:
		return false
	case len(appliedFilters.Size) != 0 && !lo.Contains(appliedFilters.Size, productItem.Size):
		return false
	case len(appliedFilters.Color) != 0 && lo.Contains(appliedFilters.Color, productItem.Color):
		return false
	}

	return true
}

func (s *SearchService) GetProductCatalog(ctx goatcontext.Context, productId int) (domain.Product, error) {
	return s.warehouseClient.GetProduct(ctx, productId)
}
