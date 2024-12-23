package search

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	goatclient "github.com/GOAT-prod/goathttp/client"
	"github.com/GOAT-prod/goathttp/headers"
	"net/http"
	"strings"
)

type Client struct {
	httpClient goatclient.BaseClient
}

func NewClient(httpClient goatclient.BaseClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetActiveFilters(ctx goatcontext.Context) (result any, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "filters/active", nil, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return result, c.httpClient.Do(request, body, &result)
}

func (c *Client) GetCatalog(ctx goatcontext.Context, filters map[string][]string) (result any, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "catalog", nil, prepareParams(filters))
	if err != nil {
		return nil, err
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return result, c.httpClient.Do(request, body, &result)
}

func (c *Client) GetProductCatalog(ctx goatcontext.Context, productId int) (result any, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, fmt.Sprintf("catalog/%d", productId), nil, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return result, c.httpClient.Do(request, body, &result)
}

func prepareParams(filters map[string][]string) map[string]string {
	result := make(map[string]string)
	for key, values := range filters {
		result[key] = strings.Join(values, fmt.Sprintf("&%s=", key))
	}

	return result
}
