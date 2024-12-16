package queries

import _ "embed"

var (
	//go:embed sql/get_products.sql
	GetProducts string

	//go:embed sql/get_product_by_id.sql
	GetProductsById string

	//go:embed sql/add_product.sql
	AddProducts string

	//go:embed sql/update_product.sql
	UpdateProducts string

	//go:embed sql/delete_product.sql
	DeleteProducts string

	//go:embed sql/get_product_items.sql
	GetProductItems string

	//go:embed sql/add_product_items.sql
	AddProductItems string

	//go:embed sql/update_product_items.sql
	UpdateProductItems string

	//go:embed sql/delete_product_items.sql
	DeleteProductItems string

	//go:embed sql/get_all_materials.sql
	GetAllMaterials string

	//go:embed sql/get_product_materials.sql
	GetProductsMaterials string

	//go:embed sql/add_product_materials.sql
	AddProductMaterials string

	//go:embed sql/update_product_materials.sql
	UpdateProductMaterials string

	//go:embed sql/delete_product_materials.sql
	DeleteProductMaterials string

	//go:embed sql/get_images.sql
	GetImages string

	//go:embed sql/add_images.sql
	AddImages string

	//go:embed sql/update_images.sql
	UpdateImages string

	//go:embed sql/delete_images.sql
	DeleteImages string

	//go:embed sql/get_product_item_info.sql
	GetProductItemInfos string
)
