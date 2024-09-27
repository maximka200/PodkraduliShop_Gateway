package product

import (
	"context"
	"fmt"

	productv1 "github.com/maximka200/protobuff_product/gen"
)

func (client *Client) NewProduct(ctx context.Context, req *productv1.NewProductRequest) (*productv1.NewProductResponse, error) {
	const op = "product.NewProduct"

	resp, err := client.Api.NewProduct(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (client *Client) DeleteProduct(ctx context.Context, req *productv1.DeleteProductRequest) (*productv1.DeleteProductResponse, error) {
	const op = "product.DeleteProduct"

	resp, err := client.Api.DeleteProduct(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (client *Client) GetProduct(ctx context.Context, req *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	const op = "product.GetProduct"

	resp, err := client.Api.GetProduct(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (client *Client) GetProducts(ctx context.Context, req *productv1.GetProductsRequest) (*productv1.GetProductsResponse, error) {
	// return multiple values
	const op = "product.GetProduct"

	resp, err := client.Api.GetProducts(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}
