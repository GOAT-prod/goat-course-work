package service

import (
	"cart-service/cluster/warehouse"
	"cart-service/database"
	"cart-service/domain"
	"cart-service/repository"
	"database/sql"
	"errors"
	"github.com/GOAT-prod/goatcontext"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"time"
)

type Cart interface {
	GetCart(ctx goatcontext.Context) (domain.Cart, error)
	AddCartItem(ctx goatcontext.Context, item domain.CartItem) (domain.CartItem, error)
	RemoveCartItem(ctx goatcontext.Context, id int) error
	UpdateCartItem(ctx goatcontext.Context, cartItem domain.CartItem) error
	ClearCartItems(ctx goatcontext.Context) error
	GetCartItems(ctx goatcontext.Context, ids []int) ([]domain.CartItem, error)
}

type CartServiceImpl struct {
	cartRepository  repository.Cart
	warehouseClient *warehouse.Client
}

func NewCartServiceImpl(cartRepository repository.Cart, warehouseClient *warehouse.Client) Cart {
	return &CartServiceImpl{
		cartRepository:  cartRepository,
		warehouseClient: warehouseClient,
	}
}

func (c *CartServiceImpl) GetCart(ctx goatcontext.Context) (domain.Cart, error) {
	cart, err := c.cartRepository.GetCart(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return domain.Cart{}, err
	}

	if cart.Id == 0 {
		cartId, cErr := c.cartRepository.CreateCart(ctx,
			//TODO: mappings
			database.Cart{
				CreateDate: time.Now(),
				UserId:     ctx.Authorize().UserId,
			})
		if cErr != nil {
			return domain.Cart{}, cErr
		}

		cart.Id = int(cartId)

		return domain.Cart{
			Id: cart.Id,
		}, nil
	}

	cartItems, err := c.cartRepository.GetCartItems(ctx, cart.Id)
	if err != nil {
		return domain.Cart{}, err
	}

	domainCartItems, err := c.fillCartItems(ctx, cartItems)
	if err != nil {
		return domain.Cart{}, err
	}

	//TODO: mappings
	domainCart := domain.Cart{
		Id: cart.Id,
		Total: lo.Sum(lo.Map(domainCartItems, func(item domain.CartItem, _ int) int {
			if !item.IsSelected {
				return 0
			}

			return int(item.Price.Mul(decimal.NewFromInt(int64(item.Count))).Ceil().IntPart())
		})),
		Items: domainCartItems,
	}

	return domainCart, nil
}

func (c *CartServiceImpl) AddCartItem(ctx goatcontext.Context, item domain.CartItem) (domain.CartItem, error) {
	cart, err := c.cartRepository.GetCart(ctx)
	if err != nil {
		return domain.CartItem{}, err
	}

	databaseCartItem := database.CartItem{
		ProductItemId: item.ProductItemId,
		Quantity:      item.Count,
		CartId:        cart.Id,
	}

	cartItemId, err := c.cartRepository.AddCartItem(ctx, databaseCartItem)
	if err != nil {
		return domain.CartItem{}, err
	}

	item.Id = int(cartItemId)

	return item, nil
}

func (c *CartServiceImpl) RemoveCartItem(ctx goatcontext.Context, id int) error {
	return c.cartRepository.DeleteCartItems(ctx, []int{id})
}

func (c *CartServiceImpl) UpdateCartItem(ctx goatcontext.Context, cartItem domain.CartItem) error {
	databaseCartItem := database.CartItem{
		Id:            cartItem.Id,
		ProductItemId: cartItem.ProductItemId,
		Quantity:      cartItem.Count,
		IsSelected:    cartItem.IsSelected,
	}

	return c.cartRepository.UpdateCartItem(ctx, databaseCartItem)
}

func (c *CartServiceImpl) ClearCartItems(ctx goatcontext.Context) error {
	cart, err := c.cartRepository.GetCart(ctx)
	if err != nil {
		return err
	}

	return c.cartRepository.ClearCartItems(ctx, cart.Id)
}

func (c *CartServiceImpl) GetCartItems(ctx goatcontext.Context, ids []int) ([]domain.CartItem, error) {
	cartItems, err := c.cartRepository.GetCartItemsByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	domainCartItems, err := c.fillCartItems(ctx, cartItems)
	if err != nil {
		return nil, err
	}

	return domainCartItems, nil
}

func (c *CartServiceImpl) fillCartItems(ctx goatcontext.Context, cartItems []database.CartItem) ([]domain.CartItem, error) {
	productItemsInfo, err := c.warehouseClient.GetProductItemsInfo(ctx, lo.Map(cartItems, func(item database.CartItem, _ int) int { return item.ProductItemId }))
	if err != nil {
		return nil, err
	}

	productItemsInfoMap := lo.Associate(productItemsInfo, func(item warehouse.ProductItemInfo) (int, warehouse.ProductItemInfo) {
		return item.Id, item
	})

	domainCartItems := make([]domain.CartItem, 0, len(cartItems))
	for _, cartItem := range cartItems {
		domainCartItems = append(domainCartItems, domain.CartItem{
			Id:            cartItem.Id,
			ProductItemId: cartItem.ProductItemId,
			Name:          productItemsInfoMap[cartItem.ProductItemId].Name,
			Price:         productItemsInfoMap[cartItem.ProductItemId].Price,
			Color:         productItemsInfoMap[cartItem.ProductItemId].Color,
			Size:          productItemsInfoMap[cartItem.ProductItemId].Size,
			Count:         cartItem.Quantity,
		})
	}

	return domainCartItems, nil
}
