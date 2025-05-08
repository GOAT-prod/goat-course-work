package report

import (
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	goatclient "github.com/GOAT-prod/goathttp/client"
	"github.com/GOAT-prod/goathttp/headers"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient goatclient.BaseClient
}

func NewClient(httpClient goatclient.BaseClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetSellReport(ctx goatcontext.Context, userId int, date time.Time) (result io.ReadCloser, err error) {
	request, _, err := c.httpClient.Request(ctx, http.MethodGet, fmt.Sprintf("report/user/%d/%s", userId, date.Format(time.DateOnly)), nil, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", response.StatusCode)
	}

	return response.Body, nil
}

func (c *Client) GetOrderReport(ctx goatcontext.Context, userId int, date time.Time) (result io.ReadCloser, err error) {
	request, _, err := c.httpClient.Request(ctx, http.MethodGet, fmt.Sprintf("report/factory/%d/%s", userId, date.Format(time.DateOnly)), nil, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set(headers.AuthorizationHeader(), ctx.AuthToken())

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", response.StatusCode)
	}

	return response.Body, nil
}
