package order

import (
	"context"
	"github.com/GOAT-prod/goathttp/client"
	"net/http"
	"report-service/database"
)

type Client struct {
	httpClient client.BaseClient
}

func NewClient(httpClient client.BaseClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetLatestOrders(ctx context.Context) (items []database.ReportItem, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "orders/latest", nil, nil)
	if err != nil {
		return
	}

	return items, c.httpClient.Do(request, body, &items)
}
