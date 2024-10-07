package handlers

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	grpc ClientMethods
	log  *slog.Logger
	ctx  context.Context
}

func NewHandler(log *slog.Logger, clientMethods ClientMethods, ctx context.Context) *Handler {
	return &Handler{clientMethods, log, ctx}
}

func (h *Handler) InitRouter() *gin.Engine {
	engine := gin.Default()

	product := engine.Group("/product")
	{
		product.GET("/productbyid/:id", h.GetProductById)
		product.GET("/productlist", h.GetProductsList)
		product.POST("/create", h.CreateNewProduct)
		product.DELETE("/delete/:id", h.DeleteProduct)
	}

	auth := engine.Group("/auth")
	{
		auth.POST("/register")
		auth.POST("/login")
		auth.DELETE("/delete")
	}

	return engine
}
