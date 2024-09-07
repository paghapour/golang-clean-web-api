package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paghapour/golang-clean-web-api/api/helper"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working!", true, 0))
}
