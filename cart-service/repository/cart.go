package repository

import (
	"cart-service/database"
	"cart-service/database/queries"
	"github.com/GOAT-prod/goatcontext"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Cart interface {
	GetCart(ctx goatcontext.Context) (cart database.Cart, err error)
	CreateCart(ctx goatcontext.Context, cart database.Cart) (id int64, err error)
	GetCartItems(ctx goatcontext.Context, cartId int) (items []database.CartItem, err error)
	AddCartItem(ctx goatcontext.Context, cartItem database.CartItem) (id int64, err error)
	UpdateCartItem(ctx goatcontext.Context, cartItem database.CartItem) error
	DeleteCartItems(ctx goatcontext.Context, cartItemIds []int) error
	ClearCartItems(ctx goatcontext.Context, cartId int) error
}

type CartRepositoryImpl struct {
	postgres *sqlx.DB
}

func NewCartRepository(postgres *sqlx.DB) Cart {
	return &CartRepositoryImpl{
		postgres: postgres,
	}
}

func (r *CartRepositoryImpl) GetCart(ctx goatcontext.Context) (cart database.Cart, err error) {
	return cart, r.postgres.GetContext(ctx, &cart, queries.GetCart, ctx.Authorize().UserId)
}

func (r *CartRepositoryImpl) CreateCart(ctx goatcontext.Context, cart database.Cart) (id int64, err error) {
	result, err := r.postgres.NamedExecContext(ctx, queries.CreateCart, cart)
	if err != nil {
		return
	}

	return result.LastInsertId()
}

func (r *CartRepositoryImpl) GetCartItems(ctx goatcontext.Context, cartId int) (items []database.CartItem, err error) {
	return items, r.postgres.SelectContext(ctx, &items, queries.GetCartItems, cartId)
}

func (r *CartRepositoryImpl) AddCartItem(ctx goatcontext.Context, cartItem database.CartItem) (id int64, err error) {
	result, err := r.postgres.NamedExecContext(ctx, queries.AddCartItem, cartItem)
	if err != nil {
		return
	}

	return result.LastInsertId()
}

func (r *CartRepositoryImpl) UpdateCartItem(ctx goatcontext.Context, cartItem database.CartItem) error {
	_, err := r.postgres.NamedExecContext(ctx, queries.UpdateCartItem, cartItem)
	return err
}

func (r *CartRepositoryImpl) DeleteCartItems(ctx goatcontext.Context, cartItemIds []int) error {
	_, err := r.postgres.ExecContext(ctx, queries.DeleteCartItems, pq.Array(cartItemIds))
	return err
}

func (r *CartRepositoryImpl) ClearCartItems(ctx goatcontext.Context, cartId int) error {
	_, err := r.postgres.ExecContext(ctx, queries.ClearCartItems, cartId)
	return err
}
