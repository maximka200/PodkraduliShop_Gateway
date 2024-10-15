package handlers

import (
	"context"
	"fmt"
	errorsgateway "gateway/internal/errors"
	"gateway/internal/libs/jwt"

	"gateway/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	authv1 "github.com/maximka200/buffpr/gen/go/sso"
	productv1 "github.com/maximka200/protobuff_product/gen"
)

type ClientMethods interface {
	ProductMethods
	AuthMethods
}

type AuthMethods interface {
	NewUser(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error)
	Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error)
	//DeleteUser(ctx context.Context, id int) (isDelete bool)
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
		c.JSON(http.StatusBadRequest, errorsgateway.ErrInvalidCredentials)
		return
	}
	// message with service gRPC
	resp, err := h.grpc.GetProduct(h.ctx, &productv1.GetProductRequest{Id: int64(id)})
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, errorsgateway.ErrInternal)
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
	result := h.parseInProductModel(resp)

	c.JSON(http.StatusOK, result)
}

func (h *Handler) CreateNewProduct(c *gin.Context) {
	const op = "handlers.GetProductById"

	var product productv1.NewProductRequest
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, errorsgateway.ErrInternal)
		return
	}

	if product.Title == "" || product.Description == "" || product.Price == 0 {
		h.log.Error(fmt.Sprintf("%s: empty field", op))
		c.JSON(http.StatusBadRequest, errorsgateway.ErrInvalidCredentials)
	}
	// message with service gRPC
	resp, err := h.grpc.NewProduct(h.ctx, &product)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, errorsgateway.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, resp.GetId())
}

func (h *Handler) GetProductsList(c *gin.Context) {
	const op = "handlers.GetProductsList"

	var req model.MainPageReqParam
	if err := c.ShouldBind(&req); err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, errorsgateway.ErrInternal)
		return
	}

	if req.Page == 0 || req.PerPageCount == 0 {
		h.log.Error(fmt.Sprintf("%s: page or perPageCount = none", op))
		c.JSON(http.StatusInternalServerError, errorsgateway.ErrInvalidCredentials)
		return
	}

	resp, err := h.grpc.GetProducts(h.ctx, &productv1.GetProductsRequest{Count: int64(req.PerPageCount)})
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, errorsgateway.ErrInternal)
		return
	}

	var productList []model.GetProductResponse

	for _, elem := range resp.ProductList {
		productList = append(productList, h.parseInProductModel(elem))
	}

	c.JSON(http.StatusOK, productList)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	const op = "handlers.GetProductById"

	id, err := strconv.Atoi(c.Param("id"))
	if id < 0 || err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusBadRequest, errorsgateway.ErrInvalidCredentials)
		return
	}
	// message with service gRPC
	resp, err := h.grpc.DeleteProduct(h.ctx, &productv1.DeleteProductRequest{Id: int64(id)})
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, errorsgateway.ErrInternal)
		return
	}

	h.log.Info(fmt.Sprintf("%t", resp.IsDelete))

	c.JSON(http.StatusOK, resp.IsDelete)
}

// gives the structure the appearance of model.GetProductResponse
func (h *Handler) parseInProductModel(resp *productv1.GetProductResponse) model.GetProductResponse {
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

func (h *Handler) Login(c *gin.Context) {
	const op = "handlers.Login"

	var usr model.User
	if err := c.ShouldBind(&usr); err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorsgateway.ErrInternal})
		return
	}

	if usr.AppId == 0 {
		usr.AppId = 1
	}

	resp, err := h.grpc.Login(c.Request.Context(), &authv1.LoginRequest{Email: usr.Email,
		Password: usr.Password, AppId: int64(usr.AppId)})
	if err != nil {
		// error validation
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorsgateway.ErrInternal})
		return
	}

	h.log.Info(op)

	c.JSON(http.StatusOK, resp.Token)
}

func (h *Handler) Register(c *gin.Context) {
	const op = "handlers.Login"

	var usr model.User
	if err := c.ShouldBind(&usr); err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, errorsgateway.ErrInternal)
		return
	}

	resp, err := h.grpc.NewUser(c.Request.Context(), &authv1.RegisterRequest{Email: usr.Email,
		Password: usr.Password})
	if err != nil {
		// error validation
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorsgateway.ErrInternal})
		return
	}

	h.log.Info(op)

	c.JSON(http.StatusOK, resp.UserId)
}

func (h *Handler) VerificationJWT(c *gin.Context) {
	const op = "handlers.VerificationJWT"

	jwtToken := c.GetHeader("Authorization")[7:] // Bearer: = 7 symbols
	h.log.Info(jwtToken)
	ok, err := jwt.CheckJWT(c.Request.Context(), h.log, jwtToken)
	if err != nil {
		// error validation
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorsgateway.ErrInternal})
		return
	}

	h.log.Info(op)

	c.JSON(http.StatusOK, gin.H{"ok": ok})
}
