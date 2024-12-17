package cart

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

func (c *Client) GetCartItems(ctx goatcontext.Context, ids []int) (items []Item, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodPost, "cart/items", ids, nil)
	if err != nil {
		return
	}

	request.Header.Add(headers.AuthorizationHeader(), ctx.AuthToken())

	return items, c.httpClient.Do(request, body, &items)
}
