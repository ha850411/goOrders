package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductList(c *gin.Context) {
	output := GetCommonOutput(c, "product")
	c.HTML(http.StatusOK, "productList", output)
}
