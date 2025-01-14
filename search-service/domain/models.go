package domain

import (
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

type ProductStatus string

const (
	Unknown        ProductStatus = "unknown"
	WaitingApprove ProductStatus = "waiting_approve"
	Approved       ProductStatus = "approved"
	Editing        ProductStatus = "editing"
	Deleted        ProductStatus = "deleted"
)

type AppliedFilters struct {
	Size     []int           `json:"size"`
	Color    []string        `json:"color"`
	Brand    []string        `json:"brand"`
	Material []string        `json:"material"`
	MinPrice decimal.Decimal `json:"minPrice"`
	MaxPrice decimal.Decimal `json:"maxPrice"`
}

type Catalog struct {
	Filters  any       `json:"filters"`
	SearchId string    `json:"searchId"`
	Products []Product `json:"products"`
}

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

type ProductItem struct {
	Id         int             `json:"id"`
	StockCount int             `json:"stockCount"`
	Size       int             `json:"size"`
	Weight     decimal.Decimal `json:"weight"`
	Color      string          `json:"color"`
}

type ProductMaterial struct {
	Id       int    `json:"Id"`
	Material string `json:"Material"`
}

func (p Product) GetMaterialNames() []string {
	return lo.Map(p.Materials, func(item ProductMaterial, _ int) string {
		return item.Material
	})
}

type ProductImages struct {
	Id       int    `json:"id"`
	ImageUrl string `json:"imageUrl"`
}

type Factory struct {
	Id          int    `db:"id"`
	FactoryName string `db:"name"`
}
