package clientservice

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

func (c *Client) GetClients(ctx goatcontext.Context) (clients []ClientInfo, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "clients", nil, nil)
	if err != nil {
		return
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return clients, c.httpClient.Do(request, body, &clients)
}

func (c *Client) GetClientById(ctx goatcontext.Context, clientId int) (client ClientInfo, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, fmt.Sprintf("client/%d", clientId), nil, nil)
	if err != nil {
		return
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return client, c.httpClient.Do(request, body, &client)
}

func (c *Client) UpdateClient(ctx goatcontext.Context, client ClientInfo) (result ClientInfo, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodPut, "client", client, nil)
	if err != nil {
		return
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return result, c.httpClient.Do(request, body, &result)
}

func (c *Client) DeleteClient(ctx goatcontext.Context, clientId int) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodDelete, fmt.Sprintf("client/%d", clientId), nil, nil)
	if err != nil {
		return err
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	return c.httpClient.Do(request, body, nil)
}
