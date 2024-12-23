package report

import (
	"github.com/GOAT-prod/goatcontext"
	goatclient "github.com/GOAT-prod/goathttp/client"
	"github.com/GOAT-prod/goathttp/headers"
	"net/http"
	"strconv"
)

type Client struct {
	httpClient goatclient.BaseClient
}

func NewClient(httpClient goatclient.BaseClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetSellReport(ctx goatcontext.Context, userId int) (result any, err error) {
	params := map[string]string{
		"userId": strconv.Itoa(userId),
	}

	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "report/sell", nil, params)
	if err != nil {
		return "", err
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return result, c.httpClient.Do(request, body, &result)
}

func (c *Client) GetOrderReport(ctx goatcontext.Context, userId int) (result any, err error) {
	params := map[string]string{
		"userId": strconv.Itoa(userId),
	}

	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "report/order", nil, params)
	if err != nil {
		return "", err
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return result, c.httpClient.Do(request, body, &result)
}
