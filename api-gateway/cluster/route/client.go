package route

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

func (c *Client) GetBestRoute(ctx goatcontext.Context, locations []ServiceLocation) (route Route, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "route/best", locations, nil)
	if err != nil {
		return
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return route, c.httpClient.Do(request, body, &route)
}
