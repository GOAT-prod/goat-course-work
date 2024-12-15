package authservice

import (
	"github.com/GOAT-prod/goatcontext"
	goatclient "github.com/GOAT-prod/goathttp/client"
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

func (c *Client) Login(ctx goatcontext.Context, loginData LoginData) (tokens Tokens, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodPost, "login", loginData, nil)
	if err != nil {
		return Tokens{}, err
	}

	return tokens, c.httpClient.Do(request, body, &tokens)
}

func (c *Client) Logout(ctx goatcontext.Context, refreshToken string) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodPost, "logout", nil, nil)
	if err != nil {
		return err
	}

	c.httpClient.SetCookie(request, map[string]string{"refresh_token": refreshToken})

	return c.httpClient.Do(request, body, nil)
}

func (c *Client) Register(ctx goatcontext.Context, registerData RegisterData) (tokens Tokens, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodPost, "sign-up", registerData, nil)
	if err != nil {
		return Tokens{}, err
	}

	return tokens, c.httpClient.Do(request, body, &tokens)
}
