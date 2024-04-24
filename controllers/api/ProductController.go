package api

import (
	"goOrders/database"
	"goOrders/models"
	"goOrders/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	db := database.GormConnect()
	result := []models.Products{}
	db.Where("status = 1").Find(&result)

	// Redis client
	_, err := service.GetRedisClient()
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, result)
}
