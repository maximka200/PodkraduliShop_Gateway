package product

import (
	"context"
	"fmt"

	productv1 "github.com/maximka200/protobuff_product/gen"
)

func (ProductClient *ProductClient) NewProduct(ctx context.Context, req *productv1.NewProductRequest) (*productv1.NewProductResponse, error) {
	const op = "product.NewProduct"

	resp, err := ProductClient.Api.NewProduct(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (ProductClient *ProductClient) DeleteProduct(ctx context.Context, req *productv1.DeleteProductRequest) (*productv1.DeleteProductResponse, error) {
	const op = "product.DeleteProduct"

	resp, err := ProductClient.Api.DeleteProduct(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (ProductClient *ProductClient) GetProduct(ctx context.Context, req *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	const op = "product.GetProduct"

	resp, err := ProductClient.Api.GetProduct(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

// return multiple Product values
func (ProductClient *ProductClient) GetProducts(ctx context.Context, req *productv1.GetProductsRequest) (*productv1.GetProductsResponse, error) {
	const op = "product.GetProduct"

	resp, err := ProductClient.Api.GetProducts(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}
