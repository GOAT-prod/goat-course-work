package cart

import "github.com/shopspring/decimal"

// Cart представляет корзину покупок пользователя.
type Cart struct {
	Id    int    `json:"id"`    // Уникальный идентификатор корзины.
	Total int    `json:"total"` // Общая сумма стоимости всех товаров в корзине.
	Items []Item `json:"items"` // Список товаров, добавленных в корзину.
}

// Item представляет информацию о товаре, добавленном в корзину.
type Item struct {
	Id            int             `json:"id"`         // Уникальный идентификатор элемента корзины.
	ProductItemId int             `json:"productId"`  // Идентификатор товара (связь с ProductItem).
	Name          string          `json:"name"`       // Название товара.
	Price         decimal.Decimal `json:"price"`      // Цена за единицу товара.
	Color         string          `json:"color"`      // Цвет товара.
	Size          int             `json:"size"`       // Размер товара (например, 42, 50 и т.д.).
	Count         int             `json:"count"`      // Количество единиц данного товара в корзине.
	IsSelected    bool            `json:"isSelected"` // Признак, выбран ли данный товар (например, для оформления заказа).
}
