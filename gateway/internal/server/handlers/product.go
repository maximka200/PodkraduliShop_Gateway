package handlers

import (
	"context"
	"fmt"
	errorsgeteway "geteway/internal/errors"

	"geteway/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	productv1 "github.com/maximka200/protobuff_product/gen"
)

type MainPageReqParam struct {
	Page         int `form:"page"`
	PerPageCount int `form:"pagePerCount"`
}

type ClientMethods interface {
	ProductMethods
}

type AuthMethod interface {
	// stub
	NewUser(ctx context.Context, req bool) (id int)
	Login(ctx context.Context, req bool) (jwt string)
	DeleteUser(ctx context.Context, req int) (isDelete bool)
}

type ProductMethods interface {
	NewProduct(ctx context.Context, req *productv1.NewProductRequest) (*productv1.NewProductResponse, error)
	DeleteProduct(ctx context.Context, req *productv1.DeleteProductRequest) (*productv1.DeleteProductResponse, error)
	GetProduct(ctx context.Context, req *productv1.GetProductRequest) (*productv1.GetProductResponse, error)
	GetProducts(ctx context.Context, req *productv1.GetProductsRequest) (*productv1.GetProductsResponse, error)
}

func (h *Handler) GetProductById(c *gin.Context) {
	const op = "handlers.GetProductById"

	id, err := strconv.Atoi(c.Param("id"))
	if id < 0 || err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusBadRequest, errorsgeteway.ErrInvalidCredentials)
		return
	}
	// message with service gRPC
	resp, err := h.grpc.GetProduct(h.ctx, &productv1.GetProductRequest{Id: int64(id)})
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, errorsgeteway.ErrInternal)
		return
	}

	/*
		getResp := model.GetProductResponse{
			Id:          resp.Id,
			ImageURL:    resp.ImageURL,
			Title:       resp.Title,
			Description: resp.Description,
			Price:       resp.Price,
			Currency:    string(resp.Currency),
			Discount:    resp.Discount,
			ProductURL:  resp.ProductURL,
		}
	*/
	result := h.giveProductModel(resp)

	c.JSON(http.StatusOK, result)
}

func (h *Handler) CreateNewProduct(c *gin.Context) {
	const op = "handlers.GetProductById"

	var product productv1.NewProductRequest
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, errorsgeteway.ErrInternal)
		return
	}

	if product.Title == "" || product.Description == "" || product.Price == 0 {
		h.log.Error(fmt.Sprintf("%s: empty field", op))
		c.JSON(http.StatusBadRequest, errorsgeteway.ErrInvalidCredentials)
	}
	// message with service gRPC
	resp, err := h.grpc.NewProduct(h.ctx, &product)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, errorsgeteway.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, resp.GetId())
}

func (h *Handler) GetProductsList(c *gin.Context) {
	const op = "handlers.GetProductsList"

	var req MainPageReqParam
	if err := c.ShouldBind(&req); err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, errorsgeteway.ErrInternal)
		return
	}

	if req.Page == 0 || req.PerPageCount == 0 {
		h.log.Error(fmt.Sprintf("%s: page or perPageCount = none", op))
		c.JSON(http.StatusInternalServerError, errorsgeteway.ErrInvalidCredentials)
		return
	}

	resp, err := h.grpc.GetProducts(h.ctx, &productv1.GetProductsRequest{Count: int64(req.PerPageCount)})
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, errorsgeteway.ErrInternal)
		return
	}

	var productList []model.GetProductResponse

	for _, elem := range resp.ProductList {
		productList = append(productList, h.giveProductModel(elem))
	}

	c.JSON(http.StatusOK, productList)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	const op = "handlers.GetProductById"

	id, err := strconv.Atoi(c.Param("id"))
	if id < 0 || err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusBadRequest, errorsgeteway.ErrInvalidCredentials)
		return
	}
	// message with service gRPC
	resp, err := h.grpc.DeleteProduct(h.ctx, &productv1.DeleteProductRequest{Id: int64(id)})
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, errorsgeteway.ErrInternal)
		return
	}

	h.log.Info(fmt.Sprintf("%t", resp.IsDelete))

	c.JSON(http.StatusOK, resp.IsDelete)
}

// gives the structure the appearance of model.GetProductResponse
func (h *Handler) giveProductModel(resp *productv1.GetProductResponse) model.GetProductResponse {
	mod := model.GetProductResponse{
		Id:          resp.GetId(),
		ImageURL:    resp.GetImageURL(),
		Title:       resp.GetTitle(),
		Description: resp.GetDescription(),
		Price:       resp.GetPrice(),
		Currency:    string(resp.GetCurrency()),
		Discount:    resp.GetDiscount(),
		ProductURL:  resp.GetProductURL(),
	}
	return mod
}
