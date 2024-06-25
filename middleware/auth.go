package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"goOrders/database"
	"goOrders/models"
	"goOrders/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 後台中間件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginToken, _ := c.Cookie("loginToken")
		if loginToken == "" {
			c.Redirect(http.StatusFound, "/admin/login")
		}
		client, _ := service.GetRedisClient()
		userInfoStr, _ := client.Get(context.Background(), loginToken).Result()
		if userInfoStr == "" {
			c.Redirect(http.StatusFound, "/admin/login")
		}

		var userInfo models.Users
		json.Unmarshal([]byte(userInfoStr), &userInfo)

		c.Set("userInfo", userInfo)
		c.Next()
	}
}

// 更新 redis 及 session
func SetUserRedis(c *gin.Context) string {

	// 取得原本的 userInfo
	originalUserInfo, exist := c.Get("userInfo")
	if !exist {
		return "userInfo is empty"
	}

	// 抓最新的 userInfo
	db := database.GormConnect()
	newUserInfo := &models.Users{}
	db.Where("username = ?", originalUserInfo.(models.Users).Username).First(&newUserInfo)

	client, err2 := service.GetRedisClient()
	if err2 != "" {
		return err2
	}

	// 將登入資訊 json encode
	jsonData, _ := json.Marshal(newUserInfo)

	// 將登入資訊寫入 redis
	loginToken, _ := c.Cookie("loginToken")
	ttl := client.TTL(context.Background(), loginToken)
	fmt.Printf("ttl: %v\n", ttl)
	client.Set(context.Background(), loginToken, string(jsonData), ttl.Val())

	// 寫入中間件
	c.Set("userInfo", newUserInfo)

	return ""
}
