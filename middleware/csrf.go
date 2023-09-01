package middleware

import (
	"fmt"
	"goOrders/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const SecretKey = "csrfToken"

// 前台中間件
func CsrfHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeStamp := strconv.Itoa(int(time.Now().Unix()))
		aes := service.Aes{Key: SecretKey}
		token, err := aes.Encrypt(timeStamp)
		if err != nil {
			fmt.Printf("token encrypt error: %v\n", err.Error())
		}
		c.SetCookie("csrfToken", token, 86400, "/", "", false, false)
		c.Next()
	}
}

// 前台中間件
func CsrfAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		decryptToken, err := c.Cookie("csrfToken")
		if err != nil {
			fmt.Println("取得 cookie csrfToken 失敗:", err)
			c.Abort()
		}
		aes := service.Aes{Key: SecretKey}
		csrfToken, err := aes.Decrypt(decryptToken)
		if err != nil {
			fmt.Println("csrfToken Decrypt error:", err)
			c.Abort()
		}
		fmt.Printf("csrfToken: %v\n", csrfToken)
		now := int(time.Now().Unix())
		intCsrfToken, _ := strconv.Atoi(csrfToken)
		if (now - intCsrfToken) > 86400 {
			fmt.Println("csrfToken 已過期")
			c.Abort()
		}
		c.Next()
	}
}
