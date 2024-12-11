package clientservice

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/client"
	"net/http"
)

type Client struct {
	httpClient client.BaseClient
}

func NewClient(httpClient client.BaseClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetClientInfoShort(ctx goatcontext.Context, clientIds []int) (info []ClientInfoShort, err error) {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, "info/short", clientIds, nil)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Add("Authorization", ctx.AuthToken())

	return info, c.httpClient.Do(httpRequest, body, &info)
}
