package mappings

import (
	"github.com/samber/lo"
	"warehouse-service/database/models"
	"warehouse-service/domain"
)

func ToDomainProductMaterial(material models.ProductMaterial) domain.ProductMaterial {
	return domain.ProductMaterial{
		Id:       material.Id,
		Material: material.Material,
	}
}

func ToDatabaseProductMaterial(materials domain.ProductMaterial) models.ProductMaterial {
	return models.ProductMaterial{
		Id:       materials.Id,
		Material: materials.Material,
	}
}

func ToDomainProductMaterials(materials []models.ProductMaterial) []domain.ProductMaterial {
	return lo.Map(materials, func(item models.ProductMaterial, _ int) domain.ProductMaterial {
		return ToDomainProductMaterial(item)
	})
}

func ToDatabaseProductMaterials(materials []domain.ProductMaterial) []models.ProductMaterial {
	return lo.Map(materials, func(item domain.ProductMaterial, _ int) models.ProductMaterial {
		return ToDatabaseProductMaterial(item)
	})
}