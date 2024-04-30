package api

import (
	"goOrders/database"
	"goOrders/middleware"
	"goOrders/models"

	"github.com/gin-gonic/gin"
)

func UnsetLineNotify(c *gin.Context) {
	db := database.GormConnect()
	user, _ := c.Get("userInfo")
	userInfo := user.(models.Users)
	result := db.Model(&models.Users{}).Where("username = ?", userInfo.Username).Update("linebot_token", "")
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
	}
	// 重新寫入 redis
	middleware.SetUserRedis(c)

	c.JSON(200, gin.H{
		"message":      "success",
		"rowsAffected": result.RowsAffected,
	})
}
