package domain

import "github.com/shopspring/decimal"

type Product struct {
	Id        int             `json:"id"`
	BrandName string          `json:"brand"`
	Name      string          `json:"name"`
	Price     decimal.Decimal `json:"price"`
	Status    string          `json:"status"`
	FactoryId int             `json:"factory_id"`

	Items     []ProductItem
	Materials []ProductMaterial
	Images    []ProductImages
}

type ProductItem struct {
	Id         int             `json:"id"`
	ProductId  int             `json:"productId"`
	StockCount int             `json:"stockCount"`
	Size       int             `json:"size"`
	Weight     decimal.Decimal `json:"weight"`
	Color      string          `json:"color"`
}

type ProductMaterial struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id"`
	Material  string `json:"material"`
}

type ProductImages struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id"`
	ImageUrl  string `json:"image_url"`
}
