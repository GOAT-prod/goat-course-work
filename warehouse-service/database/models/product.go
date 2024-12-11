package models

import "github.com/shopspring/decimal"

type Product struct {
	Id        int             `db:"id" json:"id"`
	BrandName string          `db:"brand" json:"brand"`
	Name      string          `db:"name" json:"name"`
	Price     decimal.Decimal `db:"price" json:"price"`
	Status    string          `db:"status" json:"status"`
	FactoryId int             `db:"factory_id" json:"factory_id"`

	Items     []ProductItem
	Materials []ProductMaterial
	Images    []ProductImages
}

type ProductMaterial struct {
	Id        int    `db:"id" json:"id"`
	ProductId int    `db:"product_id" json:"product_id"`
	Material  string `db:"material" json:"material"`
}

type ProductImages struct {
	Id        int    `db:"id" json:"id"`
	ProductId int    `db:"product_id" json:"product_id"`
	ImageUrl  string `db:"url" json:"image_url"`
}
