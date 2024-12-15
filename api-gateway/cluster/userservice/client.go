package userservice

import (
	"fmt"
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

func (c *Client) GetUsers(ctx goatcontext.Context) (users []User, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "users", nil, nil)
	if err != nil {
		return
	}

	return users, c.httpClient.Do(request, body, &users)
}

func (c *Client) GetUserById(ctx goatcontext.Context, userId int) (user User, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, fmt.Sprintf("user/%d", userId), nil, nil)
	if err != nil {
		return
	}

	return user, c.httpClient.Do(request, body, &user)
}

func (c *Client) AddUser(ctx goatcontext.Context, user User) (result User, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodPost, "user", user, nil)
	if err != nil {
		return
	}

	return result, c.httpClient.Do(request, body, &result)
}

func (c *Client) UpdateUser(ctx goatcontext.Context, user User) (result User, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodPut, "user", user, nil)
	if err != nil {
		return
	}

	return result, c.httpClient.Do(request, body, &result)
}

func (c *Client) DeleteUser(ctx goatcontext.Context, userId int) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodDelete, fmt.Sprintf("user/%d", userId), nil, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(request, body, &err)
}
