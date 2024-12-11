package mappings

import (
	"github.com/samber/lo"
	"warehouse-service/database/models"
	"warehouse-service/domain"
)

func ToDomainProductImage(image models.ProductImages) domain.ProductImages {
	return domain.ProductImages{
		Id:       image.Id,
		ImageUrl: image.ImageUrl,
	}
}

func ToDatabaseProductImage(image domain.ProductImages) models.ProductImages {
	return models.ProductImages{
		Id:       image.Id,
		ImageUrl: image.ImageUrl,
	}
}

func ToDomainProductImages(images []models.ProductImages) []domain.ProductImages {
	return lo.Map(images, func(item models.ProductImages, _ int) domain.ProductImages {
		return ToDomainProductImage(item)
	})
}

func ToDatabaseProductImages(images []domain.ProductImages) []models.ProductImages {
	return lo.Map(images, func(item domain.ProductImages, _ int) models.ProductImages {
		return ToDatabaseProductImage(item)
	})
}
