package userservice

import (
	"auth-service/cluster"
	"auth-service/domain"
	"context"
	"net/http"
)

const (
	_getUserUri            = "user"
	_saveUserUri           = "user/registration"
	_updateUserPasswordUri = "user/password/new"
)

type Client struct {
	httpClient cluster.BaseClient
}

func NewClient(client cluster.BaseClient) *Client {
	return &Client{
		httpClient: client,
	}
}

func (c *Client) GetUserByUserName(ctx context.Context, userName string) (user domain.User, err error) {
	params := map[string]string{
		"username": userName,
	}

	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodGet, _getUserUri, nil, params)
	if err != nil {
		return
	}

	return user, c.httpClient.Do(httpRequest, body, &user)
}

func (c *Client) RegisterUser(ctx context.Context, user domain.User) (result domain.User, err error) {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPost, _saveUserUri, user, nil)
	if err != nil {
		return
	}

	return result, c.httpClient.Do(httpRequest, body, &result)
}

func (c *Client) UpdateUserPassword(ctx context.Context, user domain.UpdatePasswordRequest) error {
	httpRequest, body, err := c.httpClient.Request(ctx, http.MethodPut, _updateUserPasswordUri, user, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(httpRequest, body, nil)
}
