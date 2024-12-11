package mappings

import (
	"warehouse-service/cluster/clientservice"
	"warehouse-service/database/models"
	"warehouse-service/domain"
)

func ToDomainProduct(product models.Product, factory clientservice.ClientInfoShort) domain.Product {
	return domain.Product{
		Id:        product.Id,
		BrandName: product.BrandName,
		Factory: domain.Factory{
			Id:          factory.Id,
			FactoryName: factory.Name,
		},
		Name:      product.Name,
		Price:     product.Price,
		Items:     ToDomainProductItems(product.Items),
		Materials: ToDomainProductMaterials(product.Materials),
		Images:    ToDomainProductImages(product.Images),
		Status:    domain.ProductStatus(product.Status),
	}
}

func ToDatabaseProduct(product domain.Product) models.Product {
	return models.Product{
		Id:        product.Id,
		BrandName: product.BrandName,
		Name:      product.Name,
		Price:     product.Price,
		FactoryId: product.Factory.Id,
		Items:     ToDatabaseProductItems(product.Items),
		Materials: ToDatabaseProductMaterials(product.Materials),
		Images:    ToDatabaseProductImages(product.Images),
		Status:    string(product.Status),
	}
}
