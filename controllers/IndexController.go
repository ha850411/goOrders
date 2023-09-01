package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// 首頁
func Index(c *gin.Context) {
	c.String(http.StatusOK, "test")
}

// 自動證書驗證
func AutoCert(c *gin.Context) {
	filename := c.Param("files")
	b, err := os.ReadFile("./.well-known/acme-challenge/" + filename)
	if err != nil {
		NotFound(c)
	}
	c.String(http.StatusOK, string(b))
}

// 404 page
func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404", nil)
}
