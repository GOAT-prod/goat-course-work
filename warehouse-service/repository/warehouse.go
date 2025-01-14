package repository

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/samber/lo"
	"warehouse-service/database/models"
	"warehouse-service/database/queries"
)

type Warehouse interface {
	GetProducts(ctx goatcontext.Context) ([]models.Product, error)
	GetProductById(ctx goatcontext.Context, id int) (models.Product, error)
	GetProductsByIds(ctx goatcontext.Context, id []int) ([]models.Product, error)
	AddProducts(ctx goatcontext.Context, products []models.Product) ([]models.Product, error)
	UpdateProducts(ctx goatcontext.Context, products []models.Product) error
	DeleteProducts(ctx goatcontext.Context, productIds []int) error
	GetProductItems(ctx goatcontext.Context, productId int) (items []models.ProductItem, err error)
	AddProductItems(ctx goatcontext.Context, productItems []models.ProductItem) error
	UpdateProductItems(ctx goatcontext.Context, productItems []models.ProductItem) error
	DeleteProductItems(ctx goatcontext.Context, id []int) error
	GetAllMaterials(ctx goatcontext.Context) (materials []models.ProductMaterial, err error)
	GetProductsMaterials(ctx goatcontext.Context, productId int) (materials []models.ProductMaterial, err error)
	AddProductMaterials(ctx goatcontext.Context, productMaterials []models.ProductMaterial) error
	UpdateProductMaterials(ctx goatcontext.Context, productMaterials []models.ProductMaterial) error
	DeleteProductMaterials(ctx goatcontext.Context, id []int) error
	GetImages(ctx goatcontext.Context, productId int) (images []models.ProductImages, err error)
	AddImages(ctx goatcontext.Context, productImages []models.ProductImages) error
	UpdateImages(ctx goatcontext.Context, productImages []models.ProductImages) error
	DeleteImages(ctx goatcontext.Context, id []int) error
	GetProductItemsInfo(ctx goatcontext.Context, ids []int) (items []models.ProductItemInfo, err error)
}

type Init struct {
	postgres *sqlx.DB
}

func NewWarehouseRepository(postgres *sqlx.DB) Warehouse {
	return &Init{
		postgres: postgres,
	}
}

func (r *Init) GetProducts(ctx goatcontext.Context) ([]models.Product, error) {
	var products []models.Product
	if err := r.postgres.SelectContext(ctx, &products, queries.GetProducts); err != nil {
		return nil, err
	}

	for _, product := range products {
		items, err := r.GetProductItems(ctx, product.Id)
		if err != nil {
			return nil, err
		}

		product.Items = items

		materials, err := r.GetProductsMaterials(ctx, product.Id)
		if err != nil {
			return nil, err
		}

		product.Materials = materials

		images, err := r.GetImages(ctx, product.Id)
		if err != nil {
			return nil, err
		}

		product.Images = images
	}

	return products, nil
}

func (r *Init) GetProductById(ctx goatcontext.Context, id int) (models.Product, error) {
	var product models.Product
	if err := r.postgres.GetContext(ctx, &product, queries.GetProductsById, pq.Array([]int{id})); err != nil {
		return models.Product{}, err
	}

	items, err := r.GetProductItems(ctx, product.Id)
	if err != nil {
		return models.Product{}, err
	}

	materials, err := r.GetProductsMaterials(ctx, product.Id)
	if err != nil {
		return models.Product{}, err
	}

	images, err := r.GetImages(ctx, product.Id)
	if err != nil {
		return models.Product{}, err
	}

	product.Images = images
	product.Items = items
	product.Materials = materials

	return product, nil
}

func (r *Init) GetProductsByIds(ctx goatcontext.Context, ids []int) ([]models.Product, error) {
	var products []models.Product
	if err := r.postgres.SelectContext(ctx, &products, queries.GetProductsById, pq.Array(ids)); err != nil {
		return nil, err
	}

	for i := range products {
		items, err := r.GetProductItems(ctx, products[i].Id)
		if err != nil {
			return nil, err
		}

		materials, err := r.GetProductsMaterials(ctx, products[i].Id)
		if err != nil {
			return nil, err
		}

		images, err := r.GetImages(ctx, products[i].Id)
		if err != nil {
			return nil, err
		}

		products[i].Images = images
		products[i].Items = items
		products[i].Materials = materials
	}

	return products, nil
}

func (r *Init) AddProducts(ctx goatcontext.Context, products []models.Product) ([]models.Product, error) {
	addedProducts := make([]models.Product, 0, len(products))

	for _, product := range products {
		var id int

		stmt, stmtErr := r.postgres.PrepareNamedContext(ctx, queries.AddProducts)
		if stmtErr != nil {
			return nil, stmtErr
		}

		if err := stmt.GetContext(ctx, &id, product); err != nil {
			return nil, err
		}

		product.Id = id
		product.Items = lo.Map(product.Items, func(item models.ProductItem, _ int) models.ProductItem {
			item.ProductId = id
			return item
		})

		product.Materials = lo.Map(product.Materials, func(item models.ProductMaterial, _ int) models.ProductMaterial {
			item.ProductId = id
			return item
		})

		product.Images = lo.Map(product.Images, func(item models.ProductImages, _ int) models.ProductImages {
			item.ProductId = id
			return item
		})

		if err := r.AddProductItems(ctx, product.Items); err != nil {
			return nil, err
		}

		if err := r.AddProductMaterials(ctx, product.Materials); err != nil {
			return nil, err
		}

		if err := r.AddImages(ctx, product.Images); err != nil {
			return nil, err
		}

		addedProducts = append(addedProducts, product)
	}

	return addedProducts, nil
}

