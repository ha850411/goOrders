package middleware

import (
	"context"
	"goOrders/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 後台中間件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginToken, _ := c.Cookie("loginToken")
		if loginToken == "" {
			c.Redirect(http.StatusMovedPermanently, "/admin/login")
		}
		client, _ := service.GetRedisClient()
		userInfo, _ := client.Get(context.Background(), loginToken).Result()
		if userInfo == "" {
			c.Redirect(http.StatusMovedPermanently, "/admin/login")
		}
		c.Set("userInfo", userInfo)
		c.Next()
	}
}
