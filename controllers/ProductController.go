package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductManager(c *gin.Context) {
	output := GetCommonOutput(c, "product")
	c.HTML(http.StatusOK, "productList", output)
}

func ProductTypeManager(c *gin.Context) {
	output := GetCommonOutput(c, "productType")
	c.HTML(http.StatusOK, "productTypeList", output)
}
