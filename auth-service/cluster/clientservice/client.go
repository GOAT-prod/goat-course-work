package clientservice

import (
	"auth-service/cluster"
	"auth-service/domain"
	"context"
	"net/http"
)

type Client struct {
	httpClient cluster.BaseClient
}

func NewClient(client cluster.BaseClient) *Client {
	return &Client{
		httpClient: client,
	}
}

func (c *Client) AddClientData(ctx context.Context, clientData domain.ClientData) (client domain.ClientData, err error) {
	req, body, err := c.httpClient.Request(ctx, http.MethodPost, "client", clientData, nil)
	if err != nil {
		return domain.ClientData{}, err
	}

	return client, c.httpClient.Do(req, body, &client)
}
