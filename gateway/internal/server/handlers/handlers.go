package handlers

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	str string
	log *slog.Logger
}

func NewHandler(log *slog.Logger) *Handler {
	return &Handler{"", log}
}

func (h *Handler) InitRouter() *gin.Engine {
	engine := gin.Default()

	engine.GET("/product/:id", h.GetProductById)
	engine.GET("/products", h.GetProductsList)

	return engine
}
