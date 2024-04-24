package controllers

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"goOrders/database"
	"goOrders/models"
	"goOrders/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 登入頁
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{})
}

// 登入驗證
func Auth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 將密碼做 md5
	hash := md5.Sum([]byte(password))
	password = hex.EncodeToString(hash[:])

	// db connect
	db := database.GormConnect()
	userData := &models.Users{}
	db.Where("username = ? AND password = ?", username, password).First(&userData)
	fmt.Printf("userData.Username: %v\n", userData.Username)

	// base64
	token := generateSessionID()

	if userData.Username != "" {
		// 記住我
		remember := c.PostForm("remember")
		expire := 3600
		if remember == "Y" {
			expire = 86400 * 30
		}
		// 設定隨機 token
		c.SetCookie("loginToken", token, expire, "/", "", false, false)
		// 建立 redis 連線
		redis, err := service.GetRedisClient()
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		// 將登入資訊 json encode
		jsonData, _ := json.Marshal(userData)

		// 將登入資訊寫入 redis
		redis.Set(context.Background(), token, string(jsonData), time.Duration(expire)*time.Second)

		// 轉導到後台首頁
		c.Redirect(http.StatusMovedPermanently, "/admin")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "帳號密碼錯誤, 請重新登入",
		})
	}
}

// 登出
func Logout(c *gin.Context) {
	c.SetCookie("loginToken", "", -1, "/", "", false, false)
	c.Redirect(http.StatusMovedPermanently, "/admin/login")
}

func generateSessionID() string {
	// Generate a random session ID
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		// Fallback to current timestamp if random generation fails
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return "auth_" + base64.URLEncoding.EncodeToString(b)
}
