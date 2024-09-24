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
	PerPageCount int `form:"address"`
}

type ClientMethods interface {
	ProductMethods
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
		c.JSON(http.StatusBadRequest, gin.H{"error": errorsgeteway.ErrInvalidCredentials})
		return
	}
	// message with service gRPC
	resp, err := h.grpc.GetProduct(h.ctx, &productv1.GetProductRequest{Id: int64(id)})
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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

	c.JSON(http.StatusOK, getResp)
}

func (h *Handler) CreateNewProduct(c *gin.Context) {
	const op = "handlers.GetProductById"

	var product productv1.NewProductRequest
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// message with service gRPC
	resp, err := h.grpc.NewProduct(h.ctx, &product)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetId())
}

func (h *Handler) GetProductsList(c *gin.Context) {
	const op = "handlers.GetProductsList"

	var req MainPageReqParam
	if err := c.ShouldBind(&req); err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.grpc.GetProducts(h.ctx, &productv1.GetProductsRequest{Count: int64(req.PerPageCount)})
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusInternalServerError, resp)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	const op = "handlers.GetProductById"

	id, err := strconv.Atoi(c.Param("id"))
	if id < 0 || err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusBadRequest, gin.H{"error": errorsgeteway.ErrInvalidCredentials})
		return
	}
	// message with service gRPC
	resp, err := h.grpc.DeleteProduct(h.ctx, &productv1.DeleteProductRequest{Id: int64(id)})
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.log.Info(fmt.Sprintf("%t", resp.IsDelete))

	c.JSON(http.StatusOK, resp.IsDelete)
}
