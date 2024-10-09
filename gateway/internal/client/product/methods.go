package product

import (
	"context"
	"fmt"

	authv1 "github.com/maximka200/buffpr/gen/go/sso"
	productv1 "github.com/maximka200/protobuff_product/gen"
)

func (ProductClient *ProductClient) NewProduct(ctx context.Context,
	req *productv1.NewProductRequest) (*productv1.NewProductResponse, error) {
	const op = "product.NewProduct"

	resp, err := ProductClient.Api.NewProduct(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (ProductClient *ProductClient) DeleteProduct(ctx context.Context,
	req *productv1.DeleteProductRequest) (*productv1.DeleteProductResponse, error) {
	const op = "product.DeleteProduct"

	resp, err := ProductClient.Api.DeleteProduct(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (ProductClient *ProductClient) GetProduct(ctx context.Context,
	req *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	const op = "product.GetProduct"

	resp, err := ProductClient.Api.GetProduct(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

// return multiple Product values
func (ProductClient *ProductClient) GetProducts(ctx context.Context,
	req *productv1.GetProductsRequest) (*productv1.GetProductsResponse, error) {
	const op = "product.GetProduct"

	resp, err := ProductClient.Api.GetProducts(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (AuthClient *AuthClient) Login(ctx context.Context,
	req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	const op = "product.Login"

	resp, err := AuthClient.Api.Login(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (AuthClient *AuthClient) NewUser(ctx context.Context,
	req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	const op = "product.Register"

	resp, err := AuthClient.Api.Register(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}
