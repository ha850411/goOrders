package api

import (
	"goOrders/database"
	"goOrders/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	db := database.GormConnect()
	productsList := []models.Products{}

	// 取得啟用的商品
	db.Where("status = 1").Find(&productsList)

	result := make([]interface{}, 0)
	for _, v := range productsList {
		db.Table("products_type").Where("pid = ?", v.Id).Find(&v.ProductType)
		result = append(result, v)
	}

	c.JSON(http.StatusOK, result)
}
