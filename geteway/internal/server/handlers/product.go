package handlers

import (
	"fmt"
	errorsGeteway "geteway/internal/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MainPageReqParam struct {
	Page         int `form:"page"`
	PerPageCount int `form:"address"`
}

func (h *Handler) GetProductById(c *gin.Context) {
	const op = "handlers.GetProductById"
	id, err := strconv.Atoi(c.Param("id"))
	if id < 0 || err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusBadRequest, errorsGeteway.ErrInvalidCredentials)
		return
	}
	// message with service gRPC
	c.JSON(http.StatusBadRequest, "")
}

func (h *Handler) GetProductsList(c *gin.Context) {
	const op = "handlers.GetProductsList"
	var req MainPageReqParam
	if err := c.ShouldBind(&req); err != nil {
		h.log.Error(fmt.Sprintf("%s: %s", op, err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// message with service gRPC
	c.JSON(http.StatusBadRequest, "")
}
