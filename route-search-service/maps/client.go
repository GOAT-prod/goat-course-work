package maps

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/GOAT-prod/goathttp/client"
	"net/http"
)

type Client struct {
	apiKey     string
	httpClient client.BaseClient
}

func NewClient(apiKey string, httpClient client.BaseClient) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (c *Client) GetGeocoderData(ctx goatcontext.Context, address string) (response Response, err error) {
	queryParams := map[string]string{
		"apikey":  c.apiKey,
		"geocode": address,
		"format":  "json",
	}

	request, _, err := c.httpClient.Request(ctx, http.MethodGet, "", nil, queryParams)
	if err != nil {
		return Response{}, err
	}
	
	return response, c.httpClient.Do(request, nil, &response)
}
