package domain

import "github.com/shopspring/decimal"

type Product struct {
	Id        int               `json:"id"`
	BrandName string            `json:"brandName"`
	Factory   Factory           `json:"factory"`
	Name      string            `json:"name"`
	Price     decimal.Decimal   `json:"price"`
	Items     []ProductItem     `json:"items"`
	Materials []ProductMaterial `json:"materials"`
	Images    []ProductImages   `json:"images"`
	Status    ProductStatus     `json:"status"`
}

type ProductMaterial struct {
	Id       int    `json:"Id"`
	Material string `json:"Material"`
}

type ProductImages struct {
	Id       int    `json:"id"`
	ImageUrl string `json:"imageUrl"`
}

type Factory struct {
	Id          int    `db:"id"`
	FactoryName string `db:"name"`
}