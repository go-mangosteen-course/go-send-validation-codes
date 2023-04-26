package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary      用来测试 API 是否正常工作
// @Description  用来测试 API 是否正常工作
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
