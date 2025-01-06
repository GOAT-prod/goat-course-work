package notifier

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

func (c *Client) SendMessage(ctx goatcontext.Context, msg MailMessage) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodPost, "mail", msg, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(request, body, nil)
}
