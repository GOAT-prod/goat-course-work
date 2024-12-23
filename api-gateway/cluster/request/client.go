package request

import (
	"fmt"
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

func (c *Client) GetRequests(ctx goatcontext.Context) (result any, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "requests", nil, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return result, c.httpClient.Do(request, body, &result)
}

func (c *Client) UpdateRequestStatus(ctx goatcontext.Context, id int, status string) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodPut, fmt.Sprintf("request/%d/status/%s", id, status), nil, nil)
	if err != nil {
		return err
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return c.httpClient.Do(request, body, nil)
}
