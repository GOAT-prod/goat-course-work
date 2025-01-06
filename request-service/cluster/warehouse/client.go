package warehouse

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/client"
	"github.com/GOAT-prod/goathttp/headers"
	"net/http"
	"request-service/domain"
)

type Client struct {
	httpClient client.BaseClient
}

func NewClient(httpClient client.BaseClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetDetailedProduct(ctx goatcontext.Context, productId int) (product domain.Product, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, fmt.Sprintf("product/%d", productId), nil, nil)
	if err != nil {
		return domain.Product{}, err
	}

	request.Header.Add(headers.AuthorizationHeader(), ctx.AuthToken())

	return product, c.httpClient.Do(request, body, &product)
}
