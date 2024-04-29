package api

import (
	"goOrders/database"
	"goOrders/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	db := database.GormConnect()
	productsList := []models.Products{}

	// 接收參數並強制轉 int
	limitQuery := c.DefaultQuery("limit", "10")
	pageQuery := c.DefaultQuery("page", "0")
	keyword := c.DefaultQuery("keyword", "")
	limit, _ := strconv.Atoi(limitQuery)
	page, _ := strconv.Atoi(pageQuery)
	var totalRows int64 // 總筆數

	// 取得啟用的商品
	queryBuilder := db.Model(&models.Products{}).Where("status = 1")
	// 關鍵字
	if keyword != "" {
		queryBuilder.Where("name LIKE ?", "%"+keyword+"%")
	}
	queryBuilder.Count(&totalRows)
	queryBuilder.Debug().Limit(limit).Offset(page*limit - limit).Find(&productsList)

	data := make([]interface{}, 0)
	for _, v := range productsList {
		db.Table("products_type").Where("pid = ?", v.Id).Find(&v.ProductType)
		data = append(data, v)
	}

	result := map[string]interface{}{
		"totlaRows": totalRows,
		"data":      data,
	}

	c.JSON(http.StatusOK, result)
}
