package warehousesevice

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

func (c *Client) GetProducts(ctx goatcontext.Context) (products []Product, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "products", nil, nil)
	if err != nil {
		return
	}

	return products, c.httpClient.Do(request, body, &products)
}

func (c *Client) GetProductById(ctx goatcontext.Context, productId int) (product Product, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, fmt.Sprintf("product/%d", productId), nil, nil)
	if err != nil {
		return
	}

	return product, c.httpClient.Do(request, body, &product)
}

func (c *Client) AddProducts(ctx goatcontext.Context, products []Product) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodPost, "products", products, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(request, body, nil)
}

func (c *Client) UpdateProducts(ctx goatcontext.Context, products []Product) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodPut, "products", products, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(request, body, nil)
}

func (c *Client) DeleteProducts(ctx goatcontext.Context, productsIds []int) error {
	request, body, err := c.httpClient.Request(ctx, http.MethodDelete, "products", productsIds, nil)
	if err != nil {
		return err
	}

	return c.httpClient.Do(request, body, nil)
}

func (c *Client) GetAllMaterials(ctx goatcontext.Context) (materials []ProductMaterial, err error) {
	request, body, err := c.httpClient.Request(ctx, http.MethodGet, "materials/list", nil, nil)
	if err != nil {
		return
	}
	
	return materials, c.httpClient.Do(request, body, &materials)
}