func (r *Init) UpdateProducts(ctx goatcontext.Context, products []models.Product) error {
	tx, err := r.postgres.Begin()
	if err != nil {
		return err
	}

	for _, product := range products {
		if _, err = tx.ExecContext(ctx, queries.UpdateProducts, product); err != nil {
			_ = tx.Rollback()
			return err
		}

		if err = r.UpdateProductItems(ctx, product.Items); err != nil {
			_ = tx.Rollback()
			return err
		}

		if err = r.UpdateProductMaterials(ctx, product.Materials); err != nil {
			_ = tx.Rollback()
			return err
		}

		if err = r.UpdateImages(ctx, product.Images); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *Init) DeleteProducts(ctx goatcontext.Context, productIds []int) error {
	_, err := r.postgres.ExecContext(ctx, queries.DeleteProducts, pq.Array(productIds))
	return err
}

func (r *Init) GetProductItems(ctx goatcontext.Context, productId int) (items []models.ProductItem, err error) {
	return items, r.postgres.SelectContext(ctx, &items, queries.GetProductItems, productId)
}

func (r *Init) AddProductItems(ctx goatcontext.Context, productItems []models.ProductItem) error {
	for _, item := range productItems {
		if _, err := r.postgres.NamedExecContext(ctx, queries.AddProductItems, item); err != nil {
			return err
		}
	}

	return nil
}

func (r *Init) UpdateProductItems(ctx goatcontext.Context, productItems []models.ProductItem) error {
	for _, item := range productItems {
		if _, err := r.postgres.NamedExecContext(ctx, queries.UpdateProductItems, item); err != nil {
			return err
		}
	}

	return nil
}

func (r *Init) DeleteProductItems(ctx goatcontext.Context, id []int) error {
	_, err := r.postgres.ExecContext(ctx, queries.DeleteProductItems, pq.Array(id))
	return err
}

func (r *Init) GetAllMaterials(ctx goatcontext.Context) (materials []models.ProductMaterial, err error) {
	return materials, r.postgres.SelectContext(ctx, &materials, queries.GetAllMaterials)
}

func (r *Init) GetProductsMaterials(ctx goatcontext.Context, productId int) (materials []models.ProductMaterial, err error) {
	return materials, r.postgres.SelectContext(ctx, &materials, queries.GetProductsMaterials, productId)
}

func (r *Init) AddProductMaterials(ctx goatcontext.Context, productMaterials []models.ProductMaterial) error {
	for _, material := range productMaterials {
		if _, err := r.postgres.NamedExecContext(ctx, queries.AddProductMaterials, material); err != nil {
			return err
		}
	}

	return nil
}

func (r *Init) UpdateProductMaterials(ctx goatcontext.Context, productMaterials []models.ProductMaterial) error {
	for _, material := range productMaterials {
		if _, err := r.postgres.NamedExecContext(ctx, queries.UpdateProductMaterials, material); err != nil {
			return err
		}
	}

	return nil
}

func (r *Init) DeleteProductMaterials(ctx goatcontext.Context, id []int) error {
	_, err := r.postgres.ExecContext(ctx, queries.DeleteProductMaterials, pq.Array(id))
	return err
}

func (r *Init) GetImages(ctx goatcontext.Context, productId int) (images []models.ProductImages, err error) {
	return images, r.postgres.SelectContext(ctx, &images, queries.GetImages, productId)
}

func (r *Init) AddImages(ctx goatcontext.Context, productImages []models.ProductImages) error {
	for _, image := range productImages {
		if _, err := r.postgres.NamedExecContext(ctx, queries.AddImages, image); err != nil {
			return err
		}
	}

	return nil
}

func (r *Init) UpdateImages(ctx goatcontext.Context, productImages []models.ProductImages) error {
	for _, image := range productImages {
		if _, err := r.postgres.NamedExecContext(ctx, queries.UpdateImages, image); err != nil {
			return err
		}
	}

	return nil
}

func (r *Init) DeleteImages(ctx goatcontext.Context, id []int) error {
	_, err := r.postgres.ExecContext(ctx, queries.DeleteImages, pq.Array(id))
	return err
}

func (r *Init) GetProductItemsInfo(ctx goatcontext.Context, ids []int) (items []models.ProductItemInfo, err error) {
	return items, r.postgres.SelectContext(ctx, &items, queries.GetProductItemInfos, pq.Array(ids))
}
