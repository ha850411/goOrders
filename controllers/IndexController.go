package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 首頁
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{})
}
