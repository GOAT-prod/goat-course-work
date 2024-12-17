package order

import (
	"github.com/GOAT-prod/goatcontext"
	goatclient "github.com/GOAT-prod/goathttp/client"
	"github.com/GOAT-prod/goathttp/headers"
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

func (c *Client) GetUserOrders(ctx goatcontext.Context) (orders []Order, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "orders", nil, nil)
	if err != nil {
		return
	}

	request.Header.Add(headers.AuthorizationHeader(), ctx.AuthToken())

	return orders, c.httpClient.Do(request, body, &orders)
}

func (c *Client) CreateOrder(ctx goatcontext.Context, cartItemIds []int) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodPost, "order", cartItemIds, nil)
	if err != nil {
		return err
	}

	request.Header.Add(headers.AuthorizationHeader(), ctx.AuthToken())

	return c.httpClient.Do(request, body, nil)
}
