package warehouse

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/client"
	"github.com/GOAT-prod/goathttp/headers"
	"net/http"
	"search-service/domain"
)

type Client struct {
	httpClient client.BaseClient
}

func NewClient(httpClient client.BaseClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetProducts(ctx goatcontext.Context) (products []domain.Product, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "products", nil, nil)
	if err != nil {
		return
	}

	request.Header.Add(headers.AuthorizationHeader(), ctx.AuthToken())

	return products, c.httpClient.Do(request, body, &products)
}

func (c *Client) GetProduct(ctx goatcontext.Context, id int) (product domain.Product, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, fmt.Sprintf("product/%d", id), nil, nil)
	if err != nil {
		return
	}

	request.Header.Add(headers.AuthorizationHeader(), ctx.AuthToken())

	return product, c.httpClient.Do(request, body, &product)
}
