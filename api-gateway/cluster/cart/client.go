package cart

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	goatclient "github.com/GOAT-prod/goathttp/client"
	"net/http"
)

type Client struct {
	httpClient goatclient.BaseClient
}

func NewClient(httpClient goatclient.BaseClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetCart(ctx goatcontext.Context) (cart Cart, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "cart", nil, nil)
	if err != nil {
		return Cart{}, err
	}

	return cart, c.httpClient.Do(request, body, &cart)
}

func (c *Client) AddCartItem(ctx goatcontext.Context, item Item) (result Item, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodPost, "cart/item", item, nil)
	if err != nil {
		return Item{}, err
	}

	return result, c.httpClient.Do(request, body, &result)
}

func (c *Client) UpdateCartItem(ctx goatcontext.Context, item Item) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodPut, "cart/item", item, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(request, body, nil)
}

func (c *Client) DeleteCartItem(ctx goatcontext.Context, itemId int) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodDelete, fmt.Sprintf("cart/item/%d", itemId), nil, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(request, body, nil)
}

func (c *Client) ClearCart(ctx goatcontext.Context) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodDelete, "cart/clear", nil, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(request, body, nil)
}
