package mappings

import (
	"github.com/samber/lo"
	"warehouse-service/database/models"
	"warehouse-service/domain"
)

func ToDomainProductItem(productItem models.ProductItem) domain.ProductItem {
	return domain.ProductItem{
		Id:         productItem.Id,
		StockCount: productItem.StockCount,
		Size:       productItem.Size,
		Weight:     productItem.Weight,
		Color:      productItem.Color,
	}
}

func ToDatabaseProductItem(productItem domain.ProductItem) models.ProductItem {
	return models.ProductItem{
		Id:         productItem.Id,
		StockCount: productItem.StockCount,
		Size:       productItem.Size,
		Weight:     productItem.Weight,
		Color:      productItem.Color,
	}
}

func ToDomainProductItems(items []models.ProductItem) []domain.ProductItem {
	return lo.Map(items, func(item models.ProductItem, _ int) domain.ProductItem {
		return ToDomainProductItem(item)
	})
}

func ToDatabaseProductItems(items []domain.ProductItem) []models.ProductItem {
	return lo.Map(items, func(item domain.ProductItem, _ int) models.ProductItem {
		return ToDatabaseProductItem(item)
	})
}
