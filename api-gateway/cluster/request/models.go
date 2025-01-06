package request

import "github.com/shopspring/decimal"

type Request struct {
	Id       int       `json:"id"`
	Type     string    `json:"type"`
	Products []Product `json:"products"`
}

// Product представляет информацию о продукте.
type Product struct {
	Id        int               `json:"id"`        // Уникальный идентификатор продукта.
	BrandName string            `json:"brandName"` // Название бренда продукта.
	Factory   Factory           `json:"factory"`   // Завод, производящий продукт.
	Name      string            `json:"name"`      // Название продукта.
	Price     decimal.Decimal   `json:"price"`     // Цена продукта.
	Items     []ProductItem     `json:"items"`     // Список товаров (размер, количество на складе и т.д.).
	Materials []ProductMaterial `json:"materials"` // Материалы, из которых состоит продукт.
	Images    []ProductImages   `json:"images"`    // Список изображений продукта.
	Status    string            `json:"status"`    // Статус продукта (например, "в наличии", "не в наличии").
}

// ProductItem представляет информацию о товаре на складе.
type ProductItem struct {
	Id         int             `json:"id"`         // Уникальный идентификатор товара.
	StockCount int             `json:"stockCount"` // Количество товара на складе.
	Size       int             `json:"size"`       // Размер товара (например, 42, 50 и т.д.).
	Weight     decimal.Decimal `json:"weight"`     // Вес товара.
	Color      string          `json:"color"`      // Цвет товара.
}

// ProductMaterial представляет информацию о материале, из которого состоит продукт.
type ProductMaterial struct {
	Id       int    `json:"Id"`       // Уникальный идентификатор материала.
	Material string `json:"Material"` // Название материала (например, "дерево", "металл").
}

// ProductImages представляет информацию об изображении продукта.
type ProductImages struct {
	Id       int    `json:"id"`       // Уникальный идентификатор изображения.
	ImageUrl string `json:"imageUrl"` // URL изображения продукта.
}

// Factory представляет информацию о заводе.
type Factory struct {
	Id          int    `db:"id"`   // Уникальный идентификатор завода.
	FactoryName string `db:"name"` // Название завода.
}
