package api

import (
	"goOrders/database"
	"goOrders/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	db := database.GormConnect()
	result := []models.Products{}
	db.Where("status = 1").Find(&result)
	c.JSON(http.StatusOK, result)
}
